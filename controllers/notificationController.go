package controllers

import (
	"JobPredictorAPI/models"
	"JobPredictorAPI/services"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// NotificationController struct
type NotificationController struct {
    NotificationService *services.NotificationService
}

// NewNotificationController creates a new controller for notifications
func NewNotificationController(ns *services.NotificationService) *NotificationController {
    return &NotificationController{
        NotificationService: ns,
    }
}

// GetNotifications - Retrieves all notifications for a user
func (nc *NotificationController) GetNotifications(c *gin.Context) {
    userID, err := strconv.Atoi(c.Param("userID"))
    if err != nil {
        log.Println("invalid user ID:", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    notifications, err := nc.NotificationService.GetNotificationsByUserID(c, userID)
    if err != nil {
        log.Println("unable to get user notification", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, notifications)
}

// CreateNotification - Handles creating a new notification
func (nc *NotificationController) CreateNotification(c *gin.Context) {
    var newNotification models.Notification
    if err := c.BindJSON(&newNotification); err != nil {
        log.Println("cannot bindJSON", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

   notify,  err := nc.NotificationService.CreateNotification(c, &newNotification)
    if err != nil {
        log.Println("unable to create notification:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "message": "Notification created successfully",
        "details":notify,
    })
}

// UpdateNotification - Updates a notification's details
func (nc *NotificationController) UpdateNotification(c *gin.Context) {
    notificationID, err := strconv.Atoi(c.Param("notificationID"))
    if err != nil {
        log.Println("invalid notification:", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid notification ID"})
        return
    }

    var updatedNotification models.Notification
    if err := c.BindJSON(&updatedNotification); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    updatedNotification.NotificationID = notificationID

    err = nc.NotificationService.UpdateNotification(c, &updatedNotification)
    if err != nil {
        log.Println("unable to update notification", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Notification updated successfully"})
}

// DeleteNotification - Deletes a notification
func (nc *NotificationController) DeleteNotification(c *gin.Context) {
    notificationID, err := strconv.Atoi(c.Param("notificationID"))
    if err != nil {
        log.Println("invalid notification ID:", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid notification ID"})
        return
    }

    err = nc.NotificationService.DeleteNotification(c, notificationID)
    if err != nil {
        log.Println("unable to delete notification", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Notification deleted successfully"})
}
