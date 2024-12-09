package env

const (
	Development = "development"
	Staging     = "staging"
	Uat         = "uat"
	Sandbox     = "sandbox"
	Production  = "production"
)

func IsDevelopment(environment string) bool { return environment == Development }
func IsStaging(environment string) bool     { return environment == Staging }
func IsUat(environment string) bool         { return environment == Uat }
func IsSandbox(environment string) bool     { return environment == Sandbox }
func IsProduction(environment string) bool  { return environment == Production }

func IsValidEnv(environment string) bool {
	return IsDevelopment(environment) || IsStaging(environment) || IsUat(environment) || IsSandbox(environment) || IsProduction(environment)
}
