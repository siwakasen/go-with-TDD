// Package concurrency
package concurrency

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	// NOTE: unbuffered channel
	resultChannel := make(chan result)

	for _, url := range urls {
		go func() {
			resultChannel <- result{url, wc(url)}
		}()
	}

	for range urls {
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}

func CheckWebsitesWithBufferedChannel(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	BufferSize := len(urls)

	// NOTE: buffered channel
	resultChannel := make(chan result, BufferSize)

	for _, url := range urls {
		go func() {
			resultChannel <- result{url, wc(url)}
		}()
	}

	for range urls {
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}
