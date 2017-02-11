package graylog

import (
    "fmt"
    "net/http"
)

/*
CommonError thrown on any graylog related error
*/
type CommonError struct {
    Response http.Response
    Request  http.Request
}

/*

 */
func (e CommonError) Error() string {
    return fmt.Sprintf("%v (%v) on %v", e.Response.Status, e.Response.StatusCode, e.Request.URL)
}
