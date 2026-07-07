package greeter

// Greet returns the service name (lives under internal/, importable only within this module).
func Greet() string {
	return "go-internal-pkg"
}
