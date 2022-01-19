package validator

import (
	"errors"
	"strings"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type (
	InvalidValidationError = validator.InvalidValidationError
	ValidationErrors       = validator.ValidationErrors
	FieldError             = validator.FieldError
)

type Errors []error

func ToErrors(err error, ts ...ut.Translator) Errors {
	if err == nil {
		return nil
	}

	if es, ok := err.(ValidationErrors); ok {
		t := ZH
		if len(ts) > 0 && ts[0] != nil {
			t = ts[0]
		}

		errs := make([]error, len(es))
		for i := 0; i < len(es); i++ {
			errs[i] = errors.New(es[i].Translate(t))
		}
		return errs
	}

	return []error{err}
}

func (es Errors) First() error {
	if len(es) == 0 {
		return nil
	}
	return es[0]
}

func (es Errors) String() string {
	sb := strings.Builder{}
	for i := 0; i < len(es); i++ {
		if i > 0 {
			sb.WriteString("\n")
		}
		sb.WriteString(es[i].Error())
	}
	return sb.String()
}
