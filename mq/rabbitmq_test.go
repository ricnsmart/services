package mq

import (
	"reflect"
	"testing"
)

const rabbitMQURL = "amqp://ricnsmart:030e6dcb31d9@139.9.170.194:5672"

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
