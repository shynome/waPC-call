package model

//go:generate go run github.com/pquerna/ffjson $GOFILE

type Input struct {
	Name string `json:"name"`
}
