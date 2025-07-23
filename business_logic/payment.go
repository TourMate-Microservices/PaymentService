package businesslogic

import (
	"context"
	"database/sql"
	"errors"
	"log"
	domain_status "tourmate/payment-service/constant/domain_status"
	mail_const "tourmate/payment-service/constant/mail_const"
	"tourmate/payment-service/constant/noti"
	"tourmate/payment-service/infrastructure/grpc/user"
	"tourmate/payment-service/infrastructure/grpc/user/pb"
	business_logic "tourmate/payment-service/interface/business_logic"
	"tourmate/payment-service/interface/repo"
	"tourmate/payment-service/model/dto/request"
	"tourmate/payment-service/model/dto/response"
	"tourmate/payment-service/model/entity"
	"tourmate/payment-service/repository"
	"tourmate/payment-service/repository/db"
	db_server "tourmate/payment-service/repository/db_server"
	"tourmate/payment-service/utils"
)

type paymentService struct {
	logger      *log.Logger
	userService business_logic.IUserService
	paymentRepo repo.IPaymentRepo
}

func InitializePaymentService(db *sql.DB, userService business_logic.IUserService, logger *log.Logger) business_logic.IPaymentService {
	return &paymentService{
		logger:      logger,
		userService: userService,
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

	return InitializePaymentService(cnn, userService, logger), nil
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

	if req.CustomerId != nil {
		user, err := p.userService.GetUser(pb.GetUserRequest{
			Id: int32(*req.CustomerId),
		}, ctx)

		if err != nil {
			return response.PaginationDataResponse{}, err
		}

		if user == nil {
			return response.PaginationDataResponse{}, errors.New(noti.GENERIC_ERROR_WARN_MSG)
		}
	}

	data, pages, err := p.paymentRepo.GetPayments(req, ctx)

	return response.PaginationDataResponse{
		Data:       data,
		Page:       req.Request.Page,
		TotalPages: pages,
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
	if req.Status != "" {
		payment.Status = req.Status
	}

	// Must validate(implement later)
	if req.Method != "" {
		payment.PaymentMethod = req.Method
	}

	return p.paymentRepo.UpdatePayment(*payment, ctx)
}

// // CreatePayment implements businesslogic.IPaymentService.
// func (p *paymentService) CreatePaymentThroughCart(req request.CreatePaymentThroughCartRequest, ctx context.Context) (string, error) {
// 	var errRes error = errors.New(noti.GENERIC_ERROR_WARN_MSG)

// 	if !isEntityExist(p.userRepo, req.UserId, id_type, ctx) {
// 		return "", errRes
// 	}

// 	cart, _, err := p.cartRepo.GetCart(req.UserId, ctx)
// 	if err != nil {
// 		return "", err
// 	}

// 	// No items in cart but execute payment
// 	if cart == nil {
// 		return "", errRes
// 	}

// 	var itemsInCart []response.CartItem = utils.JsonStringToObject[[]response.CartItem](cart.Items)
// 	var totalAmount float64
// 	var items []payos.Item
// 	var invetories []entity.ProductInventory
// 	var formatItems []response.CartItem

// 	for _, prod := range req.Items {
// 		// Item not existed in cart
// 		if !strings.Contains(cart.Items, prod.ProductId) {
// 			return "", errRes
// 		}

// 		inventory, err := p.invetoryRepo.GetProductInventory(prod.ProductId, ctx)
// 		if err != nil {
// 			return "", err
// 		}

// 		if inventory == nil {
// 			return "", errRes
// 		}

// 		product, err := p.productRepo.GetProductById(prod.ProductId, ctx)
// 		if err != nil {
// 			return "", err
// 		}

// 		if product == nil {
// 			return "", errRes
// 		}

// 		if inventory.CurrentQuantity < int64(prod.Quantity) {
// 			return "", errRes
// 		}

// 		inventory.CurrentQuantity -= int64(prod.Quantity)
// 		totalAmount += float64(prod.Quantity) * product.Price
// 		invetories = append(invetories, *inventory)

// 		items = append(items, payos.Item{
// 			Name:     product.ProductName,
// 			Quantity: prod.Quantity,
// 			Price:    int(product.Price),
// 		})

// 		formatItems = append(formatItems, response.CartItem{
// 			ProductId: prod.ProductId,
// 			Name:      product.ProductName,
// 			ImageUrl:  product.Image,
// 			Quantity:  prod.Quantity,
// 			Price:     product.Price,
// 			Currency:  product.Currency,
// 		})

// 		for index, item := range itemsInCart {
// 			if item.ProductId == prod.ProductId {
// 				item.Quantity -= prod.Quantity
// 				if item.Quantity <= 0 { // Remove item from cart
// 					itemsInCart = append(itemsInCart[:index], itemsInCart[index+1:]...)
// 				}

// 				break
// 			}
// 		}
// 	}

// 	var paymentId string = utils.GenerateId()
// 	var orderCode int = utils.GenerateNumber()

// 	// Create transaction url
// 	data, err := payos.CreatePaymentLink(payos.CheckoutRequestType{
// 		OrderCode:   int64(orderCode),
// 		Amount:      int(totalAmount),
// 		Items:       items,
// 		Description: fmt.Sprint(orderCode),
// 		ReturnUrl:   os.Getenv(payment_env.PAYMENT_CALLBACK_SUCCESS) + paymentId,
// 		CancelUrl:   os.Getenv(payment_env.PAYMENT_CALLBACK_CANCEL) + paymentId,
// 	})

// 	if err != nil {
// 		p.logger.Println("Err: ", err.Error())
// 		return "", errors.New(noti.INTERNALL_ERR_MSG)
// 	}

// 	// Update product inventories
// 	for _, inventory := range invetories {
// 		if err := p.invetoryRepo.UpdateProductInventory(inventory, ctx); err != nil {
// 			return "", err
// 		}
// 	}

// 	var curTime time.Time = time.Now()

// 	// No more items in cart
// 	if len(itemsInCart) == 0 {
// 		p.cartRepo.RemoveCart(req.UserId, ctx)
// 	} else {
// 		cart.UpdatedAt = curTime
// 		cart.ExpiredAt = curTime.AddDate(0, 0, 7)
// 		p.cartRepo.UpdateCart(*cart, ctx)
// 	}

// 	var orderId string = utils.GenerateId()

// 	// Create order
// 	if err := p.orderRepo.CreateOrder(entity.Order{
// 		OrderId:     orderId,
// 		UserId:      req.UserId,
// 		Items:       utils.ObjectToJsonString(formatItems),
// 		TotalAmount: totalAmount,
// 		Currency:    currency.VIETNAM_DONG,
// 		Note:        req.Note,
// 		Status:      domain_status.ORDER_PENDING,
// 		CreatedAt:   curTime,
// 		UpdatedAt:   curTime,
// 	}, ctx); err != nil {
// 		return "", err
// 	}

// 	// Create payment
// 	if err := p.paymentRepo.CreatePayment(entity.Payment{
// 		PaymentId:     paymentId,
// 		OrderId:       orderId,
// 		UserId:        req.UserId,
// 		TransactionId: fmt.Sprint(orderCode),
// 		Amount:        totalAmount,
// 		Currency:      currency.VIETNAM_DONG,
// 		Status:        domain_status.PAYMENT_PENDING,
// 		Method:        payment_method.PAYOS,
// 		CreatedAt:     curTime,
// 		UpdatedAt:     curTime,
// 	}, ctx); err != nil {
// 		return "", err
// 	}

// 	return data.CheckoutUrl, nil
// }

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

// CallbackPaymentSuccess implements businesslogic.IPaymentService.
func (p *paymentService) CallbackPaymentSuccess(id int, ctx context.Context) (string, error) {
	var payment *entity.Payment
	var capturedErr error

	// Get payment
	for i := 1; i <= 3; i++ {
		payment, capturedErr = p.paymentRepo.GetPaymentById(id, ctx)
		if capturedErr == nil {
			break
		}
	}

	if payment == nil {
		return "", errors.New(noti.GENERIC_ERROR_WARN_MSG)
	}

	if capturedErr != nil {
		return "", capturedErr
	}

	payment.Status = domain_status.PAYMENT_PAID
	// order.Status = domain_status.ORDER_COMPLETED

	// var curTime time.Time = time.Now()

	// // Update payment
	// payment.UpdatedAt = curTime
	if err := p.paymentRepo.UpdatePayment(*payment, ctx); err != nil {
		return "", err
	}

	// // Update order
	// order.UpdatedAt = curTime
	// if err := p.orderRepo.UpdateOrder(*order, ctx); err != nil {
	// 	return "", err
	// }

	// Get user data
	user, err := p.userService.GetUser(pb.GetUserRequest{
		Id: int32(payment.CustomerId),
	}, ctx)

	if err != nil {
		return "", err
	}

	// // Create ship
	// p.shippingRepo.CreateShipping(entity.Shipping{
	// 	OrderId: order.OrderId,
	// 	ShippingDetail: utils.ObjectToJsonString(response.ShippingDetail{
	// 		RecipientName: fullName,
	// 	}),
	// 	CreatedAt: curTime,
	// 	UpdatedAt: curTime,
	// }, ctx)

	// Send mail
	utils.SendMail(request.SendMailRequest{
		Body: request.MailBody{ // Mail body
			Subject:       noti.NOTI_PAYMENT_MAIL_SUBJECT,
			Email:         user.Email,
			Username:      user.Fullname,
			TransactionId: id,
		},

		TemplatePath: mail_const.PAYMENT_CALLBACK_SUCCESS_TEMPLATE,

		Logger: p.logger, // Logger
	})

	return "url-to-process-payment-page", nil
}

// CallbackPaymentCancel implements businesslogic.IPaymentService.
func (p *paymentService) CallbackPaymentCancel(id int, ctx context.Context) (string, error) {
	var payment *entity.Payment
	var capturedErr error

	// Get payment
	for i := 1; i <= 3; i++ {
		payment, capturedErr = p.paymentRepo.GetPaymentById(id, ctx)
		if capturedErr == nil {
			break
		}
	}

	if payment == nil {
		return "", errors.New(noti.GENERIC_ERROR_WARN_MSG)
	}

	// // Get order
	// for i := 1; i <= 3; i++ {
	// 	order, capturedErr = p.orderRepo.GetOrder(payment.OrderId, ctx)
	// 	if capturedErr == nil {
	// 		break
	// 	}
	// }

	// if capturedErr != nil {
	// 	return "", capturedErr
	// }

	payment.Status = domain_status.PAYMENT_CANCELLED
	// order.Status = domain_status.ORDER_CANCELLED

	// var curTime time.Time = time.Now()

	// // Update payment
	// payment.UpdatedAt = curTime
	if err := p.paymentRepo.UpdatePayment(*payment, ctx); err != nil {
		return "", err
	}

	// // Update order
	// order.UpdatedAt = curTime
	// if err := p.orderRepo.UpdateOrder(*order, ctx); err != nil {
	// 	return "", err
	// }

	// // Refund product amount
	// for _, item := range utils.JsonStringToObject[[]response.CartItem](order.Items) {
	// 	inventory, _ := p.invetoryRepo.GetProductInventory(item.ProductId, ctx)
	// 	if inventory != nil {
	// 		inventory.CurrentQuantity += int64(item.Quantity)
	// 		p.invetoryRepo.UpdateProductInventory(*inventory, ctx)
	// 	}
	// }

	// Get user data
	user, err := p.userService.GetUser(pb.GetUserRequest{
		Id: int32(payment.CustomerId),
	}, ctx)

	if err != nil {
		return "", err
	}

	// Send mail
	utils.SendMail(request.SendMailRequest{
		Body: request.MailBody{ // Mail body
			Subject:       noti.NOTI_PAYMENT_MAIL_SUBJECT,
			Email:         user.Email,
			Username:      user.Fullname,
			TransactionId: id,
		},

		TemplatePath: mail_const.PAYMENT_CALLBACK_CANCEL_TEMPLATE,

		Logger: p.logger, // Logger
	})

	return "url-to-process-payment-page", nil
}
