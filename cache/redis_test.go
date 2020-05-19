package cache

import (
	"testing"
)

const (
	address  = "139.9.170.194:10032"
	password = "caf2fc65a2b3"
)

func TestConnect(t *testing.T) {
	type args struct {
		address  string
		password string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "test", args: struct {
			address  string
			password string
		}{address: address, password: password}, wantErr: false}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Connect(tt.args.address, tt.args.password); (err != nil) != tt.wantErr {
				t.Errorf("Connect() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
