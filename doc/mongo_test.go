package doc

import "testing"

const (
	address = "mongodb://ricnsmart:df426e941cf1@139.9.170.194:27017"
	dbName  = "gateway"
)

func TestConnect(t *testing.T) {
	type args struct {
		address string
		dbName  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test", args{
			address: address,
			dbName:  dbName,
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Connect(tt.args.address, tt.args.dbName); (err != nil) != tt.wantErr {
				t.Errorf("Connect() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
