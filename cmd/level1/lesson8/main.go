package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"os"
	"reflect"
	"strconv"
)

type Config struct {
	ConfigFile string `envconfig:"CONFIG_FILE" json:"config_file" yaml:"configFile"`
	Count int `envconfig:"COUNT" json:"count" yaml:"count"`
	Strings []string `evnconfig:"STRINGS" json:"strings" yaml:"strings"`
	Maps map[string]string `evnconfig:"MAPS" json:"maps" yaml:"maps" `
}

// flags (arguments): ./main -config_file=config.txt -count 10
func mainFlags() {
	configFile := flag.String("config_file", "", "config.txt" )
	count := flag.Int("count", 0, "10")
	flag.Parse()

	fmt.Println(*configFile, *count)

	var cfg Config

	flag.StringVar(&cfg.ConfigFile, "config_file", "", "config.txt")
	flag.IntVar(&cfg.Count, "count", 0, "config.txt")
	flag.Parse()
	fmt.Println(cfg)

}

func mainEnv() {
	var cfg Config
	cfg.ConfigFile = os.Getenv("CONFIG_FILE")

	cfg.Count,_ = strconv.Atoi(os.Getenv("COUNT"))
	fmt.Println(cfg)
}


func mainReflect() {
	var cfg Config

	tp := reflect.TypeOf(cfg)

	for i := 0; i < tp.NumField();i++ {
		fmt.Println(tp.Field(i).Tag.Get("json"))
	}
}

func main() {
	var cfg Config
	if err := envconfig.Process("lesson8", &cfg); err != nil {
		panic(err)
	}
	f, err := os.Open("cmd/lesson8/config.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	decoder := json.NewDecoder(f)
	decoder.Decode(&cfg)
	fmt.Println(cfg)

	mainReflect()
}

// env: CONFIG_FILE=/etc/passwd ./main