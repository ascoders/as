package controllers

import (
	"newWoku/lib/response"
)

type Editor struct{}

type EditorController interface {
}

func (this *Editor) Gets() (int, []byte) {
	return response.Success("Gets success!")
}
