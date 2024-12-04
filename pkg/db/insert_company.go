package db

import (
	"xm-companies/pkg/models"
)

const createCompanySQL = `INSERT INTO companies (id, name, description, num_employees, is_registered, type) VALUES ($1, $2, $3, $4, $5, $6)`

// Insert company inserts a new company to the database
func (d *DB) InsertCompany(company models.Company) error {
	_, err := d.db.Exec(createCompanySQL,
		company.ID,
		company.Name,
		company.Description,
		company.NumEmployees,
		company.IsRegistered,
		company.Type,
	)
	return err
}
