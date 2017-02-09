package graylog

import (
    "fmt"
    "net/http"
    "time"
)

const (
    // DefaultAgentHTTPpTimeout Http Timeout in seconds
    DefaultAgentHTTPTimeout int = 2
    // DefaultAgentHTTPUserAgent http requests user agent string
    DefaultAgentHTTPUserAgent string = "check-graylog"
)

/*
Agent is the main object
*/
type Agent struct {
    Host     string
    Port     string
    User     string
    Password string

    httpClient    http.Client
    httpUserAgent string
}

/*
Init to the agent
*/
func (agent *Agent) Init(httpTimeout int) {

    if httpTimeout == 0 {
        httpTimeout = DefaultAgentHTTPTimeout
    }

    agent.httpClient = http.Client{
        Timeout: time.Second * time.Duration(httpTimeout),
    }

    agent.httpUserAgent = DefaultAgentHTTPUserAgent
}

/*
GetSystem returns api/system response
*/
func (agent *Agent) GetSystem() (*http.Response, error) {
    url := fmt.Sprintf("http://%v:%v/%v", agent.Host, agent.Port, "api/system")

    req, err := http.NewRequest(http.MethodGet, url, nil)
    if err != nil {
        return nil, err
    }

    req.Header.Set("User-Agent", agent.httpUserAgent)

    res, getErr := agent.httpClient.Do(req)
    if getErr != nil {
        return nil, err
    }

    return res, nil
}
