package util

import (
	"encoding/json"
	"reflect"
	"strconv"
)

func Marshal[T any](value *T) ([]byte, error) {
	parseSetDefault[T](value)
	return json.Marshal(value)
}

func MarshalIndent[T any](value *T, prefix, indent string) ([]byte, error) {
	parseSetDefault[T](value)
	return json.MarshalIndent(value, prefix, indent)
}

func parseSetDefault[T any](obj *T) {
	typeof := reflect.TypeOf(obj)
	valueof := reflect.ValueOf(obj)
	for i := 0; i < typeof.Elem().NumField(); i++ {
		if valueof.Elem().Field(i).IsZero() {
			f := typeof.Elem().Field(i)
			def := f.Tag.Get("default")
			// comment := f.Tag.Get("comment")
			if def != "" {
				if def == "-" {
					continue
				}
				key := typeof.Elem().Field(i).Type.String()
				switch key {
				case "int":
					result, _ := strconv.Atoi(def)
					valueof.Elem().Field(i).SetInt(int64(result))
				case "int64":
					result, _ := strconv.ParseInt(def, 10, 64)
					valueof.Elem().Field(i).SetInt(result)
				case "uint":
					result, _ := strconv.ParseUint(def, 10, 64)
					valueof.Elem().Field(i).SetUint(result)
				case "string":
					valueof.Elem().Field(i).SetString(def)
				case "bool":
					result, _ := strconv.ParseBool(def)
					valueof.Elem().Field(i).SetBool(result)
				case "float32":
					result, _ := strconv.ParseFloat(def, 32)
					valueof.Elem().Field(i).SetFloat(result)
				case "float64":
					result, _ := strconv.ParseFloat(def, 64)
					valueof.Elem().Field(i).SetFloat(result)
				}
			}
		}
	}
}
