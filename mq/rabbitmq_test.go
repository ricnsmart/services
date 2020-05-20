package mq

import (
	"reflect"
	"testing"
)

const rabbitMQURL = "amqp://"

func TestNewConnection(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want *Connection
	}{
		{"test", args{url: rabbitMQURL}, &Connection{
			url:   rabbitMQURL,
			state: ClosedState,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewConnection(tt.args.url); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConnection() = %v, want %v", got, tt.want)
			}
		})
	}
}
