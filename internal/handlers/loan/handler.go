package loan

import (
	"encoding/json"
	"log"
	"net/http"

	"loan-service/internal/entity"
	"loan-service/internal/infrastructure/middleware"
	"loan-service/internal/infrastructure/response"
	"loan-service/internal/services/loan"
)

type Handler struct {
	loanService loan.Service
}

func New(loanService loan.Service) Handler {
	return Handler{
		loanService: loanService,
	}
}
func (h *Handler) CreateLoan(w http.ResponseWriter, r *http.Request) {
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

	// Validation checks
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

	user, _ := middleware.GetUserFromContext(r.Context())

	createLoanRequest := loan.CreateLoanRequest{
		User: entity.User{
			UserID: user.UserID,
			Role:   user.Role,
		},
		PrincipalAmount:    request.PrincipalAmount,
		InterestRate:       request.InterestRate,
		ReturnOnInvestment: request.ReturnOnInvestment,
	}

	result, err := h.loanService.CreateLoan(createLoanRequest)
	if err != nil {
		log.Println(err)
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
