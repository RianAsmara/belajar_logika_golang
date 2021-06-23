package polindrome

import "testing"

func TestPolindromNumber(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "Example input 1",
			args: args{
				input: "1 10",
			},
			want:    9,
			wantErr: false,
		},
		{
			name: "Example input 2",
			args: args{
				input: "99 100",
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "Example input 3",
			args: args{
				input: "21 31",
			},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PolindromeNumber(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("PolindromNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PolindromNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
