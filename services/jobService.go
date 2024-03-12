package services

import (
	"JobPredictorAPI/models"
	"JobPredictorAPI/utils" // Ensure this path correctly points to your utils package
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"gorm.io/gorm"
)

type JobService struct {
	Db *gorm.DB
}

func NewJobService(db *gorm.DB) *JobService {
	return &JobService{Db: db}
}

// JobAPIResponse represents a job from the external API
type JobAPIResponse struct {
	CompanyName string `json:"companyName"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Link        string `json:"link"`
	Location    string `json:"location"`
	// Add other fields as needed
}

// APIResponse is the top-level structure of the API response
type APIResponse struct {
	Jobs []JobAPIResponse `json:"jobs"`
}

// FetchAndStoreJobs gets a limited number of jobs from the external API and stores them in the database
func (s *JobService) FetchAndStoreJobs(ctx context.Context) error {
	url := "https://zobjobs.com/api/jobs"

	respBody, err := utils.HttpGet(url)
	if err != nil {
		return err
	}

	var apiResponse APIResponse
	if err := json.Unmarshal(respBody, &apiResponse); err != nil {
		return err
	}

	// Process only 3 jobs each day
	for i, jobAPIResponse := range apiResponse.Jobs {
		if i >= 3 {
			break
		}
		job := models.Job{
			Title:       jobAPIResponse.Title,
			Description: jobAPIResponse.Description,
			Company:     jobAPIResponse.CompanyName,
			Location:    jobAPIResponse.Location,
		}

		// Insert the job into the database
		if err := s.Db.Raw("INSERT INTO public.job (title, description, company, location) VALUES ($1, $2, $3, $4)", job.Title, job.Description, job.Company, job.Location).WithContext(ctx).Error; err != nil {
			return err
			// Log or handle the error as needed
		}
	}

	return nil
}

// Placeholder for additional methods

// GetJobByID fetches a single job by its ID
func (s *JobService) GetJobByID(ctx context.Context, jobID int) (*models.Job, error) {
	// Create a Job instance to hold the data
	var job models.Job

	// Execute the query to find the job by ID
	result := s.Db.Model(&models.Job{}).Select("job_id, title, description, company, location").Where("job_id = $1", jobID).WithContext(ctx).First(&job)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// If no rows were returned, return a nil job and no error
			return nil, nil
		}
		// Return any other error that occurred during the query execution
		return nil, result.Error
	}

	// Return the job and no error
	return &job, nil
}

// UpdateJob updates a given job's details
func (s *JobService) UpdateJob(ctx context.Context, job *models.Job) error {
	// SQL query to update a job
	query := "UPDATE public.job SET title = $1, description = $2, company = $3, location = $4 WHERE job_id = $5"

	result := s.Db.Exec(query, job.Title, job.Description, job.Company, job.Location, job.JobID).WithContext(ctx)
	if result.Error != nil {
		return result.Error
	}
	// Check affected rows if needed
	if result.RowsAffected == 0 {
		return errors.New("no rows affected")
	}
	return nil
}

func (s *JobService) CreateJob(ctx context.Context, job *models.Job) (models.Job, error) {
	query := "INSERT INTO public.job (title, description, company, location, skills_required, salary, employment_type) VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING id, title, description, company, location, skills_required, salary, employment_type"

	var createdJob models.Job
	result := s.Db.Raw(query, job.Title, job.Description, job.Company, job.Location, job.SkillsRequired, job.Salary, job.EmploymentType).WithContext(ctx).Scan(&createdJob).WithContext(ctx)
	if result.Error != nil {
		return models.Job{}, result.Error
	}
	if result.RowsAffected == 0 {
		return models.Job{}, errors.New("no rows affected")
	}
	return createdJob, nil
}

// DeleteJob removes a job from the database
func (s *JobService) DeleteJob(ctx context.Context, jobID int) error {
	// Create a new Job instance with the jobID set
	job := models.Job{JobID: jobID}

	// Use GORM's Delete method to delete the job
	result := s.Db.Delete(&job).WithContext(ctx)
	if result.Error != nil {
		// Return any error that occurred during the delete operation
		return result.Error
	}

	// Check if no rows were affected
	if result.RowsAffected == 0 {
		return errors.New("no rows affected")
	}

	// Return no error if delete was successful
	return nil
}

// GetAllJobs retrieves all jobs from the database.
func (s *JobService) GetAllJobs(ctx context.Context) ([]models.Job, error) {
	// Create a slice to hold the jobs
	var jobs []models.Job

	result := s.Db.Find(&jobs).WithContext(ctx)
	if result.Error != nil {
		// Return any error that occurred during the query execution
		return nil, result.Error
	}
	// Check if no rows were affected
	if result.RowsAffected == 0 {
		return nil, errors.New("no rows affected")
	}
	return jobs, nil
}

// GetJobRecommendations retrieves job recommendations from the Flask app.
func (s *JobService) GetJobRecommendations(ctx context.Context, jwtToken string) ([]models.Job, error) {
	url := "http://127.0.0.1:5000/recommendations/recommend"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+jwtToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response struct {
		Recommendations []models.Job `json:"recommendations"`
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return response.Recommendations, nil
}
