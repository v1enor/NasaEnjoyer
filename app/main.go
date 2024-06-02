package main

import (
	"NasaEnjoyer/apod"
	"fmt"
	"net/http"
	"strconv"

	"NasaEnjoyer/domain"
	"NasaEnjoyer/internal/repository/filesystem"
	"NasaEnjoyer/internal/repository/postgresql"
	"NasaEnjoyer/internal/rest"
	"NasaEnjoyer/internal/rest/middleware"
	"NasaEnjoyer/internal/workers"
	"NasaEnjoyer/pkg/nasa"
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	defaultTimeout       = 10
	defaultTimeoutWorker = 24
	defaultAddress       = ":9090"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func initDB(databaseURL string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	err = db.AutoMigrate(&domain.APOD{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Println("Successfully connected to the database")
	return db
}

func main() {
	// Загрузка переменных окружения
	nasaAPIKey := os.Getenv("NASA_API_KEY")
	nasaURL := os.Getenv("NASA_URL")
	imageSaveDir := os.Getenv("IMAGE_SAVE_DIR")
	timeoutStr := os.Getenv("CONTEXT_TIMEOUT")
	timeoutWorkerStr := os.Getenv("WORKER_TIMEOUT")
	serverAddress := os.Getenv("SERVER_ADDRESS")
	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbUser := os.Getenv("DATABASE_USER")
	dbPassword := os.Getenv("DATABASE_PASSWORD")
	dbName := os.Getenv("DATABASE_NAME")
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)

	// Инициализация базы данных
	db := initDB(connStr)

	apodRepo := postgresql.NewAPODRepository(db)
	nasaCli := nasa.NewNASAClient(nasaAPIKey, nasaURL, http.DefaultClient)
	fileRepo := filesystem.NewFileRepository()

	// Инициализация сервиса
	service := apod.NewService(apodRepo, nasaCli, fileRepo, imageSaveDir)

	// Инициализация HTTP сервера
	e := echo.New()
	e.Use(middleware.CORS)
	timeout, err := strconv.Atoi(timeoutStr)
	if err != nil {
		log.Println("failed to parse timeout, using default timeout")
		timeout = defaultTimeout
	}
	timeoutContext := time.Duration(timeout) * time.Second
	e.Use(middleware.SetRequestContextWithTimeout(timeoutContext))

	// Запуск воркера
	timeoutWorker, err := strconv.Atoi(timeoutWorkerStr)
	if err != nil {
		log.Println("failed to parse worker timeout, using default timeout")
		timeoutWorker = defaultTimeoutWorker
	}
	worker := workers.NewAPODWorker(service, time.Duration(timeoutWorker)*time.Hour)
	go worker.Start(context.Background(), timeoutContext)

	// Инициализация хэндлеров
	rest.NewAPODHandler(service, e)

	if serverAddress == "" {
		serverAddress = defaultAddress
	}
	log.Fatal(e.Start(serverAddress))
}
