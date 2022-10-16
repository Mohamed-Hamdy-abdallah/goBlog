package models

type Blog struct {
	Id     uint   `Json:"id"`
	Title  string `Json:"title"`
	Desc   string `Json:"desc"`
	Image  string `Json:"image"`
	UserID string `Json:"userid"`
	User   User   `Json:"user";"gorm":"foreignkey:UserID"`
}
