package model

import (
	"github.com/google/uuid"
)

type Bairro struct {
	Model
	Nome     string    `gorm:"size:100;not null"`
	CidadeID uuid.UUID `gorm:"not null"`
	Cidade   Cidade    `gorm:"foreignKey:CidadeID"`
	//Logradouros []Logradouro `gorm:"foreignKey:BairroID"`
}
