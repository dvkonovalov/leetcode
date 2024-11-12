package hashmap

import "sort"

func groupAnagrams(strs []string) [][]string {
	mp := make(map[string][]string, 0)
	for _, str := range strs {
		sortStr := sortStrings(str)
		if _, ok := mp[sortStr]; !ok {
			mp[sortStr] = []string{str}
		} else {
			mp[sortStr] = append(mp[sortStr], str)
		}
	}
	ans := make([][]string, 0, len(mp))
	for _, mas := range mp {
		ans = append(ans, mas)
	}
	return ans
}

func sortStrings(str string) string {
	characters := []rune(str)
	sort.Slice(characters, func(i, j int) bool {
		return characters[i] < characters[j]
	})
	return string(characters)
}
