package account

import (
	"reflect"
	"testing"
)

func TestOpenAccount(t *testing.T) {
	type args struct {
		userID uint
	}
	tests := []struct {
		name string
		args args
		want NewAccountInput
	}{
		{
			name: "should return NewAccountInput",
			args: args{
				userID: 1,
			},
			want: NewAccountInput{
				Balance:  10000,
				UserID:   1,
				Currency: BRL,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OpenAccount(tt.args.userID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OpenAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}
