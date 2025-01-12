package main

import (
	"demo/password/account"
	"demo/password/encrypter"
	"demo/password/files"
	"demo/password/output"
	"fmt"
	"github.com/fatih/color"
	"github.com/joho/godotenv"
	"strings"
)

var menu = map[string]func(*account.VaultWithDb){
	"1": createAccount,
	"2": findAccountByUrl,
	"3": findAccountByLogin,
	"4": deleteAccount,
}

var menuVariants = []string{
	"1. Создать аккаунт",
	"2. Найти аккаунт по URL",
	"3. Найти аккаунт по логину",
	"4. Удалить аккаунт",
	"5. Выход",
	"Выберите вариант",
}

func main() {
	fmt.Println("__Менеджер паролей__")
	env := godotenv.Load()

	if env != nil {
		output.PrintError("Не удалось найти .env file")
	}

	vault := account.NewVault(files.NewJsonDb("db/data.vault"), *encrypter.NewEncrypter())
Menu:
	for {

		variant := promptData(menuVariants...)
		menuFunc := menu[variant]
		if menuFunc == nil {
			break Menu
		}
		menuFunc(vault)
	}

}

func deleteAccount(vault *account.VaultWithDb) {
	url := promptData("Введите URL для поиска")
	isDeleted := vault.DeleteAccountByUrl(url)
	if isDeleted {
		color.Green("Удалено")
	} else {
		output.PrintError("Не найдено")
	}
}

func findAccountByUrl(vault *account.VaultWithDb) {
	url := promptData("Введите URL для поиска")
	accounts := vault.FindAccounts(url, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Url, str)
	})

	outputAccountInfo(&accounts)
}

func findAccountByLogin(vault *account.VaultWithDb) {
	login := promptData("Введите логин для поиска")
	accounts := vault.FindAccounts(login, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Login, str)
	})

	outputAccountInfo(&accounts)
}

func outputAccountInfo(accounts *[]account.Account) {
	if len(*accounts) == 0 {
		output.PrintError("Аккаунтов не найдено")
	}

	for _, account := range *accounts {
		account.Output()
	}
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

func promptData(prompt ...string) string {
	for i, line := range prompt {
		if i == len(prompt)-1 {
			fmt.Printf("%v:", line)
		} else {
			fmt.Println(line)
		}
	}
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

//func reverse(arr *[4]int) {
//	for index, value := range *arr {
//		(*arr)[len(arr)-1-index] = value
//	}
//}
//
//func double(num *int) {
//	*num = *num * 2
//}

//func menuCounter() func() { // замыкание
//	i := 0
//	return func() {
//		i++
//		fmt.Println(i)
//	}
//}
