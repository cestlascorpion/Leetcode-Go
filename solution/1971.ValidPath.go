package solution

func ValidPath(n int, edges [][]int, start int, end int) bool {
	if start == end {
		return true
	}

	parents := make([]int, n)
	for i := range parents {
		parents[i] = i
	}

	var find func(x int) int
	find = func(x int) int {
		if parents[x] == x { // be careful
			return x
		}
		return find(parents[x])
	}

	for _, e := range edges {
		a := find(e[0])
		b := find(e[1])
		if a == b {
			continue
		}
		if a < b {
			parents[b] = a
		} else {
			parents[a] = b
		}
	}

	return find(start) == find(end)
}

/**
graph: a-b a-e b-d a-c c-d

parents a b c d e
        a b c d e

a-b ->
parents[a] < parents[b] -> parents[b] = parents[a]
        a b c d e
        a a c d e

a-e ->
parents[a] < parents[e] -> parents[e] = parents[a]
        a b c d e
        a a c e a

b-d ->
parents[b] < parents[d] -> parents[d] = parents[b]
        a b c d e
        a a c a a

a-c ->
parents[a] < parents[c] -> parents[c] = parents[a]
        a b c d e
        a a a a a

c-d ->
parents[c] = parents[d] do nothing

note: `parents[x] value` IS NOT `the smallest node index with connection`
       IN JUST ONE LOOP OF ALL EDGES
be careful with the recursive termination condition

 */