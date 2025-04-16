##Problems (solved)

[13. Roman-to-Integer](https://leetcode.com/problems/roman-to-integer/description/)
`#bruteforce`

| Run      | Time | Memory   |
| -------- | ---- | -------- |
| v1 *GO*  | 4ms  | 4.88 MB  |
| *python* | 8ms  | 12.37 MB |

Complexity: **O(n)**, where ’n’ is the length of the input string.

Space Complexity: **O(1)**, because fixed dictionary is used to map Roman numerals to their corresponding integer values, constant amount of additional memory for variables.


[1. Two Sum](https://leetcode.com/problems/two-sum/description/)

| Run      | Time  | Memory   |
| -------- | ----- | -------- |
| v1 *GO*  | 59ms  | 5.38 MB  |
| v2 *GO*  | 60ms  | 5.50 MB  |
| v3 *GO*  | 5ms   | 7.18 MB  |

> [!CAUTION]
> Using `range` rathere than `for` loop is slower