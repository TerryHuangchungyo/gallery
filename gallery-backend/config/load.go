package config

import (
	. "gallery-backend/utils"
	"os"

	toml "github.com/pelletier/go-toml"
)

type configReader struct {
	Err error
}

func (r *configReader) GetConfig(tree *toml.Tree, config interface{}) {
	if r.Err != nil {
		return
	}

	r.Err = tree.Unmarshal(config)
}

func init() {
	var err error

	projectEnv, exist := os.LookupEnv("PROJECT_ENV")

	if exist {
		var configData *toml.Tree

		switch projectEnv {
		case "local":
			configData, err = toml.LoadFile("/go/src/env/local.toml")
		case "gcp":
			configData, err = toml.LoadFile("/go/src/env/gcp.toml")
		default:
			InfoLog.Println("Server start with default config setting...")
			return
		}

		if err != nil {
			ErrorLog.Println("Error on load config file...")
			panic(err)
		}

		var reader configReader
		reader.GetConfig(configData, &App)
		reader.GetConfig(configData, &GalleryDB)
		if reader.Err != nil {
			ErrorLog.Println("Error on load config file...")
			panic(err)
		}
	} else {
		InfoLog.Println("Server start with default config setting...")
	}
}
