package loan

import (
	"encoding/json"
	"loan-service/internal/infrastructure/middleware"
	"loan-service/internal/infrastructure/response"
	"loan-service/internal/services/loan"
	"log"
	"net/http"
)

func (h *Handler) CreateLoan(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var (
		request CreateLoanRequest
		err     error
	)

	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.BuildResponse(w, http.StatusBadRequest, response.Response{
			Error: "Invalid request body",
		})
		return
	}

	if request.PrincipalAmount <= 0 {
		response.BuildResponse(w, http.StatusBadRequest, response.Response{
			Error: "Principal amount must be greater than 0",
		})
		return
	}

	if request.InterestRate <= 0 {
		response.BuildResponse(w, http.StatusBadRequest, response.Response{
			Error: "Interest rate must be greater than 0",
		})
		return
	}

	if request.ReturnOnInvestment <= 0 {
		response.BuildResponse(w, http.StatusBadRequest, response.Response{
			Error: "Return on investment must be greater than 0",
		})
		return
	}

	userID, _ := ctx.Value(middleware.UserIDContextKey).(int64)
	userRole, _ := ctx.Value(middleware.UserRoleContextKey).(int)

	createLoanRequest := loan.CreateLoanRequest{
		UserID:             userID,
		UserRole:           userRole,
		PrincipalAmount:    request.PrincipalAmount,
		InterestRate:       request.InterestRate,
		ReturnOnInvestment: request.ReturnOnInvestment,
	}

	result, err := h.loanService.CreateLoan(ctx, createLoanRequest)
	if err != nil {
		log.Println("Error creating loan:", err)
		response.BuildResponse(w, http.StatusInternalServerError, response.Response{
			Error: err.Error(),
		})
		return
	}

	response.BuildResponse(w, http.StatusOK, response.Response{
		Message: "Loan created successfully",
		Result: CreateLoanResponse{
			LoanID: result.LoanID,
		},
	})
}

func (h *Handler) GetLoans(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) GetLoanDetails(w http.ResponseWriter, r *http.Request) {

}
