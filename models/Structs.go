package models

import "github.com/lib/pq"

type Category struct {
	IDCategory int    `gorm:"primaryKey;column:idcategory" json:"id_category"`
	Name       string `gorm:"column:name;type:varchar(255)" json:"name"`
}

type CategoryMeme struct {
	IDCategory int `gorm:"primaryKey;column:idcategory" json:"id_category"`
	IDMeme     int `gorm:"primaryKey;column:idmeme" json:"id_meme"`
}

type Meme struct {
	IDMeme           int             `gorm:"primaryKey;column:idmeme;autoIncrement" json:"id_meme"`
	Name             string          `gorm:"column:name;type:varchar(255)" json:"name"`
	IDUser           int             `gorm:"column:iduser" json:"id_user"`
	ShortDescription string          `gorm:"column:shortdescription;type:varchar(1023)" json:"short_description"`
	LongDescription  string          `gorm:"column:longdescription;type:varchar(1023)" json:"long_description"`
	VectorList       pq.Float32Array `gorm:"column:vectorlist;type:real[]" json:"vector_list"`
	Image            string          `gorm:"column:image;type:varchar(255)" json:"image"`
	User             User            `gorm:"-" json:"user"`
	Tags             []Tag           `gorm:"many2many:tagmeme;" json:"tags"`
	Categories       []Category      `gorm:"many2many:categorymeme;" json:"categories"`
}

type Tag struct {
	IDTag int    `gorm:"primaryKey;column:idtag" json:"id_tag"`
	Name  string `gorm:"column:name;type:varchar(255)" json:"name"`
}

type TagMeme struct {
	IDTag  int `gorm:"primaryKey;column:idtag" json:"id_tag"`
	IDMeme int `gorm:"primaryKey;column:idmeme" json:"id_meme"`
}
type UserRole struct {
	IDRole   int    `gorm:"primaryKey;column:idrole" json:"id_role"`
	RoleName string `gorm:"column:namerole;type:varchar(255)" json:"name_role"`
}

type User struct {
	IDUser   int      `gorm:"primaryKey;column:iduser" json:"id_user"`
	Name     string   `gorm:"column:name;type:varchar(255)" json:"name"`
	RoleID   int      `gorm:"column:idrole;not null" json:"id_role"`
	Login    string   `gorm:"column:login;type:varchar(255)" json:"login"`
	Password string   `gorm:"column:password;type:varchar(255)" json:"password"`
	Role     UserRole `gorm:"foreignKey:RoleID;references:IDRole" json:"role"`
	_        []Meme   `gorm:"-"`
}

//хз зачем просто архивные структуры
/*package models

type Category struct {
	IDCategory int    ``gorm:"primaryKey;column:idcategory" json:"id_category"`
	Name       string ``gorm:"column:name;type:varchar(255)" json:"name"`
}

type CategoryMeme struct {
	IDCategory int ``gorm:"primaryKey;column:idcategory" json:"id_category"`
	IDMeme     int ``gorm:"primaryKey;column:idmeme" json:"id_meme"`
}

type Meme struct {
	IDMeme           int        ``gorm:"primaryKey;column:idmeme;autoIncrement" json:"id_meme"`
	Name             string     ``gorm:"column:name;type:varchar(255)" json:"name"`
	IDUser           int        ``gorm:"column:iduser" json:"id_user"`
	ShortDescription string     ``gorm:"column:shortdescription;type:varchar(1023)" json:"short_description"`
	LongDescription  string     ``gorm:"column:longdescription;type:varchar(1023)" json:"long_description"`
	VectorList       []float32  ``gorm:"column:vectorlist;type:_float4" json:"vector_list"`
	Image            string     ``gorm:"column:image;type:varchar(255)" json:"image"`
	User             User       ``gorm:"foreignKey:IDUser;references:IDUser" json:"user"`
	Tags             []Tag      ``gorm:"many2many:tagmeme;" json:"tags"`
	Categories       []Category ``gorm:"many2many:categorymeme;" json:"categories"`
}

type Tag struct {
	IDTag int    ``gorm:"primaryKey;column:idtag" json:"id_tag"`
	Name  string ``gorm:"column:name;type:varchar(255)" json:"name"`
}

type TagMeme struct {
	IDTag  int ``gorm:"primaryKey;column:idtag" json:"id_tag"`
	IDMeme int ``gorm:"primaryKey;column:idmeme" json:"id_meme"`
}
type UserRole struct {
	IDRole   int    ``gorm:"primaryKey;column:idrole" json:"id_role"`
	NameRole string ``gorm:"column:namerole;type:varchar(255)" json:"name_role"`
}
type User struct {
	IDUser   int      ``gorm:"primaryKey;column:iduser" json:"id_user"`
	Name     string   ``gorm:"column:name;type:varchar(255)" json:"name"`
	IDRole   int      ``gorm:"column:idrole" json:"id_role"`
	Login    string   ``gorm:"column:login;type:varchar(255)" json:"login"`
	Password string   ``gorm:"column:password;type:varchar(255)" json:"password"`
	Role     UserRole ``gorm:"foreignKey:IDRole;references:IDRole" json:"role"`
}*/
