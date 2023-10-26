package db

import (
	"gorm.io/gorm"
)

type SQLTransactionTable struct {
	db *gorm.DB
}

// player requests (aka transactions) are objects that need to be made available such that
// anyone can recreate the state

type TransactionSQLFormat struct {
	GameId string

	// unix in nano seconds
	UnixTimestamp int `gorm:"primaryKey;autoIncrement:false"`

	// which tick it was registered at
	Tick int

	// serialized data string
	Data string

	Type string
}

func NewTransactionTable(db *gorm.DB) (*SQLTransactionTable, error) {
	dst := TransactionSQLFormat{}
	err := db.AutoMigrate(&dst)
	if err != nil {
		return nil, err
	}

	txTable := SQLTransactionTable{db: db}
	return &txTable, nil
}

func (t *SQLTransactionTable) AddEntries(entries ...TransactionSQLFormat) error {
	for _, entry := range entries {
		tx := t.db.Save(entry)
		if tx.Error != nil {
			return tx.Error
		}
	}
	return nil
}

func (t *SQLTransactionTable) GetEntriesUntilTick(tickNumber int) ([]TransactionSQLFormat, error) {
	var entries []TransactionSQLFormat
	tx := t.db.Where("`Tick` < ?", tickNumber+1).Find(&entries)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return entries, nil
}
