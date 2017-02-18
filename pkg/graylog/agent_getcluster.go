package graylog

import (
    "encoding/json"
    "fmt"
)

/*
ClusterResponse response
*/
type ClusterResponse map[string]SystemOverviewResponse

/*
GetCluster return api/cluster response
*/
func (agent *Agent) GetCluster() (*ClusterResponse, error) {
    url := fmt.Sprintf("http://%v:%v/%v", agent.Host, agent.Port, "api/cluster")
    if agent.PrettyResponse {
        url = fmt.Sprintf("%v?pretty=true", url)
    }

    req, err := agent.newGetRequest(url)
    if err != nil {
        return nil, err
    }

    body, _, doErr := agent.do(req)
    if doErr != nil {
        return nil, doErr
    }

    cRes := ClusterResponse{}
    jsonErr := json.Unmarshal(body, &cRes)

    if jsonErr != nil {
        return nil, jsonErr
    }

    return &cRes, nil
}
