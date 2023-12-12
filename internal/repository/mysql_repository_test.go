package repository

import (
	"errors"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"invoice-test/internal/repository/model"
	"regexp"
	"testing"
)

func TestMysqlRepositoryListInvoices(t *testing.T) {
	db, mock, _ := sqlmock.New()
	t.Run("正常系", func(t *testing.T) {
		t.Run("DBから適切なデータが返ってくる場合", func(t *testing.T) {
			mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM invoices WHERE company_id = ? AND payment_due_date >= ? AND payment_due_date <= ? AND status IN ('PENDING', 'PROCESSING')`)).
				WithArgs(1, "2023-01-01", "2023-02-01").
				WillReturnRows(sqlmock.NewRows([]string{"invoice_id", "company_id", "suppliers_id", "issue_date", "payment_amount", "fee", "fee_rate", "tax", "tax_rate", "total_amount", "payment_due_date", "status"}).
					AddRow(1, 1, 1, "2023-01-01", 100000.00, 1000.0, 0.01, 5000.0, 0.05, 105000, "2023-02-01", "PENDING"))

			user := &model.User{
				UserId:    1,
				CompanyId: 1,
			}
			condition := &model.ListInvoiceCondition{
				From: "2023-01-01",
				To:   "2023-02-01",
			}

			repo := MysqlRepository{
				client: db,
			}

			res, _ := repo.ListInvoices(user, condition)

			assert.Equal(t, 1, len(res))

		})
	})

	t.Run("異常系", func(t *testing.T) {
		t.Run("クエリが失敗する場合", func(t *testing.T) {
			mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM invoices WHERE company_id = ? AND payment_due_date >= ? AND payment_due_date <= ? AND status IN ('PENDING', 'PROCESSING')`)).WithArgs(1, "2023-01-01", "2023-02-01").WillReturnError(errors.New("query error"))

			user := &model.User{
				UserId:    1,
				CompanyId: 1,
			}
			condition := &model.ListInvoiceCondition{
				From: "2023-01-01",
				To:   "2023-02-01",
			}

			repo := MysqlRepository{
				client: db,
			}

			_, err := repo.ListInvoices(user, condition)

			assert.EqualError(t, err, "query error")

		})

		t.Run("カラムが適切でない場合", func(t *testing.T) {
			mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM invoices WHERE company_id = ? AND payment_due_date >= ? AND payment_due_date <= ? AND status IN ('PENDING', 'PROCESSING')`)).
				WithArgs(1, "2023-01-01", "2023-02-01").
				WillReturnRows(sqlmock.NewRows([]string{"company_id", "suppliers_id", "issue_date", "payment_amount", "fee", "fee_rate", "tax", "tax_rate", "total_amount", "payment_due_date", "status"}).
					AddRow(1, 1, "2023-01-01", 100000.00, 1000.0, 0.01, 5000.0, 0.05, 105000, "2023-02-01", "PENDING"))

			user := &model.User{
				UserId:    1,
				CompanyId: 1,
			}
			condition := &model.ListInvoiceCondition{
				From: "2023-01-01",
				To:   "2023-02-01",
			}

			repo := MysqlRepository{
				client: db,
			}

			_, err := repo.ListInvoices(user, condition)

			assert.EqualError(t, err, "sql: expected 11 destination arguments in Scan, not 12")

		})

	})

}

func TestFindUser(t *testing.T) {
	db, mock, _ := sqlmock.New()
	t.Run("正常系", func(t *testing.T) {
		t.Run("適切なユーザーデータが返ってくる場合", func(t *testing.T) {
			mock.ExpectQuery(regexp.QuoteMeta(`SELECT solt FROM users WHERE email = ?`)).WithArgs("hoge@gmail.com").WillReturnRows(sqlmock.NewRows([]string{"solt"}).AddRow("solt"))
			mock.ExpectQuery(regexp.QuoteMeta(`SELECT user_id, company_id FROM users WHERE email = ? AND password = ?`)).WithArgs("hoge@gmail.com", "8e7ad3c83d2f85d84c4e080541f9bea646dc7ba747ee3d7cbfe47dd9f5198966").WillReturnRows(sqlmock.NewRows([]string{"user_id", "company_id"}).AddRow(1, 1))

			repo := MysqlRepository{
				client: db,
			}

			res, _ := repo.FindUser("hoge@gmail.com", "password")

			assert.Equal(t, 1, res.UserId)

		})
	})

	t.Run("異常系", func(t *testing.T) {
		t.Run("存在しないemailの場合", func(t *testing.T) {
			mock.ExpectQuery(regexp.QuoteMeta(`SELECT solt FROM users WHERE email = ?`)).WithArgs("hoge@gmail.com").WillReturnRows(sqlmock.NewRows([]string{"solt"}))
			repo := MysqlRepository{
				client: db,
			}
			_, err := repo.FindUser("hoge@gmail.com", "password")
			assert.EqualError(t, err, "user not found")
		})

		t.Run("パスワードが間違っている場合", func(t *testing.T) {
			mock.ExpectQuery(regexp.QuoteMeta(`SELECT solt FROM users WHERE email = ?`)).WithArgs("hoge@gmail.com").WillReturnRows(sqlmock.NewRows([]string{"solt"}).AddRow("solt"))
			mock.ExpectQuery(regexp.QuoteMeta(`SELECT user_id, company_id FROM users WHERE email = ? AND password = ?`)).WithArgs("hoge@gmail.com", "187440bac37e2d1fa33edaf4b9d76bcef94f5318131bd63d508d62dca266a267").WillReturnRows(sqlmock.NewRows([]string{"user_id", "company_id"}))

			repo := MysqlRepository{
				client: db,
			}
			res, err := repo.FindUser("hoge@gmail.com", "invalid")
			fmt.Println(res)
			assert.EqualError(t, err, "user not found")

		})

		t.Run("soltカラムが適切でない場合", func(t *testing.T) {
			mock.ExpectQuery(regexp.QuoteMeta(`SELECT solt FROM users WHERE email = ?`)).WithArgs("hoge@gmail.com").WillReturnRows(sqlmock.NewRows([]string{}).AddRow())
			repo := MysqlRepository{
				client: db,
			}
			res, err := repo.FindUser("hoge@gmail.com", "invalid")
			fmt.Println(res)
			assert.EqualError(t, err, "sql: expected 0 destination arguments in Scan, not 1")
		})

		t.Run("user_idカラムが適切でない場合", func(t *testing.T) {
			mock.ExpectQuery(regexp.QuoteMeta(`SELECT solt FROM users WHERE email = ?`)).WithArgs("hoge@gmail.com").WillReturnRows(sqlmock.NewRows([]string{"solt"}).AddRow("solt"))
			mock.ExpectQuery(regexp.QuoteMeta(`SELECT user_id, company_id FROM users WHERE email = ? AND password = ?`)).WithArgs("hoge@gmail.com", "187440bac37e2d1fa33edaf4b9d76bcef94f5318131bd63d508d62dca266a267").WillReturnRows(sqlmock.NewRows([]string{"company_id"}).AddRow(1))
			repo := MysqlRepository{
				client: db,
			}
			res, err := repo.FindUser("hoge@gmail.com", "invalid")
			fmt.Println(res)
			assert.EqualError(t, err, "sql: expected 1 destination arguments in Scan, not 2")
		})
	})
}

func TestMysqlRepositoryPostInvoice(t *testing.T) {
	db, mock, _ := sqlmock.New()
	t.Run("正常系", func(t *testing.T) {
		t.Run("請求データ生成に成功する場合", func(t *testing.T) {
			mock.ExpectPrepare(regexp.QuoteMeta(`INSERT INTO invoices(company_id, suppliers_id, issue_date, payment_amount, fee, fee_rate, tax, tax_rate, total_amount, payment_due_date, status) VALUES(?,?,?,?,?,?,?,?,?,?,?)`)).
				ExpectExec().WithArgs(1, 1, "2023-12-13", 1000.0, 40.0, 0.04, 4.0, 0.1, 1044.0, "2023-12-13", "PENDING").WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM invoices WHERE invoice_id = ?`)).
				WithArgs(1).
				WillReturnRows(sqlmock.NewRows([]string{"invoice_id", "company_id", "suppliers_id", "issue_date", "payment_amount", "fee", "fee_rate", "tax", "tax_rate", "total_amount", "payment_due_date", "status"}).
					AddRow(1, 1, 1, "2023-12-13", 1000.0, 40.0, 0.04, 4.0, 0.1, 1044.0, "2023-12-13", "PENDING"))
			repo := MysqlRepository{
				client: db,
			}
			user := &model.User{
				UserId:    1,
				CompanyId: 1,
			}
			condition := &model.PostInvoiceCondition{
				SuppliersId:    1,
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
			res, _ := repo.PostInvoice(user, condition)
			assert.Equal(t, 1, int(res.InvoiceId))
		})
	})

	t.Run("異常系", func(t *testing.T) {
		t.Run("prepareクエリが失敗する場合", func(t *testing.T) {
			mock.ExpectPrepare(regexp.QuoteMeta(`INSERT INTO invoices(company_id, suppliers_id, issue_date, payment_amount, fee, fee_rate, tax, tax_rate, total_amount, payment_due_date, status) VALUES(?,?,?,?,?,?,?,?,?,?,?)`)).WillReturnError(errors.New("prepare error"))
			repo := MysqlRepository{
				client: db,
			}
			user := &model.User{
				UserId:    1,
				CompanyId: 1,
			}
			condition := &model.PostInvoiceCondition{
				SuppliersId:    1,
				IssueDate:      "2023-12-13",
				PaymentAmount:  1000,
				Fee:            40.0,
				FeeRate:        0.04,
				Tax:            4,
				TaxRate:        0.1,
				TotalAmount:    1044.0,
				PaymentDueDate: "2023-12-13",
				Status:         "PENDING",
			}
			_, err := repo.PostInvoice(user, condition)
			assert.EqualError(t, err, "prepare error")
		})

		t.Run("insertクエリが失敗する場合", func(t *testing.T) {
			mock.ExpectPrepare(regexp.QuoteMeta(`INSERT INTO invoices(company_id, suppliers_id, issue_date, payment_amount, fee, fee_rate, tax, tax_rate, total_amount, payment_due_date, status) VALUES(?,?,?,?,?,?,?,?,?,?,?)`)).
				ExpectExec().WithArgs(1, 1, "2023-12-13", 1000.0, 40.0, 0.04, 4.0, 0.1, 1044.0, "2023-12-13", "PENDING").WillReturnError(errors.New("exec error"))
			repo := MysqlRepository{
				client: db,
			}
			user := &model.User{
				UserId:    1,
				CompanyId: 1,
			}
			condition := &model.PostInvoiceCondition{
				SuppliersId:    1,
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
			_, err := repo.PostInvoice(user, condition)
			assert.EqualError(t, err, "exec error")
		})

		t.Run("取得するカラムが適切でない場合", func(t *testing.T) {
			mock.ExpectPrepare(regexp.QuoteMeta(`INSERT INTO invoices(company_id, suppliers_id, issue_date, payment_amount, fee, fee_rate, tax, tax_rate, total_amount, payment_due_date, status) VALUES(?,?,?,?,?,?,?,?,?,?,?)`)).
				ExpectExec().WithArgs(1, 1, "2023-12-13", 1000.0, 40.0, 0.04, 4.0, 0.1, 1044.0, "2023-12-13", "PENDING").WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM invoices WHERE invoice_id = ?`)).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{}).AddRow())
			repo := MysqlRepository{
				client: db,
			}
			user := &model.User{
				UserId:    1,
				CompanyId: 1,
			}
			condition := &model.PostInvoiceCondition{
				SuppliersId:    1,
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
			_, err := repo.PostInvoice(user, condition)
			assert.EqualError(t, err, "sql: expected 0 destination arguments in Scan, not 12")
		})

	})
}
