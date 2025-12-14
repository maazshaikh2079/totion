package util

type Item struct {
	Title_ string
	Desc   string
}

func (i Item) Title() string       { return i.Title_ }
func (i Item) Description() string { return i.Desc }
func (i Item) FilterValue() string { return i.Title_ }
