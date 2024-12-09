package strings

type UniqueSlice struct {
	known map[string]bool
	items []string
}

type IUniqueSlice interface {
	Contains(item string) bool
	Append(items ...string) *UniqueSlice
	Items() []string
	Len() int
}

var _ IUniqueSlice = (*UniqueSlice)(nil)

func NewUniqueSlice(items ...string) *UniqueSlice {
	return NewUniqueSliceCap(len(items)).Append(items...)
}

func NewUniqueSliceCap(size int) *UniqueSlice {
	return &UniqueSlice{
		known: make(map[string]bool, size),
		items: make([]string, 0, size),
	}
}

func (u *UniqueSlice) Contains(item string) bool {
	return u.known[item]
}

func (u *UniqueSlice) Append(items ...string) *UniqueSlice {
	for _, s := range items {
		if u.known[s] {
			continue
		}

		u.known[s] = true
		u.items = append(u.items, s)
	}

	return u
}

func (u UniqueSlice) Items() []string {
	return u.items
}

func (u UniqueSlice) Len() int {
	return len(u.items)
}
