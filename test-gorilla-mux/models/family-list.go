package models

type Family struct {
	ID          int    `json:"-" gorm:"primary_key; not null"`
	FL_Relation string `json:"hubungan" gorm:"type : varchar(50); not null"`
	FL_Name     string `json:"nama" gorm:"type : varchar(50); not null"`
	FL_Dob      string `json:"tanggal_lahir" gorm:"type varchar(50); not null"`
	CustomerID  uint   `json:"-"`
}

type FamilyResponse struct {
	ID          int    `json:"-"`
	FL_Relation string `json:"hubungan"`
	FL_Name     string `json:"rl_nama"`
	FL_Dob      string `json:"rl_tanggal_lahir"`
	CustomerID  int
}

func (FamilyResponse) TableName() string {
	return "Families"
}
