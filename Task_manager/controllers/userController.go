package controllers

import (
	"context"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"task-manager/database"
	"task-manager/helpers"
	"task-manager/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var UserCollection *mongo.Collection = database.OpenCollection(*database.Client, "user")
var validate = validator.New()

func VerifyPassword(hashedPwd, plainPwd string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPwd))
	if err != nil {
		return false, "username or password is incorrect"
	}
	return true, ""
}

func HashPassword(userPassword string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userPassword), 12)
	if err != nil {
		log.Println("error hashing password:", err)
		return ""
	}
	return string(hashedPassword)
}

func ValidateUsername(username string) bool {
	var validUsername = regexp.MustCompile(`^[A-Za-z][A-Za-z0-9_]*$`)
	return validUsername.MatchString(username)
}

func Signup() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var user models.User
		if err := ctx.BindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := validate.Struct(user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		count, err := UserCollection.CountDocuments(c, bson.M{"username": user.Username})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while validating username"})
			return
		}
		if count > 0 {
			ctx.JSON(http.StatusConflict, gin.H{"error": "username already exists"})
			return
		}

		totalUsers, err := UserCollection.CountDocuments(context.TODO(), bson.D{})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to count users"})
			return
		}
		if totalUsers == 0 && user.UserType != "ADMIN" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Only an ADMIN can create the first user"})
			return
		}

		user.Password = HashPassword(user.Password)
		user.CreatedAt = time.Now()
		user.UpdatedAt = time.Now()
		user.ID = primitive.NewObjectID()
		user.UserID = user.ID.Hex()

		token, refreshToken, err := helpers.GenerateAllTokens(user.Username, user.FirstName, user.LastName, user.UserType, user.UserID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate tokens"})
			return
		}

		user.Token = token
		user.RefreshToken = refreshToken

		_, insertErr := UserCollection.InsertOne(c, user)
		if insertErr != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "User creation failed"})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
	}
}

func Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var user models.User
		var foundUser models.User

		if err := ctx.BindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := UserCollection.FindOne(c, bson.M{"username": user.Username}).Decode(&foundUser)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "username or password is incorrect"})
			return
		}

		passwordIsValid, msg := VerifyPassword(foundUser.Password, user.Password)
		if !passwordIsValid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": msg})
			return
		}

		token, refreshToken, err := helpers.GenerateAllTokens(foundUser.Username, foundUser.FirstName, foundUser.LastName, foundUser.UserType, foundUser.UserID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate tokens"})
			return
		}

		helpers.UpdateAllToken(token, refreshToken, foundUser.UserID)
		foundUser.Token = token
		foundUser.RefreshToken = refreshToken

		ctx.JSON(http.StatusOK, gin.H{
			"user_id":   foundUser.UserID,
			"username":  foundUser.Username,
			"firstName": foundUser.FirstName,
			"lastName":  foundUser.LastName,
			"userType":  foundUser.UserType,
			"token":     foundUser.Token,
		})
	}
}

func GetUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID := ctx.Param("user_id")

		if err := helpers.MatchUserTypeToUid(ctx, userID); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var user models.User
		err := UserCollection.FindOne(c, bson.M{"user_id": userID}).Decode(&user)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, user)
	}
}

func GetUsers() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := helpers.CheckUserType(ctx, "ADMIN"); err != nil {
			ctx.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}

		c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		recordPerPage, _ := strconv.Atoi(ctx.DefaultQuery("recordPerPage", "10"))
		page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
		startIndex := (page - 1) * recordPerPage

		matchStage := bson.D{{"$match", bson.D{}}}
		groupStage := bson.D{{"$group", bson.D{
			{"_id", nil},
			{"total_count", bson.D{{"$sum", 1}}},
			{"data", bson.D{{"$push", "$$ROOT"}}},
		}}}
		projectStage := bson.D{{"$project", bson.D{
			{"_id", 0},
			{"total_count", 1},
			{"user_items", bson.D{{"$slice", bson.A{"$data", startIndex, recordPerPage}}}},
		}}}

		result, err := UserCollection.Aggregate(c, mongo.Pipeline{matchStage, groupStage, projectStage})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while listing user items"})
			return
		}

		var allUsers []bson.M
		if err := result.All(c, &allUsers); err != nil || len(allUsers) == 0 {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not retrieve users"})
			return
		}

		ctx.JSON(http.StatusOK, allUsers[0])
	}
}
