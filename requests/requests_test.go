package requests

import (
	"reflect"
	"testing"
)
type Args struct {
	receivedFlag int
}
func TestIsAccessFlagValid(t *testing.T) {

	tests := []struct {
		name  string
		args  Args
		want  error
		want1 bool
	}{
		{"IsAccessFlagValid", Args{1}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := IsAccessFlagValid(tt.args.receivedFlag)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IsAccessFlagValid() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("IsAccessFlagValid() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestRequestTA(t *testing.T) {
	const (
		attRTA1 = iota
      	attRTA2
		attRTA3
	)


	tests := []struct {
		name    string
		args Args
		attempt int
		
alidFlag int
		wantErr bool
	}{
		{"RequestTA", Args{1}, attRTA1, 0,true},
		{"RequestTA", Args{1}, attRTA2, 0,true},
		{"RequestTA", Args{1}, attRTA3, 0,true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := RequestTA(); (err != nil) != tt.wantErr {
				t.Errorf("RequestTA() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}