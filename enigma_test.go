package main

import "testing"

func TestIntegrationEncrypt(t *testing.T)  {
    e := NewEncrypt()
    d := NewDescrypt()
    k := keygenerator.New()
    e.Key = k
    d.Key = k
    e.Date = "01012016"
    d.Date = "01012016"
    secretMessage := e.Encrypt("helloworld")
    decodedMessage := d.Decrypt(text)
    if decodedMessage != "helloworld"  {
        t.Log("The integration test failed: ", decodedMessage)
        t.Fail()
    }
}
