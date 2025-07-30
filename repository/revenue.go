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

type revenueRepo struct {
	db     *sql.DB
	logger *log.Logger
}

func InitializeRevenueRepo(db *sql.DB, logger *log.Logger) repo.IRevenueRepo {
	return &revenueRepo{
		db:     db,
		logger: logger,
	}
}

// GetRevenues implements repo.IRevenueRepo.
func (r *revenueRepo) GetRevenues(req request.GetRevenuesRequest, ctx context.Context) (*[]entity.Revenue, error) {
	var table string = entity.Payment{}.GetPaymentTable()
	var errLogMsg string = fmt.Sprintf(noti.REPO_ERR_MSG, table) + "GetRevenues - "
	var limitRecords int = *req.PageSize
	var internalErr error = errors.New(noti.INTERNALL_ERR_MSG)

	var queryCondition string = fmt.Sprintf("WHERE tourGuideId = %d", req.TourGuideId)

	if req.Year != nil {
		queryCondition += fmt.Sprintf(" AND YEAR(createdAt) = %d", *req.Year)
	}

	if req.Month != nil {
		queryCondition += fmt.Sprintf(" AND MONTH(createdAt) = %d", *req.Month)
	}

	if req.PaymentStatus != nil {
		var intStatus int = 0
		if *req.PaymentStatus {
			intStatus = 1
		}

		queryCondition += fmt.Sprintf(" AND paymentStatus = %d", intStatus)
	}

	var orderCondition string = generateOrderCondition("createdAt", "DESC")
	var query string = generateRetrieveQuery(table, queryCondition+orderCondition, limitRecords, *req.PageNumber, false)

	rows, err := r.db.Query(query)
	if err != nil {
		r.logger.Println(errLogMsg + err.Error())
		return nil, internalErr
	}

	var res []entity.Revenue
	for rows.Next() {
		var x entity.Revenue
		if err := rows.Scan(
			&x.RevenueId, &x.PaymentId, &x.TourGuideId, &x.InvoiceId,
			&x.TotalAmount, &x.ActualReceived, &x.PlatformCommission, &x.PaymentStatus, &x.CreatedAt); err != nil {

			r.logger.Println(errLogMsg + err.Error())
			return nil, internalErr
		}

		res = append(res, x)
	}

	return &res, nil
}

// GetRevenuesByMonth implements repo.IRevenueRepo.
func (r *revenueRepo) GetRevenuesByMonth(tourGuideId int, year int, month int, ctx context.Context) (*[]entity.Revenue, error) {
	var table string = entity.Revenue{}.GetRevenueTable()
	var errLogMsg string = fmt.Sprintf(noti.REPO_ERR_MSG, table) + "GetRevenuesByMonth - "
	var query string = "SELECT * FROM " + table + " WHERE tourGuideId = @p1 AND YEAR(createdAt) = @p2 AND MONTH(createdAt) = @p3 ORDER BY created_at DESC"
	var internalErr error = errors.New(noti.INTERNALL_ERR_MSG)

	rows, err := r.db.Query(query, tourGuideId, year, month)
	if err != nil {
		r.logger.Println(errLogMsg + err.Error())
		return nil, internalErr
	}

	var res []entity.Revenue
	for rows.Next() {
		var x entity.Revenue
		if err := rows.Scan(
			&x.RevenueId, &x.PaymentId, &x.TourGuideId, &x.InvoiceId,
			&x.TotalAmount, &x.ActualReceived, &x.PlatformCommission, &x.PaymentStatus, &x.CreatedAt); err != nil {

			r.logger.Println(errLogMsg + err.Error())
			return nil, internalErr
		}

		res = append(res, x)
	}

	return &res, nil
}

// GetRevenueTotalAmountByMonth implements repo.IRevenueRepo.
func (r *revenueRepo) GetRevenueTotalAmountByMonth(tourGuideId int, year int, month int, ctx context.Context) (float64, error) {
	var table string = entity.Revenue{}.GetRevenueTable()
	var errLogMsg string = fmt.Sprintf(noti.REPO_ERR_MSG, table) + "GetRevenueTotalAmountByMoth - "
	var query string = "SELECT COALESCE(SUM(totalAmount), 0) FROM " + table + " WHERE tourGuideId = @p1 AND YEAR(createdAt) = @p2 AND MONTH(createdAt) = @p3"

	var totalAmount float64
	if err := r.db.QueryRow(query, tourGuideId, year, month).Scan(&totalAmount); err != nil {
		r.logger.Println(errLogMsg + err.Error())
		return 0, err
	}

	return totalAmount, nil
}

// GetCountTotalRevenue implements repo.IRevenueRepo.
func (r *revenueRepo) GetCountTotalRevenue(req request.GetRevenuesRequest, ctx context.Context) (int, error) {
	var table string = entity.Payment{}.GetPaymentTable()
	var errLogMsg string = fmt.Sprintf(noti.REPO_ERR_MSG, table) + "GetRevenues - "
	var queryCondition string = fmt.Sprintf("WHERE tourGuideId = %d", req.TourGuideId)

	if req.Year != nil {
		queryCondition += fmt.Sprintf(" AND YEAR(createdAt) = %d", *req.Year)
	}

	if req.Month != nil {
		queryCondition += fmt.Sprintf(" AND MONTH(createdAt) = %d", *req.Month)
	}

	if req.PaymentStatus != nil {
		var intStatus int = 0
		if *req.PaymentStatus {
			intStatus = 1
		}

		queryCondition += fmt.Sprintf(" AND paymentStatus = %d", intStatus)
	}

	var query string = generateRetrieveQuery(table, queryCondition, 0, 0, true)

	var res int
	if err := r.db.QueryRow(query).Scan(&res); err != nil {
		r.logger.Println(errLogMsg + err.Error())
		return 0, errors.New(noti.INTERNALL_ERR_MSG)
	}

	return res, nil
}

// CreateRevenue implements repo.IRevenueRepo.
func (r *revenueRepo) CreateRevenue(revenue entity.Revenue, ctx context.Context) (int, error) {
	var query string = "INSERT INTO " + revenue.GetRevenueTable() +
		" (paymentId, tourGuideId, invoiceId, totalAmount, " +
		"actualReceived, platformCommission, paymentStatus, createdAt) " +
		"OUTPUT INSERTED.revenueId " +
		"values (@p1, @p2, @p3, @p4, @p5, @p6, @p7, @p8)"
	var errLogMsg string = fmt.Sprintf(noti.REPO_ERR_MSG, revenue.GetRevenueTable()) + "CreateRevenue - "

	var res int
	if err := r.db.QueryRow(query, revenue.PaymentId, revenue.TourGuideId, revenue.InvoiceId, revenue.TotalAmount,
		revenue.ActualReceived, revenue.PlatformCommission, revenue.PaymentStatus, revenue.CreatedAt).Scan(&res); err != nil {

		r.logger.Println(errLogMsg + err.Error())
		return 0, errors.New(noti.INTERNALL_ERR_MSG)
	}

	return res, nil
}

// GetRevenue implements repo.IRevenueRepo.
func (r *revenueRepo) GetRevenue(id int, ctx context.Context) (*entity.Revenue, error) {
	var res entity.Revenue
	var query string = "SELECT * FROM " + res.GetRevenueTable() + " WHERE revenueId = @p1"
	var errLogMsg string = fmt.Sprintf(noti.REPO_ERR_MSG, res.GetRevenueTable()) + "GetRevenue - "

	if err := r.db.QueryRow(query, id).Scan(
		&res.RevenueId, &res.PaymentId, &res.TourGuideId, &res.InvoiceId,
		&res.TotalAmount, &res.ActualReceived, &res.PlatformCommission, &res.PaymentStatus, &res.CreatedAt); err != nil {

		if err == sql.ErrNoRows {
			return nil, nil
		}

		r.logger.Println(errLogMsg + err.Error())
		return nil, errors.New(noti.INTERNALL_ERR_MSG)
	}

	return &res, nil
}

// RemoveRevenue implements repo.IRevenueRepo.
func (r *revenueRepo) RemoveRevenue(id int, ctx context.Context) error {
	var tmp entity.Revenue
	var query string = "DELETE FROM " + tmp.GetRevenueTable() + " WHERE revenueId = @p1"
	var errLogMsg string = fmt.Sprintf(noti.REPO_ERR_MSG, tmp.GetRevenueTable()) + "GetRevenue - "
	var internalErr error = errors.New(noti.INTERNALL_ERR_MSG)

	res, err := r.db.Exec(query, id)
	if err != nil {
		r.logger.Println(errLogMsg + err.Error())
		return internalErr
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		r.logger.Println(errLogMsg + err.Error())
		return internalErr
	}

	if rowsAffected == 0 {
		return errors.New(fmt.Sprintf(noti.UNDEFINED_OBJECT_WARN_MSG, tmp.GetRevenueTable()))
	}

	return nil
}

// UpdateRevenue implements repo.IRevenueRepo.
func (r *revenueRepo) UpdateRevenue(revenue entity.Revenue, ctx context.Context) error {
	panic("unimplemented")
}
