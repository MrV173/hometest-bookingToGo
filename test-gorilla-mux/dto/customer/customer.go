package customerdto

type CreateCustomerRequest struct {
	CST_Name      string `json:"nama" form:"nama" validate:"required"`
	CST_Dob       string `json:"tanggal_lahir" form:"tanggal_lahir" validate:"required"`
	CST_PhoneNum  string `json:"telepon" form:"telepon" validate:"required"`
	NationalityID int    `json:"nationality_id" form:"nationality_id"`
	CST_Email     string `json:"email" form:"email" validate:"required"`
}

type UpdateCustomerRequest struct {
	CST_Name     string `json:"nama" form:"nama"`
	CST_Dob      string `json:"tanggal_lahir" form:"tanggal_lahir"`
	CST_PhoneNum string `json:"telepon" form:"telepon"`
	CST_Email    string `json:"email" form:"email"`
}

type CustomerResponse struct {
	ID           int    `json:"id"`
	CST_Name     string `json:"nama"`
	CST_Dob      string `json:"tanggal_lahir"`
	CST_PhoneNum string `json:"telepon"`
	CST_Email    string `json:"email"`
}
