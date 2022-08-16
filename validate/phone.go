package validate

import (
	"fmt"
	"regexp"
)

var phoneRegex = `^\+((\(|0)?\d{1,3})?((\s|\)|\-))?(\d{10})$`

func Phone(p string) error {
	re := regexp.MustCompile(phoneRegex)

	if ok := re.MatchString(p); !ok {
		return fmt.Errorf("'%s' value doesn't match the phone regexp", p)
	}

	return nil
}
