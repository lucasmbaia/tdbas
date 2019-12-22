package datamodels

type Organizations struct {
	ID	    string	`json:",omitempty"`
	Name	    string	`json:",omitempty"`
	Description *string	`json:",omitempty"`
	Databases   []Databases	`json:",omitempty" gorm:"association_foreignkey:ID;foreignkey:Organization"`
}

func (Organizations) TableName() string {
	return "organizations"
}
