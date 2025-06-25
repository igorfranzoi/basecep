package model

import (
	"github.com/google/uuid"
)

type Logradouro struct {
	Model
	Nome     string    `gorm:"size:250;not null"`
	Tipo     string    `gorm:"size:50;not null"`
	BairroID uuid.UUID `gorm:"not null"`
	Bairro   Bairro    `gorm:"foreignKey:BairroID"`
	CEPs     []CEP     `gorm:"foreignKey:LogradouroID"`
}
