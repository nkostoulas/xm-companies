package db

import "fmt"

const deleteCompanySQL = "DELETE FROM companies WHERE id = $1"

// DeleteCompany deletes a company based on the id provided
func (r *DB) DeleteCompany(id string) error {
	result, err := r.db.Exec(deleteCompanySQL, id)
	if err != nil {
		return fmt.Errorf("error deleting company: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error checking rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("company with id '%s' does not exist", id)
	}

	return nil
}
