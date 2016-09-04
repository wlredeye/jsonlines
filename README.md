Simple JSON Lines encoding/decoding package for Go.

Using:
```go
package main

import (
	"jsonlines"
	"strings"
	"fmt"
)

func main() {
	data := `{"Name": "Bob", "Age": 20, "Cars": ["Ford", "Dodge"]}
	{"Name": "John", "Age": 30, "Cars": ["BMW", "Toyota"]}`
	type Person struct {
		Name string
		Age int64
		Cars []string
	}
	people := []Person{}
	err := jsonlines.Decode(strings.NewReader(data), &people)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(people)
}
```
Output:
```
[{Bob 20 [Ford Dodge]} {John 30 [BMW Toyota]}]
```
