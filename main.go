package main

import (
	"JobPredictorAPI/controllers"
	"JobPredictorAPI/router"
	"JobPredictorAPI/services"

	_ "github.com/lib/pq" // PostgreSQL driver
)

func main() {
	// PostgreSQL connection string
	//const connStr = "postgresql://austine:wik@localhost/job_search_db"

	Local := "postgresql://postgres:philippians@localhost:5432/postgres?sslmode=disable"
	//dsnString := os.Getenv("Local")
	// Open a DB connection
	db, err := services.ConnectToDB(Local)
	if err != nil {
		return
	}
	// Initialize services
	jobService := services.NewJobService(db)
	userService := services.NewUserService(db)
	savedJobsService := services.NewSavedJobsService(db)
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
