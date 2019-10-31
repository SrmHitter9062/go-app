package user


type User struct {
	Id int `json:user_id`
	Name  string `json:"name"`
	Phone string `json:phone`
	Email string `json:email`
	Address string `json:address`
}

