package controllers

import (
	"backend/middleware"
	"backend/models"
	"backend/requests"
	"backend/utils"
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type UserController struct {
	db          *gorm.DB
	ctx         context.Context
	redisClient *redis.Client
}

type Signup struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewUserController(db *gorm.DB, ctx context.Context, redisClient *redis.Client) *UserController {
	return &UserController{
		db:          db,
		ctx:         ctx,
		redisClient: redisClient,
	}
}

func (uc *UserController) CreateUser(c *fiber.Ctx) error {
	//Create a new signupReq Object
	signupReq := new(Signup)
	//Parse data into the object
	if err := c.BodyParser(signupReq); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Bad Request")
	}
	// Hash the password for the user
	hashedPassword, err := utils.HashPassword(signupReq.Password)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Server Error")
	}
	// Create a new user
	user := new(models.User)
	user.CreatedAt = time.Now()
	user.Username = signupReq.Username
	user.Password = hashedPassword
	user.Permissions = make([]models.Permissions, 0)
	//Save the user in PostgreSQL
	result := uc.db.Create(user)
	if result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Unable to save user")
	}
	log.Println("User Created", user.ID)
	return c.JSON(fiber.Map{"message": "Success"})
}

func (uc *UserController) Login(c *fiber.Ctx) error {
	// Reusing the same struct for login
	loginReq := new(requests.Signup) // or create a separate Login struct if needed
	if err := c.BodyParser(loginReq); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Bad Request")
	}

	user := new(models.User)
	// Finding user with permissions loaded
	err := uc.db.Preload("Permissions").Where("username = ?", loginReq.Username).First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return fiber.NewError(fiber.StatusUnauthorized, "Invalid credentials")
		}
		return fiber.NewError(fiber.StatusInternalServerError, "Database error")
	}

	// Verifying user password
	err = utils.VerifyPassword(loginReq.Password, user.Password)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid credentials")
	}

	// Create JWT token using middleware
	tokenString, err := middleware.CreateJWTToken(user, uc.redisClient)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Token generation failed")
	}

	return c.JSON(fiber.Map{
		"token":    tokenString,
		"message":  "Login successful",
		"user_id":  user.ID,
		"username": user.Username,
	})
}

func (uc *UserController) AddPermission(c *fiber.Ctx) error {
	addPermissionReq := new(requests.AddPermission)
	if err := c.BodyParser(addPermissionReq); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Bad Request")
	}

	user := new(models.User)
	//Find the user with their existing permissions
	err := uc.db.Preload("Permissions").Where("username = ?", addPermissionReq.Username).First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return fiber.NewError(fiber.StatusNotFound, "User not found")
		}
		return fiber.NewError(fiber.StatusInternalServerError, "Database error")
	}
	log.Println("User Received", user)

	//Check if permissions for the given entity already exists
	for _, v := range user.Permissions {
		if v.Entry == addPermissionReq.Permission.Entry {
			return fiber.NewError(fiber.StatusBadRequest, "Permission already exists")
		}
	}

	//Add the new permission to the user
	newPermission := models.Permissions{
		UserID:    user.ID,
		Entry:     addPermissionReq.Permission.Entry,
		AddFlag:   addPermissionReq.Permission.AddFlag,
		AdminFlag: addPermissionReq.Permission.AdminFlag,
	}

	//Create the new permission record
	err = uc.db.Create(&newPermission).Error
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Unable to add permission")
	}

	return c.JSON(fiber.Map{"message": "Success"})
}
