package main

import (
	"reflect"
	"testing"
)

func Test_solution(t *testing.T) {
	tests := []struct {
		name     string
		interval [][]int
		want     [][]int
	}{{
		name: "base case",
		interval: [][]int{
			{1, 3}, {2, 6}, {8, 10}, {15, 18},
		},
		want: [][]int{
			{1, 6}, {8, 10}, {15, 18},
		},
	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.interval); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
