package concurrency

import (
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	return url != "waat://furhurterwe.geds"
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckWebsites(mockWebsiteChecker, urls)
	}
}
func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gyspydave5.com",
		"waat://furhurterwe.geds",
	}

	want := map[string]bool{
		"http://google.com":          true,
		"http://blog.gyspydave5.com": true,
		"waat://furhurterwe.geds":    false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", got, want)
	}

}
