package main

import (
    "fmt"
  "github.com/amaxwellblair/enigma/encrypt"
  "github.com/amaxwellblair/enigma/keygenerator"
)

func main() {
    e := encrypt.NewEncrypt()
    d := encrypt.NewDecrypt()
    k := keygenerate.New()
    e.Key = k.Code
    d.Key = k.Code
    e.Date = "01012016"
    d.Date = "01012016"
    secretMessage := e.Encrypt("helloworld")
    decodedMessage := d.Decrypt(secretMessage)
    fmt.Println("E.key: ", e.Key, "D.Key: ", d.Key)
    fmt.Println("Encryted message: ", secretMessage)
    fmt.Println("Decrypted message: ", decodedMessage)
}
