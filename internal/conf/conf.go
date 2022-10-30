package conf

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type conf struct {
	ListenAddr string `json:"listen_addr"`

	JWTToken string `json:"jwt_token"`
	Token    string `json:"token"`
}

var CONF *conf

func init() {
	initConf()
}

func initConf() {
	var cf conf

	file, err := ioutil.ReadFile("configs/config.json")
	if err != nil {
		marshal, _ := json.MarshalIndent(cf, " ", "  ")
		if err2 := ioutil.WriteFile("config.json", marshal, 00666); err2 != nil {
			log.Fatalln(err2)
		}
		log.Fatalln(err)
	}

	err = json.Unmarshal(file, &cf)
	if err != nil {
		log.Fatalln(err)
	}

	CONF = &cf
}
