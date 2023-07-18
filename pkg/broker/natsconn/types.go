package natsconn

type ConnType int

const (
	ConnTypeDefault ConnType = iota
	ConnTypeEncoded
)

type NatsEncodeType string

const (
	NatsEncodeJson    NatsEncodeType = "json"
	NatsEncodeGob     NatsEncodeType = "gob"
	NatsEncodeDefault NatsEncodeType = "default"
)
