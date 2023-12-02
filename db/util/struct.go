package util

import (
	"fmt"
	"reflect"
)

// PrintStructValues prints the values of a struct in the format: "member name: member value"
func PrintStructValues(data interface{}) {
	val := reflect.ValueOf(data)

	if val.Kind() != reflect.Struct {
		fmt.Println("Not a struct")
		return
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		fieldValue := val.Field(i).Interface()
		fmt.Printf("%s: %v\n", field.Name, fieldValue)
	}
}