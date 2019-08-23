package lunopher

type Lunopher struct {
	owner string
}

func NewLunopher() Lunopher {
	return Lunopher{"Dylan Meeus"}
}

func (l Lunopher) Owner() string {
	return l.owner
}
