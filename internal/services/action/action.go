package action

import (
	"context"
	"errors"
	"log"
	"time"

	"loan-service/internal/entity"
	"loan-service/internal/infrastructure/constant"
)

func (svc *service) UpdateLoan(ctx context.Context, request UpdateLoanRequest) (UpdateLoanResult, error) {
	if request.UserRole != constant.UserRoleAdmin {
		log.Println("SVC.UL00 | [UpdateLoan] User is not eligible to update loan")
		err := errors.New("user is not eligible to update loan")
		return UpdateLoanResult{}, err
	}

	loan, err := svc.loanRepository.GetLoanByID(ctx, request.LoanID)
	if err != nil {
		log.Println("SVC.UL01 | [UpdateLoan] Error getting loan:", err)
		return UpdateLoanResult{}, err
	}

	if loan == nil {
		return UpdateLoanResult{}, errors.New("loan not found")
	}

	if (request.ActionType == constant.ActionTypeApproveLoan && loan.Status != constant.LoanStatusProposed) ||
		(request.ActionType == constant.ActionTypeDisburse && loan.Status != constant.LoanStatusInvested) {
		errorWording := "loan is not eligible to be approved"
		if request.ActionType == constant.ActionTypeDisburse {
			errorWording = "loan is not eligible to be disbursed"
		}
		log.Println("SVC.UL02 | [UpdateLoan] " + errorWording)
		return UpdateLoanResult{}, errors.New(errorWording)
	}

	tx := svc.database.BeginTx()
	defer func() {
		if err != nil {
			err = svc.database.Rollback(tx)
			if err != nil {
				log.Println("SVC.UL03 | [UpdateLoan] Error rolling back transaction:", err)
			}
		}
	}()

	status := constant.LoanStatusApproved
	if request.ActionType == constant.ActionTypeDisburse {
		status = constant.LoanStatusDisbursed

		userBalance, err := svc.fundRepository.GetBalanceByUserID(ctx, loan.BorrowerID)
		if err != nil {
			log.Println("SVC.UL07 | [UpdateLoan] Error getting user balance:", err)
			return UpdateLoanResult{}, err
		}

		err = svc.fundRepository.UpdateBalanceByUserID(ctx, tx, loan.BorrowerID, userBalance+loan.PrincipalAmount)
		if err != nil {
			log.Println("SVC.UL08 | [UpdateLoan] Error updating user balance:", err)
			return UpdateLoanResult{}, err
		}
	}

	currentTime := time.Now()
	loan.Status = status
	loan.UpdatedAt = currentTime
	loan.UpdatedBy = request.UserID
	err = svc.loanRepository.UpdateLoan(ctx, tx, loan)
	if err != nil {
		log.Println("SVC.UL04 | [UpdateLoan] Error updating loan:", err)
		return UpdateLoanResult{}, err
	}

	actionApprove := &entity.Action{
		LoanID:      request.LoanID,
		ActionType:  request.ActionType,
		ActionBy:    request.UserRole,
		DocumentURL: request.DocumentURL,
		CreatedAt:   currentTime,
		CreatedBy:   request.UserID,
	}
	_, err = svc.actionRepository.SetAction(ctx, tx, actionApprove)
	if err != nil {
		log.Println("SVC.UL05 | [UpdateLoan] Error setting action:", err)
		return UpdateLoanResult{}, err
	}

	errCommit := tx.Commit().Error
	if errCommit != nil {
		log.Println("SVC.UL06 | [UpdateLoan] Error committing transaction:", errCommit)
		return UpdateLoanResult{}, errCommit
	}

	return UpdateLoanResult{
		LoanID: loan.LoanID,
	}, nil
}
