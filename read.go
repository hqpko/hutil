package hutil

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
)

func ReadFile(fpath string) ([]byte, error) {
	f, e := os.Open(fpath)
	if e != nil {
		return nil, e
	}
	defer f.Close()
	return ioutil.ReadAll(f)
}

func ReadJson(fpath string, i interface{}) error {
	d, e := ReadFile(fpath)
	if e != nil {
		return e
	}
	return json.Unmarshal(d, i)
}

func ReadToml(fpath string, i interface{}) error {
	_, e := toml.DecodeFile(fpath, i)
	return e
}
