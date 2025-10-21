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
```

# Simple Code Example for methods

This Go program demonstrates how to calculate distances between points 
and the total length (perimeter) of a path made up of multiple points.

- **Point struct**: Represents a coordinate (X, Y).
- **Path type**: A slice of points forming a connected path.
- **Distance() function & method**: Calculate straight-line distance between two points.
- **Path.Distance()**: Sums distances between consecutive points to get total path length.
- **main()**:
  - Calculates distance between two points (1,2) and (4,6).
  - Computes total perimeter of a rectangular path.
  - 
```go

// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"math"
)

// Point represents a 2D point with X and Y coordinates
type Point struct {
	X, Y float64
}

// Path represents a sequence (slice) of Points
type Path []Point

// Distance (function) — calculates distance between two Points
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Distance (method) — same as above, but defined as a method on Point
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Distance (method on Path) — computes total distance along a path
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func main() {
	p := Point{1, 2}
	q := Point{4, 6}

	// Both calls compute the same distance using different styles
	fmt.Println(Distance(p, q)) // function style
	fmt.Println(p.Distance(q))  // method style

	// Define a path and compute its total distance (perimeter)
	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}

	fmt.Println(perim.Distance()) // total distance of the path
	p.ScaleBy(2)
	fmt.Println(p)
}


