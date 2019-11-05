TS5 Token Bucket Algorithm
=======================
#### Problem description
Implement a token bucket with given rate and capacity

#### Difficulty
N/A

#### Keywords


#### Idea


#### Complexity
- Runtime: O(N) 
- Space: O(N)
  
#### LC performance
- Runtime: N/A
- Memory usage: N/A

#### Code
```python
import time

from multiprocessing import Process, Queue

class TokenBucket(Process):

  def __init__(self, rate, capacity):
    super(Process, self).__init__()
    self.rate = rate*.001
    self.q = Queue(maxsize=capacity)

  def run(self):
    while True:
      self.q.put(1, block=False)
      time.sleep(self.rate)
  
  def fetchToken(self, n):
    return sum([self.q.get() for _ in range(n) if not self.q.empty()])

```