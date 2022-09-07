package models

type Request struct {
	ID        uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Requestor string `gorm:"size:255;not null" json:"requestor"`
	To        string `gorm:"size:100;not null" json:"to"`
	Status    string `gorm:"size:100;not null;" json:"status"`
}
