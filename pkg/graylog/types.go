package graylog

import (
    "errors"
    "flag"
)

/*
AgentFlags object
*/
type AgentFlags struct {
    Mode AgentModeType

    flagMode *string
}

/*
AgentModeType object for defining monitoring solution for the agent
@see AgentMode_* constants
*/
type AgentModeType int8

/*
Constants for defining AgentModes
*/
const (
    // AgentMode_CMK defining Check_MK mode for agent
    AgentMode_CMK AgentModeType = 0
    // AgentMode_ZBX defining Zabbix mode for agent
    AgentMode_ZBX AgentModeType = 1
)

var agentModeDescriptions = [...]string{
    "Check_MK",
    "Zabbix",
}

func (mode AgentModeType) String() string {
    return agentModeDescriptions[mode]
}

/*
Initialize just a test
*/
func (agentflags *AgentFlags) Initialize() {
    agentflags.flagMode = flag.String("mode", "cmk", "Agent Mode {cmk|zbx}. cmk = Check_MK, zbx = Zabbix")
}

/*
Parse calling flag.Parse
*/
func (agentflags *AgentFlags) Parse() error {
    flag.Parse()

    switch *agentflags.flagMode {
    case "cmk":
        agentflags.Mode = AgentMode_CMK
    case "zbx":
        agentflags.Mode = AgentMode_ZBX
    default:
        return errors.New("Unknown agentmode given")
    }

    return nil
}

/*
Usage calling flag.Usage
*/
func (agentflags *AgentFlags) Usage() {
    flag.Usage()
}
