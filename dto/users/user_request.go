package usersdto

type CreateUserRequest struct {
	ID         int    `json:"id"`
	Fullname   string `json:"fullname" gorm:"type: varchar(255)"`
	Email      string `json:"email" gorm:"type: varchar(255)"`
	Password   string `json:"password" gorm:"type: varchar(255)"`
	Username   string `json:"username" gorm:"type: varchar(255)"`
	ListAsRole string `json:"listAsRole" gorm:"type: varchar(225)"`
	Gender     string `json:"gender" gorm:"type: varchar(225)"`
	Phone      string `json:"phone" gorm:"type: varchar(225)"`
	Address    string `json:"address" gorm:"type: varchar(225)"`
}

type UpdateUserRequest struct {
	// Fullname  string    `json:"fullname" form:"fullname"`
	// Username  string    `json:"username" form:"username"`
	// Email     string    `json:"email" form:"email"`
	// Password  string    `json:"password" form:"password"`
	// RoleID    int       `json:"roleid" form:"listasid"`
	// Gender    string    `json:"gender" form:"gender"`
	// Phone     string    `json:"phone" form:"phone"`
	// Address   string    `json:"address" form:"address"`
	Image string `json:"image" form:"image"`
	// CreatedAt time.Time `json:"created_at"`
	// UpdatedAt time.Time `json:"updated_at"`
}
