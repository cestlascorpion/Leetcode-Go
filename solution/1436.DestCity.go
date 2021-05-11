package solution

func DestCity(paths [][]string) string {
	dict := make(map[string]string)
	for idx := range paths {
		dict[paths[idx][0]] = paths[idx][1]
	}

	dest := paths[0][0]
	for {
		v, ok := dict[dest]
		if !ok {
			break
		}
		dest = v
	}
	return dest
}

func DestCity2(paths [][]string) string {
	dict := make(map[string]struct{})
	for idx := range paths {
		dict[paths[idx][0]] = struct{}{}
	}
	for _, path := range paths {
		if _, ok := dict[path[1]]; !ok {
			return path[1]
		}
	}
	return ""
}
