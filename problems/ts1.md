TS1 Non-repeating random number generator
=======================
#### Problem description
Design a random number generator that generates non-repeating random numbers in the given range. It implements the following functions:

- `next()` returns a non-repeating random number
- `update(lo, hi)` updates the range from which the random numbers are picked

The random number generator takes a range between `lo` and `hi` inclusively for initialization. 

#### Difficulty
N/A

#### Keywords
- [Random](../categories/random.md)

#### Idea


#### Complexity
- Runtime: O(N) for initilization and update, O(1) for pick random number 
- Space: O(N)
  
#### LC performance
- Runtime: N/A
- Memory usage: N/A

#### Code
```python
import random 

class RandomNumberGenerator:
  
  def __init__(self, lo, hi):
    # array of numbers in the given range
    self.nums = list(range(lo, hi + 1))
    # right bound of numbers unpicked in the array
    self.right = len(self.nums) - 1

  def next(self):
    # if the right bound is negative, that means all the numbers in the range 
    # have been picked, thus throw an exception 
    if self.right < 0:
      raise Exception('All numbers in the given range are picked')
    # Fisher-Yates shuffling
    nums = self.nums
    # pick a random index from the range of unpikced numbers, which is the next 
    # random number to return
    idx = random.randint(0, self.right)
    n = nums[idx]
    # swap the number with the unpicked number at the right boundary, and move 
    # the right bound to left by 1, excluding the number just picked from the 
    # unpicked pool
    nums[idx], nums[self.right] = nums[self.right], nums[idx]
    self.right -= 1
    return n

  def update(self, lo, hi):
      # collect all the picked random number in the new range into a set
      picked = set([n for n in self.nums[self.right + 1:] if lo <= n <= hi])
      # create an array of unpicked numbers in the given range
      cands = [n for n in range(lo, hi + 1) if n not in picked]
      # concatenate the two pools as the full array of the new range, wherein 
      # picked and unpicked numbers are on the right and left, respectively
      self.nums, self.right = cands + list(picked), len(cands) - 1
```