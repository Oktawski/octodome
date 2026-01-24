package valuetype

import (
	"errors"
	"net/mail"
	"strings"
	"unicode"
)

type Email string

const (
	emailRequiredErr          = "email is required"
	emailInvalidAddressErr    = "invalid email address"
	emailTooLongErr           = "email is too long"
	emailInvalidFormatErr     = "invalid email format"
	emailInvalidCharactersErr = "email contains invalid characters"
	emailAtSymbolCountErr     = "email must contain exactly one @ symbol"
	emailDotSymbolCountErr    = "email must contain exactly one dot"
)

func NewEmail(address string) (Email, error) {
	if errs := isValidEmail(address); len(errs) > 0 {
		return "", errors.Join(errs...)
	}

	return Email(address), nil
}

func isValidEmail(address string) []error {
	errs := []error{}
	address = strings.TrimSpace(address)
	if address == "" {
		return append(errs, errors.New(emailRequiredErr))
	}
	if errs := containsInvalidCharacters(errs, address); len(errs) > 0 {
		return errs
	}
	if _, err := mail.ParseAddress(address); err != nil {
		return append(errs, errors.New(emailInvalidAddressErr))
	}
	if len(address) > 254 {
		return append(errs, errors.New(emailTooLongErr))
	}
	if errs := hasInvalidAtSymbolUsage(errs, address); len(errs) > 0 {
		return errs
	}
	if errs := hasInvalidDotUsage(errs, address); len(errs) > 0 {
		return errs
	}

	return errs
}

func containsInvalidCharacters(errs []error, address string) []error {
	if strings.IndexFunc(address, func(r rune) bool {
		return unicode.IsSpace(r) || unicode.IsControl(r)
	}) != -1 {
		return append(errs, errors.New(emailInvalidCharactersErr))
	}
	if strings.ContainsAny(address, `/\`) {
		return append(errs, errors.New(emailInvalidCharactersErr))
	}
	if strings.IndexFunc(address, func(r rune) bool {
		if r > unicode.MaxASCII {
			return true
		}
		switch {
		case r >= 'a' && r <= 'z':
			return false
		case r >= 'A' && r <= 'Z':
			return false
		case r >= '0' && r <= '9':
			return false
		}
		switch r {
		case '@', '.', '_', '-', '+':
			return false
		default:
			return true
		}
	}) != -1 {
		return append(errs, errors.New(emailInvalidCharactersErr))
	}
	return errs
}

func hasInvalidAtSymbolUsage(errs []error, address string) []error {
	if strings.Count(address, "@") != 1 {
		return append(errs, errors.New(emailAtSymbolCountErr))
	}
	parts := strings.Split(address, "@")
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return append(errs, errors.New(emailInvalidFormatErr))
	}
	return errs
}

func hasInvalidDotUsage(errs []error, address string) []error {
	if strings.Count(address, ".") != 1 {
		return append(errs, errors.New(emailDotSymbolCountErr))
	}
	return errs
}
