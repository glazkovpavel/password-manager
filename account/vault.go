package account

import (
	"demo/password/encrypter"
	"encoding/json"
	"github.com/fatih/color"
	"strings"
	"time"
)

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type VaultWithDb struct {
	Vault
	db  Db
	enc encrypter.Encrypter
}

type ByteReader interface {
	Read() ([]byte, error)
}

type ByteWriter interface {
	Write([]byte)
}

type Db interface {
	ByteWriter
	ByteReader
}

func (vault *Vault) FindAccounts(str string, checker func(Account, string) bool) []Account {
	var accounts []Account
	for _, account := range vault.Accounts {
		//isMatched := strings.Contains(account.Url, url)
		isMatched := checker(account, str)
		if isMatched {
			accounts = append(accounts, account)
		}
	}
	return accounts
}

func (vault *VaultWithDb) DeleteAccountByUrl(url string) bool {
	var accounts []Account
	isDeleted := false
	for _, account := range vault.Accounts {
		isMatched := strings.Contains(account.Url, url)
		if !isMatched {
			accounts = append(accounts, account)
			continue
		}
		isDeleted = true
	}
	vault.Accounts = accounts
	vault.save()
	return isDeleted
}

func NewVault(db Db, enc encrypter.Encrypter) *VaultWithDb {
	file, err := db.Read()
	if err != nil {
		return &VaultWithDb{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db:  db,
			enc: enc,
		}
	}
	data := enc.Decrypt(file)
	var vault Vault
	err = json.Unmarshal(data, &vault)
	color.Cyan("Найдено %d аккаунтов", len(vault.Accounts))
	if err != nil {
		color.Red("Не удалось разобрать файл data.vault", err.Error())
		return &VaultWithDb{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db:  db,
			enc: enc,
		}
	}
	return &VaultWithDb{
		Vault: vault,
		db:    db,
		enc:   enc,
	}

}

func (vault *VaultWithDb) AddAccount(acc Account) {
	vault.Accounts = append(vault.Accounts, acc)
	vault.save()
}

func (vault *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(vault)
	return file, err
}

func (vault *VaultWithDb) save() {
	vault.UpdatedAt = time.Now()
	data, err := vault.Vault.ToBytes()
	encData := vault.enc.Encrypt(data)
	if err != nil {
		color.Red("Не удалось преобразовать", err.Error())
	}
	vault.db.Write(encData)
}
