package Usecase

type ICalendar interface {
	Add() error
	Delete() error
	Edit() error
	ListAll()
}
