package registry

type Registry map[int]func() error

func NewRegistry() Registry {
	return make(Registry)
}

func (r Registry) Add(idx int, fun func() error) {
	r[idx] = fun
}

func (r Registry) Get(idx int) (func() error, bool) {
	fn, ok := r[idx]
	if !ok {
		return nil, false
	}

	return fn, true
}
