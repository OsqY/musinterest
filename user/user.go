package user

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Handler(ctx *gin.Context) {
	session := sessions.Default(ctx)
	profile := session.Get("profile")
	ctx.JSON(200, gin.H{"profile": profile})
}
