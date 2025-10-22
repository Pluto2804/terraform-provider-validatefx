package validate

import "testing"

func TestIsSemver_Valid(t *testing.T) {
	tests := []string{
		"1.2.3",
		"v1.2.3",
		"0.0.1",
		"v0.0.1",
		"10.20.30",
		"v10.20.30",
		"1.2.3-alpha",
		"v1.2.3-beta.4",
		"1.2.3+build.5",
		"v1.2.3-alpha+build.5",
	}

	for _, s := range tests {
		if !IsSemver(s) {
			t.Errorf("expected valid semver for %q but got false", s)
		}
	}
}

func TestIsSemver_Invalid(t *testing.T) {
	tests := []string{
		"",
		"1",
		"1.2",
		"1.2.3.4",
		"v1.2",
		"1.2.x",
		"v1.2.3-",
		"v1.2.3+",
		"a.b.c",
		"1.2.3beta",
	}

	for _, s := range tests {
		if IsSemver(s) {
			t.Errorf("expected invalid semver for %q but got true", s)
		}
	}
}
