package logger

import (
	"log"
)

func InitLogger() {
	log.Println("Logger initialized")
}

# File 8: .env
SERVER_PORT=:8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=product_user
DB_PASSWORD=securepassword
DB_NAME=product_system
REDIS_HOST=localhost:6379
QUEUE_HOST=amqp://guest:guest@localhost:5672/