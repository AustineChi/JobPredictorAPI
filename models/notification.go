package models

import (
    "time"
)

// Notification represents a notification sent to a user
type Notification struct {
    NotificationID int       `json:"notificationId" gorm:"column:notification_id;primary_key"`
    UserID         int       `json:"userId" gorm:"column:user_id"`
    Message        string    `json:"message"`
    DateSent       time.Time `json:"dateSent" gorm:"column:date_sent"`
}
