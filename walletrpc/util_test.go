package walletrpc

import "testing"

func TestXMRToDecimal(t *testing.T) {
	values := []struct {
		in uint64
		out string
	}{
		{0, "0"},
		{150020000, "0.00015002"},
		{23000000000000, "23"},
		{69006900000000, "69.0069"},
		{1, "0.000000000001"},
	}
	for _, v := range values {
		res := XMRToDecimal(v.in)
		if res != v.out {
			t.Fatalf(`XMRToDecimal(%d) = "%s", want "%s"`, v.in, res, v.out)
		}
	}
}

func TestXMRToFloat64(t *testing.T) {
	values := []struct {
		in uint64
		out float64
	}{
		{0, 0},
		{150020000, 0.00015002},
		{23000000000000, 23},
		{69006900000000, 69.0069},
		{1, 0.000000000001},
	}
	for _, v := range values {
		res := XMRToFloat64(v.in)
		if res != v.out {
			t.Fatalf("XMRToFloat64(%d) = %f, want %f", v.in, res, v.out)
		}
	}
}
