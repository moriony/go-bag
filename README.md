# Bag

Container for concurrent access to scalar data

## Example

```golang
package main

import (
	"fmt"
	"time"

	"github.com/moriony/go-bag"
)

func main() {
	b := bag.New()
	b.Strings().Set("time", time.Now().String())
	t, err := b.Strings().Get("time")
	fmt.Println(t, err)
}
```
