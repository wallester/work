package servicestate

type ServiceState string

const (
	// Unknown service state.
	Unknown ServiceState = "unknown"

	// Started means the service was not running when checked, but it is now started.
	Started ServiceState = "started"

	// Running means the service was already running when checked and no action was taken.
	Running ServiceState = "running"

	// NotRunning means the service was not running when checked and no action was taken.
	NotRunning ServiceState = "not running"

	// Stopped means the service was running when checked, but it is now stopped.
	Stopped ServiceState = "stopped"
)
