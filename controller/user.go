package controller

import (
	"net/http"
	"user_api/model"

	"github.com/gin-gonic/gin"
)

func Create(context *gin.Context) {
	var input model.UserInput

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := model.User{
		Username: input.Username,
		Password: input.Password,
		Admin:    input.Admin,
	}

	savedUser, err := user.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"user": savedUser})
}

func FindAll(context *gin.Context) {
	users, err := model.FindAll()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"user": users})
}

func Update(context *gin.Context) {
	var input model.UserUpdateInput
	username := context.Param("id")

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := model.Find(username)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.Password = input.Password
	user.Admin = input.Admin || false
	updatedUser, er := user.Update()

	if er != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"user": updatedUser})
}

func Delete(context *gin.Context) {

}
