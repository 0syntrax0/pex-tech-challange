package fibonacci

type Client struct {
	Previous uint
	Current  uint
	Next     uint
}

// creates new client
func New() *Client {
	return &Client{
		Current: 1,
		Next:    1,
	}
}

// returns current number in current sequence
// and calculates next set
func (f *Client) GetCurrent() uint {
	current := f.Current
	f.calculate()
	return current
}

// returns previous number in current sequence
func (f *Client) GetPrevious() uint {
	prev := f.Previous
	return prev
}

// returns next number in current sequence
func (f *Client) GetNext() uint {
	next := f.Next
	return next
}

// calculates next numbers in the sequence
func (f *Client) calculate() {
	f.Previous = f.Current
	f.Current = f.Next
	f.Next = f.Current + f.Previous
}
