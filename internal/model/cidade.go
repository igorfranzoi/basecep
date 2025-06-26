package model

import (
	"github.com/google/uuid"
)

type Cidade struct {
	Model
	Nome     string    `gorm:"size:200;not null"`
	EstadoID uuid.UUID `gorm:"not null"`
	Estado   Estado    `gorm:"foreignKey:EstadoID"`
	Bairros  []Bairro  `gorm:"foreignKey:CidadeID"`
}
