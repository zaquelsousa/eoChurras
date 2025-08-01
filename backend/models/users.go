package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Name string
	PassWorld string
	Role Role
}

type Role string
const (
	Manager       Role = "manager"
	Garcom        Role = "garcom"
	Churrasqueiro Role = "churrasqueiro"
)

