package uuid

import "testing"

func TestStringToUUID(t *testing.T) {
	cases := []struct {
		input, expected string
	}{
		{"0", "00000000-0000-0000-0000-000000000000"},
		{"1", "00000000-0000-0000-0000-000000000001"},
		{"9", "00000000-0000-0000-0000-000000000009"},
		{"10", "00000000-0000-0000-0000-00000000000a"},
		{"15", "00000000-0000-0000-0000-00000000000f"},
		{"16", "00000000-0000-0000-0000-000000000010"},
		{"100", "00000000-0000-0000-0000-000000000064"},
		{"212345435656", "00000000-0000-0000-0000-003170c63608"},
		{"9999999999999999999999", "00000000-0000-021e-19e0-c9bab23fffff"},
	}

	for _, c := range cases {
		output, err := StringToUUID(c.input)
		if err != nil || output == nil || output.String() != c.expected {
			t.Errorf("StringToUUID(%q) == %q, expected %q", c.input, output, c.expected)
		}
	}
}
