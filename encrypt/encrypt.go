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
    movement -= 1
    difference := top - bottom
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
