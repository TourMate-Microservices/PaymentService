package businesslogic

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"time"
	domain_status "tourmate/payment-service/constant/domain_status"
	payment_env "tourmate/payment-service/constant/env/payment"
	mail_const "tourmate/payment-service/constant/mail_const"
	"tourmate/payment-service/constant/noti"
	"tourmate/payment-service/infrastructure/grpc/tour"
	tour_pb "tourmate/payment-service/infrastructure/grpc/tour/pb"
	"tourmate/payment-service/infrastructure/grpc/user"
	user_pb "tourmate/payment-service/infrastructure/grpc/user/pb"

	business_logic "tourmate/payment-service/interface/business_logic"
	"tourmate/payment-service/interface/repo"
	"tourmate/payment-service/model/dto/request"
	"tourmate/payment-service/model/dto/response"
	"tourmate/payment-service/model/entity"
	"tourmate/payment-service/repository"
	"tourmate/payment-service/repository/db"
	db_server "tourmate/payment-service/repository/db_server"

	"tourmate/payment-service/utils"

	"github.com/payOSHQ/payos-lib-golang"
)

type paymentService struct {
	logger      *log.Logger
	userService business_logic.IUserService
	tourService business_logic.ITourService
	paymentRepo repo.IPaymentRepo
}

func InitializePaymentService(db *sql.DB, userService business_logic.IUserService, tourService business_logic.ITourService, logger *log.Logger) business_logic.IPaymentService {
	return &paymentService{
		logger:      logger,
		userService: userService,
		tourService: tourService,
		paymentRepo: repository.InitializePaymentRepo(db, logger),
	}
}

func GeneratePaymentService() (business_logic.IPaymentService, error) {
	var logger = utils.GetLogConfig()

	cnn, err := db.ConnectDB(logger, db_server.InitializeMsSQL())

	if err != nil {
		return nil, err
	}

	userService, _ := user.GenerateUserService(logger)
	tourService, _ := tour.GenerateTourService(logger)

	return InitializePaymentService(cnn, userService, tourService, logger), nil
}

// GetPaymentById implements businesslogic.IPaymentService.
func (p *paymentService) GetPaymentById(id int, ctx context.Context) (*entity.Payment, error) {
	return p.paymentRepo.GetPaymentById(id, ctx)
}

// GetPayments implements businesslogic.IPaymentService.
func (p *paymentService) GetPayments(req request.GetPaymentsRequest, ctx context.Context) (response.PaginationDataResponse, error) {
	if req.Request.Page < 1 {
		req.Request.Page = 1
	}

	req.Request.FilterProp = utils.AssignFilterProperty(req.Request.FilterProp)
	req.Request.Order = utils.AssignOrder(req.Request.Order)
	req.PageSize = 10

	p.logger.Println("GetPayments Request: ", req)

	if req.CustomerId != nil {
		user, err := p.userService.GetCustomerById(ctx, &user_pb.GetCustomerByIdRequest{
			CustomerId: int32(*req.CustomerId),
		})

		if err != nil {
			return response.PaginationDataResponse{}, err
		}

		if user == nil {
			return response.PaginationDataResponse{}, errors.New(noti.GENERIC_ERROR_WARN_MSG)
		}
	}

	data, pages, totalRecords, err := p.paymentRepo.GetPayments(req, ctx)

	return response.PaginationDataResponse{
		Data:        data,
		Page:        req.Request.Page,
		TotalPages:  pages,
		TotalCount:  totalRecords,
		PerPage:     req.PageSize,
		HasNext:     req.Request.Page < pages,
		HasPrevious: req.Request.Page > 1,
	}, err
}

// UpdatePayment implements businesslogic.IPaymentService.
func (p *paymentService) UpdatePayment(req request.UpdatePaymentRequest, ctx context.Context) error {
	payment, err := p.paymentRepo.GetPaymentById(req.PaymentId, ctx)
	if err != nil {
		return err
	}

	if payment == nil {
		return errors.New(noti.GENERIC_ERROR_WARN_MSG)
	}

	// Must validate(implement later)
	if req.Method != "" {
		payment.PaymentMethod = req.Method
	}

	return p.paymentRepo.UpdatePayment(*payment, ctx)
}

// CreatePayment implements businesslogic.IPaymentService.
func (p *paymentService) CreatePayment(req request.CreatePaymentRequest, ctx context.Context) (*entity.Payment, error) {
	res, err := p.paymentRepo.CreatePayment(entity.Payment{
		CustomerId:    req.CustomerId,
		InvoiceId:     req.InvoiceId,
		ServiceId:     req.ServiceId,
		Price:         req.Price,
		PaymentMethod: req.PaymentMethod,
		CreatedAt:     time.Now(),
		Status:        domain_status.PAYMENT_PAID,
	}, ctx)

	if err != nil {
		return nil, err
	}

	userInfo, _ := p.userService.GetCustomerById(ctx, &user_pb.GetCustomerByIdRequest{
		CustomerId: int32(req.CustomerId),
	})

	if userInfo != nil {
		utils.SendMail(request.SendMailRequest{
			Body: request.MailBody{ // Mail body
				Subject:       noti.NOTI_PAYMENT_MAIL_SUBJECT,
				Email:         userInfo.Email,
				Username:      userInfo.FullName,
				TransactionId: res.InvoiceId,
			},
			TemplatePath: mail_const.PAYMENT_CALLBACK_CANCEL_TEMPLATE,
			Logger:       p.logger, // Logger
		})
	}

	return res, nil
}

// // CreatePaymentDirect implements businesslogic.IPaymentService.
// func (p *paymentService) CreatePaymentDirect(req request.CreatePaymentDirectRequest, ctx context.Context) (string, error) {
// 	var errRes error = errors.New(noti.GENERIC_ERROR_WARN_MSG)

// 	if !isEntityExist(p.userRepo, req.UserId, id_type, ctx) {
// 		return "", errRes
// 	}

// 	product, err := p.productRepo.GetProductById(req.Product.ProductId, ctx)
// 	if err != nil {
// 		return "", err
// 	}

// 	if product == nil {
// 		return "", errRes
// 	}

// 	inventory, err := p.invetoryRepo.GetProductInventory(product.ProductId, ctx)
// 	if err != nil {
// 		return "", err
// 	}

// 	if inventory == nil {
// 		return "", errRes
// 	}

// 	if inventory.CurrentQuantity < int64(req.Product.Quantity) {
// 		return "", errRes
// 	}

// 	var paymentId string = utils.GenerateId()
// 	var orderCode int = utils.GenerateNumber()
// 	var totalAmount float64 = product.Price * float64(req.Product.Quantity)

// 	// Create transaction url
// 	data, err := payos.CreatePaymentLink(payos.CheckoutRequestType{
// 		OrderCode: int64(orderCode),
// 		Amount:    int(totalAmount),
// 		Items: []payos.Item{
// 			{
// 				Name:     product.ProductName,
// 				Quantity: req.Product.Quantity,
// 				Price:    int(product.Price),
// 			},
// 		},
// 		Description: fmt.Sprint(orderCode),
// 		ReturnUrl:   os.Getenv(payment_env.PAYMENT_CALLBACK_SUCCESS) + paymentId,
// 		CancelUrl:   os.Getenv(payment_env.PAYMENT_CALLBACK_CANCEL) + paymentId,
// 	})

// 	if err != nil {
// 		p.logger.Println(fmt.Sprintf(noti.PAYMENT_GENERATE_TRANSACTION_URL_ERR_MSG, payment_method.PAYOS) + err.Error())
// 		return "", errors.New(noti.INTERNALL_ERR_MSG)
// 	}

// 	inventory.CurrentQuantity -= int64(req.Product.Quantity)

// 	// Update inventory
// 	if err := p.invetoryRepo.UpdateProductInventory(*inventory, ctx); err != nil {
// 		return "", err
// 	}

// 	var orderId string = utils.GenerateId()
// 	var curTime time.Time = time.Now()

// 	// Create order
// 	if err := p.orderRepo.CreateOrder(entity.Order{
// 		OrderId: orderId,
// 		UserId:  req.UserId,
// 		Items: utils.ObjectToJsonString([]response.CartItem{
// 			{
// 				ProductId: req.Product.ProductId,
// 				Name:      product.ProductName,
// 				ImageUrl:  product.Image,
// 				Quantity:  req.Product.Quantity,
// 				Price:     product.Price,
// 				Currency:  product.Currency,
// 			},
// 		}),
// 		TotalAmount: totalAmount,
// 		Currency:    product.Currency,
// 		Status:      domain_status.ORDER_PENDING,
// 		Note:        req.Note,
// 		CreatedAt:   curTime,
// 		UpdatedAt:   curTime,
// 	}, ctx); err != nil {
// 		return "", err
// 	}

// 	// Create payment
// 	if err := p.paymentRepo.CreatePayment(entity.Payment{
// 		PaymentId:     paymentId,
// 		UserId:        req.UserId,
// 		OrderId:       orderId,
// 		TransactionId: orderId,
// 		Amount:        totalAmount,
// 		Currency:      product.Currency,
// 		Status:        domain_status.PAYMENT_PENDING,
// 		Method:        payment_method.PAYOS,
// 		CreatedAt:     curTime,
// 		UpdatedAt:     curTime,
// 	}, ctx); err != nil {
// 		return "", err
// 	}

// 	// Transaction URL
// 	return data.CheckoutUrl, nil
// }

// // CallbackPaymentSuccess implements businesslogic.IPaymentService.
// func (p *paymentService) CallbackPaymentSuccess(component response.PaymentCallbackComponent, ctx context.Context) (string, error) {
// 	if err := p.paymentRepo.CreatePayment(entity.Payment{
// 		Price:         component.Price,
// 		PaymentMethod: component.PaymentMethod,
// 		CustomerId:    component.CustomerId,
// 		AccountId:     component.AccountId,
// 		CreatedAt:     time.Now(),
// 	}, ctx); err != nil {
// 		return "", err
// 	}
// 	// Get user data
// 	user, err := p.userService.GetUser(ctx, &pb.GetUserRequest{
// 		Id: int32(component.CustomerId),
// 	})

// 	if err != nil {
// 		return "", err
// 	}

// 	// Send mail
// 	utils.SendMail(request.SendMailRequest{
// 		Body: request.MailBody{ // Mail body
// 			Subject:       noti.NOTI_PAYMENT_MAIL_SUBJECT,
// 			Email:         user.Email,
// 			Username:      user.Fullname,
// 			TransactionId: component.OrderCode,
// 		},

// 		TemplatePath: mail_const.PAYMENT_CALLBACK_SUCCESS_TEMPLATE,

// 		Logger: p.logger, // Logger
// 	})

// 	return "url-to-process-payment-page", nil
// }

// // CallbackPaymentCancel implements businesslogic.IPaymentService.
// func (p *paymentService) CallbackPaymentCancel(component response.PaymentCallbackComponent, ctx context.Context) (string, error) {
// 	// Get user data
// 	user, err := p.userService.GetUser(ctx, &pb.GetUserRequest{
// 		Id: int32(component.CustomerId),
// 	})

// 	if err != nil {
// 		return "", err
// 	}

// 	// Send mail
// 	utils.SendMail(request.SendMailRequest{
// 		Body: request.MailBody{ // Mail body
// 			Subject:       noti.NOTI_PAYMENT_MAIL_SUBJECT,
// 			Email:         user.Email,
// 			Username:      user.Fullname,
// 			TransactionId: component.OrderCode,
// 		},

// 		TemplatePath: mail_const.PAYMENT_CALLBACK_CANCEL_TEMPLATE,

// 		Logger: p.logger, // Logger
// 	})

// 	return "url-to-process-payment-page", nil
// }

// func generateCallbackUrl(data response.PaymentCallbackComponent, domainUrl string) string {
// 	return fmt.Sprintf(
// 		"%s?customerId=%d&accountId=%d&paymentMethod=%s&price=%.2f&orderCode=%d",
// 		domainUrl,
// 		data.CustomerId,
// 		data.AccountId,
// 		data.PaymentMethod,
// 		data.Price,
// 		data.OrderCode,
// 	)
// }

// CreatePayosTransaction implements businesslogic.IPaymentService.
func (p *paymentService) CreatePayosTransaction(req request.CreatePayosTransactionRequest, ctx context.Context) (response.UrlResponse, error) {
	var description string = fmt.Sprintf("Invoice %d", req.InvoiceId)
	p.logger.Println("Description: ", description)
	p.logger.Printf("Request data - Amount: %f, InvoiceId: %d", req.Amount, req.InvoiceId)

	// Validate input data
	if req.Amount <= 0 {
		p.logger.Println("Invalid amount: amount must be greater than 0")
		return response.UrlResponse{}, errors.New("amount must be greater than 0")
	}

	if req.InvoiceId <= 0 {
		p.logger.Println("Invalid invoice ID: invoice ID must be greater than 0")
		return response.UrlResponse{}, errors.New("invoice ID must be greater than 0")
	}

	// Convert amount to integer (PayOS expects amount in VND, not cents for VN)
	amount := int(req.Amount)

	// Generate unique order code
	orderCode := int64(utils.GenerateNumber())
	p.logger.Printf("Generated OrderCode: %d", orderCode)

	returnUrl := os.Getenv(payment_env.PAYMENT_CALLBACK_SUCCESS)
	cancelUrl := os.Getenv(payment_env.PAYMENT_CALLBACK_CANCEL)
	p.logger.Printf("PayOS Request: Amount=%d, OrderCode=%d, Description=%s, ReturnUrl=%s, CancelUrl=%s", amount, orderCode, description, returnUrl, cancelUrl)
	p.logger.Printf("PayOS Items: %+v", []payos.Item{{Name: description, Quantity: 1, Price: amount}})
	data, err := payos.CreatePaymentLink(payos.CheckoutRequestType{
		Amount:    amount,
		OrderCode: orderCode,
		Items: []payos.Item{
			{
				Name:     description,
				Quantity: 1,
				Price:    amount,
			},
		},
		Description: description,
		ReturnUrl:   returnUrl,
		CancelUrl:   cancelUrl,
	})

	if err != nil {
		p.logger.Printf("Failed to create PayOS link: %v", err)
		return response.UrlResponse{}, fmt.Errorf("failed to create payment link: %v", err)
	}

	p.logger.Println("Payos link: ", data.CheckoutUrl)

	return response.UrlResponse{
		Url: data.CheckoutUrl,
	}, nil
}

// GetPaymentWithService implements businesslogic.IPaymentService.
func (p *paymentService) GetPaymentWithService(id int, ctx context.Context) (*response.PaymentWithServiceNameResponse, error) {
	payment, err := p.paymentRepo.GetPaymentById(id, ctx)
	if err != nil {
		return nil, err
	}

	if payment == nil {
		return nil, errors.New(noti.GENERIC_ERROR_WARN_MSG)
	}

	serviceInfo, err := p.tourService.GetTourById(ctx, &tour_pb.TourServiceIdRequest{
		ServiceId: int32(payment.ServiceId),
	})

	if err != nil {
		return nil, err
	}

	if serviceInfo == nil {
		return nil, errors.New(noti.GENERIC_ERROR_WARN_MSG)
	}

	return &response.PaymentWithServiceNameResponse{
		PaymentId:   payment.PaymentId,
		Price:       payment.Price,
		ServiceId:   payment.ServiceId,
		ServiceName: serviceInfo.ServiceName,
		CreatedAt:   payment.CreatedAt,
	}, nil
}
