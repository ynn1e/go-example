package main

type BloodType int

//go:generate stringer -type BloodType blood_type.go
const (
	A BloodType = iota
	B
	O
	AB
)
