package domainstatus

const (
	ORDER_PENDING    string = "PENDING"    // ĐƠN HÀNG VỪA ĐƯỢC TẠO
	ORDER_CONFIRMED  string = "CONFIRMED"  // ĐÃ XÁC NHẬN ĐƠN HÀNG
	ORDER_PROCESSING string = "PROCESSING" // ĐƠN ĐANG ĐƯỢC XỬ LÝ
	ORDER_SHIPPED    string = "SHIPPED"    // ĐÃ GỬI VẬN CHUYỂN
	ORDER_DELIVERED  string = "DELIVERED"  // ĐÃ GIAO HÀNG
	ORDER_COMPLETED  string = "COMPLETED"  // ĐƠN HOÀN TẤT
	ORDER_CANCELLED  string = "CANCELLED"  // BỊ HỦY BỞI KHÁCH HÀNG HOẶC HỆ THỐNG
	ORDER_FAILED     string = "FAILED"     // LỖI XỬ LÝ ĐƠN
	ORDER_REFUNDED   string = "REFUNDED"   // ĐÃ HOÀN TIỀN
	ORDER_RETURNED   string = "RETURNED"   // KHÁCH TRẢ HÀNG
)
