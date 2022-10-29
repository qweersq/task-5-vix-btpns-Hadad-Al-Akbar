package models

type Photo struct {
	ID      uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Title   string `gorm:"type:varchar(255)" json:"title"`
	Caption string `gorm:"type:varchar(255)" json:"caption"`
	Url     string `gorm:"type:varchar(255)" json:"url"`
	UserID  uint64 `gorm:"not null" json:"-"`
	User    User   `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
}
