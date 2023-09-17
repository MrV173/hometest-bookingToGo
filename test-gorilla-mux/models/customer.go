package models

type Customer struct {
	ID            int                 `json:"-" gorm:"primary_key"`
	Families      []Family            `json:"Keluarga" gorm:"foreignKey:CustomerID; not null"`
	CST_Name      string              `json:"nama" gorm:"type: varchar(255);not null"`
	CST_Dob       string              `json:"tanggal_lahir" gorm:"type: date; not null"`
	CST_PhoneNum  string              `json:"telepon" gorm:"type: varchar(20); not null"`
	NationalityID int                 `json:"-" gorm:"not null"`
	Nationality   NationalityResponse `json:"kewarganegaraan"`
	CST_Email     string              `json:"email" gorm:"type: varchar(50);not null"`
}

type CustomerResponse struct {
	ID           int            `json:"id"`
	Family       FamilyResponse `json:"hubungan"`
	CST_Name     string         `json:"nama"`
	CST_Dob      string         `json:"tanggal_lahir"`
	CST_PhoneNum string         `json:"telepon"`
	CST_Email    string         `json:"email"`
}

func (CustomerResponse) TableName() string {
	return "customers"
}
