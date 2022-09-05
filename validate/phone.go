package validate

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type PhoneOption func(*PhoneOptions) error

type PhoneOptions struct {
	DialCode        string
	PhoneNumberSize string
}

func WithDialCode(v string) PhoneOption {
	return func(o *PhoneOptions) error {
		if v == "" {
			return errors.New("empty dial code")
		}
		o.DialCode = v
		return nil
	}
}

func WithPhoneNumberSize(v string) PhoneOption {
	return func(o *PhoneOptions) error {
		if v == "" {
			return errors.New("empty phone number size")
		}
		o.PhoneNumberSize = v
		return nil
	}
}

func Phone(p string, opts ...PhoneOption) error {
	options := PhoneOptions{
		PhoneNumberSize: "10",
	}

	for _, opt := range opts {
		if err := opt(&options); err != nil {
			return err
		}
	}

	dialCodeRegex := `\d{1,3}`
	if options.DialCode != "" {
		options.DialCode = strings.ReplaceAll(options.DialCode, "-", `\-`)
		dialCodeRegex = fmt.Sprintf("(%s)", options.DialCode)
	}

	phoneRegex := fmt.Sprintf(
		`^\+((\(|0)?%s)?((\s|\)|\-))?(\d{%s})$`,
		dialCodeRegex,
		options.PhoneNumberSize,
	)

	re := regexp.MustCompile(phoneRegex)

	if ok := re.MatchString(p); !ok {
		return fmt.Errorf("'%s' value doesn't match the phone regexp", p)
	}

	return nil
}
