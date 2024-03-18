package main

import (
	"JobPredictorAPI/controllers"
	"JobPredictorAPI/models"
	"JobPredictorAPI/router"
	"JobPredictorAPI/services"
	"log"
	"os"
	_ "github.com/lib/pq" // PostgreSQL driver
)

func main() {
	// PostgreSQL connection string
	//const connStr = "postgresql://austine:wik@localhost/job_search_db"

	// Load environment variables from .env file
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }
	dsnString := os.Getenv("Local")
	log.Println(dsnString)
	// Open a DB connection
	db, err := models.ConnectToDB(dsnString)
	if err != nil {
		return
	}
	// Initialize services
	jobService := services.NewJobService(db)
	userService := services.NewUserService(db)
	savedJobsService := services.NewSavedJobsService(db, jobService)
	notificationService := services.NewNotificationService(db)
	interactionService := services.NewInteractionService(db)

	// Initialize controllers
	userCtrl := controllers.NewUserController(userService)
	jobCtrl := controllers.NewJobController(jobService)
	savedJobsCtrl := controllers.NewSavedJobsController(savedJobsService)
	notificationCtrl := controllers.NewNotificationController(notificationService)
	interactionCtrl := controllers.NewInteractionController(interactionService)

	// Set up and start the Gin router
	r := router.SetupRouter(userCtrl, jobCtrl, savedJobsCtrl, notificationCtrl, interactionCtrl)
	r.Run() // By default, it runs on http://localhost:8080
}
