package decrypt

import "testing"

func TestCreateDecrypt(t *testing.T)  {
    var d interface{}
    d = New()
    switch d.(type){
    case *Decrypt:

    default:
        t.Log("Decrypt was never created")
        t.Fail()
    }
}

func TestCreateDecrypt()  {

}
