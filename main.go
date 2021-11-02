package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Seja bem vindo ao gerador de senhas!")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Digite quantos caracteres você deseja para a sua senha: ")
	scanner.Scan()
	input, err := strconv.ParseInt(scanner.Text(), 10, 64)

	if err != nil {
		log.Fatal(err)
	}

	var i int64
	chars := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "1", "2", "3", "4", "5", "6", "7", "8", "9", "!", "@", "#", "$", "%", "¨", "&", "*", "(", ")", "[", "]", "{", "}", ";", ":", "/", "|", ",", ".", "<", ">", "¹", "²", "³", "£", "¢", "¬", "§", "´", "`", "~", "^", "ª", "º", "°", "?", "'", "-", "+", "=", "_"}
	password := ""

	for i < input {
		index := rand.Intn(len(chars))
		password += chars[index]
		i++
	}

	fmt.Println("==========================================================================\nSua senha é:", password)
}