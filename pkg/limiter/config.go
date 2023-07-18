package limiter

type Config struct {
	CachePrefix string
}

func (c *Config) Validate() error {
	return nil
}
