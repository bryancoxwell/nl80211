package nl80211

import (
	"fmt"
)

// wifiFreqToChannel converts a WiFi frequency to a channel number.
// Adapted from from https://github.com/kismetwireless/kismet/blob/6f57161569d431abce00a135c816066e901ff164/capture_linux_wifi/capture_linux_wifi.c#L592
func wifiFreqToChannel(freq uint32) (uint32, error) {
	switch {
	case freq < 2412:
		return 0, fmt.Errorf("wifiFreqToChannel: invalid frequency %d", freq)
	case freq == 2484:
		return 14, nil
	case freq < 2484:
		return (freq - 2407) / 5, nil
	case freq >= 4910 && freq <= 4980:
		return (freq - 4000) / 5, nil
	case freq < 5950:
		return (freq - 5000) / 5, nil
	case freq <= 45000: 
		return freq - 5950 / 5, nil
	case freq >= 58320 && freq <= 70200:
		return (freq - 56160) / 2160, nil
	default:
		return 0, fmt.Errorf("wifiFreqToChannel: invalid frequency %d", freq)
	}
}

// wifiChannelToFreq converts a 2.4GHz or 5GHz WiFi channel number to a frequency.
// Adapted from https://github.com/kismetwireless/kismet/blob/6f57161569d431abce00a135c816066e901ff164/capture_linux_wifi/capture_linux_wifi.c#L556
func wifiChannelToFreq(channel uint32) (uint32, error) {
	switch {
	case channel == 0:
		return 0, fmt.Errorf("wifiChannelToFreq: invalid channel %d", channel)
	case channel == 14:
		return 2484, nil
	case channel < 14:
		return 2407 + channel*5, nil
	case channel >= 182 && channel <= 196:
		return 4000 + channel*5, nil
	case channel > 14:
		return 5000 + channel*5, nil
	default:
		return 0, fmt.Errorf("wifiChannelToFreq: invalid channel %d", channel)
	}
}