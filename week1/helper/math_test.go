package helper

import (
	"testing"
)

func TestIndexOf(t *testing.T) {
	type args struct {
		arr  []int
		find int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "unhappy case",
			args: args{
				arr:  []int{1, 2, 3},
				find: -1,
			},
			want: -1,
		},
		{
			name: "happy case",
			args: args{
				arr:  []int{1, 2, 3},
				find: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IndexOf(tt.args.arr, tt.args.find); got != tt.want {
				t.Errorf("IndexOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndexOfString(t *testing.T) {
	type args struct {
		arr  []string
		find string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "happy case",
			args: args{
				arr:  []string{"abc", "def", "g"},
				find: "g",
			},
			want: 2,
		},
		{
			name: "unhappy case",
			args: args{
				arr:  []string{"abc", "def", "g"},
				find: "z",
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IndexOfString(tt.args.arr, tt.args.find); got != tt.want {
				t.Errorf("IndexOfString() = %v, want %v", got, tt.want)
			}
		})
	}
}
