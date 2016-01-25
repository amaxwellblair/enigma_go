package encrypt

import (
  "strconv"
)

type Encrypt struct {
    Key string
    Date string
}

func New() (*Encrypt) {
    return new(Encrypt)
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

func modulo(movement int, bottom int, top int) (int) {
    difference := top - bottom + 1
    division := movement / difference
    remainder := movement - division*difference
    return bottom + remainder
}

func (e *Encrypt) Rotate(number int) (tumbler int) {
    chars := []rune(e.Key)
    if (number > (len(chars)-2)) {
        return
    }
    tens, _ := strconv.Atoi(string(chars[number]))
    ones, _ := strconv.Atoi(string(chars[number+1]))
    tumbler = tens*10 + ones
    return
}

func (e *Encrypt) dateConv()(dateNums string) {
    dateInt, _ := strconv.Atoi(e.Date)
    dateRune := []rune(strconv.Itoa(pow(dateInt,2)))
    runeLength := len(dateRune)
    dateNums = string(dateRune[(runeLength-4):])
    return
}

func (e *Encrypt) Offset(index int) (tumbler int) {
    encodedDate := e.dateConv()
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
        totalTumble = e.Offset(index) + e.Rotate(index) + int(rawRunes[i]) - int('a')
        rawRunes[i] = rune(modulo(totalTumble, 97, 122))
    }
    return string(rawRunes)
}
