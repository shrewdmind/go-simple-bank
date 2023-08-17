package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/shrewdmind/simplebank/util"
)

var validCurrency validator.Func = func(fieLdLevel validator.FieldLevel) bool {
	if currency, ok := fieLdLevel.Field().Interface().(string); ok {
		return util.IsSuppoertedCurrency(currency)
	}
	return false
}
