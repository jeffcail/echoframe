package boot

import (
	"github.com/echo-scaffolding/common/db"
	"github.com/echo-scaffolding/conf"
	leveldb1 "github.com/jeffcail/leveldb"
)

// InitLevelDB
func InitLevelDB() {
	levelDB, err := leveldb1.CreateLevelDB(conf.Config.LevelDBPath)
	if err != nil {
		panic(err)
	}
	db.SetLevelDB(levelDB)
}
