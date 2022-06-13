#### First Bad Version

1. [The Problem](#the-problem)
2. [Implementaiton Notes](#implementation-notes)


**Problem link** [here](https://leetcode.com/problems/two-sum/)

**Algorithm Type:** `2 Pointers`

Given an array of integers `nums` and an integer `target`, return indices of the two numbers such that they add up to `target`.

You may assume that each input would have **exactly one solution**, and you may not use the same element twice.

You can return the answer in any order.

**Challenge:** Can you come up with an algorithm that is less than O(n2) time complexity?

 

**Example 1:**

```
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
```

**Example 2:**
```
Input: nums = [3,2,4], target = 6
Output: [1,2]
```

**Example 3:**
```
Input: nums = [3,3], target = 6
Output: [0,1]
``` 

**Constraints:**

* `2 <= nums.length <= 104`
* `-109 <= nums[i] <= 109`
* `-109 <= target <= 109`

**Only one valid answer exists.**

### Implementation Notes

Implemented the twosums problem using 2 different approaches:
1.  One less efficient with complexity O(n log n), in which we are first sortering the list first, then always summing the first and last elements to identify if we need to reduce or increase the sum to hit the traget.


2. A more efficient approach in O(n) by laveraging maps.