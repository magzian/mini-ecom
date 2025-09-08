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
var ctx context.Context

func ConnectDatabase() (*gorm.DB, error) {
	ctx = context.Background()
	dsn := os.Getenv("DATABASE_DSN")
	if dsn == "" {
		dsn = "host=localhost user=myuser password=mypassword dbname=ecommerce port=5432 sslmode=disable"
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0})
	status := redisClient.Ping(ctx)
	log.Println(status)

	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
