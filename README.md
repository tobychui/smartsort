# smartsort
A smart sorting algorithm for Go to sort filename containing digits that is not zero padded

## Usage
```
import "github.com/tobychui/smartsort/sorter"
sortedSlice := sorter.SortString(stringSlice)
```

## Examples
```
original := []string{"a1c.o", "A1c.o", "a2016c.o", "a10c.o", "a2c.o", "A2c.o", "a11c.o"}
fmt.Println("Before: ", original)
results := sorter.SortString(original)
fmt.Println("After: ", results)
```

#### Outputs

```
Before:  [a1c.o A1c.o a2016c.o a10c.o a2c.o A2c.o a11c.o]
After:  [A1c.o A2c.o a1c.o a2c.o a10c.o a11c.o a2016c.o]
```

