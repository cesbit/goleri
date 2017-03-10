package goleri

const modeRequired = 0
const modeOptional = 1

type expecting struct {
	required []Element
	optional []Element
	pos      int
	modes    map[int]uint8
}

func newExpecting() *expecting {
	modes := make(map[int]uint8)
	modes[0] = modeRequired
	return &expecting{
		pos:   0,
		modes: modes,
	}
}

func (e *expecting) empty() {
	e.required = []Element{}
	e.optional = []Element{}
}

func (e *expecting) update(elem Element, pos int) {
	if pos > e.pos {
		e.empty()
		e.pos = pos
	}
	if pos == e.pos {
		if e.modes[pos] == modeRequired {
			e.required = appendIfMissing(e.required, elem)

		} else {
			e.optional = appendIfMissing(e.optional, elem)
		}
	}
}

func (e *expecting) setMode(pos int, mode uint8) {
	// do nothing when mode is already set to optional
	if m, ok := e.modes[pos]; ok && m == modeOptional {
		return
	}
	e.modes[pos] = mode
}

func (e *expecting) getExpecting() []Element {
	if e.optional != nil {
		for _, elem := range e.optional {
			e.required = appendIfMissing(e.required, elem)
		}
		e.optional = nil
	}
	return e.required
}

func appendIfMissing(slice []Element, elem Element) []Element {
	for _, e := range slice {
		if e == elem {
			return slice
		}
	}
	return append(slice, elem)
}
