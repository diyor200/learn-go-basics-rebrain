package main

import (
	"net/http/httptest"
	"testing"
)

func BenchmarkHelloWorld(b *testing.B) {
	for i := -0; i < b.N; i++ {
		w := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/hello", nil)

		helloWorld(w, req)

	}
}
