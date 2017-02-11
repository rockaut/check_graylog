package main

import (
    "fmt"
    "github.com/rockaut/check_graylog/pkg/graylog"
    "os"
    //"strconv"
)

var agentflags graylog.AgentFlags
var exitCode int
var agent graylog.Agent

func init() {
    agentflags := graylog.AgentFlags{}
    agentflags.Initialize()

    err := agentflags.Parse()
    if err != nil {
        agentflags.Usage()
        exitCode = 1
        os.Exit(exitCode)
    }

    fmt.Printf("mode: %v (%v)\n", agentflags.Mode, int(agentflags.Mode))

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
