package controllers

import (
    "net/http"
    "strconv"
    "JobPredictorAPI/models"
    "JobPredictorAPI/services"
    "github.com/gin-gonic/gin"
)

// InteractionController struct
type InteractionController struct {
    InteractionService *services.InteractionService
}

// NewInteractionController creates a new controller for interactions
func NewInteractionController(is *services.InteractionService) *InteractionController {
    return &InteractionController{
        InteractionService: is,
    }
}

// LogInteraction - Handles logging a new interaction
func (ic *InteractionController) LogInteraction(c *gin.Context) {
    var newInteraction models.Interaction
    if err := c.BindJSON(&newInteraction); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err := ic.InteractionService.CreateInteraction(c, &newInteraction)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "Interaction logged successfully"})
}

// GetInteractions - Retrieves all interactions for a user
func (ic *InteractionController) GetInteractions(c *gin.Context) {
    userID, err := strconv.Atoi(c.Param("userID"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    interactions, err := ic.InteractionService.GetInteractionsByUserID(c, userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, interactions)
}

// UpdateInteraction - Updates an interaction's details (if applicable)
func (ic *InteractionController) UpdateInteraction(c *gin.Context) {
    interactionID, err := strconv.Atoi(c.Param("interactionID"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid interaction ID"})
        return
    }

    var updatedInteraction models.Interaction
    if err := c.BindJSON(&updatedInteraction); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    updatedInteraction.InteractionID = interactionID

    err = ic.InteractionService.UpdateInteraction(c, &updatedInteraction)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Interaction updated successfully"})
}

// DeleteInteraction - Deletes an interaction
func (ic *InteractionController) DeleteInteraction(c *gin.Context) {
    interactionID, err := strconv.Atoi(c.Param("interactionID"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid interaction ID"})
        return
    }

    err = ic.InteractionService.DeleteInteraction(c, interactionID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Interaction deleted successfully"})
}
