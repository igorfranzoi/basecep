package model

import (
	"github.com/igorfranzoi/basecep/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Cidade struct {
	Model
	Nome     string    `gorm:"size:200;not null"`
	EstadoID uuid.UUID `gorm:"not null"`
	Estado   Estado    `gorm:"foreignKey:EstadoID"`
	Bairros  []Bairro  `gorm:"foreignKey:CidadeID"`
}

func (o *Cidade) BeforeCreate(tx *gorm.DB) (err error) {
	utils.ReturnUUID(o)
	return
}
