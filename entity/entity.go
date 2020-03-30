package entity

type User struct {
	Type         string `json:"type"  binding:"required"`
	Firstname    string `json:"firstname" binding:"required"`
	Lastname     string `json:"lastname"  binding:"required"`
	Email        string `json:"email"  binding:"required"`
	OfficeId     string `json:"officeId"`
	Image        string `json:"image"`
	Active       string `json:"active"`
	Company      string `json:"company"  binding:"required"`
	SecretAnswer string `json:"secretAnswer"`
}

type UserDB struct {
	ID           string `bson:"_id, omitempty"`
	Type         string `json:"type"  binding:"required"`
	Firstname    string `json:"firstname" binding:"required"`
	Lastname     string `json:"lastname"  binding:"required"`
	Email        string `json:"email"  binding:"required"`
	Company      string `json:"company"  binding:"required"`
	OfficeId     string `json:"officeId"`
	SecretAnswer string `json:"secretAnswer"`
}

type UpdateUser struct {
	ID           string `json:"_id" binding:"required"`
	Type         string `json:"type"  binding:"required"`
	Firstname    string `json:"firstname" binding:"required"`
	Lastname     string `json:"lastname"  binding:"required"`
	OfficeId     string `json:"officeId"`
	Company      string `json:"company"  binding:"required"`
	SecretAnswer string `json:"secretAnswer"`
}

type UserEmail struct {
	Email string `json:"email"  binding:"required"`
}

type UpdateImage struct {
	Image string `json:"image" binding:"required"`
}

type Response struct {
	Data    UserDB `json: "data"`
	Message string `json:"message"`
}

type ResponseUpdate struct {
	Data    UpdateUser `json: "data"`
	Message string     `json:"message"`
}
