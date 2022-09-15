package controller

import (
	"github.com/elysiumyun/elysium/internal/app/model"
	"github.com/elysiumyun/elysium/internal/pkg/datasource"
	"github.com/elysiumyun/elysium/internal/pkg/response"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func UserRegister(ctx *gin.Context) {
	// get user post data
	userDataCtx, ok := ctx.Get("user")
	if !ok {
		response.Fail(ctx, nil, "Get ctx prams failed!")
		return
	}
	userData := userDataCtx.(model.UserData)

	// valid user register data
	if (userData.UserName == "") || (userData.Password == "") {
		response.Fail(ctx, nil, "The username or password cannot be empty")
		return
	}

	if len(userData.Password) < 8 {
		response.Fail(ctx, nil, "The number of password digits is too low")
		return
	}

	// encrypt user passwd
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.DefaultCost)
	if err != nil {
		response.Abnormal(ctx, nil, "Password encryption processing failed")
		return
	}

	newUserData := userData
	newUserData.Password = string(hasedPassword)

	// user register record save
	var newUser model.User = model.User{
		UserData: newUserData,
	}

	db, dbCtx := datasource.Get.Mysql.GetClient()
	tx := db.WithContext(dbCtx)

	if err := tx.Create(&newUser).Error; err != nil {
		response.Abnormal(ctx, nil, "user register failed")
		return
	}

	response.Success(ctx, gin.H{
		"user": newUser,
	}, "Registration was successful")
}
