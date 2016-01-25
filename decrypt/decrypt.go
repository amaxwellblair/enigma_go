package decrypt

type Decrypt struct {
    Key string
    Date string
}

func New() (*Decrypt)  {
    return new(Decrypt)
}
