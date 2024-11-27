package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	"github.com/nsqio/go-nsq"

	"loan-service/internal/infrastructure/constant"
	nsqInfra "loan-service/internal/infrastructure/nsq"

	"loan-service/internal/services/action"
)

type NSQHandler struct {
	actionService action.Service
}

func NewNSQHandler(actionService action.Service) *NSQHandler {
	return &NSQHandler{
		actionService: actionService,
	}
}

func (h *NSQHandler) HandleMessage(message *nsq.Message) error {
	var wrappedMessage nsqInfra.Message

	err := json.Unmarshal(message.Body, &wrappedMessage)
	if err != nil {
		log.Printf("Error unmarshaling message: %v", err)
		return err
	}

	switch wrappedMessage.Topic {
	case constant.NSQTopicLoanInvestmentCompleted:
		payload, ok := wrappedMessage.Payload.(map[string]interface{})
		if !ok {
			log.Printf("Error: Payload is not a map[string]interface{}")
			return errors.New("payload is not a map[string]interface{}")
		}

		investmentCompletedMessage := nsqInfra.InvestmentCompletedMessage{}
		investmentCompletedMessage.LoanID = int64(payload["loan_id"].(float64))
		investmentCompletedMessage.InvestorID = int64(payload["investor_id"].(float64))

		h.actionService.SendAgreementLetter(context.Background(), action.SendAgreementLetterRequest{
			LoanID:     investmentCompletedMessage.LoanID,
			InvestorID: investmentCompletedMessage.InvestorID,
		})
	}

	return nil
}
