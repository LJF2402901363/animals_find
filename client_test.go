/*
@author: 陌意随影
@since: 2023/11/6 01:37:01
@desc:
*/
package main

import "testing"

func TestIsNotFind(t *testing.T) {
	type args struct {
		params []interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "test0",
			args: args{
				params: []interface{}{"ht5", 0},
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "test1",
			args: args{
				params: []interface{}{"ab2", 0},
			},
			want:    false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsNotFind(tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsNotFind() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsNotFind() got = %v, want %v", got, tt.want)
			}
		})
	}
}
