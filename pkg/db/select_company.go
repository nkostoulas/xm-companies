package db

import (
	"xm-companies/pkg/models"
)

const selectCompanySQL = "SELECT id, name, description, num_employees, is_registered, type FROM companies WHERE id = $1"

// SelectCompany fetches a company based on the id provided
func (r *DB) SelectCompany(id string) (models.Company, error) {
	var company models.Company
	row := r.db.QueryRow(selectCompanySQL, id)
	err := row.Scan(
		&company.ID,
		&company.Name,
		&company.Description,
		&company.NumEmployees,
		&company.IsRegistered,
		&company.Type,
	)
	if err != nil {
		return models.Company{}, err
	}
	return company, nil
}
