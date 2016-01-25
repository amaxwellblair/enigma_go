package keygenerate

import "math/rand"

type Key struct {
    Code string
}

func New() (k *Key) {
    k = new(Key)
    k.Code = createRandom(5)
    return
}

func createRandom(length int) string  {
    rand.Seed(1)
    code := []rune{}
    for i := 0; i < length; i++ {
        code = append(code, rune(rand.Intn(10)+48))
    }
    return string(code)
}
