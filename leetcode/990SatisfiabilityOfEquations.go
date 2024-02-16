package main

var N = 26

func find(p [26]int, x int) int {
	if x != p[x] {
		p[x] = find(p, p[x])
	}
	return p[x]
}

func equationsPossible(equations []string) bool {
	var p [26]int
	for i := 0; i < 26; i++ {
		p[i] = i
	}

	for _, v := range equations {
		a := int(v[0] - 'a')
		b := int(v[3] - 'a')
		if v[1] == '=' {
			p[find(p, a)] = find(p, b)
		}
	}
	for _, v := range equations {
		a := int(v[0] - 'a')
		b := int(v[3] - 'a')
		if v[1] == '!' {
			if find(p, a) == find(p, b) {
				return false
			}
		}

	}
	return true
}
