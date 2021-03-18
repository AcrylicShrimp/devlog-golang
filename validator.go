package main

import "github.com/go-playground/validator/v10"

type Validator struct {
	validate *validator.Validate
}

func NewValidator() *Validator {
	return &Validator{validate: validator.New()}
}

func (v *Validator) Validate(i interface{}) (err error) {
	if err = v.validate.Struct(i); err != nil {
		return err
	}

	return nil
}
