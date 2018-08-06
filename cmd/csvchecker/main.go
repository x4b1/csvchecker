package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

var usage = func() {
	fmt.Fprint(os.Stderr, `usage: csvchecker (configuration JSON path) (CSV path)
csvchecker given a configuration json and csv validate structure`)
}

func main() {
	h := commandHandler{}
	h.fromArgs(os.Args[1:])
}

type commandHandler struct {
}

func (c *commandHandler) fromArgs(args []string) {
	p := c.getArgs(args)
	if len(p) < 2 {
		fmt.Fprint(os.Stderr, "Invalid argument number")
	}
	conf := c.processConfiguration(p[0])
	fmt.Printf("%+v\n", conf)
}

func (c *commandHandler) getArgs(args []string) []string {
	flagSet := flag.NewFlagSet("csvchecker", flag.ExitOnError)
	flagSet.Usage = usage
	flagSet.Parse(args)

	return flagSet.Args()
}

func (c *commandHandler) processConfiguration(path string) *configuration {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
	}
	fmt.Println(string(file))
	conf := new(configuration)
	err = json.Unmarshal(file, &conf)

	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
	}

	return conf
}

type configuration struct {
	separator rune `json:"separator"`
	columns   []struct {
		position  int `json:"position"`
		validator struct {
			validatorType string `json:"type"`
			checkRange    struct {
				min int
				max int
			} `json:"range"`
			regex string
		} `json:"validation"`
	} `json:"columns"`
}
