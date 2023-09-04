In Cache
================================

```go
package main

import (
	"fmt"
	"github.com/nailrads/handler_cache"
)

func main() {
	cache := cache.New()

	cache.Set("userId", 42)
	userId := cache.Get("userId")

	fmt.Println("Step 1", userId)

	cache.Delete("userId")
	userId = cache.Get("userId")

	fmt.Println("Step 2", userId)
}

```