package utils

type Validator struct{}

func (v *Validator) IsEmail(value string) bool {
	return true
}
