package main

import "testing"

func Test_toBytes(t *testing.T) {
	cases := map[string]int64{
		"1K": 1024,
		"1k": 1024,
		"2K": 2048,
		"1M": 1048576,
		"2m": 1048576 * 2,
		"1g": 1024 << 20,
		"1T": 1024 << 30,
	}

	for k, v := range cases {
		if n, err := toBytes(k); err != nil {
			t.Error("toBytes failed: ", err)
		} else if n != v {
			t.Errorf("expect %v, actually got %v\n", v, n)
		}
	}
}
