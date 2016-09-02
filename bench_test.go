package bench

import (
	"testing"
	"fmt"
)

func TestUrl(t *testing.T) {
	report, err := Url("aukbit.com")
	if err != nil {
		t.Error("Expected HTTP report for", err)
	}
	if len(report) == 0 {
		t.Error("Expected some report !=", report)
	}
	fmt.Println(report)
}

func TestUrls(t *testing.T) {
	urls := []string{"expresso.pt", "publico.pt", "observador.pt"}
	report, err := Urls(urls)
	if err != nil {
		t.Error("Expected HTTP report", err)
	}
	if len(report) == 0 {
		t.Error("Expected some report !=", report)
	}
	fmt.Println(report)
}
