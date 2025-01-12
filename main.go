package main

import (
	"demo/password/account"
	"demo/password/files"
	"fmt"
	"github.com/fatih/color"
)

func main() {
	fmt.Println("__Менеджер паролей__")
	vault := account.NewVault(files.NewJsonDb("db/data.json"))
Menu:
	for {

		variant := getMenu()
		switch variant {
		case 1:
			createAccount(vault)
		case 2:
			findAccount(vault)
		case 3:
			deleteAccount(vault)
		default:
			break Menu
		}
	}

}

func deleteAccount(vault *account.VaultWithDb) {
	url := promptData("Введите URL для поиска")
	isDeleted := vault.DeleteAccountByUrl(url)
	if isDeleted {
		color.Green("Удалено")
	} else {
		color.Red("Не найдено")
	}
}

func findAccount(vault *account.VaultWithDb) {
	url := promptData("Введите URL для поиска")
	accounts := vault.FindAccountsByUrl(url)

	if len(accounts) == 0 {
		color.Red("Аккаунтов не найдено")
	}

	for _, account := range accounts {
		account.Output()
	}
}

func getMenu() int {
	var variant int
	fmt.Println("Выберите вариант:")
	fmt.Println("1. Создать аккаунт")
	fmt.Println("2. Найти аккаунт")
	fmt.Println("3. Удалить аккаунт")
	fmt.Println("4. Выход")

	fmt.Scan(&variant)
	return variant
}

func createAccount(vault *account.VaultWithDb) {
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")
	myAccount, err := account.NewAccount(login, password, url)

	if err != nil {
		color.Cyan("Не верный формат URL или Логин", err)
		return
	}

	vault.AddAccount(*myAccount)

	//myAccount.Output()
}

func reverse(arr *[4]int) {
	for index, value := range *arr {
		(*arr)[len(arr)-1-index] = value
	}
}

func double(num *int) {
	*num = *num * 2
}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var input string
	fmt.Scanln(&input)
	return input
}

//str := "Привет!"
//str := []rune("Привет!")

//for _, char := range str {
//	fmt.Println(char, string(char))
//}

//a := 5
//pointerA := &a
//double(nil)
//fmt.Println(a)
//fmt.Println(pointerA)
//
//c := [4]int{1, 2, 3, 4}
//reverse(&c)

//
////account1 := account{login, password, url} // порядок важен
//myAccount := account{ // порядок не важен
//	url:   url,
//	login: login,
//}
