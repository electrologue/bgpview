# BGPView API Go Client

A simple API Client written in Go for https://bgpview.io/

API Documentation: https://bgpview.docs.apiary.io

## Examples

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/electrologue/bgpview"
)

func main() {
	client := bgpview.NewClient()

	data, err := client.GetASN(context.Background(), 61138)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data)
}
```

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/electrologue/bgpview"
)

func main() {
	client := bgpview.NewClient()

	data, err := client.GetASNPrefixes(context.Background(), 61138)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data)
}
```

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/electrologue/bgpview"
)

func main() {
	client := bgpview.NewClient()

	data, err := client.GetASNPeers(context.Background(), 61138)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data)
}
```

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/electrologue/bgpview"
)

func main() {
	client := bgpview.NewClient()

	data, err := client.GetASNUpstreams(context.Background(), 61138)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data)
}
```

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/electrologue/bgpview"
)

func main() {
	client := bgpview.NewClient()

	data, err := client.GetASNDownstreams(context.Background(), 61138)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data)
}
```

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/electrologue/bgpview"
)

func main() {
	client := bgpview.NewClient()

	data, err := client.GetASNIxs(context.Background(), 61138)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data)
}
```

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/electrologue/bgpview"
)

func main() {
	client := bgpview.NewClient()

	data, err := client.GetPrefix(context.Background(), "192.209.63.0", 24)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data)
}
```

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/electrologue/bgpview"
)

func main() {
	client := bgpview.NewClient()

	data, err := client.GetIP(context.Background(), "2a05:dfc7:60::")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data)
}
```

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/electrologue/bgpview"
)

func main() {
	client := bgpview.NewClient()

	data, err := client.GetIX(context.Background(), 492)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data)
}
```

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/electrologue/bgpview"
)

func main() {
	client := bgpview.NewClient()

	data, err := client.GetSearch(context.Background(), "digitalocean")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data)
}

```
