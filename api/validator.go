package api

import (
	"github.com/eliasmanj/go-bank/util"
	"github.com/go-playground/validator/v10"
)

var validCurrency validator.Func = func(fieldfLevel validator.FieldLevel) bool {
	if currency, ok := fieldfLevel.Field().Interface().(string); ok {
		return util.IsCurrencySupported(currency)
	}
	return false
}
