package main

import (
	"fmt"
	"os"

	"github.com/3-shake/reckoner-presto/presto"
)

type QueryResult *presto.QueryResult

func main() {
	baseUrl := os.Getenv("PRESTO_BASE_URL")

	config := presto.Config{
		Host:     baseUrl,
		User:     "hive",
		Catalog:  "hive",
		Schema:   "sample",
		TimeZone: "UTC",
	}
	presto := presto.New(&config)
	query, err := presto.Query(fmt.Sprintf("SELECT * FROM %q LIMIT 1", "sample"))
	if err != nil {
		fmt.Println(err)
		return
	}
	next, err := query.Next()
	fmt.Println(next, err)
}
