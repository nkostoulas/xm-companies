package db

import (
	"fmt"
	"xm-companies/pkg/models"
)

const (
	updateCompanySQL            = "UPDATE companies SET "
	updateCompanyWhereClauseSQL = " WHERE id = $"
	companyExistsSQL            = "SELECT EXISTS (SELECT 1 FROM companies WHERE id = $1)"
)

// UpdateCompany company updates a company based on the id and updates provided
func (r *DB) UpdateCompany(id string, updates models.UpdateCompany) error {
	// Check if the company exists
	var exists bool
	err := r.db.QueryRow(companyExistsSQL, id).Scan(&exists)
	if err != nil {
		return fmt.Errorf("error checking company existence: %w", err)
	}
	if !exists {
		return fmt.Errorf("company with id '%s' does not exist", id)
	}

	// Build dynamic query based on non-nil fields
	query := updateCompanySQL
	args := []interface{}{}
	i := 1
	if updates.Name != nil {
		query += fmt.Sprintf("name = $%d, ", i)
		args = append(args, *updates.Name)
		i++
	}
	if updates.Description != nil {
		query += fmt.Sprintf("description = $%d, ", i)
		args = append(args, *updates.Description)
		i++
	}
	if updates.NumEmployees != nil {
		query += fmt.Sprintf("num_employees = $%d, ", i)
		args = append(args, *updates.NumEmployees)
		i++
	}
	if updates.IsRegistered != nil {
		query += fmt.Sprintf("is_registered = $%d, ", i)
		args = append(args, *updates.IsRegistered)
		i++
	}
	if updates.Type != nil {
		query += fmt.Sprintf("type = $%d, ", i)
		args = append(args, *updates.Type)
		i++
	}

	// Remove trailing comma and space, add WHERE clause
	query = query[:len(query)-2] + updateCompanyWhereClauseSQL + fmt.Sprintf("%d", i)
	args = append(args, id)

	_, err = r.db.Exec(query, args...)
	return err
}
