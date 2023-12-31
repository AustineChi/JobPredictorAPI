package router

import (
    "JobPredictorAPI/controllers"
    "github.com/gin-gonic/gin"
)

func SetupRouter(
    userController *controllers.UserController,
    jobController *controllers.JobController,
    savedJobsController *controllers.SavedJobsController,
    recommendationController *controllers.RecommendationController,
    notificationController *controllers.NotificationController,
    interactionController *controllers.InteractionController,
) *gin.Engine {
    
    router := gin.Default()

    // User routes
    router.POST("/register", userController.Register)
    router.POST("/login", userController.Login)
    router.GET("/user/:id", userController.GetUser)
    router.PUT("/user/:id", userController.UpdateUser)

    // Job routes
    router.GET("/jobs", jobController.GetAllJobs)
    router.GET("/jobs/:id", jobController.GetJob)
    router.POST("/jobs", jobController.CreateJob)
    router.PUT("/jobs/:id", jobController.UpdateJob)
    router.DELETE("/jobs/:id", jobController.DeleteJob)

    // Saved Jobs routes
    router.POST("/save-job", savedJobsController.SaveJob)
    router.GET("/saved-jobs/:userID", savedJobsController.GetSavedJobs)
    router.DELETE("/saved-job/:savedJobID", savedJobsController.DeleteSavedJob)

    // Recommendation routes (if any specific routes are needed)
    router.GET("/recommendations/:userID", jobController.GetJobRecommendations)

    // Notification routes
    router.GET("/notifications/:userID", notificationController.GetNotifications)
    router.POST("/notifications", notificationController.CreateNotification)
    router.PUT("/notifications/:notificationID", notificationController.UpdateNotification)
    router.DELETE("/notifications/:notificationID", notificationController.DeleteNotification)

    // Interaction routes
    router.POST("/interactions", interactionController.LogInteraction)
    router.GET("/interactions/:userID", interactionController.GetInteractions)
    router.PUT("/interactions/:interactionID", interactionController.UpdateInteraction)
    router.DELETE("/interactions/:interactionID", interactionController.DeleteInteraction)


    return router
}
