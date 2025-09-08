package middleware

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"backend/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/redis/go-redis/v9"
)

var SecretKey = []byte("SecretKey")

// CreateJWTToken creates a JWT token for a user with their permissions
func CreateJWTToken(user *models.User, redisClient *redis.Client) (string, error) {
	// Creating hash for user permissions
	objStr := fmt.Sprintf("%+v", user.Permissions)
	data := []byte(objStr)
	hasher := sha256.New()
	_, err := hasher.Write(data)
	if err != nil {
		log.Printf("Error creating hash: %v", err)
		return "", err
	}

	hash := hasher.Sum(nil)
	hashString := hex.EncodeToString(hash)

	// Converting the hash to a JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["hash"] = hashString
	claims["user_id"] = user.ID
	claims["username"] = user.Username
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	// Setting the hash based value in redis
	permissionsJSON, err := json.Marshal(user.Permissions)
	if err != nil {
		log.Printf("Error marshaling permissions: %v", err)
		return "", err
	}

	result, err := redisClient.SetNX(context.Background(), hashString, permissionsJSON, time.Hour*1).Result()
	if err != nil {
		log.Printf("Redis error: %v", err)
		return "", err
	}
	log.Println("Result from redis", result)

	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		log.Printf("Error signing token: %v", err)
		return "", err
	}

	log.Println("JWT Token:", tokenString)
	return tokenString, nil
}

// ValidateJWTToken validates and parses a JWT token
func ValidateJWTToken(tokenString string) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return SecretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return &claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

// JWTMiddleware validates JWT tokens for protected routes
func JWTMiddleware(redisClient *redis.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get token from Authorization header
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return fiber.NewError(fiber.StatusUnauthorized, "Missing authorization header")
		}

		// Remove "Bearer " prefix
		tokenString := ""
		if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
			tokenString = authHeader[7:]
		} else {
			return fiber.NewError(fiber.StatusUnauthorized, "Invalid authorization format")
		}

		// Validate token
		claims, err := ValidateJWTToken(tokenString)
		if err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, "Invalid token")
		}

		// Check if token is expired
		if exp, ok := (*claims)["exp"].(float64); ok {
			if time.Now().Unix() > int64(exp) {
				return fiber.NewError(fiber.StatusUnauthorized, "Token expired")
			}
		}

		// Store claims in context for use in handlers
		c.Locals("user_id", (*claims)["user_id"])
		c.Locals("username", (*claims)["username"])
		c.Locals("hash", (*claims)["hash"])

		return c.Next()
	}
}
