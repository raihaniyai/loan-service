package action

import (
	"encoding/json"
	"loan-service/internal/entity"
	"loan-service/internal/infrastructure/middleware"
	"loan-service/internal/infrastructure/response"
	"loan-service/internal/infrastructure/validator"
	"loan-service/internal/services/action"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

	if !validator.IsValidURL(request.DocumentURL) {
		response.BuildResponse(w, http.StatusBadRequest, response.Response{
			Error: "Document URL is not valid",
		})
		return
	}

	user, _ := middleware.GetUserFromContext(ctx)
	approveLoanRequest := action.ApproveLoanRequest{
		User: entity.User{
			UserID: user.UserID,
			Role:   user.Role,
		},
		LoanID:      loanID,
		DocumentURL: request.DocumentURL,
	}
	result, err := h.actionService.ApproveLoan(ctx, approveLoanRequest)
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

func (h *Handler) InvestLoan(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) DisburseLoan(w http.ResponseWriter, r *http.Request) {

}
