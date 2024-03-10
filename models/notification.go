package models

import (
    "time"
)

// Notification represents a notification sent to a user
type Notification struct {
    NotificationID int       `gorm:"column:notification_id;primary_key" json:"notificationId"`
    UserID         int       `gorm:"column:user_id" json:"userId"`
    Message        string    `gorm:"column:message" json:"message"`
    DateSent       time.Time `gorm:"column:date_sent" json:"dateSent"`
}
