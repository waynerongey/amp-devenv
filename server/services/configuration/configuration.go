package configuration

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

//Configuration .
type Configuration struct {
	Application  string
	Domain       string
	Address      string
	TimeOffset   time.Duration
	ReadTimeout  int64
	WriteTimeout int64
	Production   bool
	Static       string
	Templates    string
	AmpStylesFile string
	SessionKey   string //sessionStore storage
	Ssl          Ssl
}

//Ssl .
type Ssl struct {
	Enabled bool
}

//Config .
var Config *Configuration

func init() {
	file, err := os.Open("config.json")
	defer file.Close()
	if err != nil {
		log.Fatalln("Cannot open config file", err)
	}
	decoder := json.NewDecoder(file)
	c := Configuration{}
	err = decoder.Decode(&c)
	if err != nil {
		log.Fatalln("Cannot get configuration from file", err)
	}
	Config = &c
}
