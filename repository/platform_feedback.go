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

type platformFeedbackRepo struct {
	db     *sql.DB
	logger *log.Logger
}

func InitializePlatformFeedbackRepo(db *sql.DB, logger *log.Logger) repo.IPlatformFeedbackRepo {
	return &platformFeedbackRepo{
		db:     db,
		logger: logger,
	}
}

// CreatePlatformFeedback implements repo.IPlatformFeedbackRepo.
func (p *platformFeedbackRepo) CreatePlatformFeedback(platformFeedback entity.PlatformFeedback, ctx context.Context) error {
	var query string = "INSERT INTO " + platformFeedback.GetPlatformFeedbackTable() +
		" (customerId, paymentId, rating, " +
		"content, createdAt) " +
		"values (@p1, @p2, @p3, @p4, @p5)"
	var errLogMsg string = fmt.Sprintf(noti.REPO_ERR_MSG, platformFeedback.GetPlatformFeedbackTable()) + "CreatePlatformFeedback - "

	if _, err := p.db.Exec(query, platformFeedback.CustomerId, platformFeedback.PaymentId,
		platformFeedback.Rating, platformFeedback.Content, platformFeedback.CreatedAt); err != nil {

		p.logger.Println(errLogMsg + err.Error())
		return errors.New(noti.INTERNALL_ERR_MSG)
	}

	return nil
}

// GetplatformFeedbackById implements repo.IplatformFeedbackRepo.
func (p *platformFeedbackRepo) GetPlatformFeedbackById(id int, ctx context.Context) (*entity.PlatformFeedback, error) {
	var res entity.PlatformFeedback
	var query string = "SELECT * FROM " + res.GetPlatformFeedbackTable() + " WHERE feedbackId = @p1"
	var errLogMsg string = fmt.Sprintf(noti.REPO_ERR_MSG, res.GetPlatformFeedbackTable()) + "GetplatformFeedbackById - "

	if err := p.db.QueryRow(query, id).Scan(
		&res.FeedbackId, &res.CustomerId, &res.PaymentId,
		&res.Rating, &res.Content, &res.CreatedAt); err != nil {

		if err == sql.ErrNoRows {
			return nil, nil
		}

		p.logger.Println(errLogMsg + err.Error())
		return nil, errors.New(noti.INTERNALL_ERR_MSG)
	}

	return &res, nil
}

// GetplatformFeedbacks implements repo.IPlatformFeedbackRepo.
func (p *platformFeedbackRepo) GetPlatformFeedbacks(req request.GetPlatformFeedbacksRequest, ctx context.Context) (*[]entity.PlatformFeedback, int, int, error) {
	var table string = entity.PlatformFeedback{}.GetPlatformFeedbackTable()
	var limitRecords int = *req.PageSize

	var errLogMsg string = fmt.Sprintf(noti.REPO_ERR_MSG, table) + "GetplatformFeedbacks - "
	var queryCondition string = "WHERE "
	var isHavePreviousCond bool = false
	if req.CustomerId != nil {
		queryCondition += fmt.Sprintf("customerId = %d", *req.CustomerId)
		isHavePreviousCond = true
	}
	if req.Rating != nil {
		if isHavePreviousCond {
			queryCondition += " AND "
		}

		queryCondition += fmt.Sprintf("rating = %d", *req.Rating)
	}

	if queryCondition == "WHERE " {
		queryCondition = ""
	}

	var orderCondition string = generateOrderCondition(req.Request.FilterProp, req.Request.Order)
	var query string = generateRetrieveQuery(table, queryCondition+orderCondition, limitRecords, *req.PageIndex, false)

	rows, err := p.db.Query(query)
	if err != nil {
		p.logger.Println(errLogMsg + err.Error())
		return nil, 0, 0, errors.New(noti.INTERNALL_ERR_MSG)
	}

	var res []entity.PlatformFeedback
	for rows.Next() {
		var x entity.PlatformFeedback
		if err := rows.Scan(
			&x.FeedbackId, &x.CustomerId, &x.PaymentId,
			&x.Rating, &x.Content, &x.CreatedAt,
		); err != nil {

			p.logger.Println(errLogMsg + err.Error())
			return nil, 0, 0, errors.New(noti.INTERNALL_ERR_MSG)
		}

		res = append(res, x)
	}

	// Track total records in table
	var totalRecords int
	p.db.QueryRow(generateRetrieveQuery(table, queryCondition, limitRecords, *req.PageIndex, true)).Scan(&totalRecords)

	return &res, caculateTotalPages(totalRecords, limitRecords), totalRecords, nil
}

// UpdatePlatformFeedback implements repo.IPlatformFeedbackRepo.
func (p *platformFeedbackRepo) UpdatePlatformFeedback(platformFeedback entity.PlatformFeedback, ctx context.Context) error {
	var errLogMsg string = fmt.Sprintf(noti.REPO_ERR_MSG, platformFeedback.GetPlatformFeedbackTable()) + "UpdatePlatformFeedback - "
	var query string = "UPDATE " + platformFeedback.GetPlatformFeedbackTable() + " SET content = @p1, rating = @p2 WHERE feedbackId = @p3"

	res, err := p.db.Exec(query, platformFeedback.Content, platformFeedback.Rating, platformFeedback.FeedbackId)

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
		return errors.New(fmt.Sprintf(noti.UNDEFINED_OBJECT_WARN_MSG, platformFeedback.GetPlatformFeedbackTable()))
	}

	return nil
}
