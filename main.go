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
    aflags := agentFlags{}
    aflags.initializeFlags()

    err := aflags.parseFlags()
    if err != nil {
        flag.Usage()
        exitCode = 1
        os.Exit(exitCode)
    }

    agent = graylog.Agent{
        Host:     *aflags.Host,
        Port:     *aflags.Port,
        User:     *aflags.User,
        Password: *aflags.Password,
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

/* =================
 Flag handling
================= */

type agentFlags struct {
    Mode     *string
    Host     *string
    Port     *int
    User     *string
    Password *string
}

func (flags *agentFlags) initializeFlags() {
    flags.Mode = flag.String("mode", "cmk", "Agent Mode {cmk|zbx}. cmk = Check_MK, zbx = Zabbix")
    flags.Host = flag.String("host", "127.0.0.1", "Graylog node to fetch data")
    flags.Port = flag.Int("port", 9000, "Graylog port to fetch data")
    flags.User = flag.String("user", "admin", "Graylog user or \"token\"")
    flags.Password = flag.String("password", "admin", "Users password or token")
}

func (flags *agentFlags) parseFlags() error {
    flag.Parse()

    switch *flags.Mode {
    case "cmk":
    case "zbx":
        //pass
    default:
        return errors.New("Unknown agentmode given")
    }

    switch *flags.Host {
    case "":
        return errors.New("a graylog host has to be defined")
    }

    switch *flags.User {
    case "":
        return errors.New("a user or token has to be provided")
    }

    switch *flags.Password {
    case "":
        return errors.New("a password/token has to be provided")
    }

    return nil
}
