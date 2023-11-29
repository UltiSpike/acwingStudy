package main

func longestCommonPrefix(strs []string) string {
	var son [200][26]int
	var idx int
	var ans = strs[0]
	var p string
	for i := 0; i < len(strs); i++ {
		p = query(&son, strs[0])
		insert(&son, idx, strs[0])
		if len(p) < len(ans) {
			ans = p
		}
	}
	return ""
}

func insert(son *[200][26]int, idx int, str string) {
	var p int
	var u int
	for i := 0; i < len(str); i++ {
		u = int(str[i] - 'a')
		if (*son)[p][u] == 0 {
			idx++
			(*son)[p][u] = idx
		}
		p = (*son)[p][u]
	}

	return
}

func query(son *[200][26]int, str string) string {
	var p int
	var u int
	var temp string
	for i := 0; i < len(str); i++ {
		u = int(str[i] - 'a')
		if (*son)[p][u] == 0 {
			return temp
		}
		temp += string(str[i])
		p = (*son)[p][u]
	}

	return temp
}
