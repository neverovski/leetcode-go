# [513. Find Bottom Left Tree Value](https://leetcode.com/problems/find-bottom-left-tree-value)

Given the root of a binary tree, return the leftmost value in the last row of the tree.

**Example 1:**

```
Input: root = [2,1,3]
Output: 1
```

**Example 2:**

```
Input: root = [1,2,3,4,null,5,6,null,null,7]
Output: 7
```

**Example 3:**

```
Input: root = [1,2,3,4,null,5,6,null,null]
Output: 4
```

**Example 4:**

```
Input: root = [2,1,3,null,5,null,6,null,8]
Output: 8
```

**Example 5:**

```
Input: root = [0,null,-1]
Output: -1
```

**Example 6:**

```
Input: root = [0,-1,null]
Output: -1
```

**Constraints:**

- The number of nodes in the tree is in the range `[1, 10^4]`.
- `-2^31 <= Node.val <= 2^31 - 1`
