package models

import "time"

type Organization_type string

const (
	IE  Organization_type = "IE"
	LLC Organization_type = "LLC"
	JSC Organization_type = "JSC"
)

type Organization struct {
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Type        Organization_type `json:"type"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
}

type Tender struct {
	ID              string    `json:"id"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	ServiceType     string    `json:"serviceType"`
	Status          string    `json:"status"`
	OrganizationID  string    `json:"organizationId"`
	CreatorID       string    `json:"creatorId"`
	Version         int       `json:"version"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
	CreatorUsername string    `json:"creatorUsername"`
}
