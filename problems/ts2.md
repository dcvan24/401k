TS2 Convert Array to Circular Linked List
=======================
#### Problem description
Convert an array to a circular linked list. For example, given `[1, 2, 3]`, convert it to `1 -> 2 -> 3 -> 1`. 

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

  def convertArrayToCircularLinkedList(self, arr):
    if not arr:
      return None
    p = head = ListNode(arr[0])
    for i in range(1, len(arr)):
      p.next = ListNode(arr[i])    
      p = p.next
    p.next = head
    return head
```