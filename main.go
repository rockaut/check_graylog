package main

import (
    "errors"
    "flag"
    "fmt"
    "github.com/rockaut/check_graylog/pkg/graylog"
    "os"
    //"strconv"
)

var exitCode int
var agent graylog.Agent

func init() {

    err := initializeFlags()
    if err != nil {
        flag.Usage()
        exitCode = 1
        os.Exit(exitCode)
    }

    agent = graylog.Agent{
        Host:     "127.0.0.1",
        Port:     9000,
        User:     "admin",
        Password: "admin",
    }
    agent.Init(2)

    exitCode = 0
}

func main() {
    fmt.Println("Say hello to Graylog")

    res, err := agent.GetSystem()
    if err != nil {
        fmt.Println(err)
    }

    fmt.Print(res)

    os.Exit(exitCode)
}

func initializeFlags() error {
    flagMode := flag.String("mode", "cmk", "Agent Mode {cmk|zbx}. cmk = Check_MK, zbx = Zabbix")
    flagHost := flag.String("host", "", "Graylog node to fetch data")

    flag.Parse()

    switch *flagMode {
    case "cmk":
    case "zbx":
        //pass
    default:
        return errors.New("Unknown agentmode given")
    }

    switch *flagHost {
    case "":
        return errors.New("a graylog host has to be defined")
    }

    return nil
}
