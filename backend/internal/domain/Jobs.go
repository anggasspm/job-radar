package domain

import "time"

type Job struct {
	ID         int64   `json:"id" db:"id"`
	SourceID   int     `json:"source_id" db:"source_id"`
	ExternalID *string `json:"external_id,omitempty" db:"external_id"`

	Title       string  `json:"title" db:"title"`
	Company     string  `json:"company" db:"company"`
	Location    string  `json:"location" db:"location"`
	Category    *string `json:"category,omitempty" db:"category"`
	Description *string `json:"description,omitempty" db:"description"`

	SalaryMin int64  `json:"salary_min" db:"salary_min"`
	SalaryMax int64  `json:"salary_max" db:"salary_max"`
	Currency  string `json:"currency" db:"currency"`

	MinExp    int16   `json:"min_exp" db:"min_exp"`
	MaxExp    int16   `json:"max_exp" db:"max_exp"`
	Education *string `json:"education,omitempty" db:"education"`

	RawURL string `json:"raw_url" db:"raw_url"`

	PostedDate *time.Time `json:"posted_date,omitempty" db:"posted_date"`
	ScrapedAt  time.Time  `json:"scraped_at" db:"scraped_at"`

	IsActive bool `json:"is_active" db:"is_active"`

	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
