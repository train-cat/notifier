package helper

// List of exit code
const (
	ExitCodeSuccess = iota
	ExitCodeErrorInitConfig
	ExitCodeErrorListenServer
	ExitCodeErrorStopServer
)
