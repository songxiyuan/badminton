package third

import "testing"

func Test_check(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "",
			args: args{
				id: "63365bd8c0e88cd4478b46a7",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			check(tt.args.id)
		})
	}
}
