package utils

import (
	"github.com/charmbracelet/charm/kv"
)

const KV_STORE = "sonr-dev-db"

func Set(key []byte, value []byte) error {
	// Open a database (or create one if it doesn’t exist)
	db, err := kv.OpenWithDefaults(KV_STORE)
	if err != nil {
		return err
	}
	defer db.Close()

	// Fetch updates and easily define your own syncing strategy
	if err := db.Sync(); err != nil {
		return err
	}

	// Save some data
	if err := db.Set(key, value); err != nil {
		return err
	}
	return nil
}

func Get(key []byte) ([]byte, error) {
	// Open a database (or create one if it doesn’t exist)
	db, err := kv.OpenWithDefaults(KV_STORE)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Fetch updates and easily define your own syncing strategy
	if err := db.Sync(); err != nil {
		return nil, err
	}

	// Quickly get a value
	return db.Get(key)
}
