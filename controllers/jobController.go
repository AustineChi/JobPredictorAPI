package controllers

import (
	"JobPredictorAPI/models"
	"JobPredictorAPI/services"
	"net/http"
	"strconv"

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

func (jc *JobController) GetJob(c *gin.Context) {
	jobID, err := strconv.Atoi(c.Param("jobID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unable to get JobID", "details": err.Error()})
		return
	}
	job, err := jc.JobService.GetJobByID(c, jobID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"job": job})
}

func (jc *JobController) CreateJob(c *gin.Context) {
	var JobModel models.Job

	err := c.ShouldBindJSON(JobModel)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid Request Payload", "details": err.Error()})
	}
	newJob, err := jc.JobService.CreateJob(c, &JobModel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to create"})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"JobCreated": newJob})
}

func (jc *JobController) UpdateJob(c *gin.Context) {
	jobID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	var updatedJob models.Job
	if err := c.ShouldBindJSON(&updatedJob); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedJob.JobID = jobID
	err = jc.JobService.UpdateJob(c, &updatedJob)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Job updated successfully"})

}

func (jc *JobController) DeleteJob(c *gin.Context) {
	jobID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var currentJob models.Job
	currentJob.JobID = jobID
	err = jc.JobService.DeleteJob(c, jobID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Job deleted Successfully"})

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
	userIDInt, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
    //convert back to string
    userID := strconv.Itoa(userIDInt)
	jobs, err := jc.JobService.GetJobRecommendations(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"recommendations": jobs})
}

// GetJobByID - Retrieves a single job by its ID
func (jc *JobController) GetJobByID(c *gin.Context) {
	jobID, err := strconv.Atoi(c.Param("id"))
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
