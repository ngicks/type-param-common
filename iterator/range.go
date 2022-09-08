package iterator

// Range is numeric range iterator.
// Range is [start, end) if it has end, else, start to max of int.
type Range struct {
	start int
	end   int
}

func NewRange(start, end int) *Range {
	return &Range{
		start: start,
		end:   end,
	}
}

func NewInfiniteRange(start int) *Range {
	return &Range{
		start: start,
	}
}

func (r *Range) Next() (next int, ok bool) {
	if r.start >= r.end {
		return 0, false
	}
	next = r.start
	r.start++
	return next, true
}

func (r *Range) NextBack() (next int, ok bool) {
	if r.start >= r.end {
		return 0, false
	}
	next = r.end - 1
	r.end--
	return next, true
}

func (r *Range) ToIterator() Iterator[int] {
	return Iterator[int]{
		SeIterator: r,
	}
}
