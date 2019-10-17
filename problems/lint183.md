Lint183 Wood Cut 
=======================
[Problem description](https://www.lintcode.com/problem/wood-cut/description)

#### Difficulty
<span style="color:red">Hard</span>

#### Keywords
- [Binary search](../categories/binary_search.md)
  
#### Idea


#### Complexity
- Runtime: O(NlogN)
- Space: O(1)
  
#### LC performance
- Runtime: 201 ms
- Memory usage: 

#### Code
```python
class Solution:
    """
    @param L: Given n pieces of wood with length L[i]
    @param k: An integer
    @return: The maximum length of the small pieces
    """
    def woodCut(self, L, k):
        sum_ = sum(L)
        if sum_ < k:
            return 0
        
        lo, hi = 0, sum_//k
        while lo + 1 < hi:
            mid = (lo + hi)//2
            n_cut = sum(l//mid for l in L)
            if n_cut >= k:
                lo = mid
            else:
                hi = mid 
        return lo if sum(l//hi for l in L) < k else hi
```