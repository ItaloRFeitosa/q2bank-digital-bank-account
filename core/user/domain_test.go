package user

import "testing"

func TestKind_IsSeller(t *testing.T) {
	tests := []struct {
		name string
		k    Kind
		want bool
	}{
		{
			name: "when seller should be true",
			k:    Seller,
			want: true,
		},
		{
			name: "when not seller should be false",
			k:    Customer,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.k.IsSeller(); got != tt.want {
				t.Errorf("Kind.IsSeller() = %v, want %v", got, tt.want)
			}
		})
	}
}
