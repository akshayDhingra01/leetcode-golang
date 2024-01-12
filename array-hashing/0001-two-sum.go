// https://leetcode.com/problems/two-sum/description/

// Optimized Solution || Time O(N) || Space O(N) || Hash Map

func twoSum(nums []int, target int) []int {

    hash := make(map[int] int) // hash map of Value to obtain and it's index
    for index, num := range nums {
        get, ok := hash[num]
        if ok {
            ans := []int {get, index}
            return ans
        } else {
            hash[target - num] = index
        }
    }
    return []int{-1,-1}
}