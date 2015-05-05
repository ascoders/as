package models

type Base struct{}

type BaseModel interface {
	Gets() []*Base
	Get(id int) *Base
	Add(this *Base) error
	Update(this *Base) error
	Delete(id int) error
}

func (this *Base) Gets() []*Base {
	var r []*Base

	return r
}

func (this *Base) Get(id int) *Base {
	return this
}

func (this *Base) Add() error {
	return nil
}

func (this *Base) Update() error {
	return nil
}

func (this *Base) Delete(id int) error {
	return nil
}
