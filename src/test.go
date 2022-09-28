package main

import (
	"bufio"
	"fmt"
	"os"
	"rsc.io/quote"
)

func name_greeter(name string) {
	fmt.Println("Hello", name)
}

func main() {
	fmt.Println(quote.Go())
	scanner := bufio.NewScanner(os.Stdin)
	var test int = 244
	var testString string = "Hello there general"
	const constante int = 145564
	fmt.Println("valeur de la constante constante : ", constante)
	fmt.Println(test)
	fmt.Println(testString)
	fmt.Print("Please enter your name: ")
	scanner.Scan()
	nom := scanner.Text()
	name_greeter(nom)

}
