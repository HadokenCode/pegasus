package mnetamqp

import (
	"github.com/cpapidas/pegasus/netamqp"
	"github.com/streadway/amqp"
)

// MockConnection mock for amqp.Connection
type MockConnection struct {
	CloseMock   func() error
	ChannelMock func() (netamqp.IChannel, error)
}

// Close mock for amqp.Connection Close
func (m MockConnection) Close() error {
	if m.CloseMock != nil {
		return m.CloseMock()
	}
	return nil
}

// Channel mock for amqp.Connection Channel
func (m MockConnection) Channel() (netamqp.IChannel, error) {
	if m.ChannelMock != nil {
		return m.ChannelMock()
	}
	return &amqp.Channel{}, nil
}
