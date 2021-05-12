package solution

type OrderedStream struct {
	ptr     int
	content []string
}

func ConstructorOrderedStream(n int) OrderedStream {
	return OrderedStream{0, make([]string, n)}
}

func (o *OrderedStream) Insert(idKey int, value string) []string {
	o.content[idKey-1] = value
	res := make([]string, 0)
	for o.ptr < len(o.content) && len(o.content[o.ptr]) > 0 {
		res = append(res, o.content[o.ptr])
		o.content[o.ptr] = ""
		o.ptr++
	}
	return res
}
