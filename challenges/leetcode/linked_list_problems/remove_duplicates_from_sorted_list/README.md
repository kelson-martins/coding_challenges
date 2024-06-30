# Problem 

Given the `head` of a sorted linked list, delete all duplicates such that each element appears only once. Return the linked list **sorted** as well. [Problem Link](https://leetcode.com/problems/remove-duplicates-from-sorted-list/)
 
Example 1:
```
Input: head = [1,1,2]
Output: [1,2]
```

Example 2:
```
Input: head = [1,1,2,3,3]
Output: [1,2,3]
```

# Solution Notes

The propose solution runs on runtime O(n) by making use of a map that serves as a buffer for identifying items already on the linked list.