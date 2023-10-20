package taskmanager

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

func genTaskID(str string, payload []byte) string {
	mac := hmac.New(sha256.New, []byte(str))
	mac.Write([]byte(payload))
	return fmt.Sprintf("%x", mac.Sum(nil))
}
