package controller

import (
	userModel "vijju/user/model"

	"github.com/gin-gonic/gin"
)

var users []*userModel.User

func NewController(app *gin.Engine) {
	app.GET("/users", GetAllUser)
	app.POST("/users", CreateAllUSer)
	app.GET("/users/{name}", GetUserByName)
}

// GetAllUsers   Return the List of all Users
// GetAllUsers                godoc
// @Summary      Get All Users
// @Description  Returns the all Users.
// @Tags         users
// @Produce      json
// @Success      200  {object}  userModel.User
// @Router       /users [get]
func GetAllUser(ctx *gin.Context) {
	ctx.JSON(200, users)
}

// CreateAllUSer   Create User
// GetAllUsers                godoc
// @Summary      Create User
// @Description  Create Users.
// @Tags         users
// @Produce      json
// @Param user body userModel.User true "userModel.User"
// @Success      200  {object}  userModel.User
// @Router       /users [post]
func CreateAllUSer(ctx *gin.Context) {
	user := &userModel.User{}
	ctx.BindJSON(&user)
	users = append(users, user)
	ctx.JSON(200, user)
}

// GetUserByName   Get User By Name
// GetUserByName                godoc
// @Summary      GetUserByName
// @Description  GetUserByName.
// @Tags         users
// @Produce      json
// @Param        name  path      string  true  "search user by name"
// @Success      200  {object}  userModel.User
// @Router       /users/{name} [get]
func GetUserByName(ctx *gin.Context) {
	user := &userModel.User{}
	name := ctx.Param("name")
	for _, val := range users {
		if val.Name == name {
			ctx.JSON(200, val)
			return
		}
	}
	ctx.JSON(200, user)
}
