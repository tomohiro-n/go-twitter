package server
import (
	"path/filepath"
	"io/ioutil"
	"encoding/json"
)

type Config struct {
	Database DatabaseConfig `json:"database"`
	Twitter  TwitterConfig `json:"twitter"`
	SearchTerms []string `json:"searchTerms"`
}

type TwitterConfig struct {
	ConsumerKey    string `json:"consumerKey"`
	ConsumerSecret string `json:"consumerSecret"`
	AccessToken    string `json:"accessToken"`
	AccessSecret   string `json:"accessSecret"`
}

type DatabaseConfig struct {
	Host         string `json:"host"`
	DatabaseName string `json:"databaseName"`
}

var ServerConfig Config

func LoadConfig() {
	absPath, _ := filepath.Abs("src/resources/params.json")
	file, err := ioutil.ReadFile(absPath)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(file, &ServerConfig)
}