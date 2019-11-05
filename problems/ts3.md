TS3 Generate Sublists of A Circular Linked List
=======================
#### Problem description
Given a circular linked list, generate all its sublists. For example, given `1 -> 2 -> 3 -> 1`. return `['1', '1 -> 2', '1 -> 2 -> 3', '2', '2 -> 3', '2 -> 3 -> 1', '3', '3 -> 1', '3 -> 1 -> 2']`

#### Difficulty
N/A

#### Keywords
- [Linked list](../categories/linked_list.md)

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

  def generateSublists(self, head): 
    res = []
    # in case of an empty list
    if not head:
      return res

    def copyList(start, end):
      """
      Function for copying a sublist given the start and end
      """
      p = head = ListNode(start.val)
      cur = start.next
      while cur and cur != end.next:
        p.next = ListNode(cur.val)
        p = p.next
        cur = cur.next
      return head
    
    # find the length of the circular linked list
    len_, p = 0, head
    while len_ == 0 or p != head:
      len_ += 1
      p = p.next
    
    for i in range(len_):
      # find the start of the sublist 
      start = head
      for _ in range(i):
        start = start.next
      # find the end of the sublist
      end = start
      for j in range(len_):
        # copy the sublist every time when the end moves forward until a full 
        # circular list with the current start has been copied
        res.append(copyList(start, end))
        end = end.next
    
    return res
```