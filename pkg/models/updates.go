package models

import (
	"errors"
	"fmt"
)

// UpdateCompany defines updates to the company model
type UpdateCompany struct {
	Name         *string      `json:"name,omitempty"`
	Description  *string      `json:"description,omitempty"`
	NumEmployees *int         `json:"num_employees,omitempty"`
	IsRegistered *bool        `json:"is_registered,omitempty"`
	Type         *CompanyType `json:"type,omitempty"`
}

// Validate runs validation on the company updates
func (u *UpdateCompany) Validate() error {
	if u.Name != nil {
		if len(*u.Name) == 0 {
			return errors.New("name cannot be empty")
		}
		if len(*u.Name) > 15 {
			return errors.New("name cannot exceed 15 characters")
		}
	}

	if u.NumEmployees != nil && *u.NumEmployees <= 0 {
		return errors.New("num_employees must be greater than zero")
	}

	if u.Type != nil {
		validTypes := []CompanyType{
			CompanyTypeCorporations,
			CompanyTypeNonProfit,
			CompanyTypeCooperative,
			CompanyTypeSoleProprietorship,
		}
		isValidType := false
		for _, t := range validTypes {
			if *u.Type == t {
				isValidType = true
				break
			}
		}
		if !isValidType {
			return fmt.Errorf("invalid company type: %s", *u.Type)
		}
	}

	return nil
}
