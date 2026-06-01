package utils

import (
	"net/http"
	"strings"
	"testing"
)

func TestMSEdgeUserAgentKeepsProductSeparator(t *testing.T) {
	if !strings.Contains(MSEdgeUA, " Safari/537.36 Edg/") {
		t.Fatalf("MSEdgeUA should separate Safari and Edg products with a space, got %q", MSEdgeUA)
	}
}

func TestTryDefaultHeadersWithEdgeKeepsProductSeparator(t *testing.T) {
	header := http.Header{}
	header.Set("User-Agent", "edge")

	TryDefaultHeadersWith(header, "fetch")

	ua := header.Get("User-Agent")
	if !strings.Contains(ua, " Safari/537.36 Edg/") {
		t.Fatalf("edge header should separate Safari and Edg products with a space, got %q", ua)
	}
}
