package main

import "fmt"

func main() {

	var phone string

	fmt.Println("Введите номер (Прим. 78005353535):")

	fmt.Scanln(&phone)

	vkSpam(phone)

}
