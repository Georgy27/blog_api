package validation

import (
	"context"
	"github.com/Georgy27/blogger_api/pkg/errors/validate"
)

func ValidateID(id int64) validate.Condition {
	return func(ctx context.Context) error {
		if id <= 0 {
			return validate.NewValidationErrors("id must be greater than 0")
		}

		return nil
	}
}

func ValidateTitle(title string, method string) validate.Condition {
	return func(ctx context.Context) error {
		if method == "create" {
			if title == "" {
				return validate.NewValidationErrors("title must not be empty")
			}
		}

		if len(title) > 30 {
			return validate.NewValidationErrors("title must be less than 30 characters")
		}

		return nil
	}
}

func ValidateContent(content string, method string) validate.Condition {
	return func(ctx context.Context) error {
		if method == "create" {
			if content == "" {
				return validate.NewValidationErrors("content must not be empty")
			}
		}

		if len(content) > 1000 {
			return validate.NewValidationErrors("content must be less than 1000 characters")
		}

		return nil
	}
}

func ValidateShortDescription(shortDescription string, method string) validate.Condition {
	return func(ctx context.Context) error {
		if method == "create" {
			if shortDescription == "" {
				return validate.NewValidationErrors("short_description must not be empty")
			}
		}

		if len(shortDescription) > 100 {
			return validate.NewValidationErrors("short_description must be less than 100 characters")
		}

		return nil
	}
}
