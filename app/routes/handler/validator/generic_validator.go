package validator

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

// RegisterTagName used to replace the field name with json tag for the error message
func RegisterTagName() {
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := fld.Tag.Get(JSON)
		if name == Underscore || name == EmptyString {
			return fld.Name
		}

		return name
	})
}

// RegisterCustomValidation use add custom validator
func RegisterCustomValidation(validate *validator.Validate) {
	validate.RegisterValidation(alpha, func(fl validator.FieldLevel) bool {
		return containsOnly(fl.Field().String(), alphaRegex)
	})

	validate.RegisterValidation(alphaNumeric, func(fl validator.FieldLevel) bool {
		return containsOnly(fl.Field().String(), alphaNumericRegex)
	})

	validate.RegisterValidation(year, func(fl validator.FieldLevel) bool {
		return !containsOnly(strconv.Itoa(int(fl.Field().Int())), yearRegex)
	})

	validate.RegisterValidation(timestamp, func(fl validator.FieldLevel) bool {
		return containsOnly(fl.Field().String(), timestampRegex)
	})

	validate.RegisterValidation(intWithPlus, func(fl validator.FieldLevel) bool {
		return containsOnly(fl.Field().String(), positiveIntegerWithPlusRegex)
	})

	validate.RegisterValidation(alphaNumericWithHyphenSpace, func(fl validator.FieldLevel) bool {
		return containsOnly(fl.Field().String(), alphaNumericWithHyphenSpaceRegex)
	})

	validate.RegisterValidation(date, func(fl validator.FieldLevel) bool {
		return !containsOnly(fl.Field().String(), dateRegex)
	})

	validate.RegisterValidation(oneOfPriority, func(fl validator.FieldLevel) bool {
		return validateOneOf(fl.Field().String(), AllowedPriorities)
	})

	validate.RegisterValidation(oneOfBool, func(fl validator.FieldLevel) bool {
		return validateOneOf(fl.Field().String(), boolValues)
	})

	validate.RegisterValidation(oneOfStatus, func(fl validator.FieldLevel) bool {
		return validateOneOf(fl.Field().String(), statusValues)
	})

	validate.RegisterValidation(oneOfPeriod, func(fl validator.FieldLevel) bool {
		return validateOneOf(fl.Field().String(), periodValues)
	})

	validate.RegisterValidation(omitEmpty, func(fl validator.FieldLevel) bool {
		return trimAndCheckLength(fl.Field().String(), 0)
	})

	validate.RegisterValidation(oneOfSortFields, func(f1 validator.FieldLevel) bool {
		return validateOneOf(f1.Field().String(), validSortFields)
	})

	validate.RegisterValidation(oneOfSortOrders, func(f1 validator.FieldLevel) bool {
		return validateOneOf(f1.Field().String(), validSortOrders)
	})
}

// RegisterCustomTranslation use add custom validator translation
func RegisterCustomTranslation(validate *validator.Validate, trans ut.Translator) {
	validate.RegisterTranslation(alpha, trans, func(ut ut.Translator) error {
		return ut.Add(alpha, "{0} must contain alpha characters", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(alpha, fe.Field())
		return t
	})

	validate.RegisterTranslation(alphaNumeric, trans, func(ut ut.Translator) error {
		return ut.Add(alphaNumeric, "{0} must contain alpha-numeric characters", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(alphaNumeric, fe.Field())
		return t
	})

	validate.RegisterTranslation(timestamp, trans, func(ut ut.Translator) error {
		return ut.Add(timestamp, "{0} must be a valid UTC timestamp", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(timestamp, fe.Field())
		return t
	})

	validate.RegisterTranslation(intWithPlus, trans, func(ut ut.Translator) error {
		return ut.Add(intWithPlus, "{0} must contains with only + and integers", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(intWithPlus, fe.Field())
		return t
	})

	validate.RegisterTranslation(alphaNumericWithHyphenSpace, trans, func(ut ut.Translator) error {
		return ut.Add(alphaNumericWithHyphenSpace, "{0} must contain alpha-numeric characters with hyphen and space", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(alphaNumericWithHyphenSpace, fe.Field())
		return t
	})

	validate.RegisterTranslation(year, trans, func(ut ut.Translator) error {
		return ut.Add(year, "{0} must contain valid year", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(year, fe.Field())
		return t
	})

	validate.RegisterTranslation(date, trans, func(ut ut.Translator) error {
		return ut.Add(date, "{0} must contain valid date format YYYY-MM-DD", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(date, fe.Field())
		return t
	})

	validate.RegisterTranslation(oneOfPriority, trans, func(ut ut.Translator) error {
		return ut.Add(oneOfPriority, fmt.Sprintf("Priority list must only contain: %v", AllowedPriorities), true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(oneOfPriority, fe.Field())
		return t
	})

	validate.RegisterTranslation(oneOfBool, trans, func(ut ut.Translator) error {
		return ut.Add(oneOfBool, fmt.Sprintf("{0} must contains %v boolean value", boolValues), true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(oneOfBool, fe.Field())
		return t
	})

	validate.RegisterTranslation(oneOfStatus, trans, func(ut ut.Translator) error {
		return ut.Add(oneOfStatus, fmt.Sprintf("{0} must contains %v status value", statusValues), true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(oneOfStatus, fe.Field())
		return t
	})

	validate.RegisterTranslation(oneOfPeriod, trans, func(ut ut.Translator) error {
		return ut.Add(oneOfPeriod, fmt.Sprintf("{0} must contains %v period value", periodValues), true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(oneOfPeriod, fe.Field())
		return t
	})

	validate.RegisterTranslation(omitEmpty, trans, func(ut ut.Translator) error {
		return ut.Add(omitEmpty, "{0} can not be empty", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(omitEmpty, fe.Field())
		return t
	})

	validate.RegisterTranslation(oneOfSortFields, trans, func(ut ut.Translator) error {
		return ut.Add(oneOfSortFields, fmt.Sprintf("{0} must be one of the valid sorting fields: "+strings.Join(validSortFields, ", ")), true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(oneOfSortFields, fe.Field())
		return t
	})

	validate.RegisterTranslation(oneOfSortOrders, trans, func(ut ut.Translator) error {
		return ut.Add(oneOfSortOrders, fmt.Sprintf("{0} must be one of the valid sorting orders: "+strings.Join(validSortOrders, ", ")), true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(oneOfSortOrders, fe.Field())
		return t
	})
}
