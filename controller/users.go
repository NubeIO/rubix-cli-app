package controller

import (
	"gthub.com/NubeIO/rubix-cli-app/pkg/model"
	"math/rand"
	"net/http"
	"time"

	"github.com/NubeIO/lib-uuid/uuid"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func getUserBody(c *gin.Context) (dto *model.User, err error) {
	err = c.ShouldBindJSON(&dto)
	return dto, err
}

func (inst *Controller) GetUser(c *gin.Context) {
	data, err := inst.DB.GetUser(c.Params.ByName("uuid"))
	if err != nil {
		reposeHandler(data, err, c)
		return
	}
	reposeHandler(data, err, c)
}

func (inst *Controller) GetUsers(c *gin.Context) {
	data, err := inst.DB.GetUsers()
	if err != nil {
		reposeHandler(data, err, c)
		return
	}
	reposeHandler(data, err, c)
}

func (inst *Controller) UpdateUser(c *gin.Context) {
	body, _ := getUserBody(c)
	data, err := inst.DB.UpdateUser(c.Params.ByName("uuid"), body)
	if err != nil {
		reposeHandler(data, err, c)
		return
	}
	reposeHandler(data, err, c)
}

func (inst *Controller) DeleteUser(c *gin.Context) {
	data, err := inst.DB.DeleteUser(c.Params.ByName("uuid"))
	if err != nil {
		reposeHandler(data, err, c)
		return
	}
	reposeHandler(data, err, c)
}

func (inst *Controller) DropUsers(c *gin.Context) {
	data, err := inst.DB.DropUsers()
	if err != nil {
		reposeHandler(data, err, c)
		return
	}
	reposeHandler(data, err, c)
}

func (inst *Controller) Login(c *gin.Context) (interface{}, error) {
	var loginVals model.LoginUser
	var user model.User
	if err := c.ShouldBindJSON(&loginVals); err != nil {
		return "", jwt.ErrMissingLoginValues
	}
	email := loginVals.Email
	if result := inst.DB.DB.Where("email = ?", email).First(&user); result.Error != nil {
		return "", jwt.ErrFailedAuthentication
	} else {
		if err := bcrypt.CompareHashAndPassword([]byte(user.Hash), []byte(loginVals.Password)); err != nil {
			return "", jwt.ErrFailedAuthentication
		}
		return user, nil
	}
}

const charset string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const length int = 8

func GenerateUID() string {
	var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func (inst *Controller) AddUser(c *gin.Context) {
	var user model.User
	var newUser model.NewUser

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		panic(err.Error())
	}

	if result := inst.DB.DB.Where("email = ?", newUser.Email).First(&user); result.Error != nil {
		// TODO: Differentiate between server error and user user not found error
		hash, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
		if err != nil {
			panic(err)
		}
		user = model.User{Username: newUser.Username, Email: newUser.Email, Hash: string(hash), UID: GenerateUID()}
		user.UUID = uuid.ShortUUID("usr")
		if err := inst.DB.DB.Create(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, &user)
		return
	} else {
		c.JSON(http.StatusConflict, gin.H{"Error": "User already registered"})
		return
	}
}
