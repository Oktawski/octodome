package domain

type PasswordHasher interface {
	HashPassword(password string) (string, error)
	CompareHashAndPassword(hashedPassword, password string) error
}
