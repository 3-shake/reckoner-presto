package presto

type Config struct {
	Host string

	// HTTP-Request-Header
	User     string `default:"hive"`
	Catalog  string `default:"hive"`
	Schema   string `default:"default"`
	TimeZone string `default:"UTC"`
}
