package http

type config struct {
	Port        uint   `default:"8080"`
	ContentType string `split_words:"true" default:"text/plain"`
}
