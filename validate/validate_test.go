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
		options []validate.PhoneOption
		isValid bool
	}{
		{"", nil, false},
		{"+79991224312", nil, true},
		{"+995593482054", nil, true},
		{"+9995593482054", nil, true},
		{"+7999122431q", nil, false},
		{"79991224312", nil, false},
		{"+7999-1224312", nil, false},
		{"+9995000009991224312", nil, false},
		{"+6512345678", []validate.PhoneOption{
			validate.WithDialCode("65"),
			validate.WithPhoneNumberSize("8"),
		}, true},
		{"+6512345678", []validate.PhoneOption{
			validate.WithDialCode("66"),
			validate.WithPhoneNumberSize("8"),
		}, false},
		{"+65123456789012", []validate.PhoneOption{
			validate.WithPhoneNumberSize("12"),
		}, true},
		{"+651123456789012", []validate.PhoneOption{
			validate.WithPhoneNumberSize("12"),
		}, true},
		{"+6512123456789012", []validate.PhoneOption{
			validate.WithPhoneNumberSize("12"),
		}, false},
		{"+1-2229567843451", []validate.PhoneOption{
			validate.WithDialCode("1-222"),
		}, true},
		{"+1-2229567843451", []validate.PhoneOption{
			validate.WithDialCode(""),
		}, false},
		{"+1-2229567843451", []validate.PhoneOption{
			validate.WithPhoneNumberSize(""),
		}, false},
	}

	for _, tc := range tt {
		t.Run("test mobile number: "+tc.input, func(t *testing.T) {
			opts := []validate.PhoneOption{}
			if tc.options != nil {
				opts = tc.options
			}

			err := validate.Phone(tc.input, opts...)
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
