package boot

import (
	"errors"
	"log"
	"os"

	_procs "github.com/echo-scaffolding/common/procs"
	"github.com/spf13/viper"
)

const (
	DevYaml  = "dev.yaml"
	PreYaml  = "pre.yaml"
	ProdYaml = "prod.yaml"
	DevIni   = "dev.ini"
	PreIni   = "pre.ini"
	ProdIni  = "prod.ini"
)

var (
	LoadErr = errors.New("the main configuration file is missing a parameter")
)

// Boot
func Boot() {
	InitLogger()
	InitMysql()
	InstanceRedis()
	InitLevelDB()
	InitMongoDB()
	InitEs()
}

type ApplicationConf struct {
	Local              bool
	ExtFormat          string
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
		cmd()
		return
	}

	em := app.EnvModel
	ext := app.ExtFormat
	if !app.Local {
		log.Fatal(LoadErr)
	}

	if app.ExtFormat == "" && app.EnvModel == "" {
		InitYaml(DevYaml)
		return
	}

	if app.ExtFormat == "" && app.EnvModel != "" {
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

	if app.ExtFormat != "" && app.EnvModel == "" {
		switch app.ExtFormat {
		case "ini":
			InitIni(DevIni)
			return
		case "yaml":
			InitYaml(DevYaml)
			return
		default:
			InitYaml(DevYaml)
			return
		}
	}

	if app.Local {
		switch em {
		case "dev":
			switch ext {
			case "yaml":
				InitYaml(DevYaml)
				return
			case "ini":
				InitIni(DevIni)
				return
			default:
				InitYaml(DevYaml)
				return
			}
		case "pre":
			switch ext {
			case "yaml":
				InitYaml(PreYaml)
				return
			case "ini":
				InitIni(PreIni)
				return
			default:
				InitYaml(PreYaml)
				return
			}
		case "prod":
			switch ext {
			case "yaml":
				InitYaml(ProdYaml)
				return
			case "ini":
				InitIni(ProdIni)
				return
			default:
				InitYaml(ProdYaml)
				return
			}
		}
	}
}

func cmd() {
	if len(os.Args) <= 1 && len(os.Args) < 9 {
		return
	}

	if os.Args[1] == "-ip" && os.Args[3] == "-p" && os.Args[5] == "-c" && os.Args[7] == "-g" {
		InitNacos()
	}
}
