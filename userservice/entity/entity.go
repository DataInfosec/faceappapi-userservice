package entity

type User struct {
	Type            string `json:"type"  binding:"required"`
	Firstname       string `json:"firstname" binding:"required"`
	Lastname        string `json:"lastname"  binding:"required"`
	Email           string `json:"email"  binding:"required"`
	Password        string `json:"password"  binding:"required"`
	Confirmpassword string `json:"confirmpassword"  binding:"required"`
}

type UserDB struct {
	ID        string `bson:"_id, omitempty"`
	Type      string `json:"type"  binding:"required"`
	Firstname string `json:"firstname" binding:"required"`
	Lastname  string `json:"lastname"  binding:"required"`
	Email     string `json:"email"  binding:"required"`
}

type UpdateUser struct {
	ID        string `json:"_id" binding:"required"`
	Type      string `json:"type"  binding:"required"`
	Firstname string `json:"firstname" binding:"required"`
	Lastname  string `json:"lastname"  binding:"required"`
}

type UserEmail struct {
	Email string `json:"email"  binding:"required"`
}

type Response struct {
	Data    UserDB `json: "data"`
	Message string `json:"message"`
}

type ResponseUpdate struct {
	Data    UpdateUser `json: "data"`
	Message string `json:"message"`
}