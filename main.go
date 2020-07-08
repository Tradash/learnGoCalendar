package main

import "os"

type PlaceInfo string

type Events struct {
	DateStart  string
	DateEnd    string
	Members    []string
	Place      PlaceInfo
	InfoEvents string
}

type ICalendar interface {
	Add() error
	Delete() error
	Edit() error
	ListAll()
}

func main() {
	println("Запущено ...", os.Args[0])
}
