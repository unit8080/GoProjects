// Given an ascending order linked list with integers and an array with random numbers, write a function to add the first number from array into the LL and delete the next number from the LL. Continue till all the array numbers are done.

// Ex: |2| --> |10| --> |16| --> |25| --> NULL
// Arr: [1, 16, 4, 25]
// Result: |1| --> |2| --> |4| --> |10| --> NULL




type Node struct {
    data int
    next *Node
}

var arr []int 

func insertDelete(head *Node, arr []int) *Node  {
    
    size := len(arr)
    if size == 0 {
        return head
    }
    
   
    for i := 0; i < size; i++ {
        elem := arr[i]

         curr := head
         prev := nil
         
        // even or odd operation
        if (i % 2 == 0). {
            node := new(Node)
            node.data = elem
            if curr == nil {
                head = node
            } else {
                // add in sorted order
                for curr != nil {
                    if curr.data < elem {
                        prev = curr
                        curr = curr.next
                    }
                }
                prev.next = node
                node.next = curr
            }
        } else {
            if curr != nil {
                for curr != nil {
                    found := false
                    // find the matching elem
                    if curr.data == elem {
                        found = true
                        break
                    }
                    
                    prev = curr
                    curr = curr.next
                }
                if found {
                    
                    if head == curr {
                        head = curr.next
                    }
                    // head ptr also ??
                    if curr != nil {
                        prev.next = curr.next
                        // curr will be GC.
                    } else {
                        // element is not found
                    }
                }
            } 
            
        }
    }
    return head
    //Complexity: 
    //TC : = O(K * N)
}


