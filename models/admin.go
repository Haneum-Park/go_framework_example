package models

type Admin struct {
	BaseEntity
	Username    string `gorm:"type:varchar(100);not null;comment:관리자 아이디(이메일)" json:"username"`
	Password    string `gorm:"type:text;not null;comment:비밀번호" json:"password"`
	Salt        string `gorm:"type:text;not null" json:"salt"`
	AdminPermit uint   `gorm:"size:1;not null;comment:관리자 권한;default:0" json:"admin_permit"`
} // NOTE 관리자
