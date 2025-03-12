package storage

import (
	"github.com/charmbracelet/log"
	badger "github.com/dgraph-io/badger/v4"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

var logger log.Logger

func init() {
	logger = *log.WithPrefix("Storage")
	logger.Debug("Storage init")
}

func mustGetDB() *badger.DB {
	path := viper.GetString("storage.path")

	db, err := badger.Open(badger.DefaultOptions(path))

	if err != nil {
		logger.Fatalf("Failed open connection to key-value storage file '%s': %s", path, err)
	}
	return db
}
func hashPassword(password string) string {
	cost := viper.GetInt("password.cost")
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		logger.Fatalf("Failed hash password: %s", err)
	}
	return string(hashed)
}

func bytesToString(str *string) func([]byte) error {
	return func(b []byte) error {
		strValue := string(b)
		str = &strValue
		return nil
	}
}

func getDefaultPassword() string {
	pass := viper.GetString("system.default-password")

	return hashPassword(pass)
}

// Returns hashed password
func GetPassword() string {
	db := mustGetDB()
	defer db.Close()
	txn := db.NewTransaction(false)
	item, err := txn.Get([]byte("passwordHash"))
	if err != nil {
		logger.Warnf("Failed get password hash from key-value-store: %s", err)
		logger.Warnf("Default password will be used")
		return getDefaultPassword()
	}
	var value string
	err = item.Value(bytesToString(&value))
	if err != nil {
		logger.Warnf("Failed get password hash from key-value-store: %s", err)
		logger.Warnf("Default password will be used")
		return getDefaultPassword()
	}
	return value
}
