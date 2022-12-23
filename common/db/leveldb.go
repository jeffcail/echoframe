package db

import leveldb1 "github.com/jeffcail/leveldb"

var Ldb *leveldb1.LevelDB

func SetLevelDB(_leveldb *leveldb1.LevelDB) {
	Ldb = _leveldb
}
