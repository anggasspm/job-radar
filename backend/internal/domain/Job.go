package domain

import "time"

// still representing db 100%, need to be simplified
type Job struct {
	ID         int64   
	SourceID   int     
	ExternalID *string 

	Title       string  
	Company     string 
	Location    string  
	Category    *string 
	Description *string 

	SalaryMin int64  
	SalaryMax int64  
	Currency  string 

	MinExp    int16  
	MaxExp    int16   
	Education *string 

	RawURL string 

	PostedDate *time.Time 
	ScrapedAt  time.Time  

	IsActive bool 

	CreatedAt time.Time 
	UpdatedAt time.Time 
}
