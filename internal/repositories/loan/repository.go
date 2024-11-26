package loan

import (
	"database/sql"
	"log"

	"loan-service/internal/entity"
)

const (
	queryInsertLoan = `
		INSERT INTO loans (
			loan_id, borrower_id, status, created_at, updated_at
		) VALUES ($1, $2, $3, $4) RETURNING id
	`
)

type Repository interface {
	CreateLoan(loan *entity.Loan) error
}

type repository struct {
	database *sql.DB
}

func New(database *sql.DB) Repository {
	return &repository{
		database: database,
	}
}

func (r *repository) CreateLoan(loan *entity.Loan) error {
	err := r.database.QueryRow(queryInsertLoan, loan.BorrowerID, loan.Status, loan.CreatedAt, loan.UpdatedAt).Scan(&loan.ID)
	if err != nil {
		log.Println("Error inserting loan:", err)
		return err
	}
	return nil
}
