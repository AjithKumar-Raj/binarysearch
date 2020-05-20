Given a list of integers, return whether the largest number is bigger than the second-largest number by more than two times.

For example, given the list [3, 9, 6], you should return false, since 9 is not bigger than 12 (2 times 6).

Given the list [6, 3, 15], you should return true, since 15 is bigger than 12 (2 times 6).

Example 1
Input

nums = [3, 6, 9]
Output

false
Explanation

9 is not bigger than 2 * 6.

Example 2
Input

nums = [3, 6, 15]
Output

true
Explanation

15 is bigger than 12 (2 * 6).

Example 3
Input

nums = [3, 6, 12]
Output

false
Explanation

12 is not bigger than 2 * 6, they're equal.

