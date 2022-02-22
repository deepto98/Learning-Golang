package main

import (
	"fmt"
	"os"
)

type Calculation struct {
	operand1  float64
	operand2  float64
	operation string
}

func main() {
	operands, operation := acceptInputFromCli()
	cal := Calculation{operand1: operands[0], operand2: operands[1], operation: operation}
	fmt.Println(getResult(cal))

}
func getResult(cal Calculation) float64 {
	operation := cal.operation
	switch operation {
	case "+":
		return cal.operand1 + cal.operand2
	case "-":
		return cal.operand1 - cal.operand2
	case "*":
		return cal.operand1 * cal.operand2
	case "/":
		return cal.operand1 / cal.operand2
	case "%":
		modulo := int(cal.operand1) % int(cal.operand2)
		return float64(modulo)
	}
	return 1
}
func acceptInputFromCli() ([]float64, string) {
	var i, j float64
	var k string

	fmt.Println("Enter First Operand")
	_, err := fmt.Scan(&i)
	checkOperandType(err)

	fmt.Println("Enter Second Operand")
	_, err = fmt.Scan(&j)
	checkOperandType(err)

	fmt.Println("Enter Operation")
	_, _ = fmt.Scan(&k)
	checkOperationType(k)

	arr := []float64{i, j} // slice
	return arr, k
}

func checkOperandType(err error) bool {
	if err != nil {
		fmt.Println("Wrong type Entered")
		os.Exit(0)
	}
	return true
}
func checkOperationType(operation string) bool {
	switch operation {
	case "+", "-", "*", "/", "%":
		return true
	}
	fmt.Println("Wrong operation Entered")
	os.Exit(0)
	return false
}
