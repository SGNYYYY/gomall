package service

import (
	"context"
	"strconv"
	"time"

	"github.com/SGNYYYY/gomall/app/payment/biz/dal/mysql"
	"github.com/SGNYYYY/gomall/app/payment/biz/model"
	payment "github.com/SGNYYYY/gomall/rpc_gen/kitex_gen/payment"
	"github.com/cloudwego/kitex/pkg/kerrors"
	creditcart "github.com/durango/go-credit-card"
	"github.com/google/uuid"
)

type ChargeService struct {
	ctx context.Context
} // NewChargeService new ChargeService
func NewChargeService(ctx context.Context) *ChargeService {
	return &ChargeService{ctx: ctx}
}

// Run create note info
func (s *ChargeService) Run(req *payment.ChargeReq) (resp *payment.ChargeResp, err error) {
	// Finish your business logic.
	card := creditcart.Card{
		Number: req.CreditCard.CreditCardNumber,
		Cvv:    strconv.Itoa(int(req.CreditCard.CreditCardCvv)),
		Month:  strconv.Itoa(int(req.CreditCard.CreditCardExpirationMonth)),
		Year:   strconv.Itoa(int(req.CreditCard.CreditCardExpirationYear)),
	}

	err = card.Validate(true) // true表示允许使用特殊卡号，比如测试卡号
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(4004001, err.Error())
	}
	transactionId, err := uuid.NewRandom()
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(4005001, err.Error())
	}
	err = model.CreatePaymentLog(s.ctx, mysql.DB, &model.PaymentLog{
		UserId:        req.UserId,
		OrderId:       req.OrderId,
		TransactionId: transactionId.String(),
		Amount:        req.Amount,
		PayAt:         time.Now(),
	})
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(4005002, err.Error())
	}
	return &payment.ChargeResp{
		TransactionId: transactionId.String(),
	}, nil
}
