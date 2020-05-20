package doc

import "testing"

const (
	address = "mongodb://"
	dbName  = ""
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
