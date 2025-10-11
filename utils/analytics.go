package utils

import (
	"crypto/rand"
	"encoding/hex"
	"strings"
)

// GenerateSessionID generates a unique session ID
func GenerateSessionID() string {
	bytes := make([]byte, 16)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

// ParseDevice parses user agent to determine device type
func ParseDevice(userAgent string) string {
	ua := strings.ToLower(userAgent)

	if strings.Contains(ua, "mobile") || strings.Contains(ua, "android") || strings.Contains(ua, "iphone") {
		return "mobile"
	}
	if strings.Contains(ua, "tablet") || strings.Contains(ua, "ipad") {
		return "tablet"
	}
	return "desktop"
}

// ParseBrowser parses user agent to determine browser
func ParseBrowser(userAgent string) string {
	ua := strings.ToLower(userAgent)

	if strings.Contains(ua, "chrome") && !strings.Contains(ua, "edg") {
		return "chrome"
	}
	if strings.Contains(ua, "firefox") {
		return "firefox"
	}
	if strings.Contains(ua, "safari") && !strings.Contains(ua, "chrome") {
		return "safari"
	}
	if strings.Contains(ua, "edg") {
		return "edge"
	}
	if strings.Contains(ua, "opera") {
		return "opera"
	}
	return "unknown"
}

// ParseOS parses user agent to determine operating system
func ParseOS(userAgent string) string {
	ua := strings.ToLower(userAgent)

	if strings.Contains(ua, "windows") {
		return "windows"
	}
	if strings.Contains(ua, "mac") {
		return "macos"
	}
	if strings.Contains(ua, "linux") {
		return "linux"
	}
	if strings.Contains(ua, "android") {
		return "android"
	}
	if strings.Contains(ua, "ios") || strings.Contains(ua, "iphone") || strings.Contains(ua, "ipad") {
		return "ios"
	}
	return "unknown"
}
