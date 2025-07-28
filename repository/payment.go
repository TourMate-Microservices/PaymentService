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

	_ "github.com/lib/pq"
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

// CreatePaymentWithScopeId implements repo.IPaymentRepo.
func (p *paymentRepo) CreatePaymentWithScopeId(payment entity.Payment, ctx context.Context) (int, error) {
	var errLogMsg string = fmt.Sprintf(noti.REPO_ERR_MSG, payment.GetPaymentTable()) + "CreatePaymentWithScopeId - "
	var internalErr error = errors.New(noti.INTERNALL_ERR_MSG)
	var query string = "INSERT INTO " + payment.GetPaymentTable() +
		" (customerId, invoiceId, " +
		"price, paymentMethod, createdAt, serviceId, status) " +
		"values (@p1, @p2, @p3, @p4, @p5, @p6, @p7)"

	tx, err := p.db.BeginTx(ctx, nil)
	if err != nil {
		p.logger.Println(errLogMsg + err.Error())
		return -1, internalErr
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, query, payment.CustomerId, payment.InvoiceId,
		payment.Price, payment.PaymentMethod, payment.CreatedAt, payment.ServiceId, payment.Status)
	if err != nil {
		p.logger.Println(errLogMsg + err.Error())
		return -1, internalErr
	}

	var res int
	if err := tx.QueryRowContext(ctx, `SELECT SCOPE_IDENTITY()`).Scan(&res); err != nil {
		p.logger.Println(errLogMsg + err.Error())
		return -1, err
	}

	if err := tx.Commit(); err != nil {
		p.logger.Println(errLogMsg + err.Error())
		return -1, err
	}

	return res, nil
}

// CreatePayment implements repo.IPaymentRepo.
func (p *paymentRepo) CreatePayment(payment entity.Payment, ctx context.Context) (*entity.Payment, error) {
	var query string = "INSERT INTO " + payment.GetPaymentTable() +
		" (customerId, invoiceId, " +
		"price, paymentMethod, createdAt, serviceId, status) " +
		"OUTPUT INSERTED.paymentId " +
		"values (@p1, @p2, @p3, @p4, @p5, @p6, @p7)"
	var errLogMsg string = fmt.Sprintf(noti.REPO_ERR_MSG, payment.GetPaymentTable()) + "CreatePayment - "

	var paymentId int
	if err := p.db.QueryRow(query, payment.CustomerId, payment.InvoiceId,
		payment.Price, payment.PaymentMethod, payment.CreatedAt, payment.ServiceId, payment.Status).Scan(&paymentId); err != nil {

		p.logger.Println(errLogMsg + err.Error())
		return nil, errors.New(noti.INTERNALL_ERR_MSG)
	}

	// Set the generated ID and return the payment
	payment.PaymentId = paymentId
	return &payment, nil
}

// GetAllPayments implements repo.IPaymentRepo.
func (p *paymentRepo) GetPayments(req request.GetPaymentsRequest, ctx context.Context) (*[]entity.Payment, int, int, error) {
	var table string = entity.Payment{}.GetPaymentTable()
	var errLogMsg string = fmt.Sprintf(noti.REPO_ERR_MSG, table) + "GetPayments - "
	var limitRecords int = req.PageSize

	var queryCondition string = "WHERE "
	var isHavePreviousCond bool = false

	if req.Method != "" {
		queryCondition += fmt.Sprintf("LOWER(paymentMethod) LIKE LOWER('%%%%%s%%%%')", req.Method)
		isHavePreviousCond = true
	}
	if req.CustomerId != nil {
		if isHavePreviousCond {
			queryCondition += " AND "
		}

		queryCondition += fmt.Sprintf("customerId = %d", *req.CustomerId)
	}

	if queryCondition == "WHERE " {
		queryCondition = ""
	}

	var orderCondition string = generateOrderCondition(req.Request.FilterProp, req.Request.Order)
	var query string = generateRetrieveQuery(table, queryCondition+orderCondition, limitRecords, req.Request.Page, false)

	p.logger.Println("Query: ", query)

	rows, err := p.db.Query(query)
	if err != nil {
		p.logger.Println(errLogMsg + err.Error())
		return nil, 0, 0, errors.New(noti.INTERNALL_ERR_MSG)
	}

	var res []entity.Payment
	for rows.Next() {
		var x entity.Payment
		if err := rows.Scan(
			&x.PaymentId, &x.Price,
			&x.CreatedAt, &x.PaymentMethod, &x.InvoiceId, &x.CustomerId, &x.ServiceId, &x.Status); err != nil {

			p.logger.Println(errLogMsg + err.Error())
			return nil, 0, 0, errors.New(noti.INTERNALL_ERR_MSG)
		}

		res = append(res, x)
	}

	// Track total records in table
	var totalRecords int
	p.db.QueryRow(generateRetrieveQuery(table, queryCondition, limitRecords, req.Request.Page, true)).Scan(&totalRecords)

	return &res, caculateTotalPages(totalRecords, limitRecords), totalRecords, nil
}

// GetPaymentById implements repo.IPaymentRepo
func (p *paymentRepo) GetPaymentById(id int, ctx context.Context) (*entity.Payment, error) {
	var res entity.Payment
	var table string = res.GetPaymentTable()
	var query string = "SELECT * FROM " + table + " WHERE paymentId = @p1"
	var errLogMsg string = fmt.Sprintf(noti.REPO_ERR_MSG, table) + "GetPaymentById - "

	if err := p.db.QueryRow(query, id).Scan(
		&res.PaymentId, &res.Price, &res.CreatedAt,
		&res.PaymentMethod, &res.InvoiceId, &res.CustomerId, &res.ServiceId, &res.Status); err != nil {

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
	var query string = "UPDATE " + payment.GetPaymentTable() + " SET paymentMethod = @p1 WHERE paymentId = @p2"

	res, err := p.db.Exec(query, payment.PaymentMethod, payment.PaymentId)

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
