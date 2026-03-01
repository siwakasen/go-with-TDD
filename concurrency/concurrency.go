// Package concurrency
package concurrency

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func() {
			resultChannel <- result{url, wc(url)}
		}()
	}
	// TODO: example todo comment (showing hightlight on lazyvim)

	for range urls {
		r := <-resultChannel
		results[r.string] = r.bool
	}
	// FIX: example fix comment

	return results
}
