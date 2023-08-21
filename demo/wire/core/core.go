package core

type Config struct {
	DBSource string
}

func NewConfig() *Config {
	return &Config{
		DBSource: "aabc",
	}
}

type DB struct {
	table string
}

func NewDB(cfg *Config) *DB {
	return &DB{
		table: cfg.DBSource,
	}
}
func (d *DB) Find() string {
	return d.table
}
