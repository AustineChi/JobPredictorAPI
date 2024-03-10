package services

import (
	"JobPredictorAPI/models"
	"context"
	"errors"

	"gorm.io/gorm"
)

type NotificationService struct {
    Db *gorm.DB
}

func NewNotificationService(db *gorm.DB) *NotificationService {
    return &NotificationService{Db: db}
}

// CreateNotification creates a new notification in the database
func (s *NotificationService) CreateNotification(ctx context.Context, notification *models.Notification) error {
    if err := s.Db.Create(notification).WithContext(ctx).Error; err !=nil{
        return err
    }
    return nil
}

// GetNotificationsByUserID retrieves all notifications for a specific user
func (s *NotificationService) GetNotificationsByUserID(ctx context.Context, userID int) ([]models.Notification, error) {
	var notifications []models.Notification

	// Use GORM's Find method to fetch notifications for the given user
	result := s.Db.Where("user_id = ?", userID).Find(&notifications)
	if result.Error != nil {
		return nil, result.Error
	}
    if result.RowsAffected == 0{
        return nil, errors.New("No rows affected")
    }
	return notifications, nil
}

// UpdateNotification updates a notification's details
func (s *NotificationService) UpdateNotification(ctx context.Context, notification *models.Notification) error {
	result := s.Db.Save(notification)
	if result.Error != nil {
		return result.Error
	}
    if result.RowsAffected == 0{
        return  errors.New("No rows affected")
    }
	return nil
}

// DeleteNotification removes a notification from the database
func (s *NotificationService) DeleteNotification(ctx context.Context, notificationID int) error {
	// Create a new Notification instance with the notificationID set
	notification := models.Notification{NotificationID: notificationID}

	result := s.Db.Delete(&notification)
	if result.Error != nil {
		return result.Error
	}
    if result.RowsAffected == 0{
        return  errors.New("No rows affected")
    }
	return nil
}
