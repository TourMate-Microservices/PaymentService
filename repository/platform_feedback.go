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
		"(AccountId, PaymentId, Rating, " +
		"Content, CreatedAt) " +
		"values ($1, $2, $3, $4, $5)"
	var errLogMsg string = fmt.Sprintf(noti.REPO_ERR_MSG, platformFeedback.GetPlatformFeedbackTable()) + "CreatePlatformFeedback - "

	if _, err := p.db.Exec(query, platformFeedback.AccountId, platformFeedback.PaymentId,
		platformFeedback.Rating, platformFeedback.Content, platformFeedback.CreatedAt); err != nil {

		p.logger.Println(errLogMsg + err.Error())
		return errors.New(noti.INTERNALL_ERR_MSG)
	}

	return nil
}

// GetplatformFeedbackById implements repo.IplatformFeedbackRepo.
func (p *platformFeedbackRepo) GetPlatformFeedbackById(id int, ctx context.Context) (*entity.PlatformFeedback, error) {
	var res entity.PlatformFeedback
	var query string = "SELECT * FROM " + res.GetPlatformFeedbackTable() + " WHERE FeedbackId = $1"
	var errLogMsg string = fmt.Sprintf(noti.REPO_ERR_MSG, res.GetPlatformFeedbackTable()) + "GetplatformFeedbackById - "

	if err := p.db.QueryRow(query, id).Scan(
		&res.FeedbackId, &res.AccountId, &res.PaymentId,
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
	var res []entity.PlatformFeedback
	var table string = res[0].GetPlatformFeedbackTable()
	var limitRecords int = res[0].GetPlatformFeedbackLimitRecords()

	var errLogMsg string = fmt.Sprintf(noti.REPO_ERR_MSG, table) + "GetplatformFeedbacks - "
	var queryCondition string = fmt.Sprintf(
		" WHERE AccountId = '%d' AND Rating = '%d' LOWER(content) LIKE LOWER('%%%s%%')",
		req.AccountId,
		req.Rating,
		req.Request.Keyword,
	)
	var orderCondition string = generateOrderCondition(req.Request.FilterProp, req.Request.Order)
	var query string = generateRetrieveQuery(table, queryCondition+orderCondition, limitRecords, req.Request.Page, false)

	rows, err := p.db.Query(query)
	if err != nil {
		p.logger.Println(errLogMsg + err.Error())
		return nil, 0, 0, errors.New(noti.INTERNALL_ERR_MSG)
	}

	for rows.Next() {
		var x entity.PlatformFeedback
		if err := rows.Scan(
			&x.FeedbackId, &x.AccountId, &x.PaymentId,
			&x.Rating, &x.Content, &x.CreatedAt,
		); err != nil {

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

// UpdatePlatformFeedback implements repo.IPlatformFeedbackRepo.
func (p *platformFeedbackRepo) UpdatePlatformFeedback(platformFeedback entity.PlatformFeedback, ctx context.Context) error {
	var errLogMsg string = fmt.Sprintf(noti.REPO_ERR_MSG, platformFeedback.GetPlatformFeedbackTable()) + "UpdatePlatformFeedback - "
	var query string = "UPDATE " + platformFeedback.GetPlatformFeedbackTable() + " SET content = $1, rating = $2 WHERE FeedbackId = $3"

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
