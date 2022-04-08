package types

type User struct {
	ID   int64  `gorm:"primarykey" json:"id"`
	Name string `json:"name"`

	Screeches []Screech `json:"screeches"`
}

type Screech struct {
	ID int64 `gorm:"primarykey" json:"id"`

	UserID int64 `json:"user_id"`
	User   User

	Content string `json:"content"`
}
