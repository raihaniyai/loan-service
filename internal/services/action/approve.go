package action

import (
	"context"
	"errors"
	"log"
	"time"

	"loan-service/internal/entity"
	"loan-service/internal/infrastructure/constant"
)

func (svc *service) ApproveLoan(ctx context.Context, request ApproveLoanRequest) (ApproveLoanResult, error) {
	if request.User.Role != constant.UserRoleAdmin {
		log.Println("SVC.AL00 | [ApproveLoan] User is not eligible to approve loan")
		err := errors.New("user is not eligible to approve loan")
		return ApproveLoanResult{}, err
	}

	loan, err := svc.loanRepository.GetLoanByID(ctx, request.LoanID)
	if err != nil {
		log.Println("SVC.AL01 | [ApproveLoan] Error getting loan:", err)
		return ApproveLoanResult{}, err
	}

	if loan == nil {
		return ApproveLoanResult{}, errors.New("loan not found")
	}

	if loan.Status != constant.LoanStatusProposed {
		return ApproveLoanResult{}, errors.New("loan is not eligible to be approved")
	}

	tx := svc.database.BeginTx()
	defer func() {
		if err != nil {
			err = svc.database.Rollback(tx)
			if err != nil {
				log.Println("SVC.AL02 | [ApproveLoan] Error rolling back transaction:", err)
			}
		}
	}()

	currentTime := time.Now()
	loan.Status = constant.LoanStatusApproved
	loan.UpdatedAt = currentTime
	loan.UpdatedBy = request.User.UserID
	err = svc.loanRepository.UpdateLoan(ctx, tx, loan)
	if err != nil {
		log.Println("SVC.AL03 | [ApproveLoan] Error updating loan:", err)
		return ApproveLoanResult{}, err
	}

	actionApprove := &entity.Action{
		LoanID:      request.LoanID,
		ActionType:  constant.ActionTypeApproveLoan,
		ActionBy:    request.User.Role,
		DocumentURL: request.DocumentURL,
		CreatedAt:   currentTime,
		CreatedBy:   request.User.UserID,
	}
	_, err = svc.actionRepository.SetAction(ctx, tx, actionApprove)
	if err != nil {
		return ApproveLoanResult{}, err
	}

	err = tx.Commit().Error
	if err != nil {
		log.Println("SVC.AL04 | [ApproveLoan] Error committing transaction:", err)
		return ApproveLoanResult{}, err
	}

	return ApproveLoanResult{}, nil
}
