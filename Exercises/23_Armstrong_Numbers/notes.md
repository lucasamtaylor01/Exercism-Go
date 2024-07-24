
# Notes about Armstrong Numbers Armstrong Numbers Exercise  

My main difficulty was with this line:

```go
    for _, digit := range digits
```

I asked ChatGPT because my loop was getting too complicated. I asked for a simplification, and it provided this loop that I have seen a few times but never understood very well. Let's analyze it:

1. `_, digit` is the part that defines the variables that will be used in the loop.

2. `_` is used to ignore the index value. If we don't ignore it, the code would be like this:
    ```go
    for index, value := range slice {
        // loop
    }
    ```

3. `digit` stores the value of each element of the slice that we defined before.

4. `range digits` returns the index and the value of the element at the current position. So, the expression `range digits` will return each index and value of the `digits` slice. 

## Example

```go
package main

import "fmt"

func main() {
    digits := []int{1, 2, 3}
    
    for _, digit := range digits {
        fmt.Println(digit)
    }
}
```
