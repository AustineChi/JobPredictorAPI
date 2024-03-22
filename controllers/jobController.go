package controllers

import (
	"JobPredictorAPI/models"
	"JobPredictorAPI/services"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type JobController struct {
	JobService *services.JobService
	UserService *services.UserService
}

func NewJobController(jobService *services.JobService, userService *services.UserService) *JobController {
	return &JobController{
		JobService: jobService,
		UserService: userService,
	}
}

func (jc *JobController) GetJob(c *gin.Context) {
	//get Job from URL parameters
	idParam := c.Param("id")
	if idParam == "" {
		log.Println("job ID is empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Job ID param is empty"})
		return
	}
	jobID, err := strconv.Atoi(idParam)
	if err != nil {
		log.Println("unable to access JobID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "unable to get JobID", "details": err.Error()})
		return
	}
	job, err := jc.JobService.GetJobByID(c, jobID)
	if err != nil {
		log.Println("unable to get jobID:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"job": job})
}

func (jc *JobController) CreateJob(c *gin.Context) {
	var JobModel models.Job

	err := c.ShouldBindJSON(&JobModel)
	if err != nil {
		log.Println("invalid request Payload:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid Request Payload", "details": err.Error()})
	}
	newJob, err := jc.JobService.CreateJob(c, &JobModel)
	if err != nil {
		log.Println("unable to create Job:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to create"})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"JobCreated": newJob})
}

func (jc *JobController) UpdateJob(c *gin.Context) {
	jobID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println("invalid user ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	var updatedJob models.Job
	if err := c.ShouldBindJSON(&updatedJob); err != nil {
		log.Println("cannot Bind JSON", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedJob.JobID = jobID
	updated, err := jc.JobService.UpdateJob(c, &updatedJob)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Job updated successfully",
		"details": updated,
	})
}

func (jc *JobController) DeleteJob(c *gin.Context) {
	jobID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println("invalid user ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var currentJob models.Job
	currentJob.JobID = jobID
	err = jc.JobService.DeleteJob(c, jobID)
	if err != nil {
		log.Println("unable to delete Job:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Job deleted Successfully"})

}

// GetAllJobs - Retrieves all jobs
func (jc *JobController) GetAllJobs(c *gin.Context) {
	jobs, err := jc.JobService.GetAllJobs(c)
	if err != nil {
		log.Println("unable to get AllJobs:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"jobs": jobs})
	log.Println("jobs gotten successfully")
}

// GetJobRecommendations - Retrieves job recommendations for a user
func (jc *JobController) GetJobRecommendations(c *gin.Context) {
	userIDInt, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		log.Println("invalid user ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	id, preference, err := jc.UserService.GetUserByIdAndJobPreference(c,userIDInt)
	if err !=nil {
		log.Println("unable to get ID and prefrence:",err)
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
	}
	//convert back to string
	userID := strconv.Itoa(id)
	jobs, err := jc.JobService.GetJobRecommendations(c, userID, preference)
	if err != nil {
		log.Println("unable to get jobrecommedation:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"recommendations": jobs})
}

// GetJobByID - Retrieves a single job by its ID
func (jc *JobController) GetJobByID(c *gin.Context) {
	jobID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println("invalid job ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid job ID"})
		return
	}

	job, err := jc.JobService.GetJobByID(c, jobID)
	if err != nil {
		log.Println("unable to get jobID:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"job": job})
}

// populate job to db
func (jc *JobController) Seed(c *gin.Context) {
	jobs, err := jc.JobService.Seed(c)
	if err != nil {
		log.Println("cannot get Api data", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "cannot get Api data",
			"details": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{"jobs": jobs})
	log.Println("jobs obtained")
}
