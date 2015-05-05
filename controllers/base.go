package controllers

import (
	"newWoku/lib/response"
)

type Base struct{}

type BaseController interface {
	Gets() (int, []byte)
	Get() (int, []byte)
	Add() (int, []byte)
	Update() (int, []byte)
	Delete() (int, []byte)
}

func (this *Base) Gets() (int, []byte) {
	return response.Success("Gets success!")
}

func (this *Base) Get() (int, []byte) {
	return response.Success("Get success!")
}

func (this *Base) Add() (int, []byte) {
	return response.Success("Add success!")
}

func (this *Base) Update() (int, []byte) {
	return response.Success("Update success!")
}

func (this *Base) Delete() (int, []byte) {
	return response.Success("Delete success!")
}
