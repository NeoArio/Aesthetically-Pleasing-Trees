package main

import (
	"log"
	"testing"
	"time"
)

func TestGardens(t *testing.T) {
	tests := []struct {
		name string
		Garden Garden
		want int
	}{
		{
			name: "Already aesthetically pleasing garden without removal",
			Garden: Garden{1, 3, 1, 2},
			want: 0,
		},
		{
			name: "Already aesthetically pleasing garden without removal",
			Garden: Garden{8, 2, 7, 5},
			want: 0,
		},
		{
			name: "Many ways to solve the problem",
			Garden: Garden{3, 4, 5, 3, 7},
			want: 3,
		},
		{
			name: "Not possible to achieve desired result",
			Garden: Garden{1, 2, 3, 4},
			want: -1,
		},
		{
			name: "Not possible to achieve desired result",
			Garden: Garden{1, 2, 2, 4},
			want: -1,
		},

		// Can add more cases here to test the implementation.
	}
	for i, tt := range tests {
		start := time.Now()
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.Garden.SolveByFirstSolution(); got != tt.want {
				t.Errorf("Answer = %v, Want %v", got, tt.want)
			}
		})
		elapsed := time.Since(start)
		log.Printf("First Solution Test %v\ninput:\n\t%v\nExecution time: %v",i,tt.Garden,elapsed)
	}

	for i, tt := range tests {
		start := time.Now()
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.Garden.SolveBySecondSolution(); got != tt.want {
				t.Errorf("Answer = %v, Want %v", got, tt.want)
			}
		})
		elapsed := time.Since(start)
		log.Printf("Second Solution Test %v\ninput:\n\t%v\nExecution time: %v",i,tt.Garden,elapsed)
	}
}
