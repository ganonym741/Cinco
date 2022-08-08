package param

type User struct {
	Id         string `gorm:"type:uuid;primary_key" json:"id"`
	Username   string `json:"username"`
	Fullname   string `json:"fullname"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	BirthDate  string `json:"birthdate"`
	Domicile   string `json:"domicile"`
	Occupation string `json:"occupation"`
}
