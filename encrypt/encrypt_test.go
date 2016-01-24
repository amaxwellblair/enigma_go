package encrypt

import (
  "testing"
)

func TestCreateEncryptStruct(t *testing.T)  {
    var e interface{}
    e = New()
    switch e.(type) {
      case *Encrypt:

      default:
        t.Log("The encrypt struct isn't being created correctly")
        t.Fail()
    }
}

func TestZeroRotatingOnAGivenIndex(t *testing.T)  {
    e := New()
    e.Key = "12345"
    rotation := e.Rotate(0)
    if rotation != 12 {
      t.Log("The 0 rotation didn't yield the proper number")
      t.Fail()
    }
}

func TestOneRotatingOnAGivenIndex(t *testing.T)  {
    e := New()
    e.Key = "12345"
    rotation := e.Rotate(1)
    if rotation != 23 {
      t.Log("The 1 rotation didn't yield the proper number")
      t.Fail()
    }
}

func TestArotationIsOutOfRange(t *testing.T) {
    e := New()
    e.Key = "12345"
    rotation := e.Rotate(4)
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
    e := New()
    e.Date = "01012016"
    lastFourDigits := e.dateConv()
    if lastFourDigits != "4256" {
      t.Log("The date convertor isnt working quite right!", lastFourDigits)
      t.Fail()
    }
}

func TestAOffsetOnAGivenIndex(t *testing.T) {
    e := New()
    e.Date = "01012016"
    offset := e.Offset(0)
    if offset != 4 {
      t.Log("Offset doesn't work on the 0 index given")
      t.Fail()
    }
}

func TestOffsetCanNotGoOutOfBoundsOnItsIndex(t *testing.T) {
    e := New()
    e.Date = "01012016"
    offset := e.Offset(5)
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
        t.Log("Modulo does not work on numbers with in the first cycle through")
        t.Fail()
    }
}
