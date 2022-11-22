package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Please enter the first number: ");

	firstInput, _ := reader.ReadString('\n');

	firstOperand, err := strconv.ParseFloat(strings.TrimSpace(firstInput), 64)
		
	if err != nil {
		panic("Please enter a number")
	}

	fmt.Print("Please enter the second number: ")

	secondInput, _:= reader.ReadString('\n');
	
	secondOperand, err := strconv.ParseFloat(strings.TrimSpace(secondInput), 64);

	if err != nil {
		panic("Please enter a number")
	}

	
	sum := firstOperand + secondOperand

	sum = math.Round(sum * 100) / 100

	fmt.Printf("%v + %v = %v\n\n", firstOperand, secondOperand, sum)
	
}


