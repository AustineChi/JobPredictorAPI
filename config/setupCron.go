package config

import (
	"JobPredictorAPI/services"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

func SetupCron(jobService *services.JobService) {
    var g *gin.Context
    c := cron.New()
    _, err := c.AddFunc("@daily", func() {
        err := jobService.FetchAndStoreJobs(g)
        if err != nil {
            log.Printf("Error fetching and storing jobs: %v", err)
        }
    })

    if err != nil {
        log.Fatalf("Error scheduling FetchAndStoreJobs: %v", err)
    }

    c.Start()
}
