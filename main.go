package main

import (
	"codebase-service/config"
	productHandler "codebase-service/handlers/products"
	userHandler "codebase-service/handlers/users"
	"codebase-service/repository/products"
	"codebase-service/repository/users"
	"codebase-service/routes"
	productSvc "codebase-service/usecases/products"
	userSvc "codebase-service/usecases/users"
	"context"
	"database/sql"
	"log"

	"github.com/go-playground/validator"
	"github.com/redis/go-redis/v9"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		return
	}

	dbConn, err := config.ConnectToDatabase(config.Connection{
		Host:     cfg.DBHost,
		Port:     cfg.DBPort,
		User:     cfg.DBUser,
		Password: cfg.DBPassword,
		DBName:   cfg.DBName,
	})
	if err != nil {
		return
	}
	defer dbConn.Close()

	redisConn, err := config.ConnectToRedis(config.RedisConnection{
		Host: cfg.RedisHost,
		Port: cfg.RedisPort,
		Pass: cfg.RedisPass,
		DB:   cfg.RedisDB,
	})
	if err != nil {
		log.Fatalf("cannot connect to redis: %v", err)
		return
	}

	// checj if redis is connected
	err = redisConn.Ping(context.Background()).Err()
	if err != nil {
		log.Fatalf("cannot connect to redis: %v", err)
		return
	} else {
		log.Println("connected to redis")
	}

	validator := validator.New()

	routes := setupRoutes(dbConn, redisConn, validator)
	routes.Run(cfg.AppPort)
}

func setupRoutes(
	db *sql.DB,
	rdb *redis.Client,
	validator *validator.Validate,
) *routes.Routes {
	userStore := users.NewStore(db)
	userSvc := userSvc.NewUserSvc(userStore)
	userHandler := userHandler.NewHandler(userSvc, validator)

	productStore := products.NewStore(db, rdb)
	productSvc := productSvc.NewProductSvc(productStore)
	productHandler := productHandler.NewHandler(productSvc, validator)

	return &routes.Routes{
		User:    userHandler,
		Product: productHandler,
	}
}
