package handler

import (
	"fmt"
	"strings"
	"time"
)

// data transfer object
type CreateListingRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	City        string `json:"city"`
}

type CreateListingResponse struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       int64     `json:"price"`
	City        string    `json:"city"`
	CreatedAt   time.Time `json:"created_at"`
}

type ValidationError struct {
	Field string
	Msg string
}

func (v *ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", v.Field, v.Msg)
}

func (req CreateListingRequest) Validate() error {
	if strings.TrimSpace(req.Title) == "" {
		return &ValidationError{Field: "title", Msg: "must not be empty"}
	}

	if strings.TrimSpace(req.Description) == "" {
		return &ValidationError{Field: "Description", Msg: "must not be empty"}
	}
	
	if req.Price < 0 {
		return &ValidationError{Field: "Price", Msg: "must be greater than zero"}
	}
	
	if strings.TrimSpace(req.City) == "" {
		return &ValidationError{Field: "City", Msg: "must not be empty"}
	}

	return nil
}