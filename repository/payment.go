package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"tourmate/payment-service/constant/noti"
	"tourmate/payment-service/interface/repo"
	"tourmate/payment-service/model/dto/request"
	"tourmate/payment-service/model/entity"
)

type paymentRepo struct {
	db     *sql.DB
	logger *log.Logger
}

func InitializePaymentRepo(db *sql.DB, logger *log.Logger) repo.IPaymentRepo {
	return &paymentRepo{
		db:     db,
		logger: logger,
	}
}

const (
	payment_limit_records int = 10
)

// CreatePayment implements repo.IPaymentRepo.
func (p *paymentRepo) CreatePayment(payment entity.Payment, ctx context.Context) error {
	var query string = "INSERT INTO " + payment.GetPaymentTable() +
		"(paymentId, customerId, invoiceId, " +
		"price, status, paymentMethod, createdAt) " +
		"values ($1, $2, $3, $4, $5, $6, $7)"
	var errLogMsg string = fmt.Sprintf(noti.REPO_ERR_MSG, payment.GetPaymentTable()) + "CreatePayment - "

	if _, err := p.db.Exec(query, payment.PaymentId, payment.CustomerId, payment.InvoiceId,
		payment.Price, payment.Status, payment.PaymentMethod, payment.CreatedAt); err != nil {

		p.logger.Println(errLogMsg + err.Error())
		return errors.New(noti.INTERNALL_ERR_MSG)
	}

	return nil
}

// GetAllPayments implements repo.IPaymentRepo.
func (p *paymentRepo) GetPayments(req request.GetPaymentsRequest, ctx context.Context) (*[]entity.Payment, int, error) {
	var table string = entity.Payment{}.GetPaymentTable()
	var errLogMsg string = fmt.Sprintf(noti.REPO_ERR_MSG, table) + "GetPayments - "

	var queryCondition string = "WHERE "
	var isHavePreviousCond bool = false
	if req.Method != "" {
		queryCondition += fmt.Sprintf("LOWER(paymentMethod) LIKE LOWER('%%%s%%') ", req.Method)
		isHavePreviousCond = true
	}
	if req.Status != "" {
		if isHavePreviousCond {
			queryCondition += " AND "
		}

		queryCondition += fmt.Sprintf("LOWER(status) LIKE LOWER('%%%s%%') ", req.Status)
		isHavePreviousCond = true
	}
	if req.CustomerId != nil {
		if isHavePreviousCond {
			queryCondition += " AND "
		}

		queryCondition += fmt.Sprintf("customerId = '%d'", *req.CustomerId)
	}

	if queryCondition == "WHERE" {
		queryCondition = ""
	}

	var orderCondition string = generateOrderCondition(req.Request.FilterProp, req.Request.Order)
	var query string = generateRetrieveQuery(table, queryCondition+orderCondition, payment_limit_records, req.Request.Page, false)

	rows, err := p.db.Query(query)
	if err != nil {
		p.logger.Println(errLogMsg + err.Error())
		return nil, 0, errors.New(noti.INTERNALL_ERR_MSG)
	}

	var res []entity.Payment
	for rows.Next() {
		var x entity.Payment
		if err := rows.Scan(
			&x.PaymentId, &x.Price, &x.Status,
			&x.CreatedAt, &x.PaymentMethod, &x.CustomerId, &x.InvoiceId); err != nil {

			p.logger.Println(errLogMsg + err.Error())
			return nil, 0, errors.New(noti.INTERNALL_ERR_MSG)
		}

		res = append(res, x)
	}

	// Track total records in table
	var totalRecords int
	p.db.QueryRow(generateRetrieveQuery(table, queryCondition, payment_limit_records, req.Request.Page, true)).Scan(&totalRecords)

	return &res, caculateTotalPages(totalRecords, payment_limit_records), nil
}

// [paymentId] [int] IDENTITY(1,1) NOT NULL,
//
//	[price] [real] NOT NULL,
//	[status] [nvarchar](50) NOT NULL,
//	[createdAt] [datetime] NOT NULL,
//	[paymentMethod] [varchar](50) NOT NULL,
//	[customerId] [int] NOT NULL,
//	[invoiceId] [int] NULL
func (p *paymentRepo) GetPaymentById(id int, ctx context.Context) (*entity.Payment, error) {
	var res entity.Payment
	var table string = res.GetPaymentTable()
	var query string = "SELECT * FROM " + table + " WHERE paymentId = $1"
	var errLogMsg string = fmt.Sprintf(noti.REPO_ERR_MSG, table) + "GetPaymentById - "

	if err := p.db.QueryRow(query, id).Scan(
		&res.PaymentId, &res.Price, &res.Status, &res.CreatedAt,
		&res.PaymentMethod, &res.CustomerId, &res.CreatedAt); err != nil {

		if err == sql.ErrNoRows {
			return nil, nil
		}

		p.logger.Println(errLogMsg + err.Error())
		return nil, errors.New(noti.INTERNALL_ERR_MSG)
	}

	return &res, nil
}

// UpdatePayment implements repo.IPaymentRepo.
func (p *paymentRepo) UpdatePayment(payment entity.Payment, ctx context.Context) error {
	var errLogMsg string = fmt.Sprintf(noti.REPO_ERR_MSG, payment.GetPaymentTable()) + "UpdatePayment - "
	var query string = "UPDATE " + payment.GetPaymentTable() + " SET status = $1,  method = $2, updated_at = $3 WHERE paymentId = $4"

	res, err := p.db.Exec(query, payment.Status, payment.PaymentMethod, payment.PaymentId)

	var INTERNALL_ERR_MSGMsg error = errors.New(noti.INTERNALL_ERR_MSG)

	if err != nil {
		p.logger.Println(errLogMsg + err.Error())
		return INTERNALL_ERR_MSGMsg
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		p.logger.Println(errLogMsg + err.Error())
		return INTERNALL_ERR_MSGMsg
	}

	if rowsAffected == 0 {
		return errors.New(fmt.Sprintf(noti.UNDEFINED_OBJECT_WARN_MSG, payment.GetPaymentTable()))
	}

	return nil
}
