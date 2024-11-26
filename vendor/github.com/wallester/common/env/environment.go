package env

const (
	// Development represents "development" environment.
	Development = "development"

	// Staging represents "staging" environment.
	Staging = "staging"

	// Sandbox represents "sandbox" environment.
	Sandbox = "sandbox"

	// Production represents "production" environment.
	Production = "production"
)

// IsDevelopment indicates whether it is development environment or not.
func IsDevelopment(environment string) bool {
	return environment == Development
}

// IsStaging indicates whether it is staging environment or not.
func IsStaging(environment string) bool {
	return environment == Staging
}

// IsSandbox indicates whether it is sandbox environment or not.
func IsSandbox(environment string) bool {
	return environment == Sandbox
}

// IsProduction indicates whether it is production environment or not.
func IsProduction(environment string) bool {
	return environment == Production
}
