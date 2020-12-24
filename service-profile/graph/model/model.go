package model

type Profile struct {
	UserID 	string `json:"user"`
	Age   	string `json:"age"`
	Phone 	string `json:"phone"`
	Job   	string `json:"job"`
}

type User struct {
	ID      string   `json:"id"`
}

func (User) IsEntity() {}
