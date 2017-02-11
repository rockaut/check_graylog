package graylog

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "time"
)

const (
    // DefaultAgentHTTPTimeout Http Timeout in seconds
    DefaultAgentHTTPTimeout int = 2
    // DefaultAgentHTTPUserAgent http requests user agent string
    DefaultAgentHTTPUserAgent string = "check-graylog"
)

/*
Agent is the main object
*/
type Agent struct {
    Host     string
    Port     int
    User     string
    Password string

    token string

    httpClient    http.Client
    httpUserAgent string

    PrettyResponse bool
}

type loginResponse struct {
    Name       string `json:"name"`
    Token      string `json:"token"`
    LastAccess string `json:"last_access"`
}

/*
SystemOverviewResponse urn:jsonschema:org:graylog2:rest:models:system:responses:SystemOverviewResponse
*/
type SystemOverviewResponse struct {
    Facility       string `json:"facility"`
    Codename       string `json:"codename"`
    NodeID         string `json:"node_id"`
    ClusterID      string `json:"cluster_id"`
    Version        string `json:"version"`
    StartedAt      string `json:"started_at"`
    IsProcessing   bool   `json:"is_processing"`
    Hostname       string `json:"hostname"`
    Lifecycle      string `json:"lifecycle"`
    LbStatus       string `json:"lb_status"`
    Timezone       string `json:"timezone"`
    OperatingSyste string `json:"operating_system"`
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
    agent.PrettyResponse = true
}

func (agent *Agent) login() error {
    if agent.httpUserAgent == "" {
        agent.Init(DefaultAgentHTTPTimeout)
    }

    url := fmt.Sprintf("http://%v:%v/%v/%v/tokens/%v", agent.Host, agent.Port, "api/users", agent.User, agent.httpUserAgent)

    if agent.PrettyResponse {
        url = fmt.Sprintf("%v?pretty=true", url)
    }

    req, err := http.NewRequest(http.MethodPost, url, nil)
    if err != nil {
        return err
    }

    req.SetBasicAuth(agent.User, agent.Password)

    res, getErr := agent.httpClient.Do(req)
    if getErr != nil {
        return err
    }

    body, readErr := ioutil.ReadAll(res.Body)
    if readErr != nil {
        return readErr
    }
    defer res.Body.Close()

    loginRes := loginResponse{}
    jsonErr := json.Unmarshal(body, &loginRes)
    if jsonErr != nil {
        return jsonErr
    }

    agent.token = loginRes.Token

    return nil
}

/*
GetSystem returns api/system response
*/
func (agent *Agent) GetSystem() (*SystemOverviewResponse, error) {
    if agent.httpUserAgent == "" {
        agent.Init(DefaultAgentHTTPTimeout)
    }

    if agent.User != "token" {
        if agent.token == "" {
            agent.login()
        }
    } else {
        agent.token = agent.Password
    }

    url := fmt.Sprintf("http://%v:%v/%v", agent.Host, agent.Port, "api/system")
    if agent.PrettyResponse {
        url = fmt.Sprintf("%v?pretty=true", url)
    }

    req, err := http.NewRequest(http.MethodGet, url, nil)
    if err != nil {
        return nil, err
    }

    req.Header.Set("User-Agent", agent.httpUserAgent)
    req.Header.Set("Accept", "application/json")
    req.SetBasicAuth(agent.token, agent.User)

    res, getErr := agent.httpClient.Do(req)
    if getErr != nil {
        return nil, err
    }

    if res.StatusCode != 200 {
        err := CommonError{
            Response: *res,
            Request:  *req,
        }
        return nil, err
    }

    body, readErr := ioutil.ReadAll(res.Body)
    if readErr != nil {
        return nil, readErr
    }
    defer res.Body.Close()

    sysRes := SystemOverviewResponse{}
    jsonErr := json.Unmarshal(body, &sysRes)
    if jsonErr != nil {
        return nil, jsonErr
    }

    return &sysRes, nil
}
