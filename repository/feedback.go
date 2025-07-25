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

type feedbackRepo struct {
	db     *sql.DB
	logger *log.Logger
}

func InitializeFeedbackRepo(db *sql.DB, logger *log.Logger) repo.IFeedbackRepo {
	return &feedbackRepo{
		db:     db,
		logger: logger,
	}
}

// GetFeedbacksDetailByService implements repo.IFeedbackRepo.
func (f *feedbackRepo) GetFeedbacksDetailByService(serviceId int, ctx context.Context) (float64, int, error) {
	var query string = "SELECT AVG(rating), COUNT(*) FROM " + entity.Feedback{}.GetFeedbackTable() + " WHERE service_id = $1"
	var errLogMsg string = fmt.Sprintf(noti.REPO_ERR_MSG, entity.Feedback{}.GetFeedbackTable()) + "GetFeedbacksDetailByService - "

	var avgRating sql.NullFloat64
	var totalCount int

	if err := f.db.QueryRow(query, serviceId).Scan(&avgRating, &totalCount); err != nil {
		if err == sql.ErrNoRows {
			return 0, 0, nil
		}

		f.logger.Println(errLogMsg + err.Error())
		return -1, -1, errors.New(noti.INTERNALL_ERR_MSG)
	}

	var avg float64
	if avgRating.Valid {
		avg = avgRating.Float64
	}

	return avg, totalCount, nil
}

// CreateFeedback implements repo.IFeedbackRepo.
func (f *feedbackRepo) CreateFeedback(feedback entity.Feedback, ctx context.Context) error {
	var query string = "INSERT INTO " + feedback.GetFeedbackTable() +
		" (customer_id, service_id, tour_guide_id, created_date, " +
		"content, rating, is_deleted, " +
		"updated_at, invoice_id) " +
		"values ($1, $2, $3, $4, $5, $6, $7, $8)"
	var errLogMsg string = fmt.Sprintf(noti.REPO_ERR_MSG, feedback.GetFeedbackTable()) + "CreateFeedback - "

	if _, err := f.db.Exec(query, feedback.CustomerId, feedback.ServiceId, feedback.TourGuideId, feedback.CreatedDate,
		feedback.Content, feedback.Rating, feedback.IsDeleted,
		feedback.UpdatedAt, feedback.InvoiceId); err != nil {

		f.logger.Println(errLogMsg + err.Error())
		return errors.New(noti.INTERNALL_ERR_MSG)
	}

	return nil
}

// GetFeedbackById implements repo.IFeedbackRepo.
func (f *feedbackRepo) GetFeedbackById(id int, ctx context.Context) (*entity.Feedback, error) {
	var res entity.Feedback
	var query string = "SELECT * FROM " + res.GetFeedbackTable() + " WHERE feedback_id = $1"
	var errLogMsg string = fmt.Sprintf(noti.REPO_ERR_MSG, res.GetFeedbackTable()) + "GetFeedbackById - "

	if err := f.db.QueryRow(query, id).Scan(
		&res.FeedbackId, &res.CustomerId, &res.ServiceId, &res.TourGuideId, &res.CreatedDate,
		&res.Content, &res.Rating, &res.IsDeleted, &res.UpdatedAt, &res.InvoiceId); err != nil {

		if err == sql.ErrNoRows {
			return nil, nil
		}

		f.logger.Println(errLogMsg + err.Error())
		return nil, errors.New(noti.INTERNALL_ERR_MSG)
	}

	return &res, nil
}

// GetFeedbacks implements repo.IFeedbackRepo.
func (f *feedbackRepo) GetFeedbacks(req request.GetFeedbacksRequest, ctx context.Context) (*[]entity.Feedback, int, int, error) {
	var tmp entity.Feedback
	var table string = tmp.GetFeedbackTable()
	var limitRecords int = tmp.GetFeedbackLimitRecords()

	var errLogMsg string = fmt.Sprintf(noti.REPO_ERR_MSG, table) + "GetFeedbacks - "
	var queryCondition string = fmt.Sprintf(
		" WHERE customer_id = '%d' AND tour_guide_id = '%d' AND invoice_id = '%d' LOWER(content) LIKE LOWER('%%%s%%') AND is_deleted = '%b'",
		req.CustomerId,
		req.TourGuideId,
		req.InvoiceId,
		req.Request.Keyword,
		req.IsDeleted,
	)
	var orderCondition string = generateOrderCondition(req.Request.FilterProp, req.Request.Order)
	var query string = generateRetrieveQuery(table, queryCondition+orderCondition, limitRecords, req.Request.Page, false)

	rows, err := f.db.Query(query)
	if err != nil {
		f.logger.Println(errLogMsg + err.Error())
		return nil, 0, 0, errors.New(noti.INTERNALL_ERR_MSG)
	}

	var res []entity.Feedback
	for rows.Next() {
		var x entity.Feedback
		if err := rows.Scan(
			&x.FeedbackId, &x.CustomerId, &x.ServiceId, &x.TourGuideId, &x.CreatedDate,
			&x.Content, &x.Rating, &x.IsDeleted, &x.UpdatedAt, &x.InvoiceId,
		); err != nil {

			f.logger.Println(errLogMsg + err.Error())
			return nil, 0, 0, errors.New(noti.INTERNALL_ERR_MSG)
		}

		res = append(res, x)
	}

	// Track total records in table
	var totalRecords int
	f.db.QueryRow(generateRetrieveQuery(table, queryCondition, limitRecords, req.Request.Page, true)).Scan(&totalRecords)

	return &res, caculateTotalPages(totalRecords, limitRecords), totalRecords, nil
}

// UpdateFeedback implements repo.IFeedbackRepo.
func (f *feedbackRepo) UpdateFeedback(feedback entity.Feedback, ctx context.Context) error {
	var errLogMsg string = fmt.Sprintf(noti.REPO_ERR_MSG, feedback.GetFeedbackTable()) + "Updatefeedback - "
	var query string = "UPDATE " + feedback.GetFeedbackTable() + " SET content = $1, rating = $2, is_deleted = $3, updated_at = $4 WHERE feedback_id = $5"

	res, err := f.db.Exec(query, feedback.Content, feedback.Rating, feedback.IsDeleted, feedback.UpdatedAt, feedback.FeedbackId)

	var INTERNALL_ERR_MSGMsg error = errors.New(noti.INTERNALL_ERR_MSG)

	if err != nil {
		f.logger.Println(errLogMsg + err.Error())
		return INTERNALL_ERR_MSGMsg
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		f.logger.Println(errLogMsg + err.Error())
		return INTERNALL_ERR_MSGMsg
	}

	if rowsAffected == 0 {
		return errors.New(fmt.Sprintf(noti.UNDEFINED_OBJECT_WARN_MSG, feedback.GetFeedbackTable()))
	}

	return nil
}
