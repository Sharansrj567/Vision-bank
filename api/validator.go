package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/sharansrj567/golang_bank_app/util"
)

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		// check currency is supported
		return util.IsSupportedCurrency(currency)
	}

	return false
}
