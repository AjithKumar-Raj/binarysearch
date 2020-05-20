Given an array of integers, return a new array where each element in the new array is the number of smaller elements to the right of that element in the original input array.

For example, given the array [3, 4, 9, 6, 1], return [1, 1, 2, 1, 0], since:

There is 1 smaller element to the right of 3
There is 1 smaller element to the right of 4
There are 2 smaller elements to the right of 9
There is 1 smaller element to the right of 6
There are no smaller elements to the right of 1
Bonus: Can you do this in O(n log n) time?

Example 1
Input

lst = [3, 4, 9, 6, 1]
Output

[1, 1, 2, 1, 0]
Example 2
Input

lst = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
Output

[0, 0, 0, 0, 0, 0, 0, 0, 0, 0]
Example 3
Input

lst = [9, 8, 7, 6, 5, 4, 3, 2, 1, 0]
Output

[9, 8, 7, 6, 5, 4, 3, 2, 1, 0]