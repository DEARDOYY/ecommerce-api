package response

import "github.com/gin-gonic/gin"

const (
	ErrCodeInvalidRequest = "INVALID_REQUEST"  // request ผิดรูปแบบ เช่น JSON พัง หรือ param ไม่ถูกต้อง
	ErrCodeValidation     = "VALIDATION_ERROR" // ข้อมูลไม่ผ่านเงื่อนไข เช่น ขาด field บังคับ หรือค่าไม่อยู่ในช่วงที่กำหนด
	ErrCodeUnauthorized   = "UNAUTHORIZED"     // ยังไม่ได้ login หรือ token ไม่ถูกต้อง/หมดอายุ
	ErrCodeConflict       = "CONFLICT"         // ข้อมูลซ้ำหรือขัดแย้งกับข้อมูลที่มีอยู่ เช่น email ซ้ำ
	ErrCodeInternal       = "INTERNAL_ERROR"   // ข้อผิดพลาดภายในระบบ ที่ไม่ได้เกิดจากฝั่ง client
	ErrCodeNotFound       = "NOT_FOUND"        // ไม่พบข้อมูลที่ร้องขอ เช่น product หรือ order id ที่ไม่มีอยู่
	ErrCodeForbidden      = "FORBIDDEN"        // login แล้วแต่ไม่มีสิทธิ์ทำรายการนี้ เช่น customer เรียก API ของ admin
	ErrCodeOutOfStock     = "OUT_OF_STOCK"     // checkout/add to cart แต่สต็อกไม่พอ
	ErrCodeEmptyCart      = "EMPTY_CART"       // checkout ด้วย cart ว่าง
	ErrCodeInvalidStatus  = "INVALID_STATUS"   // เปลี่ยนสถานะ order ผิดลำดับ เช่น cancel order ที่ส่งแล้ว
)

func Error(c *gin.Context, status int, code, message string) {
	c.JSON(status, gin.H{
		"error": gin.H{
			"code":    code,
			"message": message,
		},
	})
}

func Success(c *gin.Context, status int, data any) {
	c.JSON(status, gin.H{"data": data})
}

func Message(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{"message": message})
}

func Paginated(c *gin.Context, status int, data any, page, limit, total int64) {
	c.JSON(status, gin.H{
		"meta": gin.H{"page": page, "limit": limit, "total": total},
		"data": data,
	})
}
