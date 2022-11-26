package natsconn

type NatsEncodeType string

const (
	JSON_ENCODER    NatsEncodeType = "json"
	GOB_ENCODER     NatsEncodeType = "gob"
	DEFAULT_ENCODER NatsEncodeType = "default"
)
