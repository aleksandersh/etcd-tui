package domain

type Config struct {
	Title string
}

func NewConfig(title string) *Config {
	return &Config{Title: title}
}
