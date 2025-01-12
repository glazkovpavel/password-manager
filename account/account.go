package account

import (
	"errors"
	"github.com/fatih/color"
	"math/rand/v2"
	"net/url"
	"time"
)

type Account struct {
	Login     string    `json:"login"` // - это теги или мета информация для дальнейшего маппинга полей
	Password  string    `json:"password"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (acc *Account) Output() {
	color.Cyan(acc.Login)
	color.Cyan(acc.Password)
	color.Cyan(acc.Url)
	//fmt.Println(acc.Login, acc.Password, acc.Url)
}

func (acc *Account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.Password = string(res)
}

func NewAccount(login, password, urlString string) (*Account, error) {
	if login == "" {
		return nil, errors.New("login required")
	}

	_, err := url.ParseRequestURI(urlString)

	if err != nil {
		return nil, errors.New("INVALID_URL")
	}

	newAcc := &Account{

		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Password:  password,
		Url:       urlString,
		Login:     login,
	}

	if password == "" {
		newAcc.generatePassword(12)
	}

	return newAcc, nil
}

var letterRunes = []rune("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM1234567890-+*&^%$#@!")
