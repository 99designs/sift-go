package sift

import (
	"encoding/json"
	"reflect"
	"strings"
)

// parseTag splits a struct field's json tag into its name and
// comma-separated options.
func parseTag(tag string) (string, string) {
	if idx := strings.Index(tag, ","); idx != -1 {
		return tag[:idx], string(tag[idx+1:])
	}
	return tag, ""
}

func parseField(field *reflect.StructField, immutable *reflect.Value) (string, interface{}) {
	if !field.Anonymous {
		fieldName := field.Name
		name, option := parseTag(field.Tag.Get("json"))
		value := immutable.FieldByName(field.Name)

		if option == "omitempty" && isEmptyValue(value) {
			return "", ""
		}

		if name != "" {
			fieldName = name
		}
		return fieldName, value.Interface()
	}

	return "", ""
}

func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return false
}

func marshalEvent(te TypedEvent, cf CustomFields, e Event) ([]byte, error) {
	data := make(map[string]interface{})

	for key, value := range cf {
		data[key] = value
	}

	typ := reflect.TypeOf(te)
	// if a pointer to a struct is passed, get the type of the dereferenced object
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	immutable := reflect.ValueOf(te).Elem()
	// loop through the struct's fields and set the map
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		if key, value := parseField(&field, &immutable); key != "" {
			data[key] = value
		}
	}

	im := reflect.ValueOf(e)
	ty := reflect.TypeOf(e)
	// loop through the struct's fields and set the map
	for i := 0; i < ty.NumField(); i++ {
		field := ty.Field(i)
		if key, value := parseField(&field, &im); key != "" {
			data[key] = value
		}
	}
	return json.Marshal(data)
}
