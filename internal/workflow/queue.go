package workflow

type Queue struct {
	items []Event
}

func (q *Queue) Push(e Event) {
	q.items = append(q.items, e)
}

func (q *Queue) Pop() (Event, bool) {
	if len(q.items) == 0 {
		return Event{}, false
	}
	e := q.items[0]
	q.items = q.items[1:]
	return e, true
}
