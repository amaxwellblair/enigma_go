package encrypt

import (
  "strconv"
)

type Encrypt struct {
    Key string
    Date string
}

func (e *Encrypt) GetKey()(string)  {
    return e.Key
}


func (e *Encrypt) GetDate()(string)  {
    return e.Date
}

type Decrypt struct {
    Key string
    Date string
}

func (d *Decrypt) GetKey()(string)  {
    return d.Key
}


func (d *Decrypt) GetDate()(string)  {
    return d.Date
}

type Cryptic interface {
    GetKey() string
    GetDate() string
}

func NewEncrypt() (*Encrypt) {
    return new(Encrypt)
}


func NewDecrypt() (*Decrypt) {
    return new(Decrypt)
}

func pow(base int, power int) (value int) {
    for i := 0; i < power; i++ {
        if i == 0 {
            value = base
        } else {
            value *= base
        }
    }
    return
}

func abs(number int) int  {
    if number < 0 {
        return -number
    } else {
        return number
    }
}

func modulo(movement int, bottom int, top int) (int) {
    if movement >= 0 {
        difference := top - bottom + 1
        division := movement / difference
        remainder := movement - division*difference
        return bottom + remainder
    } else {
        difference := top - bottom + 1
        division := abs(movement) / difference
        remainder := division*difference + movement
        return modulo(remainder + difference, bottom, top)
    }
}

func Rotate(c Cryptic, number int) (tumbler int) {
    chars := []rune(c.GetKey())
    if (number > (len(chars)-2)) {
        return
    }
    tens, _ := strconv.Atoi(string(chars[number]))
    ones, _ := strconv.Atoi(string(chars[number+1]))
    tumbler = tens*10 + ones
    return
}

func dateConv(c Cryptic)(dateNums string) {
    dateInt, _ := strconv.Atoi(c.GetDate())
    dateRune := []rune(strconv.Itoa(pow(dateInt,2)))
    runeLength := len(dateRune)
    dateNums = string(dateRune[(runeLength-4):])
    return
}

func Offset(c Cryptic, index int) (tumbler int) {
    encodedDate := dateConv(c)
    dateRune := []rune(encodedDate)
    if (len(dateRune)-1) < index {
        return
    }
    tumbler, _ = strconv.Atoi(string(dateRune[index]))
    return
}

func (e *Encrypt) Encrypt(rawText string) (string) {
    rawRunes := []rune(rawText)
    var totalTumble int
    var index int
    for i := 0; i < len(rawRunes); i++ {
        index = modulo(i, 0, 3)
        totalTumble = Offset(e, index) + Rotate(e, index) + int(rawRunes[i]) - int('a')
        rawRunes[i] = rune(modulo(totalTumble, 97, 122))
    }
    return string(rawRunes)
}

func (d *Decrypt) Decrypt(rawText string) (string) {
    rawRunes := []rune(rawText)
    var totalTumble int
    var index int
    for i := 0; i < len(rawRunes); i++ {
        index = modulo(i, 0, 3)
        totalTumble = -Offset(d, index) - Rotate(d, index) + int(rawRunes[i]) - int('a')
        rawRunes[i] = rune(modulo(totalTumble, 97, 122))
    }
    return string(rawRunes)
}
