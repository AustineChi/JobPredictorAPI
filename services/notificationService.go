package services

import (
    "context"
    "database/sql"
    "JobPredictorAPI/models"
)

type NotificationService struct {
    Db *sql.DB
}

func NewNotificationService(db *sql.DB) *NotificationService {
    return &NotificationService{Db: db}
}

// CreateNotification creates a new notification in the database
func (s *NotificationService) CreateNotification(ctx context.Context, notification *models.Notification) error {
    query := "INSERT INTO public.notification (user_id, message, date_sent) VALUES ($1, $2, $3)"
    _, err := s.Db.ExecContext(ctx, query, notification.UserID, notification.Message, notification.DateSent)
    return err
}

// GetNotificationsByUserID retrieves all notifications for a specific user
func (s *NotificationService) GetNotificationsByUserID(ctx context.Context, userID int) ([]models.Notification, error) {
    query := "SELECT notification_id, user_id, message, date_sent FROM public.notification WHERE user_id = $1"
    rows, err := s.Db.QueryContext(ctx, query, userID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var notifications []models.Notification
    for rows.Next() {
        var notification models.Notification
        if err := rows.Scan(&notification.NotificationID, &notification.UserID, &notification.Message, &notification.DateSent); err != nil {
            return nil, err
        }
        notifications = append(notifications, notification)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return notifications, nil
}

// UpdateNotification updates a notification's details
func (s *NotificationService) UpdateNotification(ctx context.Context, notification *models.Notification) error {
    query := "UPDATE public.notification SET message = $1, date_sent = $2 WHERE notification_id = $3"
    _, err := s.Db.ExecContext(ctx, query, notification.Message, notification.DateSent, notification.NotificationID)
    return err
}

// DeleteNotification removes a notification from the database
func (s *NotificationService) DeleteNotification(ctx context.Context, notificationID int) error {
    query := "DELETE FROM public.notification WHERE notification_id = $1"
    _, err := s.Db.ExecContext(ctx, query, notificationID)
    return err
}
