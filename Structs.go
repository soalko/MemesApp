package main

type Role struct {
	IDRole   string `gorm:"primaryKey" json:"id_role"`
	NameRole string `json:"name"`
}

type User struct {
	IDUser       string `gorm:"primaryKey" json:"id_user"`
	IDRole       string `gorm:"primaryKey" json:"id_role"`
	NameUser     string `json:"user_name"`
	Login        string `json:"login"`
	PasswordHash string `json:"password_hash"`
}

type Meme struct {
	IDMeme       int     `gorm:"primaryKey;column:idmeme" json:"id_meme"`
	Name         string  `gorm:"column:name;type:varchar(255)" json:"name_meme"`
	IDUser       int     `gorm:"column:iduser" json:"id_user"`
	ShortDescrij string  `gorm:"column:shortdescription;type:varchar(1023)" json:"short_description"`
	LongDescrip  string  `gorm:"column:longdescription;type:varchar(1023)" json:"long_description"`
	VectorList   float32 `gorm:"column:vectorlist" json:"vector_list"`
}

type TagMeme struct {
	IDTag  string `gorm:"primaryKey" json:"id_tag"`
	IDMeme string `gorm:"primaryKey" json:"id_meme"`
}

type Tag struct {
	IDTag   string `gorm:"primaryKey" json:"id_tag"`
	NameTag string `json:"name_tag"`
}

type CategoryMeme struct {
	IDCategory string `gorm:"primaryKey" json:"id_category"`
	IDMeme     string `gorm:"primaryKey" json:"id_meme"`
}

type Category struct {
	IDCategory   string `gorm:"primaryKey" json:"id_category"`
	NameCategory string `json:"name_category"`
}
