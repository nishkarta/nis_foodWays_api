package productdto

type ProductRequest struct {
	Title  string `json:"title" form:"name" gorm:"type: varchar(255)" validate:"required"`
	Image  string `json:"image" form:"name" gorm:"type: varchar(255)"`
	Price  int    `json:"price" form:"name" gorm:"type: varchar(255)" validate:"required"`
	Qty    int    `json:"qty" form:"qty" gorm:"type: int" validate:"required"`
	UserID int    `json:"user_id"`
}
