package hashmap

func RunTask() {
	// Task - Ransom Note
	// https://leetcode.com/problems/ransom-note/?envType=study-plan-v2&envId=top-interview-150
	ransomNote := "a"
	magazine := "b"
	canConstruct(ransomNote, magazine)

	// Isomorphic Strings
	// https://leetcode.com/problems/isomorphic-strings/description/?envType=study-plan-v2&envId=top-interview-150
	s := "egg"
	t := "add"
	isIsomorphic(s, t)

	// Two sum
	// https://leetcode.com/problems/two-sum/description/?envType=study-plan-v2&envId=top-interview-150
	nums := []int{2, 7, 11, 15}
	target := 9
	twoSum(nums, target)

	isHappy(3)

	groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
}
