package model

import (
	"github.com/igorfranzoi/basecep/utils"

	"gorm.io/gorm"
)

type Estado struct {
	Model
	Nome    string   `gorm:"size:250;not null"`
	UF      string   `gorm:"size:2;unique;not null"`
	Cidades []Cidade `gorm:"foreignKey:EstadoID"`
}

func (o *Estado) BeforeCreate(tx *gorm.DB) (err error) {
	utils.ReturnUUID(o)
	return
}
