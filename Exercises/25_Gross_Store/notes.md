# Comments about exercise

## A more concise version of the function `RemoveItem()`

My first attempt was:

```go
// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
 bill_value, bill_exists := bill[item]
 unit_value, unit_exists := units[unit]
    
    if !bill_exists || bill_value == 0 {
        return false
    }else if !unit_exists || unit_value == 0{
        return false
    }

    if (bill[item] - unit_value) < 0 {
        return false
    }else if (bill[item]- unit_value) == 0{
        delete(bill, item)
        return true
    }else{
     bill[item] -= unit_value
        return true
    }  
}
```

I was bothered by the number of `if`, `else if` and `else` and the number of `return`. So I asked to CHATGPT for a sugestion and the final version was:

```go
func RemoveItem(bill, units map[string]int, item, unit string) bool {
 billValue, billExists := bill[item]
 unitValue, unitExists := units[unit]
    
 if !billExists || !unitExists || billValue < unitValue {
  return false
 }
    
 if billValue == unitValue {
  delete(bill, item)
 } else {
  bill[item] -= unitValue
 }
    
 return true
}
```

**The differences:** *Write after*

## Camel case and Snake case

Today, I've learned Camel Case is default on Go.
