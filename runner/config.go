package runner

// Config defines the common configuration of a runner
type Config struct {
	ExportingPath   string `split_words:"true" default:"./plugin.so"`
	ExportingSymbol string `split_words:"true" default:"Exported"`
}
