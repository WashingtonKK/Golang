package main

import (
	"fmt"
)

type MyInterface interface {
	Method1(param1 int, param2 string, param3 bool) string
	Method2(param5 bool, param3 int) bool
}

type data struct {
	plenumReadings
	ambientReadings
}

func (t *data) Method1 (param1 int, param2 string, param3 bool) string {
	fmt.Println("Done %s", 4)
}

func (t *data) Method2 (param5 bool, param3 int)) bool {
	fmt.Println("Did %s this", 6)
}

var wk MyInterface = &data{}

func main {
	wk.Method1(2, "Wash", true)
}