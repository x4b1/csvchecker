package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

var usage = func() {
	fmt.Fprint(os.Stderr, `usage: csvchecker (configuration JSON path) (CSV path)
csvchecker given a configuration json and csv validate structure`)
}

func main() {
	command := commandHandler{}
	command.fromArgs(os.Args[1:])
}

type commandHandler struct {
}

func (c *commandHandler) fromArgs(args []string) {
	p := c.getArgs(args)
	if len(p) < 2 {
		c.fail("Invalid argument number")
	}
	conf := c.processConfigurationFile(p[0])

	checker := conf.createChecker()

	errors := checker.Check(c.getFileReader(p[1]))
	if len(errors) > 0 {
		for _, err := range errors {
			fmt.Fprint(os.Stderr, err.ToString(), "\n")
		}
		c.fail("CSV file has some errors")
	}

}

func (c *commandHandler) getArgs(args []string) []string {
	flagSet := flag.NewFlagSet("csvchecker", flag.ExitOnError)
	flagSet.Usage = usage
	flagSet.Parse(args)

	return flagSet.Args()
}

func (c *commandHandler) fail(text string) {
	fmt.Fprint(os.Stderr, text)
	os.Exit(-1)
}

func (c *commandHandler) processConfigurationFile(path string) *configuration {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		c.fail(err.Error())
	}

	conf := new(configuration)
	err = json.Unmarshal(file, &conf)

	if err != nil {
		c.fail(err.Error())
	}

	err = conf.validate()

	if err != nil {
		c.fail(err.Error())
	}

	return conf
}

func (c *commandHandler) getFileReader(path string) io.Reader {
	var r io.Reader
	var err error
	r, err = os.Open(path)

	if err != nil {
		c.fail(err.Error())
	}

	return r
}
