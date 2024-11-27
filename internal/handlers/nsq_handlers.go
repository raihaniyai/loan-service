package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/nsqio/go-nsq"

	"loan-service/internal/infrastructure/constant"
	nsqInfra "loan-service/internal/infrastructure/nsq"

	"loan-service/internal/services/action"
)

type NSQHandler struct {
	actionService action.Service
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
		payload, ok := wrappedMessage.Payload.(nsqInfra.InvestmentCompletedMessage)
		if !ok {
			log.Printf("Error casting Payload to InvestmentCompletedMessage")
			return fmt.Errorf("error casting Payload to InvestmentCompletedMessage")
		}

		h.actionService.SendAgreementLetter(context.Background(), action.SendAgreementLetterRequest{
			LoanID:     payload.LoanID,
			InvestorID: payload.InvestorID,
		})
	}

	return nil
}
