package validation

import (
	"context"
	"github.com/Georgy27/blogger_api/pkg/errors/validate"
)

const (
	create = "create"
)

func ValidateID(id int64) validate.Condition {
	return func(ctx context.Context) error {
		if id <= 0 {
			return validate.NewValidationErrors("id must be greater than 0")
		}

		return nil
	}
}

func ValidateName(name string, method string) validate.Condition {
	return func(ctx context.Context) error {
		if method == create {
			if name == "" {
				return validate.NewValidationErrors("name must not be empty")
			}
		}

		if len(name) > 15 {
			return validate.NewValidationErrors("name must be less than 15 characters")
		}

		return nil

	}
}

func ValidateDescription(description string, method string) validate.Condition {
	return func(ctx context.Context) error {
		if method == create {
			if description == "" {
				return validate.NewValidationErrors("description must not be empty")
			}
		}

		if len(description) > 500 {
			return validate.NewValidationErrors("description must be less than 500 characters")
		}

		return nil

	}
}

func ValidateWebsiteUrl(websiteUrl string, method string) validate.Condition {
	return func(ctx context.Context) error {
		if method == create {
			if websiteUrl == "" {
				return validate.NewValidationErrors("website_url must not be empty")
			}
		}

		if len(websiteUrl) > 100 {
			return validate.NewValidationErrors("website_url must be less than 100 characters")
		}

		return nil

	}
}
