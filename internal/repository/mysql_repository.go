package repository

import (
	"crypto/sha256"
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"invoice-test/internal/repository/model"
)

type MysqlRepository struct {
	client *sql.DB
}

func NewMysqlRepository(viper *viper.Viper) (*MysqlRepository, error) {
	conf := mysql.Config{
		User:   viper.GetString("MYSQL_DB_USER"),
		Passwd: viper.GetString("MYSQL_DB_PASSWORD"),
		Net:    "tcp",
		Addr:   viper.GetString("MYSQL_DB_URL"),
		DBName: viper.GetString("MYSQL_DB_NAME"),
	}

	db, err := sql.Open("mysql", conf.FormatDSN())
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &MysqlRepository{
		client: db,
	}, nil

}

func (m *MysqlRepository) ListInvoices(user *model.User, condition *model.ListInvoiceCondition) ([]model.Invoice, error) {
	rows, err := m.client.Query("SELECT * FROM invoices WHERE company_id = ? AND payment_due_date >= ? AND payment_due_date <= ? AND status IN ('PENDING', 'PROCESSING')", user.CompanyId, condition.From, condition.To)
	if err != nil {
		return nil, err
	}

	invoices := make([]model.Invoice, 0)
	for rows.Next() {
		var invoiceId uint64
		var companyId uint64
		var suppliersId uint64
		var issueDate string
		var paymentAmount float64
		var fee float64
		var feeRate float64
		var tax float64
		var taxRate float64
		var totalAmount float64
		var paymentDueDate string
		var status string
		err = rows.Scan(&invoiceId, &companyId, &suppliersId, &issueDate, &paymentAmount, &fee, &feeRate, &tax, &taxRate, &totalAmount, &paymentDueDate, &status)
		if err != nil {
			return nil, err
		}

		invoice := model.Invoice{
			InvoiceId:      invoiceId,
			CompanyId:      companyId,
			SuppliersId:    suppliersId,
			IssueDate:      issueDate,
			PaymentAmount:  paymentAmount,
			Fee:            fee,
			FeeRate:        feeRate,
			Tax:            tax,
			TaxRate:        taxRate,
			TotalAmount:    totalAmount,
			PaymentDueDate: paymentDueDate,
			Status:         status,
		}

		invoices = append(invoices, invoice)

	}

	return invoices, err

}

func (m *MysqlRepository) FindUser(email string, password string) (*model.User, error) {
	rows, err := m.client.Query("SELECT solt FROM users WHERE email = ?", email)
	if err != nil {
		return nil, err
	}

	var solt string
	if rows.Next() {
		err = rows.Scan(&solt)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("user not found")
	}

	h := sha256.New()
	h.Write([]byte(password + solt))
	hashedPassword := fmt.Sprintf("%x", h.Sum(nil))
	rows, err = m.client.Query("SELECT user_id, company_id FROM users WHERE email = ? AND password = ?", email, hashedPassword)
	if err != nil {
		return nil, fmt.Errorf("password not valid: %s", err)
	}

	var userId int
	var companyId int
	if rows.Next() {
		err = rows.Scan(&userId, &companyId)
		if err != nil {
			return nil, err
		}
		return &model.User{
			UserId:    userId,
			CompanyId: companyId,
		}, nil
	} else {
		return nil, errors.New("user not found")
	}

}
