package main

import (
	"fmt"
	"reflect"
)

type CPUMeta struct {
	Model string
	Cores int `ns:"/system/cpu/cores"`
}

type CPU struct {
	Meta  CPUMeta
	Usage int `ns:"/system/cpu/usage"`
}

func printMetrics(s interface{}) {
	rt := reflect.TypeOf(s)
	rv := reflect.ValueOf(s)

	if rt.Kind() == reflect.Struct {
		// ok
	} else if rt.Kind() == reflect.Ptr {
		if rt.Elem().Kind() == reflect.Struct {
			// ok
		} else {
			panic(fmt.Sprintf("wrong type %v", rt.Name()))
		}
	} else {
		panic(fmt.Sprintf("wrong type %v", rt.Name()))
	}

	for i := 0; i < rt.NumField(); i++ {
		rtEl := rt.Field(i)
		rtV := rv.Field(i)
		switch rtEl.Type.Kind() {
		case reflect.Int:
			tag := rtEl.Tag.Get("ns")
			fmt.Printf("%s=%d\n", tag, rtV.Int())
		case reflect.Struct:
			printMetrics(rtV.Interface())
		}
	}
}

func main() {
	cpu := CPU{
		Meta: CPUMeta{
			Model: "Intel",
			Cores: 4,
		},
		Usage: 34,
	}
	printMetrics(cpu)
}
