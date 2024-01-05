package db

import (
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type SQLTransactionTable struct {
	db *gorm.DB
}

type CelestiaConn struct {
	conn *grpc.ClientConn
}

func NewCelestiaConn(conn *grpc.ClientConn) *CelestiaConn {
	return &CelestiaConn{conn: conn}
}

// player requests (aka transactions) are objects that need to be made available such that
// anyone can recreate the state

type TransactionSQLFormat struct {
	GameId string `json:"gameId"`

	// unix in nano seconds
	UnixTimestamp int `json:"unixTimestamp" gorm:"primaryKey;autoIncrement:false"`

	// which tick it was registered at
	Tick int `json:"tick"`

	// serialized data string
	Data string `json:"data"`

	Type string `json:"type"`
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
