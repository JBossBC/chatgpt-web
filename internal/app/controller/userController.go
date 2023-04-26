package controller

import "github.com/gin-gonic/gin"

func loginHandler(context *gin.Context) {
	// get
	username := context.Query("username")
	password := context.Query("password")

}
