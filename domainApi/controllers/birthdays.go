package controllers

import (
	"fmt"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/domainApi/repositories"
	"github.com/JackMaarek/go-bot-utils/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateUserBirthday receives informations for user birthday and saves it to database
func CreateUserBirthday(c *gin.Context) {
	var birthday models.Birthday
	if err := c.ShouldBindJSON(&birthday); err != nil {
		c.JSON(http.StatusBadRequest, "request should contain valid properties")
		return
	}

	foundBirthday := models.Birthday{
		UserID: birthday.UserID,
	}

	if err := repositories.FindBirthdayByUserID(&foundBirthday); err == nil {
		c.JSON(http.StatusOK, fmt.Sprintf("birthday already exists for user id %s", foundBirthday.UserID))
		return
	}

	if err := repositories.PersistBirthday(&birthday); err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("could not create birthday for user id %s", birthday.UserID))
		return
	}

	c.JSON(http.StatusOK, "user's birthday added successfully")
}
