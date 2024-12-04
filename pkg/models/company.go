package models

import (
	"errors"
	"fmt"
)

// Company defines the company model
type Company struct {
	ID           string      `json:"id"`
	Name         string      `json:"name"`
	Description  string      `json:"description,omitempty"`
	NumEmployees int         `json:"num_employees"`
	IsRegistered bool        `json:"is_registered"`
	Type         CompanyType `json:"type"`
}

// Validate runs validation on the Company struct
func (c *Company) Validate() error {
	if c.Name == "" {
		return errors.New("name is missing")
	}
	if len(c.Name) > 15 {
		return errors.New("name cannot exceed 15 characters")
	}

	if c.NumEmployees <= 0 {
		return errors.New("num_employees must be greater than zero")
	}

	validTypes := []CompanyType{
		CompanyTypeCorporations,
		CompanyTypeNonProfit,
		CompanyTypeCooperative,
		CompanyTypeSoleProprietorship,
	}
	isValidType := false
	for _, t := range validTypes {
		if c.Type == t {
			isValidType = true
			break
		}
	}
	if !isValidType {
		return fmt.Errorf("invalid company type: %s", c.Type)
	}

	return nil
}

// CompanyType defines the company type
type CompanyType string

const (
	CompanyTypeCorporations       CompanyType = "Corporations"
	CompanyTypeNonProfit          CompanyType = "NonProfit"
	CompanyTypeCooperative        CompanyType = "Cooperative"
	CompanyTypeSoleProprietorship CompanyType = "Sole Proprietorship"
)
