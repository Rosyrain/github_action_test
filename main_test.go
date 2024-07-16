package main

import (
	"testing"
)

func TestConnect(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: "success connect", args: args{url: "https://www.baidu.com"}, want: true},
		{name: "failed connect", args: args{url: "www.baidu.com"}, want: false},
		{name: "bad connect", args: args{url: "http://www.baidu.com"}, want: true},
		//{name: "misguided attempt", args: args{url: "http://www.baidu.com"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := Connect(tt.args.url); got != tt.want {
				t.Errorf("Connect() = %v, want %v", got, tt.want)
			}
		})
	}
}
