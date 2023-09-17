package models

type Nationality struct {
	ID               int    `json:"id"`
	Nationality_Name string `json:"kewarganegaraan" gorm:"type: varchar(50); not null"`
	Nationality_Code string `json:"code" gorm:"type: varchar(2); not null"`
}

type NationalityResponse struct {
	ID               int    `json:"-"`
	Nationality_Name string `json:"kewarganegaraan"`
	Nationality_Code string `json:"code"`
}

func (NationalityResponse) TableName() string {
	return "nationalities"
}
