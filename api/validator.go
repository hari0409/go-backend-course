package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/hari0409/backend-go/util"
)

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		// Check if currency is supported or not.
		return util.IsSupportedCurrency(currency)
	}
	return false
}
