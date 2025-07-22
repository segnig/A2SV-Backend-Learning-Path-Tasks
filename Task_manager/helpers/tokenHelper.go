package helpers

import (
	"context"
	"log"
	"os"
	"task-manager/database"
	"time"

	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SignedDetails struct {
	Username  string
	FirstName string
	LastName  string
	Uid       string
	UserType  string
	jwt.StandardClaims
}

var UserCollection *mongo.Collection = database.OpenCollection(*database.Client, "user")
var SECRET_KEY string = os.Getenv("SECRET_KEY")

func init() {
	if SECRET_KEY == "" {
		log.Fatal("SECRET_KEY environment variable is not set")
	}
}

func UpdateAllToken(signedToken, signedRefreshToken, userID string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	updateObj := bson.D{
		{"token", signedToken},
		{"refresh_token", signedRefreshToken},
		{"updated_at", time.Now()},
	}

	upsert := true
	filter := bson.M{"user_id": userID}
	opt := options.UpdateOptions{Upsert: &upsert}

	_, err := UserCollection.UpdateOne(
		ctx,
		filter,
		bson.D{{"$set", updateObj}},
		&opt,
	)
	if err != nil {
		log.Println("Failed to update tokens:", err)
	}
}

func GenerateAllTokens(username, firstName, lastName, userType, userId string) (signedToken, signedRefreshToken string, err error) {
	claims := &SignedDetails{
		Username:  username,
		FirstName: firstName,
		LastName:  lastName,
		Uid:       userId,
		UserType:  userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}

	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(200 * time.Hour).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Println("Failed to sign token:", err)
		return "", "", err
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Println("Failed to sign refresh token:", err)
		return "", "", err
	}

	return token, refreshToken, nil
}

func ValidateToken(signedToken string) (claims *SignedDetails, msg string) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)

	if err != nil {
		msg = err.Error()
		return nil, msg
	}

	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		msg = "the token is invalid"
		return nil, msg
	}

	if claims.ExpiresAt < time.Now().Unix() {
		msg = "token is expired"
		return nil, msg
	}

	return claims, ""
}
