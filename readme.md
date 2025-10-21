# Simple Go `Values` Type Example

This example demonstrates a custom Go type, `Values`, which is an alias for `map[string][]string`. This structure is commonly used to manage data where a single key can map to multiple string values (like HTTP query parameters).

It also serves as a crucial illustration of **Go's nil map behavior**, specifically the difference between reading from and writing to a `nil` map.

---

## Code

The code defines the `Values` type and two methods, `Get` and `Add`.

```go
package main

import "fmt"

type Values map[string][]string

// Get returns the first value associated with the given key.
// If the key is not present or the map is nil, it returns an empty string.
func (v Values) Get(key string) string {
	// Reading from a nil map (v[key]) is safe and returns the zero value ([]string in this case).
	if vs := v[key]; len(vs) > 0 {
		return vs[0]
	}
	return ""
}

// Add appends a value to the list associated with the given key.
// NOTE: This method will cause a **runtime panic** if the map receiver 'v' is nil.
func (v Values) Add(key, value string) {
	// Attempting to write/assign to a nil map panics.
	v[key] = append(v[key], value)
}

func main() {
	// 1. Initialization and basic usage
	m := Values{"lang": {"en"}}

	fmt.Println(m["item"])     
	fmt.Println(m.Get("lang")) 

	// 2. Adding multiple values
	m.Add("item", "1")
	m.Add("item", "2")

	fmt.Println(m["item"])     
	fmt.Println(m.Get("item")) 

	// 3. Demonstrating nil map behavior for Get (SAFE)
	m = nil
	fmt.Println(m.Get("item")) 

	// 4. Demonstrating nil map behavior for Add (PANIC)
	m.Add("item", "3")         
}

