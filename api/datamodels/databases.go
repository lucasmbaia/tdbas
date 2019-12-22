package datamodels

type Databases struct {
	ID            string          `json:",omitempty"`
	Organization  string          `json:",omitempty"`
	Name          string          `json:",omitempty"`
	Description   string          `json:",omitempty"`
	Replicas      int             `json:",omitempty"`
	Status        string          `json:",omitempty"`
	Error         NullString      `json:",omitempty"`
	Username      string          `json:",omitempty" gorm:"-"`
	Password      string          `json:",omitempty" gorm:"-"`
}

func (Databases) TableName() string {
	return "databases_tdbas"
}
