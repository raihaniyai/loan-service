package user

import (
	"encoding/json"
	"log"
	"net/http"

	"loan-service/internal/infrastructure/constant"
	"loan-service/internal/infrastructure/middleware"
	"loan-service/internal/infrastructure/response"
	"loan-service/internal/infrastructure/validator"
	"loan-service/internal/services/user"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var (
		request CreateUserRequest
		err     error
	)

	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.BuildResponse(w, http.StatusBadRequest, response.Response{
			Error: "Invalid request body",
		})
		return
	}

	if constant.UserRoleText[request.Role] == "" || request.Role == constant.UserRoleSystem {
		response.BuildResponse(w, http.StatusBadRequest, response.Response{
			Error: "Invalid role",
		})
		return
	}

	if request.Email == "" {
		response.BuildResponse(w, http.StatusBadRequest, response.Response{
			Error: "Email is required",
		})
		return
	}

	if !validator.IsValidEmail(request.Email) {
		response.BuildResponse(w, http.StatusBadRequest, response.Response{
			Error: "Invalid email format",
		})
		return
	}

	if request.Name == "" {
		response.BuildResponse(w, http.StatusBadRequest, response.Response{
			Error: "Name is required",
		})
		return
	}

	result, err := h.userService.CreateUser(ctx, user.CreateUserRequest{
		Name:        request.Name,
		Role:        request.Role,
		Email:       request.Email,
		PhoneNumber: request.PhoneNumber,
	})
	if err != nil {
		log.Println("Error creating loan:", err)
		response.BuildResponse(w, http.StatusInternalServerError, response.Response{
			Error: err.Error(),
		})
		return
	}

	response.BuildResponse(w, http.StatusOK, response.Response{
		Message: "User created successfully",
		Result: CreateUserResponse{
			UserID: result.UserID,
			Role:   constant.UserRoleText[result.Role],
		},
	})
}

func (h *Handler) TopUpUserBalance(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var (
		request TopUpUserBalanceRequest
		err     error
	)

	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.BuildResponse(w, http.StatusBadRequest, response.Response{
			Error: "Invalid request body",
		})
		return
	}

	if request.TopUpAmount <= 0 {
		response.BuildResponse(w, http.StatusBadRequest, response.Response{
			Error: "Top up amount must be greater than 0",
		})
		return
	}

	userID, _ := ctx.Value(middleware.UserIDContextKey).(int64)
	userRole, _ := ctx.Value(middleware.UserRoleContextKey).(int)

	result, err := h.userService.TopUpUserBalance(ctx, user.TopUpUserBalanceRequest{
		UserID:      userID,
		UserRole:    userRole,
		TopUpAmount: request.TopUpAmount,
	})
	if err != nil {
		log.Println("Error creating loan:", err)
		response.BuildResponse(w, http.StatusInternalServerError, response.Response{
			Error: err.Error(),
		})
		return
	}

	response.BuildResponse(w, http.StatusOK, response.Response{
		Message: "Top up successfull",
		Result: TopUpUserBalanceResponse{
			BalanceAmount: result.TotalBalanceAmount,
		},
	})
}
