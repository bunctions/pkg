package runner

// Runner defines the interface of a runner
type Runner interface {
	Start()
}

// Start starts a generic runner
func Start() {
	Generic.Start()
}
