package nl80211

import (
	"fmt"
	"testing"
)

func TestWifiFreqToChannel(t *testing.T) {
	type passingTest struct {
		input uint32
		want uint32
	}

	type failingTest struct {
		input uint32
		want error 
	}

	passingTests := []passingTest{
		{2412, 1},
		{2417, 2},
		{2422, 3},
		{2427, 4},
		{2432, 5},
		{2437, 6},
		{2442, 7},
		{2447, 8},
		{2452, 9},
		{2457, 10},
		{2462, 11},
		{2467, 12},
		{2472, 13},
		{2484, 14},
		{5160, 32},
		{5180, 36},
		{5200, 40},
		{5220, 44},
		{5240, 48},
		{5260, 52},
		{5280, 56},
		{5300, 60},
		{5320, 64},
		{5480, 96},
		{5500, 100},
		{5520, 104},
		{5540, 108},
		{5560, 112},
		{5580, 116},
		{5600, 120},
		{5620, 124},
		{5640, 128},
		{5660, 132},
		{5680, 136},
		{5700, 140},
		{5720, 144},
		{5745, 149},
		{5765, 153},
		{5785, 157},
		{5805, 161},
		{5825, 165},
	}

	failingTests := []failingTest{
		{0, fmt.Errorf("wifiFreqToChannel: invalid frequency 0")},
		{75000, fmt.Errorf("wifiFreqToChannel: invalid frequency 2411")},
	}

	for _, test := range passingTests {
		got, err := wifiFreqToChannel(test.input)
		if err != nil {
			t.Fatal(err)
		}
		if got != test.want {
			t.Fatalf("wifiFreqToChannel(%d) = %d, want %d", test.input, got, test.want)
		}
	}

	for _, test := range failingTests {
		got, err := wifiFreqToChannel(test.input)
		if err == nil {
			t.Fatalf("wifiFreqToChannel(%d) should return non-nil error", test.input)
		}
		if got != 0 {
			t.Fatalf("wifiFreqToChannel(%d) = %d, want error %s", test.input, err, test.want)
		}
	}
}