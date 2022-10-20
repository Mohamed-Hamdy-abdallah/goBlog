package models

type Comment struct {
	Id      uint   `Json:"id"`
	Content string `Json:"content"`
	UserID  string `Json:"userid"`
	BlogID  string `Json:"blogid"`
	User    User   `Json:"user";"gorm":"foreignkey:UserID"`
	Blog    Blog   `Json:"user";"gorm":"foreignkey:BlogID"`
}
