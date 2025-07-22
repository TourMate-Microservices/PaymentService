package utils

import domain_status "tourmate/payment-service/constant/domain_status"

func IsOrderStatusValid(status string) bool {
	var res bool = true

	switch status {
	case domain_status.ORDER_PENDING:
	case domain_status.ORDER_CONFIRMED:
	case domain_status.ORDER_PROCESSING:
	case domain_status.ORDER_SHIPPED:
	case domain_status.ORDER_DELIVERED:
	case domain_status.ORDER_COMPLETED:
	case domain_status.ORDER_CANCELLED:
	case domain_status.ORDER_FAILED:
	case domain_status.ORDER_REFUNDED:
	case domain_status.ORDER_RETURNED:
	default:
		res = false
	}

	return res
}

func IsPaymentStatusValid(status string) bool {
	var res bool = true

	switch status {
	case domain_status.PAYMENT_INITIATED:
	case domain_status.PAYMENT_PENDING:
	case domain_status.PAYMENT_AUTHORIZED:
	case domain_status.PAYMENT_CAPTURED:
	case domain_status.PAYMENT_PAID:
	case domain_status.PAYMENT_FAILED:
	case domain_status.PAYMENT_CANCELLED:
	case domain_status.PAYMENT_REFUNDED:
	case domain_status.PAYMENT_CHARGEBACK:
	case domain_status.PAYMENT_EXPIRED:
	default:
		res = false
	}

	return res
}
