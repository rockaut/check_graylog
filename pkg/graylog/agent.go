package graylog

import (
    "bytes"
    "encoding/json"
    "errors"
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

    authKey  string
    authUser string

    httpClient    http.Client
    httpUserAgent string

    PrettyResponse bool
}

type loginResponse struct {
    ValidUntil string `json:"valid_until"`
    SesionID   string `json:"session_id"`
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

func (agent *Agent) getAuth() error {
    if agent.httpUserAgent == "" {
        agent.Init(DefaultAgentHTTPTimeout)
    }

    if agent.User == "token" {
        if agent.Password == "" {
            return errors.New("no token provided")
        }

        agent.authKey = agent.Password
        agent.authUser = "token"
        return nil
    }

    url := fmt.Sprintf("http://%v:%v/api/system/sessions", agent.Host, agent.Port)

    loginString := fmt.Sprintf("{ \"username\":\"%v\", \"password\":\"%v\", \"host\":\"\" }", agent.User, agent.Password)
    loginMap := []byte(loginString)
    req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(loginMap))
    if err != nil {
        return err
    }
    req.Header.Set("Accept", "application/json")
    req.Header.Set("Content-Type", "application/json")

    res, getErr := agent.httpClient.Do(req)
    if getErr != nil {
        return err
    }
    if res.StatusCode != 200 {
        err := CommonError{
            Response: *res,
            Request:  *req,
        }
        return err
    }

    body, readErr := ioutil.ReadAll(res.Body)
    if readErr != nil {
        return readErr
    }
    loginRes := loginResponse{}
    jsonErr := json.Unmarshal(body, &loginRes)
    if jsonErr != nil {
        return jsonErr
    }

    agent.authKey = loginRes.SesionID
    agent.authUser = "session"

    return nil
}

func (agent *Agent) newGetRequest(url string) (*http.Request, error) {
    if agent.authKey == "" || agent.authUser == "" {
        err := agent.getAuth()
        if err != nil {
            return nil, err
        }
    }

    req, err := http.NewRequest(http.MethodGet, url, nil)
    if err != nil {
        return nil, err
    }

    req.Header.Set("User-Agent", agent.httpUserAgent)
    req.Header.Set("Accept", "application/json")
    req.SetBasicAuth(agent.authKey, agent.authUser)

    return req, nil
}

func (agent *Agent) do(request *http.Request) ([]byte, *http.Response, error) {
    res, doErr := agent.httpClient.Do(request)
    if doErr != nil {
        return nil, nil, doErr
    }
    if res == nil {
        return nil, nil, errors.New("no response")
    }

    if res.StatusCode != 200 {
        err := CommonError{
            Response: *res,
            Request:  *request,
        }
        return nil, nil, err
    }

    body, readErr := ioutil.ReadAll(res.Body)
    if readErr != nil {
        return nil, nil, readErr
    }
    defer res.Body.Close()

    return body, res, nil
}
