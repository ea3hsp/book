package store

import (
	"github.com/go-kit/kit/log"
	db "github.com/sonyarouje/simdb/db"
)

type simdbRepository struct {
	db     *db.Driver
	Logger log.Logger
}

// NewSimDB creates simDB
func NewSimDB(dbname string, logger log.Logger) IStore {
	db, err := db.New(dbname)
	if err != nil {
		logger.Log("[error]", "New SimDB", "msg", err.Error())
		return nil
	}
	return &simdbRepository{
		db: db,
	}
}
