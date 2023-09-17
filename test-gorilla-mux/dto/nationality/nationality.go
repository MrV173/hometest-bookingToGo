package nationalitydto

type CreateNationalityRequest struct {
	Nationality_Name string `json:"kewarganegaraan" form:"kewarganegaraan" validate:"required"`
	Nationality_Code string `json:"code" form:"code" validate:"required"`
}

type UpdateNationalityRequest struct {
	Nationality_Name string `json:"kewarganegaraan" form:"kewarganegaraan"`
	Nationality_Code string `json:"code" form:"code"`
}

type NationalityResponse struct {
	ID               int    `json:"id"`
	Nationality_Name string `json:"kewarganegaraan"`
	Nationality_Code string `json:"code"`
}
