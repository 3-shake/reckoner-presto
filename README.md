# Reckoner Presto

## Presto Client For GO

### Installing

```
go get -u github.com/3-shake/reckoner-presto/presto
```

### Example

```go
import (
	"os"

	"github.com/3-shake/reckoner-presto/presto"
)

func main() {
	baseUrl := os.Getenv("PRESTO_BASE_URL")
	config := presto.Config{
		Host:     baseUrl,
		User:     "hive",
		Catalog:  "hive",
		Schema:   "xxxxxxxxxxx",
		TimeZone: "UTC",
	}
	presto, _ := presto.New(&config)
	query, _ := presto.Query("SELECT client_id FROM sample")
	next, _ := query.Next()
}

```

