package services

import (
	"JobPredictorAPI/models"
	"JobPredictorAPI/utils" // Ensure this path correctly points to your utils package
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
)

type JobService struct {
	Db *sql.DB
}

func NewJobService(db *sql.DB) *JobService {
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
		_, err := s.Db.ExecContext(ctx, "INSERT INTO public.job (title, description, company, location) VALUES ($1, $2, $3, $4)", job.Title, job.Description, job.Company, job.Location)
		if err != nil {
			return err
			// Log or handle the error as needed
		}
	}

	return nil
}

// Placeholder for additional methods

// GetJobByID fetches a single job by its ID
func (s *JobService) GetJobByID(ctx context.Context, jobID int) (*models.Job, error) {
	// SQL query to fetch a job by ID
	query := "SELECT job_id, title, description, company, location FROM public.job WHERE job_id = $1"

	// Execute the query
	row := s.Db.QueryRowContext(ctx, query, jobID)

	// Create a Job instance to hold the data
	var job models.Job

	// Scan the row into the Job struct
	err := row.Scan(&job.JobID, &job.Title, &job.Description, &job.Company, &job.Location)
	if err != nil {
		// If no rows were returned, return a nil job and the error
		if err == sql.ErrNoRows {
			return nil, nil
		}
		// Return any other error that occurred during the query execution
		return nil, err
	}

	// Return the job and no error
	return &job, nil
}

// UpdateJob updates a given job's details
func (s *JobService) UpdateJob(ctx context.Context, job *models.Job) error {
	// SQL query to update a job
	query := "UPDATE public.job SET title = $1, description = $2, company = $3, location = $4 WHERE job_id = $5"

	// Execute the update query
	_, err := s.Db.ExecContext(ctx, query, job.Title, job.Description, job.Company, job.Location, job.JobID)
	if err != nil {
		// Return any error that occurred during the query execution
		return err
	}

	// Return no error if update was successful
	return nil
}

func (s *JobService) CreateJob(ctx context.Context, job *models.Job) (models.Job, error) {
    query := "INSERT INTO public.job (title, description, company, location, skills_required, salary, employment_type) VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING id, title, description, company, location, skills_required, salary, employment_type"
    
    var createdJob models.Job
    err := s.Db.QueryRowContext(ctx, query, job.Title, job.Description, job.Company, job.Location, job.SkillsRequired, job.Salary, job.EmploymentType).
        Scan(&createdJob.JobID, &createdJob.Title, &createdJob.Description, &createdJob.Company, &createdJob.Location, &createdJob.SkillsRequired, &createdJob.Salary, &createdJob.EmploymentType)
    
    if err != nil {
        return models.Job{}, err
    }

    return createdJob, nil
}


// DeleteJob removes a job from the database
func (s *JobService) DeleteJob(ctx context.Context, jobID int) error {
	// SQL query to delete a job
	query := "DELETE FROM public.job WHERE job_id = $1"

	// Execute the delete query
	_, err := s.Db.ExecContext(ctx, query, jobID)
	if err != nil {
		// Return any error that occurred during the query execution
		return err
	}

	// Return no error if delete was successful
	return nil
}

// GetAllJobs retrieves all jobs from the database.
func (s *JobService) GetAllJobs(ctx context.Context) ([]models.Job, error) {
	query := "SELECT job_id, title, description, company, location FROM public.job"

	rows, err := s.Db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var jobs []models.Job
	for rows.Next() {
		var job models.Job
		if err := rows.Scan(&job.JobID, &job.Title, &job.Description, &job.Company, &job.Location); err != nil {
			return nil, err
		}
		jobs = append(jobs, job)
	}

	if err := rows.Err(); err != nil {
		return nil, err
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
