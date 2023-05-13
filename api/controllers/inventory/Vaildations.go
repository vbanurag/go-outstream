package inventory

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

func inventoryInfo(request map[string]interface{}) (err error) {
	var errorValidations []string

	for k, v := range request {
		if v == "" {
			errorValidations = append(errorValidations, fmt.Sprintf("%s cannnot be empty", k))
		}
	}
	validate := validator.New()
	err = validate.RegisterValidation("inventory-info", func(fl validator.FieldLevel) bool {
		_, ok := fl.Field().Interface().(map[string]interface{})
		if !ok {
			return false
		}
		return true
	})
	return
}
