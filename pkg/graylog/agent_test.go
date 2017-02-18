package graylog

import (
    "testing"
)

const (
    testHost           string = "127.0.0.1"
    testPort           int    = 9000
    testLoginUsername  string = "admin"
    testLoginPassword  string = "admin"
    testLoginTokenUser string = "token"
    testLoginToken     string = "sgakbhbl6jcj19rbu41o2fkvm8fme5l394966vortb29bqvad5o" // //create with http://127.0.0.1:9000/api/users/admin/tokens/check-graylog?pretty=true
    testHTTPTimeout    int    = 2
)

func TestLoginVariant_SessionLogin_Success(t *testing.T) {
    t.Log("Test trying to GetSystem (should success)")
    testAgent := Agent{
        Host:     testHost,
        Port:     testPort,
        User:     testLoginUsername,
        Password: testLoginPassword,
    }
    sor, err := testAgent.GetSystem()
    if err != nil || sor == nil {
        t.Log(err)
        t.Fail()
    }

    t.Log(sor)
}

func TestLoginVariant_SessionLogin_Fail(t *testing.T) {
    t.Log("Test trying to GetSystem (should fail with 401)")
    testAgent := Agent{
        Host:     testHost,
        Port:     testPort,
        User:     testLoginUsername[:len(testLoginUsername)-1],
        Password: testLoginPassword,
    }
    sor, err := testAgent.GetSystem()
    if err != nil {
        t.Log(err)
        t.Fail()
    }

    t.Log(sor)
}

func TestLoginVariant_TokenLogin_Success(t *testing.T) {
    t.Log("Test trying to GetSystem (should success)")
    testAgent := Agent{
        Host:     testHost,
        Port:     testPort,
        User:     testLoginTokenUser,
        Password: testLoginToken,
    }
    sor, err := testAgent.GetSystem()
    if sor == nil || err != nil {
        t.Log(err)
        t.Fail()
    }

    t.Log(sor)
}

func TestLoginVariant_TokenLogin_Fail(t *testing.T) {
    t.Log("Test trying to GetSystem (should fail w. 401)")
    testAgent := Agent{
        Host:     testHost,
        Port:     testPort,
        User:     testLoginTokenUser,
        Password: testLoginToken[:len(testLoginToken)-1],
    }
    sor, err := testAgent.GetSystem()
    if sor != nil || err == nil {
        t.Log(sor)
        t.Fail()
    }

    t.Logf("sor: %v, %v", sor, err)
}
