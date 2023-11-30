package controllers

import (
    "net/http"
    "strconv"
    "JobPredictorAPI/models"
    "JobPredictorAPI/services"

    "github.com/gin-gonic/gin"
)

type JobController struct {
    JobService *services.JobService
}

func NewJobController(jobService *services.JobService) *JobController {
    return &JobController{
        JobService: jobService,
    }
}

// GetAllJobs - Retrieves all jobs
func (jc *JobController) GetAllJobs(c *gin.Context) {
    jobs, err := jc.JobService.GetAllJobs(c)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"jobs": jobs})
}

// GetJobRecommendations - Retrieves job recommendations for a user
func (jc *JobController) GetJobRecommendations(c *gin.Context) {
    userID, err := strconv.Atoi(c.Param("userID"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    jobs, err := jc.JobService.GetJobRecommendations(c, userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"recommendations": jobs})
}

// GetJobByID - Retrieves a single job by its ID
func (jc *JobController) GetJobByID(c *gin.Context) {
    jobID, err := strconv.Atoi(c.Param("jobID"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid job ID"})
        return
    }

    job, err := jc.JobService.GetJobByID(c, jobID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"job": job})
}

