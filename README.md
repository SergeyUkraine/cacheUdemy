cacheUdemy
================================

cacheUdemy helps you to use your memory for set/get/delete some data.

See it in action:

## Example
```go
package main

import (
	"fmt"
	"github.com/SergeyUkraine/cacheUdemy"
)

func main() {
	cache := cacheUdemy.NewCache()

	cache.SetKey("userId", 42)

	userId := cache.GetKey("userId")

	fmt.Println(userId)

	cache.Delete("userId")

	userId = cache.GetKey("userId")

	fmt.Println(userId)
}

```