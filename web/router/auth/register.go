package auth

import (
	"IdentifyTemplate/web/extension"
	"IdentifyTemplate/web/model"
	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	var user model.User
	err := ctx.Bind(&user)
	if err != nil {
		ctx.JSON(400, "Illegal ContentType")
		return
	}
	if user.Name == "" || user.Password == "" || user.Mailbox == "" {
		ctx.JSON(200, gin.H{
			"status":  0,
			"message": "InvalidArgument",
			"data":    gin.H{},
		})
		return
	}

	var tempUser model.User
	extension.DB.First(&tempUser, "mailbox = ?", user.Mailbox)
	if tempUser.Name != "" || tempUser.Password != "" || tempUser.Mailbox != "" {
		ctx.JSON(200, gin.H{
			"status":  0,
			"message": "UserExist",
			"data":    gin.H{},
		})
		return
	}

	user.Password, err = extension.GeneratePasswordHash(user.Password)
	if err != nil {
		ctx.JSON(200, gin.H{
			"status":  0,
			"message": "UnknownPasswordContentType",
			"data":gin.H{},
		})
	}
	extension.DB.Create(&user)
	ctx.JSON(200, gin.H{
		"status":  1,
		"message": "Succeed",
		"data":    gin.H{},
	})
	return
}
