package greputils

import (
	"regexp"
	"testing"
)

func TestCompilePatternFixedCaseSensitive(t *testing.T) {
	re, fixed, err := compilePattern("Error", Options{Fixed: true})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if re != nil && fixed == "" {
		t.Error("expected fixed string, got empty")
	}
	if fixed != "Error" {
		t.Errorf("expected 'Error', got '%s'", fixed)
	}
}

func TestCompilePatternInsensitive(t *testing.T) {
	_, fixed, _ := compilePattern("Error", Options{Fixed: true, Insensitive: true})
	if fixed != "error" {
		t.Errorf("expected 'error', got '%s'", fixed)
	}
}

func TestMatchLineFixed(t *testing.T) {
	opts := Options{Fixed: true}
	if !matchLine("Error: something broke", nil, "Error", opts) {
		t.Error("expected match for fixed string")
	}
	if matchLine("Info: all good", nil, "Error", opts) {
		t.Error("did not expect match")
	}
}

func TestMatchLineInsensitive(t *testing.T) {
	opts := Options{Fixed: true, Insensitive: true}
	if !matchLine("error: disk full", nil, "error", opts) {
		t.Error("expected match ignoring case")
	}
}

func TestMatchLineRegexp(t *testing.T) {
	re := regexp.MustCompile(`error|warning`)
	opts := Options{}
	if !matchLine("warning: low memory", re, "", opts) {
		t.Error("expected regexp match")
	}
	if matchLine("info: all good", re, "", opts) {
		t.Error("did not expect regexp match")
	}
}

func TestMatchLineInvert(t *testing.T) {
	opts := Options{Fixed: true, Invert: true}
	if !matchLine("info: all good", nil, "Error", opts) {
		t.Error("expected invert match (line without 'Error')")
	}
	if matchLine("Error: broken", nil, "Error", opts) {
		t.Error("did not expect match (invert)")
	}
}
