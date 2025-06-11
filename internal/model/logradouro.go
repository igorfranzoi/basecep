package model

import (
	"basecep/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Logradouro struct {
	Model
	Nome     string    `gorm:"size:250;not null"`
	Tipo     string    `gorm:"size:50;not null"`
	BairroID uuid.UUID `gorm:"not null"`
	Bairro   Bairro    `gorm:"foreignKey:BairroID"`
	CEPs     []CEP     `gorm:"foreignKey:LogradouroID"`
}

func (o *Logradouro) BeforeCreate(tx *gorm.DB) (err error) {
	utils.ReturnUUID(o)
	return
}
