package limiter

type Config struct {
	CachePrefix string `yaml:"cache_prefix"`
}

func (c *Config) Validate() error {
	return nil
}
