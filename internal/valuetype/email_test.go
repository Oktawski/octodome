package valuetype

import (
	"strings"
	"testing"
)

func TestNewEmail_Required(t *testing.T) {
	_, err := NewEmail("   ")
	assertErrContains(t, err, emailRequiredErr)
}

func TestNewEmail_InvalidCharacters(t *testing.T) {
	cases := []string{
		"user*name@example.com",
		"user/name@example.com",
		`user\name@example.com`,
		"user name@example.com",
		"user\tname@example.com",
		"user\nname@example.com",
		"user:name@example.com",
		"user;name@example.com",
		"user,name@example.com",
		"user\"name@example.com",
		"uñer@example.com",
		"user@example!.com",
	}

	for _, address := range cases {
		t.Run(address, func(t *testing.T) {
			_, err := NewEmail(address)
			assertErrContains(t, err, emailInvalidCharactersErr)
		})
	}
}

func TestNewEmail_TooLong(t *testing.T) {
	address := "a@" + strings.Repeat("b", 250) + ".com"
	_, err := NewEmail(address)
	assertErrContains(t, err, emailTooLongErr)
}

func TestNewEmail_Valid(t *testing.T) {
	email, err := NewEmail("user@example.com")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if email != Email("user@example.com") {
		t.Fatalf("expected email value to match input")
	}
}

func assertErrContains(t *testing.T, err error, expected string) {
	t.Helper()
	if err == nil || !strings.Contains(err.Error(), expected) {
		t.Fatalf("expected error containing %q, got %v", expected, err)
	}
}
