package validator

import (
	"fmt"
	"reflect"
	"regexp"
)

// Validate will read the target's StructTag to determine if each field pass validation test
func Validate(target interface{}) (bool, error) {

	v := reflect.ValueOf(target).Elem()
	t := v.Type()

	for idx := 0; idx < t.NumField(); idx++ {
		f := t.Field(idx)
		fv := v.Field(idx)

		regexStr, ok := f.Tag.Lookup("ffsz-validator")
		if !ok {
			continue
		}

		r, compileErr := regexp.Compile(regexStr)
		if compileErr != nil {
			return false, fmt.Errorf("Failed to compile regex for validator with field %v and type %v", f, t)
		}

		if !r.MatchString(reflect.ValueOf(fv.Interface()).String()) {
			return false, nil
		}
	}

	return true, nil
}
