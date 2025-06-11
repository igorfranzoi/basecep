package model

import (
	"github.com/google/uuid"
)

type CEP struct {
	Model
	Codigo       string     `gorm:"size:8;unique;not null"`
	LogradouroID uuid.UUID  `gorm:"not null"`
	Logradouro   Logradouro `gorm:"foreignKey:LogradouroID"`
	NumeroInicio string     `gorm:"size:10"`
	NumeroFim    string     `gorm:"size:10"`
	Complemento  string     `gorm:"size:250"`
}
