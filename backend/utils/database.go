package utils

import (
	"context"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var RedisClient *redis.Client // Add this global variable
var ctx context.Context

func ConnectDatabase() (*gorm.DB, error) {
	ctx = context.Background()

	// PostgreSQL connection
	dsn := os.Getenv("DATABASE_DSN")
	if dsn == "" {
		dsn = "host=localhost user=myuser password=mypassword dbname=ecommerce port=5432 sslmode=disable"
	}

	// Redis connection
	RedisClient = redis.NewClient(&redis.Options{ // Assign to global variable
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	status := RedisClient.Ping(ctx)
	if status.Err() != nil {
		log.Printf("Redis connection failed: %v (continuing without Redis)", status.Err())
		RedisClient = nil
	} else {
		log.Println("Redis connected successfully:", status.Val())
	}

	// PostgreSQL connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Database connection failed: %v", err)
		return nil, err
	}

	DB = db // Assign to global variable
	log.Println("Database connected successfully")
	return db, nil
}
