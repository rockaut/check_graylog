package main

import (
    "fmt"
    "github.com/rockaut/check_graylog/pkg"
    "os"
    //"strconv"
)

var agentflags graylog.AgentFlags
var exitCode int

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

    exitCode = 0
}

func main() {
    fmt.Println("Say hello to Graylog")

    os.Exit(exitCode)
}
