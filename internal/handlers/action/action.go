package action

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"loan-service/internal/infrastructure/constant"
	"loan-service/internal/infrastructure/middleware"
	"loan-service/internal/infrastructure/response"
	"loan-service/internal/infrastructure/validator"
	"loan-service/internal/services/action"
)

func (h *Handler) ApproveLoan(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var (
		request ApproveLoanRequest
		err     error
	)

	vars := mux.Vars(r)
	loanIDStr := vars["loanID"]
	loanID, err := strconv.ParseInt(loanIDStr, 10, 64)
	if err != nil {
		response.BuildResponse(w, http.StatusBadRequest, response.Response{
			Error: "Invalid loan ID",
		})
		return
	}

	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.BuildResponse(w, http.StatusBadRequest, response.Response{
			Error: "Invalid request body",
		})
		return
	}

	if loanID <= 0 {
		response.BuildResponse(w, http.StatusBadRequest, response.Response{
			Error: "Loan ID must be greater than 0",
		})
		return
	}

	if request.DocumentURL == "" {
		response.BuildResponse(w, http.StatusBadRequest, response.Response{
			Error: "Document URL is required",
		})
		return
	}

	// assumption: the picture proof of validator has visited the borrower is stored in other service
	if !validator.IsValidURL(request.DocumentURL) {
		response.BuildResponse(w, http.StatusBadRequest, response.Response{
			Error: "Document URL is not valid",
		})
		return
	}

	userID, _ := ctx.Value(middleware.UserIDContextKey).(int64)
	userRole, _ := ctx.Value(middleware.UserRoleContextKey).(int)

	updateLoanRequest := action.UpdateLoanRequest{
		UserID:      userID,
		UserRole:    userRole,
		LoanID:      loanID,
		DocumentURL: request.DocumentURL,
		ActionType:  constant.ActionTypeApproveLoan,
	}
	result, err := h.actionService.UpdateLoan(ctx, updateLoanRequest)
	if err != nil {
		response.BuildResponse(w, http.StatusInternalServerError, response.Response{
			Error: err.Error(),
		})
		return
	}

	response.BuildResponse(w, http.StatusOK, response.Response{
		Message: "Loan approved successfully",
		Result: ApproveLoanResponse{
			LoanID: result.LoanID,
		},
	})
}

func (h *Handler) DisburseLoan(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var (
		request DisburseLoanRequest
		err     error
	)

	vars := mux.Vars(r)
	loanIDStr := vars["loanID"]
	loanID, err := strconv.ParseInt(loanIDStr, 10, 64)
	if err != nil {
		response.BuildResponse(w, http.StatusBadRequest, response.Response{
			Error: "Invalid loan ID",
		})
		return
	}

	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.BuildResponse(w, http.StatusBadRequest, response.Response{
			Error: "Invalid request body",
		})
		return
	}

	if loanID <= 0 {
		response.BuildResponse(w, http.StatusBadRequest, response.Response{
			Error: "Loan ID must be greater than 0",
		})
		return
	}

	if request.DocumentURL == "" {
		response.BuildResponse(w, http.StatusBadRequest, response.Response{
			Error: "Document URL is required",
		})
		return
	}

	// assumption: the picture proof of validator has visited the borrower is stored in other service
	if !validator.IsValidURL(request.DocumentURL) {
		response.BuildResponse(w, http.StatusBadRequest, response.Response{
			Error: "Document URL is not valid",
		})
		return
	}

	userID, _ := ctx.Value(middleware.UserIDContextKey).(int64)
	userRole, _ := ctx.Value(middleware.UserRoleContextKey).(int)

	updateLoanRequest := action.UpdateLoanRequest{
		UserID:      userID,
		UserRole:    userRole,
		LoanID:      loanID,
		DocumentURL: request.DocumentURL,
		ActionType:  constant.ActionTypeDisburse,
	}
	result, err := h.actionService.UpdateLoan(ctx, updateLoanRequest)
	if err != nil {
		response.BuildResponse(w, http.StatusInternalServerError, response.Response{
			Error: err.Error(),
		})
		return
	}

	response.BuildResponse(w, http.StatusOK, response.Response{
		Message: "Loan disbursed successfully",
		Result: DisburseLoanResponse{
			LoanID: result.LoanID,
		},
	})
}

func (h *Handler) InvestLoan(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var (
		request InvestLoanRequest
		err     error
	)

	vars := mux.Vars(r)
	loanIDStr := vars["loanID"]
	loanID, err := strconv.ParseInt(loanIDStr, 10, 64)
	if err != nil {
		response.BuildResponse(w, http.StatusBadRequest, response.Response{
			Error: "Invalid loan ID",
		})
		return
	}

	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.BuildResponse(w, http.StatusBadRequest, response.Response{
			Error: "Invalid request body",
		})
		return
	}

	if loanID <= 0 {
		response.BuildResponse(w, http.StatusBadRequest, response.Response{
			Error: "Loan ID must be greater than 0",
		})
		return
	}

	if request.InvestmentAmount <= 0 {
		response.BuildResponse(w, http.StatusBadRequest, response.Response{
			Error: "Investment amount must be greater than 0",
		})
		return
	}

	userID, _ := ctx.Value(middleware.UserIDContextKey).(int64)
	userRole, _ := ctx.Value(middleware.UserRoleContextKey).(int)

	investLoanRequest := action.InvestLoanRequest{
		UserID:           userID,
		UserRole:         userRole,
		LoanID:           loanID,
		InvestmentAmount: request.InvestmentAmount,
	}
	result, err := h.actionService.InvestLoan(ctx, investLoanRequest)
	if err != nil {
		response.BuildResponse(w, http.StatusInternalServerError, response.Response{
			Error: err.Error(),
		})
		return
	}

	response.BuildResponse(w, http.StatusOK, response.Response{
		Message: "Investment successful",
		Result: InvestLoanResponse{
			InvestmentID: result.InvestmentID,
			LoanID:       result.LoanID,
		},
	})
}
