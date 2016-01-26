package encrypt

import (
  "testing"
)

func TestCreateEncryptStruct(t *testing.T)  {
    var e interface{}
    e = NewEncrypt()
    switch e.(type) {
      case *Encrypt:

      default:
        t.Log("The encrypt struct isn't being created correctly")
        t.Fail()
    }
}

func TestZeroRotatingOnAGivenIndex(t *testing.T)  {
    e := NewEncrypt()
    e.Key = "12345"
    rotation := Rotate(e, 0)
    if rotation != 12 {
      t.Log("The 0 rotation didn't yield the proper number")
      t.Fail()
    }
}

func TestOneRotatingOnAGivenIndex(t *testing.T)  {
    e := NewEncrypt()
    e.Key = "12345"
    rotation := Rotate(e, 1)
    if rotation != 23 {
      t.Log("The 1 rotation didn't yield the proper number")
      t.Fail()
    }
}

func TestArotationIsOutOfRange(t *testing.T) {
    e := NewEncrypt()
    e.Key = "12345"
    rotation := Rotate(e, 4)
    if rotation != 0 {
      t.Log("The rotation function didn't recognize an overflow value")
      t.Fail()
    }
}

func TestPowerFunctionWorks(t *testing.T) {
  value := pow(2, 2)
  if value != 4 {
    t.Log("The power function is not working correctly")
    t.Fail()
  }
}

func TestDateChangerProperlyConvertDateIntoFourCharacterString(t *testing.T) {
    e := NewEncrypt()
    e.Date = "01012016"
    lastFourDigits := dateConv(e)
    if lastFourDigits != "4256" {
      t.Log("The date convertor isnt working quite right!", lastFourDigits)
      t.Fail()
    }
}

func TestAOffsetOnAGivenIndex(t *testing.T) {
    e := NewEncrypt()
    e.Date = "01012016"
    offset := Offset(e, 0)
    if offset != 4 {
      t.Log("Offset doesn't work on the 0 index given")
      t.Fail()
    }
}

func TestOffsetCanNotGoOutOfBoundsOnItsIndex(t *testing.T) {
    e := NewEncrypt()
    e.Date = "01012016"
    offset := Offset(e, 5)
    if offset != 0 {
        t.Log("Offset is not going to 0 when you go out of bounds on the string")
        t.Fail()
    }
}

func TestModulo(t *testing.T)  {
    k := modulo(3, 75, 77)
    if k != 75 {
        t.Log("Modulo still needs some work and is acting up: ", k)
        t.Fail()
    }
    m := modulo(2, 75, 77)
    if m != 77 {
        t.Log("Modulo does not work on numbers with in the first cycle through", m)
        t.Fail()
    }
    m = modulo(9, 75, 77)
    if m != 75 {
        t.Log("Modulo does not work on numbers with in the third cycle through", m)
        t.Fail()
    }
}

func TestEncryptOnASingleLetter(t *testing.T)  {
    e := NewEncrypt()
    e.Date = "01012016"
    e.Key = "12345"
    secret := e.Encrypt("a")
    if secret != "q" {
        t.Log("Encrypt didn't work properly on one letter: ", secret)
        t.Fail()
    }
}

func TestEncryptOnMultipleLetters(t *testing.T)  {
    e := NewEncrypt()
    e.Date = "01012016"
    e.Key = "12345"
    secret := e.Encrypt("aa")
    if secret != "qz" {
        t.Log("Encrypt didn't work properly on two letters: ", secret)
        t.Fail()
    }
    secret = e.Encrypt("aaa")
    if secret != "qzn" {
        t.Log("Encrypt didn't work properly on three letters: ", secret)
        t.Fail()
    }
    secret = e.Encrypt("aaaa")
    if secret != "qznz" {
        t.Log("Encrypt didn't work properly on four letters: ", secret)
        t.Fail()
    }
    secret = e.Encrypt("aaaaaaaaoverandout")
    if secret != "qznzqznzeurqqmqnks" {
        t.Log("Encrypt didn't work properly on a sentence: ", secret)
        t.Fail()
    }
}

func TestDecrypt(t *testing.T)  {
    // t.Skip("Not ready yet")
    d := NewDecrypt()
    d.Date = "01012016"
    d.Key = "12345"
    secret := d.Decrypt("qz")
    if secret != "aa" {
        t.Log("Decrypt didn't work properly on two letters: ", secret)
        t.Fail()
    }
    secret = d.Decrypt("qzn")
    if secret != "aaa" {
        t.Log("Decrypt didn't work properly on three letters: ", secret)
        t.Fail()
    }
    secret = d.Decrypt("qznz")
    if secret != "aaaa" {
        t.Log("Decrypt didn't work properly on four letters: ", secret)
        t.Fail()
    }
    secret = d.Decrypt("qznzqznz")
    if secret != "aaaaaaaa" {
        t.Log("Decrypt didn't work properly on a sentence: ", secret)
        t.Fail()
    }

}

func TestCrackTheCode(t *testing.T)  {
    t.Skip("never got around to it")
    key := Crack(encodedText)
    encodedText := "qznzqznzeurqqmqnks"
    d := NewDecrypt("")
    if key != "aaaaaaaaoverandout"
}
