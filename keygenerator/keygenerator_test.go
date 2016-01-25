package keygenerate

import "testing"

func TestKeyGenerateCreation(t *testing.T)  {
    var typeReader interface{}
    typeReader = New()
    switch typeReader.(type) {
    case *Key:

    default:
        t.Log("Didn't create a key")
        t.Fail()
    }
    key := New()
    if len(key.Code) != 5 {
        t.Log("Actual keystring not created: ", key.Code)
        t.Fail()
    }
}

func TestCreateRandomNumber(t *testing.T)  {
    randomString := createRandom(9)
    if len(randomString) != 9 {
        t.Log("Seems the random number generator doesn't quite work: ", randomString)
        t.Fail()
    }
}
