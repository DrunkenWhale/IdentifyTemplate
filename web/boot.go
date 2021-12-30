package web

import (
	"IdentifyTemplate/web/extension"
	"IdentifyTemplate/web/model"
	auth3 "IdentifyTemplate/web/router/auth"
	chat3 "IdentifyTemplate/web/router/chat"
	"github.com/gin-gonic/gin"
)

func Init() (*gin.Engine, error) {


	router := gin.Default()

	tableInit()

	routerInit(router)

	return router, nil
}

func routerInit(router *gin.Engine) {
	r := router.Group("/api")

	auth := r.Group("/auth")
	{
		auth.POST("/register", auth3.Register)
		auth.POST("/login", auth3.Login)
	}
	chat := r.Group("/chat")
	{
		chat.POST("/join", chat3.JoinRoom)
		chat.POST("/create", chat3.CreateRoom)
	}


}

func tableInit() {
	_ = extension.DB.AutoMigrate(&model.User{})
}
