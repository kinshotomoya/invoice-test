package model

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	mock_model "invoice-test/internal/mock"
	"invoice-test/internal/repository/model"
	"testing"
)

func TestPostInvoiceCondition_NewPostInvoiceCondition(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	t.Run("正常にConditionが生成できる", func(t *testing.T) {
		time := mock_model.NewMockCustomTimeInterface(ctrl)
		suppliersId := uint64(1)
		paymentAmount := 1000.0
		paymentDueDate := "2023-12-13"
		condition := &PostInvoiceCondition{
			SuppliersId:    &suppliersId,
			PaymentAmount:  &paymentAmount,
			PaymentDueDate: &paymentDueDate,
		}

		time.EXPECT().NowDateOnly().Return("2023-12-13")
		actual := condition.NewPostInvoiceCondition(time)

		expected := &model.PostInvoiceCondition{
			SuppliersId:    suppliersId,
			IssueDate:      "2023-12-13",
			PaymentAmount:  1000.0,
			Fee:            40.0,
			FeeRate:        0.04,
			Tax:            4.0,
			TaxRate:        0.1,
			TotalAmount:    1044.0,
			PaymentDueDate: "2023-12-13",
			Status:         "PENDING",
		}

		assert.Equal(t, expected, actual)

	})
}
