package familylistdto

type CreateFamilyRequest struct {
	FL_Relation string `json:"hubungan" form:"hubungan" validate:"required"`
	FL_Name     string `json:"rl_nama" form:"rl_nama" validate:"required"`
	FL_Dob      string `json:"rl_tanggal_lahir" form:"rl_tanggal_lahir" validate:"required"`
	CustomerID  uint   `json:"customer_id" form:"customer_id" validate:"required"`
}

type UpdateFamilyRequest struct {
	FL_Relation string `json:"hubungan" form:"hubungan"`
	FL_Name     string `json:"rl_nama" form:"rl_nama"`
	FL_Dob      string `json:"rl_tanggal_lahir" form:"rl_tanggal_lahir"`
}

type FamilyResponse struct {
	ID          int    `json:"id"`
	FL_Relation string `json:"hubungan"`
	FL_Name     string `json:"rl_nama"`
	FL_Dob      string `json:"rl_tanggal_lahir"`
}
