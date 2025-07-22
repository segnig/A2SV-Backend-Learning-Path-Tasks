package helpers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func CheckUserType(ctx *gin.Context, role string) (err error) {
	UserType := ctx.GetString("user_type")

	if UserType != role {
		return errors.New("Unauthorized to access this resource")
	}
	return nil
}

func MatchUserTypeToUid(ctx *gin.Context, userID string) (err error) {
	Uid := ctx.GetString("user_id")
	UserType := ctx.GetString("user_type")

	if UserType == "USER" && userID != Uid {
		return errors.New("Unauthorized to access this resource")
	}

	err = CheckUserType(ctx, "ADMIN")
	return err
}
