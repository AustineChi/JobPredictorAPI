package config

import (
    "log"
    "github.com/robfig/cron/v3"
    "JobPredictorAPI/services" 
)

func SetupCron(jobService *services.JobService) {
    c := cron.New()
    _, err := c.AddFunc("@daily", func() {
        err := jobService.FetchAndStoreJobs()
        if err != nil {
            log.Printf("Error fetching and storing jobs: %v", err)
        }
    })

    if err != nil {
        log.Fatalf("Error scheduling FetchAndStoreJobs: %v", err)
    }

    c.Start()
}
