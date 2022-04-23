package test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func init() {
	fmt.Println("get config init")
}

type Cfg struct {
	Name      string `yaml:name`
	Auto      bool   `yaml:auto`
	Port      int    `yaml:port`
	Blackip   []string
	Clusterip []string
	Health    Health
}

type Health struct {
	Url      string
	Cmd      string
	Interval string
	Timeout  string
	Disable  bool
}

func GetCfg() {

	cfg := Cfg{}
	fmt.Println(os.Getwd())
	fmt.Println(filepath.Abs(filepath.Dir(os.Args[0])))
	open, err := ioutil.ReadFile("./cfg.yaml")
	if err != nil {
		fmt.Println("Open err", err)
	}
	//all, err := ioutil.ReadAll(open)
	//if err != nil {
	//	fmt.Println("read all err",err)
	//}
	space := strings.TrimSpace(string(open))
	json.Unmarshal([]byte(space), &cfg)
	fmt.Println(cfg)
}
