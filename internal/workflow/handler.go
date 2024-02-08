package workflow

func Handle(e Event) bool {
	return e.Payload != ""
}
