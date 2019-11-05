TS4 Convert to Non-Decreasing Array with Minimum Changes
=======================
#### Problem description
Given an unsorted integer array, e.g., `[1, 3, 2, 5, 4]`, we are allowed to change the array by performing `+1` or `-1` operations to any number. Find the minimum changes needed to make the array non-decreasing.

#### Difficulty
N/A

#### Keywords
- [Dynamic programming](../categories/dp.md)

#### Idea


#### Complexity
- Runtime: O(N) 
- Space: O(N)
  
#### LC performance
- Runtime: N/A
- Memory usage: N/A

#### Code
```python
class Solution:

  def minChangesToNonDecreasingArray(arr):
    # if the array size is smaller than 2, no changes are needed
    if len(arr) < 2:
        return 0
    # find all the unique numbers
    vals = sorted(set(arr))
    # the DP keeps track of the minimum changes needed to make the array arr[:i]
    # non-decreasing by changing a number to another 
    dp = {v: abs(v - arr[0]) for v in vals}
    for i in range(1, len(arr)):
        min_changes, tmp = float('inf'), {}
        for v in vals:
            # minimum changes needed to make the subarray ending before the 
            # current position non-decreasing to change the values to the 
            # current value
            min_changes = min(min_changes, dp[v])
            # the cost for changing the current number to make the subarray 
            # ending here non-decreasing
            tmp[v] = min_changes + abs(arr[i] - v)
        dp = tmp
    return min(dp.values())
```