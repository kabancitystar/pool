package counter

type Counter struct {
	A int
	B int
}

func NewCounter() *Counter {
	p := Counter{0, 0}
	return &p
}

func Inc(c *Counter) {
	c.A++
	c.B++
}
