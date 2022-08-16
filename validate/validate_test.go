package validate_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tkcrm/modules/validate"
)

func Test_Phone(t *testing.T) {
	tt := []struct {
		input   string
		isValid bool
	}{
		{"", false},
		{"+79991224312", true},
		{"+995593482054", true},
		{"+9995593482054", true},
		{"+7999122431q", false},
		{"79991224312", false},
		{"+7999-1224312", false},
		{"+9995000009991224312", false},
	}

	for _, tc := range tt {
		t.Run("test mobile number: "+tc.input, func(t *testing.T) {
			err := validate.Phone(tc.input)
			if !tc.isValid {
				require.Error(t, err)
			}
			assert.Equal(t, tc.isValid, err == nil)
		})
	}
}

func Test_Email(t *testing.T) {
	tt := []struct {
		input   string
		isValid bool
	}{
		{"", false},
		{"test@test.com", true},
		{"info@finchpay.io", true},
		{"test@test.ru", true},
		{"t@com", false},
		{"t@.com", false},
		{"@test.com", false},
		{"test.com", false},
	}

	for _, tc := range tt {
		t.Run("test email: "+tc.input, func(t *testing.T) {
			err := validate.Phone(tc.input)
			if !tc.isValid {
				require.Error(t, err)
			}
			assert.Equal(t, tc.isValid, err == nil)
		})
	}
}
