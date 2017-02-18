package graylog

import (
    "encoding/json"
    "fmt"
)

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
GetSystem returns api/system response
*/
func (agent *Agent) GetSystem() (*SystemOverviewResponse, error) {
    url := fmt.Sprintf("http://%v:%v/%v", agent.Host, agent.Port, "api/system")
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

    sysRes := SystemOverviewResponse{}

    jsonErr := json.Unmarshal(body, &sysRes)
    if jsonErr != nil {
        return nil, jsonErr
    }

    return &sysRes, nil
}
