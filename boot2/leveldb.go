package boot2

import (
	"github.com/echoframe/common/db"
	"github.com/echoframe/conf"
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
