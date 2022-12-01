package models

// gorm berfungsi untuk mengatur tipe data atau custom tipe data

type User struct {
	ID       int    `json:"id" gorm:"User_Id"`
	FullName string `json:"name"`
	Email    string `json:"email" gorm:"type: varchar(255)"`
	Password string `json:"password" gorm:"type: varchar(255)"`
	Npp      string `json:"npp" gorm:"type: varchar(255)"`
	NppSup   string `json:"nppsup" gorm:"type: varchar(255)"`
}

// berfungsi untuk relasi

type UserResponse struct {
	ID       int    `json:"id" gorm:"User_Id"`
	FullName string `json:"name"`
	Email    string `json:"email" gorm:"type: varchar(255)"`
	Npp      string `json:"npp" gorm:"type: varchar(255)"`
	NppSup   string `json:"nppsup" gorm:"type: varchar(255)"`
	Addres   string `json:"addres"`
}

func (UserResponse) TableName() string {
	return "users"
}
