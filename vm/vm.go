package vm

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/jeffcail/echoframe/g"
	"github.com/jeffcail/gtools"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"os"
	"path/filepath"
	"strconv"
)

var err error

type StoreMem struct {
	Ldb   *gtools.LevelDB
	Db    *xorm.Engine
	Rdb   *redis.Client
	Mongo *gtools.MongoDb
	Log   *zap.Logger
}

var Box *StoreMem

func NewStore() {
	Box = new(StoreMem)
	if Box.Ldb == nil {
		Box.newLevelDB()
	}
	if Box.Db == nil {
		Box.newOrm()
	}

	if Box.Rdb == nil {
		Box.newRedis()
	}

	if Box.Mongo == nil {
		Box.newMongo()
	}

	Box.Log = newLogger()
}

func (s *StoreMem) newLevelDB() {
	pr, err := findProjectRoot()
	if err != nil {
		panic(err)
	}

	val := g.GM.Get("leveldb").(string)
	var p string
	if val == "" {
		p = fmt.Sprintf("%s%s", pr, "./leveldb_data")
	} else {
		p = fmt.Sprintf("%s%s", pr, val)
	}

	s.Ldb, err = gtools.CreateLevelDB(p)
	if err != nil {
		panic(err)
	}
}

func (s *StoreMem) newOrm() {
	m := g.GM.Get("mysql").(map[string]interface{})
	d, ok := m["dsn"].(string)
	if !ok {
		panic(ok)
	}
	show, ok := m["show"].(bool)
	if !ok {
		panic(ok)
	}
	s.Db, err = gtools.NewXrm(d, show)
	if err != nil {
		panic(err)
	}
}

func (s *StoreMem) newRedis() {
	m := g.GM.Get("redis").(map[string]interface{})
	url, ok := m["url"].(string)
	if !ok {
		panic(ok)
	}
	password, ok := m["password"].(int)
	if !ok {
		panic(ok)
	}
	_password := strconv.Itoa(password)

	channel := m["db"].(int)
	s.Rdb, err = gtools.NewRedis(url, _password, channel)
	if err != nil {
		panic(err)
	}
}

func (s *StoreMem) newMongo() {
	m := g.GM.Get("mongodb").(string)
	s.Mongo, err = gtools.NewMongoDb(nil, m)
	if err != nil {
		panic(err)
	}
}

func findProjectRoot() (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if _, err := os.Stat(filepath.Join(currentDir, "go.mod")); err == nil {
			return currentDir, nil
		}

		parentDir := filepath.Dir(currentDir)
		if parentDir == currentDir {
			return "", fmt.Errorf("could not find go.mod file")
		}

		currentDir = parentDir
	}
}
