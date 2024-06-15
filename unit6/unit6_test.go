package unit6

import "testing"

func TestBreadthSearch(t *testing.T) {
	type args struct {
		startPerson string
		graph       map[string][]string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"standart",
			args{"you",
				map[string][]string{
					"you":    {"alice", "bob", "claire"},
					"bob":    {"anuj", "peggy"},
					"alice":  {"peggy"},
					"claire": {"thom", "jonny"},
					"anuj":   {},
					"peggy":  {},
					"thom":   {},
					"jonny":  {},
				},
			},
			"thom",
		},
		{"nodata",
			args{"",
				map[string][]string{},
			},
			"noone",
		},
		{"nocommunications",
			args{"",
				map[string][]string{
					"you":    {},
					"bob":    {"anuj", "peggy"},
					"alice":  {"peggy"},
					"claire": {"thom", "jonny"},
					"anuj":   {},
					"peggy":  {},
					"thom":   {},
					"jonny":  {},
				},
			},
			"noone",
		},
	}
	for _, tt := range tests {
		t.Run("findIterative "+tt.name, func(t *testing.T) {
			ans := findIterative(tt.args.startPerson, tt.args.graph)
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
		t.Run("findRecursive "+tt.name, func(t *testing.T) {
			ans := findRecursive(tt.args.startPerson, tt.args.graph)
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}
