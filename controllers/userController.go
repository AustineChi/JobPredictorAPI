package controllers

import (
	"JobPredictorAPI/middleware"
	"JobPredictorAPI/models"
	"JobPredictorAPI/services"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

// Register - Handles new user registration
func (uc *UserController) Register(c *gin.Context) {
	var newUser models.User
	if err := c.BindJSON(&newUser); err != nil {
		log.Println("unable to bindJSON", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//UserService.CreateUser handles the creation logic
	user, err := uc.UserService.CreateUser(c, &newUser)
	if err != nil {
		log.Println("unable to create user", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"details": user,
	})
	log.Println("User registered successfully")
}

// Login - Handles user login
func (uc *UserController) Login(c *gin.Context) {
	// Implement login logic here
	// This will depend on your authentication strategy (e.g., JWT, OAuth)
	var UserPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	err := c.ShouldBindJSON(&UserPayload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
		log.Println("invalid request payload", err)
	}
	user, err := uc.UserService.GetEmail(c, UserPayload.Email)
	if err != nil {
		log.Println("Invalid Email")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "invalid credentials",
			"details": err.Error(),
		})
		return
	}
	valid, err := user.CheckPassword(UserPayload.Password)
	if err != nil || !valid {
		log.Println("invalid password", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid credentials"})
		return
	}
	//JWT Token generation
	token, err := middleware.GenerateToken(user.UserID, user.Email, true)
	if err != nil {
		log.Println("unable to generate token:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Println("token generated")
	_, err = json.Marshal(token)
	if err != nil {
		log.Println("unable to unmarshall data", err)
		return
	}
	log.Println("Token generated:", token)
	c.JSON(http.StatusOK, gin.H{"token": token})

	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
	log.Println("login successful")
}

// GetUser - Retrieves a user's profile
func (uc *UserController) GetUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := uc.UserService.GetUserByID(c, userID)
	if err != nil {
		log.Println("unable to get user ID:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

// UpdateUser - Updates a user's information
func (uc *UserController) UpdateUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println("invalid ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var updatedUser models.User
	if err := c.BindJSON(&updatedUser); err != nil {
		log.Println("unable to bind JSON", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedUser.UserID = userID // Ensure the ID is set to the one from the path

	err = uc.UserService.UpdateUser(c, &updatedUser)
	if err != nil {
		log.Println("unable to update user", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}
