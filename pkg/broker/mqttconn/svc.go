package mqttconn

import "context"

func (s *Mqtt) Name() string { return "mqtt" }

func (s *Mqtt) Start(_ context.Context) error { return nil }

func (s *Mqtt) Stop(_ context.Context) error { return nil }

func (s *Mqtt) Ping(_ context.Context) error { return nil }
