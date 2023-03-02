# [1680. Concatenation of Consecutive Binary Numbers](https://leetcode.com/problems/concatenation-of-consecutive-binary-numbers)

Given an integer `n`, return the **decimal value** of the binary string formed by concatenating the binary representations of `1` to `n` in order, **modulo** 10^9 + 7.

| Decimal | Binary | Len |
|--------:|:------:|:---:|
|       1 | 000001 |  1  |
|       2 | 000010 |  2  |
|       3 | 000011 |  2  |
|       4 | 000100 |  3  |
|       5 | 000101 |  3  |
|       6 | 000110 |  3  |
|       7 | 000111 |  3  |
|       8 | 001000 |  4  |
|       9 | 001001 |  4  |
|      10 | 001010 |  4  |
|      11 | 001011 |  4  |
|      12 | 001100 |  4  |

len = log<sub>2</sub>(n) + 1

**Example 1:**

```
Input: n = 1
Output: 1
Explanation: "1" in binary corresponds to the decimal value 1.
```

**Example 2:**

```
Input: n = 3
Output: 27
Explanation: In binary, 1, 2, and 3 corresponds to "1", "10", and "11".
After concatenating them, we have "11011", which corresponds to the decimal value 27.
```

**Example 2:**

```
Input: n = 12
Output: 505379714
Explanation: The concatenation results in "1101110010111011110001001101010111100".
The decimal value of that is 118505380540.
After modulo 109 + 7, the result is 505379714.
```

**Constraints:**

- `1 <= n <= 10^5`
