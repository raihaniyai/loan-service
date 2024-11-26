package action

import (
	"loan-service/internal/services/action"
)

type Handler struct {
	actionService action.Service
}

func New(actionService action.Service) Handler {
	return Handler{
		actionService: actionService,
	}
}
