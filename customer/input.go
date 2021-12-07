package customer

type RegisterCustomerInput struct {
	Nama  string `json:"nama" binding:"required"`
	NoHp  string `json:"no_hp" binding:"required"`
	Email string `json:"email"`
}

type InputDetailCustomer struct {
	ID int `uri:"id" binding:"required"`
}

type InputDeleteCustomer struct {
	ID int `json:"id" binding:"required"`
}
