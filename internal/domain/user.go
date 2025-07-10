package domain

type AuthProvider string

const (
	LocalProvider  AuthProvider = "local"
	GoogleProvider AuthProvider = "google"
	GithubProvider AuthProvider = "github"
)

type PlanType string

const (
	BasicPlan   PlanType = "basic"
	PremiumPlan PlanType = "premium"
)

type User struct {
	ID           int64
	Name         string
	Email        string
	PasswordHash string
	Provider     AuthProvider
	Plan         PlanType
}
