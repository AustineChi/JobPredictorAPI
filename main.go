package main

import (
	"JobPredictorAPI/config"
	"JobPredictorAPI/controllers"
	"JobPredictorAPI/router"
	"JobPredictorAPI/services"
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq" // PostgreSQL driver
)

func main() {
    // PostgreSQL connection string
    //const connStr = "postgresql://austine:wik@localhost/job_search_db"

    dsnString := os.Getenv("DSN")
    // Open a DB connection
    db, err := sql.Open("postgres", dsnString)
    if err != nil {
        log.Fatalf("Error opening database: %v", err)
    }

    // Check if the database is connected
    err = db.Ping()
    if err != nil {
        log.Fatalf("Error connecting to the database: %v", err)
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
