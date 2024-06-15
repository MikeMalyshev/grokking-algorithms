package unit7

import (
	"testing"
)

func TestFindBestWay(t *testing.T) {

	type args struct {
		departure   string
		destination string
		roads       map[string]float64
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"default",
			args{
				"A",
				"F",
				map[string]float64{
					"A-B": 7,
					"B-C": 10,
					"C-D": 11,
					"D-E": 6,
					"E-F": 9,
					"A-F": 14,
					"A-C": 9,
					"B-D": 15,
					"C-F": 2,
				},
			},
			"A->C->F",
		},
	}

	for _, tt := range tests {
		t.Run("test1", func(t *testing.T) {
			res := FindBestWay(tt.args.departure, tt.args.destination, tt.args.roads)
			if res != tt.want {
				t.Errorf("got %s, want %s", res, tt.want)
			}
		})
	}
}
