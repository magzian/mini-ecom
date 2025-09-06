package controllers

import (
	"backend/models"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Claims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// Register - Converted to Fiber with password hashing
func Register(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Invalid request body",
		})
	}

	// Basic validation
	if data["username"] == "" || data["password"] == "" {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Username and password are required",
		})
	}

	// Basic password strength validation
	if len(data["password"]) < 6 {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Password must be at least 6 characters long",
		})
	}

	// Clean username (remove spaces, convert to lowercase)
	username := strings.ToLower(strings.TrimSpace(data["username"]))

	db := c.Locals("db").(*gorm.DB)

	// Check if user already exists
	var existingUser models.User
	if err := db.Where("username = ?", username).First(&existingUser).Error; err == nil {
		return c.Status(409).JSON(fiber.Map{
			"success": false,
			"message": "Username already exists",
		})
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data["password"]), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Error processing password",
		})
	}

	// Create user
	user := models.User{
		Username: username,
		Password: string(hashedPassword),
	}

	if err := db.Create(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Failed to create user",
		})
	}

	// Return user info (without password)
	return c.Status(201).JSON(fiber.Map{
		"success": true,
		"message": "User created successfully",
		"data": fiber.Map{
			"user_id":  user.ID,
			"username": user.Username,
		},
	})
}

// Login - Updated with password hashing verification
func Login(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Invalid request body",
		})
	}

	// Validate input
	if data["username"] == "" {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Username is required",
		})
	}

	if data["password"] == "" {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Password is required",
		})
	}

	// Clean username
	username := strings.ToLower(strings.TrimSpace(data["username"]))

	db := c.Locals("db").(*gorm.DB)

	// Find user
	var user models.User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		// Debug: User not found
		fmt.Println("DEBUG: User not found in database for username:", username)
		return c.Status(401).JSON(fiber.Map{
			"success": false,
			"message": "Invalid credentials",
		})
	}

	fmt.Println("DEBUG: Found user:", user.Username, "with ID:", user.ID)
	fmt.Println("DEBUG: Stored password hash:", user.Password)
	fmt.Println("DEBUG: Login attempt with password:", data["password"])

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"])); err != nil {
		fmt.Println("DEBUG: Password verification failed:", err)
		return c.Status(401).JSON(fiber.Map{
			"success": false,
			"message": "Invalid credentials",
		})
	}

	fmt.Println("DEBUG: Password verification successful")

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":    user.ID,
		"username":   user.Username,
		"issued_at":  time.Now().Unix(),
		"expires_at": time.Now().Add(time.Hour * 24).Unix(),
	})

	// Get JWT secret
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "default-secret-key" // Fallback for development
	}

	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Could not generate token",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Login successful",
		"data": fiber.Map{
			"token":    tokenString,
			"user_id":  user.ID,
			"username": user.Username,
		},
	})
}

// Middleware to verify JWT token
func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(401).JSON(fiber.Map{
				"success": false,
				"message": "Authorization header required",
			})
		}

		// Extract token from "Bearer <token>"
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			return c.Status(401).JSON(fiber.Map{
				"success": false,
				"message": "Invalid authorization format",
			})
		}

		// Parse and validate token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			jwtSecret := os.Getenv("JWT_SECRET")
			if jwtSecret == "" {
				jwtSecret = "default-secret-key"
			}
			return []byte(jwtSecret), nil
		})

		if err != nil || !token.Valid {
			return c.Status(401).JSON(fiber.Map{
				"success": false,
				"message": "Invalid or expired token",
			})
		}

		// Extract claims
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Locals("user_id", claims["user_id"])
			c.Locals("username", claims["username"])
		}

		return c.Next()
	}
}

// Protected route example - Get current user info
func GetProfile(c *fiber.Ctx) error {
	userID := c.Locals("user_id")
	username := c.Locals("username")

	return c.JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"user_id":  userID,
			"username": username,
		},
	})
}
