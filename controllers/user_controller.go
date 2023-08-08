package controllers

import (
	"net/http"
	"users-itsva/models"
	"users-itsva/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Returns all users
func GetUsers(ctx *gin.Context) {
	// get db from the middleware
	db := ctx.MustGet("DB").(*gorm.DB)
	var users []models.User

	// Get all records
	// SELECT * FROM users;
	if result := db.Find(&users); result.Error != nil {
		// returns error
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}
	// returns all users
	ctx.JSON(http.StatusOK, gin.H{
		"users": &users,
	})
}

// Create a single user
func CreateUser(ctx *gin.Context) {
	// get db from the middleware
	db := ctx.MustGet("DB").(*gorm.DB)
	var newUser models.User

	// if json object matches user type
	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		// returns error
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error JSON": err.Error(),
		})
		return
	}

	// Add the username to newUser
	newUser.Username = utils.GetUserNameFromEmail(newUser.Email)

	var err error
	// if password is valid an encrypted
	if newUser.Password, err = utils.GetEncryptPasswordAndValidates(newUser.Password); err != nil {
		// returns error
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error password": err.Error(),
		})
		return
	}

	// if we can Create newUser
	if result := db.Create(&newUser); result.Error != nil {
		// returns error
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error create user": result.Error.Error(),
		})
		return
	}

	// Returns Succes!
	ctx.JSON(http.StatusOK, gin.H{
		"ok": "Usuario creado",
	})
}

// Route return user from username
func GetUserByUsername(ctx *gin.Context) {
	// get db from the middleware
	db := ctx.MustGet("DB").(*gorm.DB)
	// Get the value of the "username" parameter from the URL
	username := ctx.Param("username")

	// Createa  user in blank
	var user = models.User{}

	// if we can find the user
	// SELECT * FROM users WHERE username = username;
	if result := db.First(&user, "username = ? ", username); result.Error != nil {
		// returns error
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error ": result.Error.Error(),
		})
		return
	}

	// Returns the user found
	ctx.JSON(http.StatusOK, gin.H{
		"ok": user,
	})
}

func UpdateUserByUsername(ctx *gin.Context) {
	// get db from the middleware
	db := ctx.MustGet("DB").(*gorm.DB)
	// Get the value of the "username" parameter from the URL
	username := ctx.Param("username")

	// Createa user in blank
	var user = models.User{}

	// if we can find the user
	// SELECT * FROM users WHERE username = username;
	if result := db.First(&user, "username = ? ", username); result.Error != nil {
		// returns error
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error ": result.Error.Error(),
		})
		return
	}

	//  Create updatedUser
	var updatedUser = models.UpdateUser{}
	// if json object matches user type
	if err := ctx.ShouldBindJSON(&updatedUser); err != nil {
		// returns error
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error JSON": err.Error(),
		})
		return
	}

	var err error
	// if password is valid an encrypted
	if user.Password, err = utils.GetEncryptPasswordAndValidates(updatedUser.Password); err != nil {
		// returns error
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error password": err.Error(),
		})
		return
	}

	// if save in the database
	if result := db.Save(&user); result.Error != nil {
		// returns error
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error ": result.Error.Error(),
		})
		return
	}

	// Returns a message
	ctx.JSON(http.StatusOK, gin.H{
		"message": "usuario actualizado con exito!",
	})
}

func DeleteUserByUsername(ctx *gin.Context) {
	// get db from the middleware
	db := ctx.MustGet("DB").(*gorm.DB)
	// Get the value of the "username" parameter from the URL
	username := ctx.Param("username")

	// Createa user in blank
	var user = models.User{}

	// if we can find the user
	// SELECT * FROM users WHERE username = username;
	if result := db.First(&user, "username = ? ", username); result.Error != nil {
		// returns error
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error ": result.Error.Error(),
		})
		return
	}

	// DELETE from users where id = user.id;
	if result := db.Unscoped().Delete(&user); result.Error != nil {
		// returns error
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error ": result.Error.Error(),
		})
		return
	}

	// Returns a message
	ctx.JSON(http.StatusOK, gin.H{
		"message": "usuario elminado!",
	})

}
