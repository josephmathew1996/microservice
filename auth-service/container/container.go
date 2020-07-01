package container

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

//User map
var User map[string]string

//LoadData loads
func LoadData() {
	raw, err := ioutil.ReadFile("master.json")
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal(raw, &User)
	if err != nil {
		log.Fatalln(err)
	}
}
