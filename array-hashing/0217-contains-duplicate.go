// https://leetcode.com/problems/contains-duplicate/description/

// Optimized Solution || Time O(N) || Space O(N) || Hash Map

func containsDuplicate(nums []int) bool {
    hash := make(map[int]int)

    for _, num := range nums {
        _, ok := hash[num]
        if ok {
            return true
        }
        hash[num] = 1
    }
    return false
}