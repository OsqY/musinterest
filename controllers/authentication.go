package controllers

import (
	"net/http"
	"oscar/musinterest/helpers"
	"oscar/musinterest/initializers"
	"oscar/musinterest/models"
	"oscar/musinterest/utils"

	"github.com/gin-gonic/gin"
	"github.com/thanhpk/randstr"
	"gorm.io/gorm"
)

type AuthController struct {
	DB *gorm.DB
}

func NewAuthController(DB *gorm.DB) AuthController {
	return AuthController{DB}
}

func (ac *AuthController) Register(context *gin.Context) {
	var input models.AuthenticationInput

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := models.User{
		Email:    input.Email,
		Username: input.Username,
		Password: input.Password,
		Verified: true,
	}
	savedUser, err := user.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	config, _ := initializers.LoadConfig(".")

	code := randstr.String(20)

	username := savedUser.Username

	emailData := utils.EmailData{
		URL:       config.ClientOrigin + "/verifyemail/" + code,
		FirstName: username,
		Subject:   "Code verification for your Musinterest account",
	}

	utils.SendEmail(savedUser, &emailData)
	context.JSON(http.StatusCreated, gin.H{"data": savedUser})
}

func (ac *AuthController) Login(context *gin.Context) {
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

	if !user.Verified {
		context.JSON(http.StatusForbidden, gin.H{"error": "Please verify your email"})
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

func (authController *AuthController) VerifyEmail(context *gin.Context) {
	code := context.Params.ByName("verificationCode")
	verificationCode := utils.Encode(code)

	var updatedUser models.User
	if err := authController.DB.First(&updatedUser, "verificationCode = ?", verificationCode).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid verification code or user is invalid!"})
		return
	}

	if updatedUser.Verified {
		context.JSON(http.StatusConflict, gin.H{"error": "User already verified"})
		return
	}

	updatedUser.Verified = true
	authController.DB.Save(&updatedUser)
	context.JSON(http.StatusOK, gin.H{"status": "User verified!"})
}
