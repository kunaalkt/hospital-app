package models

import "gorm.io/gorm"

type Patient struct {
    gorm.Model
    Name         string `gorm:"not null"`
    Age          int `gorm:"not null"`
    Gender       string `gorm:"not null"`
    Diagnosis    string `gorm:"not null"`
    Prescription string
}