package boot2

import (
	"errors"
	_procs "github.com/echoframe/common/procs"
	"github.com/spf13/viper"
	"log"
)

const (
	DevYaml  = "dev.yaml"
	PreYaml  = "pre.yaml"
	ProdYaml = "prod.yaml"
)

var (
	LoadErr = errors.New("the main configuration file is missing a parameter")
)

// Boot 启动
func Boot() {
	InitLogger()
	InitMysql()
	InstanceRedis()
	InitLevelDB()
	InitMongoDB()
	//InitEs()
}

type ApplicationConf struct {
	Local              bool
	Remote             bool
	EnvModel           string
	IsEnableGOMAXPROCS bool
}

func init() {
	app := &ApplicationConf{}
	viper.SetConfigFile("application.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(app)
	if err != nil {
		panic(err)
	}

	if app.IsEnableGOMAXPROCS {
		_procs.GroRuntimeMaxCpu()
	}

	if app.Remote {
		return
	}

	em := app.EnvModel
	if !app.Local {
		log.Fatal(LoadErr)
	}

	if app.EnvModel == "" {
		InitYaml(DevYaml)
		return
	}

	if app.EnvModel != "" {
		switch app.EnvModel {
		case "dev":
			InitYaml(DevYaml)
			return
		case "pre":
			InitYaml(PreYaml)
			return
		case "prod":
			InitYaml(ProdYaml)
			return
		default:
			InitYaml(DevYaml)
			return
		}
	}

	if app.EnvModel == "" {
		InitYaml(DevYaml)
	}

	if app.Local {
		switch em {
		case "dev":
			InitYaml(DevYaml)
		case "pre":
			InitYaml(PreYaml)
		case "prod":
			InitYaml(ProdYaml)
		}
	}
}
