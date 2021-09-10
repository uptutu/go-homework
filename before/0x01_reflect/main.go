package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	o := order{
		ordId:      456,
		customerId: 56,
	}

	createQuery(o)

	e := employee{
		name:    "Naveen",
		id:      565,
		address: "Coimbatore",
		salary:  90000,
		country: "India",
	}
	createQuery(e)

	i := 90
	createQuery(i)
}

type order struct {
	ordId      int
	customerId int
}

type employee struct {
	name    string
	id      int
	address string
	salary  int
	country string
}

func createQuery(item interface{}) {
	irType := reflect.TypeOf(item)
	if irType.Kind() != reflect.Struct {
		fmt.Println("unsupported type")
		return
	}
	irValue := reflect.ValueOf(item)

	buffer := make([]byte, 0, 1<<10)
	buffer = append(buffer, []byte("insert into ")...)
	buffer = append(buffer, []byte(irType.Name())...)
	buffer = append(buffer, []byte(" values ")...)
	buffer = append(buffer, '(')
	for i := 0; i < irType.NumField(); i++ {
		field := irType.Field(i)
		val := irValue.Field(i)
		switch field.Type.Kind() {
		case reflect.String:
			buffer = append(buffer, '"')
			buffer = append(buffer, []byte(val.String())...)
			buffer = append(buffer, '"')
			buffer = append(buffer, ',')
		case reflect.Int:
			buffer = append(buffer, []byte(strconv.Itoa(int(val.Int())))...)
			buffer = append(buffer, ',')
		}
	}
	buffer[len(buffer)-1] = ')'

	fmt.Println(string(buffer))
}
