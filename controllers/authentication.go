package controllers

import (
	"net/http"
	"oscar/musinterest/helpers"
	"oscar/musinterest/models"

	"github.com/gin-gonic/gin"
)

func Register(context *gin.Context) {
	var input models.AuthenticationInput

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := models.User{
		Email:    input.Email,
		Username: input.Username,
		Password: input.Password,
	}
	savedUser, err := user.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": savedUser})
}

func Login(context *gin.Context) {
	var input models.AuthenticationInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Input doesn't the requirements'"})
		return
	}

	user, err := models.FindByUsername(input.Username)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "That user doesn't exists"})
		return
	}

	if err := user.ValidatePassword(input.Password); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Wrong password"})
		return
	}

	jwt, err := helpers.GenerateJWT(user)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	context.JSON(http.StatusOK, gin.H{"jwt": jwt})
}
