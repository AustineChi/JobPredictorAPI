package controllers

import (
	"JobPredictorAPI/models"
	"JobPredictorAPI/services"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SavedJobsController struct {
	SavedJobsService *services.SavedJobsService
}

func NewSavedJobsController(sjs *services.SavedJobsService) *SavedJobsController {
	return &SavedJobsController{
		SavedJobsService: sjs,
	}
}

// SaveJob - Handles saving a job for a user
func (sjc *SavedJobsController) SaveJob(c *gin.Context) {
	var savedJob models.SavedJob
	if err := c.BindJSON(&savedJob); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := sjc.SavedJobsService.SaveJob(c, savedJob.UserID, savedJob.JobID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Job saved successfully"})
}

// GetSavedJobs - Retrieves saved jobs for a user
func (sjc *SavedJobsController) GetSavedJobs(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		log.Println("invalid user ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	savedJobs, err := sjc.SavedJobsService.GetSavedJobs(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, savedJobs)
}

// DeleteSavedJob - Removes a saved job for a user
func (sjc *SavedJobsController) DeleteSavedJob(c *gin.Context) {
	savedJobID, err := strconv.Atoi(c.Param("savedJobID"))
	if err != nil {
		log.Println("invalid user ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid saved job ID"})
		return
	}
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		log.Println("invalid user ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID for saved Job instance"})
	}

	err = sjc.SavedJobsService.DeleteSavedJob(c, userID, savedJobID)
	if err != nil {
		log.Println("unable to delete saved job:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Job deleted successfully"})
}
