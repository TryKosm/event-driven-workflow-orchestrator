package workflow

type DeadLetter struct {
	items []Event
}

func (d *DeadLetter) Add(e Event) {
	d.items = append(d.items, e)
}
