package name

import "testing"

func TestUpper(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Upper(tt.args.s); got != tt.want {
				t.Errorf("Upper() = %v, want %v", got, tt.want)
			}
		})
	}
}


func benchUpper(i string, b *testing.B) {
	for n := 0; n < b.N; n++ {
		Upper(i)
	}
}

func BenchmarkUpper1(b *testing.B) { benchUpper("a", b)}
func BenchmarkUpper10(b *testing.B) { benchUpper("Hola que tal estas", b)}
func BenchmarkUpper20(b *testing.B) { benchUpper("Este va a ser mas largo porque se necesita asi", b)}