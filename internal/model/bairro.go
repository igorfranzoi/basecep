package model

import (
	"basecep/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Bairro struct {
	Model
	Nome        string       `gorm:"size:100;not null"`
	CidadeID    uuid.UUID    `gorm:"not null"`
	Cidade      Cidade       `gorm:"foreignKey:CidadeID"`
	Logradouros []Logradouro `gorm:"foreignKey:BairroID"`
}

func (o *Bairro) BeforeCreate(tx *gorm.DB) (err error) {
	utils.ReturnUUID(o)
	return
}
