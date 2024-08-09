package core

import "testing"

func TestVersionSerialize(t *testing.T) {
	ts := []struct {
		v *Version
		s string
	}{
		{NewVersion(1, 2, 3), "v1.2.3"},
		{NewVersion(0, 0, 0), "v0.0.0"},
		{NewVersion(1, 0, 0), "v1.0.0"},
		{NewVersion(0, 1, 0), "v0.1.0"},
		{NewVersion(0, 0, 1), "v0.0.1"},
		{NewVersion(1, 10, 0), "v1.10.0"},
	}

	for _, tc := range ts {
		if tc.v.String() != tc.s {
			t.Errorf("expected %s, got %s", tc.s, tc.v.String())
		}
		v, err := GetVersion(tc.s)
		if err != nil {
			t.Errorf("failed to parse version: %v", err)
		}
		if v.Compare(tc.v) != 0 {
			t.Errorf("expected %v, got %v", tc.v, v)
		}
	}
}

func TestVersionCompare(t *testing.T) {
	ts := []struct {
		v1, v2 *Version
		c      int
	}{
		{NewVersion(1, 2, 3), NewVersion(1, 2, 3), 0},
		{NewVersion(1, 2, 3), NewVersion(1, 2, 4), -1},
		{NewVersion(1, 2, 3), NewVersion(1, 3, 3), -1},
		{NewVersion(1, 2, 3), NewVersion(2, 2, 3), -1},
		{NewVersion(1, 2, 3), NewVersion(0, 2, 3), 1},
		{NewVersion(1, 2, 3), NewVersion(1, 1, 3), 1},
		{NewVersion(1, 2, 3), NewVersion(1, 2, 2), 1},
	}

	for _, tc := range ts {
		if tc.v1.Compare(tc.v2) != tc.c {
			t.Errorf("expected %d, got %d", tc.c, tc.v1.Compare(tc.v2))
		}
	}
}
