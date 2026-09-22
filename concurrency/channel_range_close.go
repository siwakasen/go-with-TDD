package concurrency

func sendMessage(ch chan<- string, messages []string) {
	for _, message := range messages {
		ch <- message
	}
	close(ch)
}

func receiveMessage(ch <-chan string) (messages []string) {
	for message := range ch {
		messages = append(messages, message)
	}
	return
}

func getMessage(messages []string) []string {
	ch := make(chan string)
	go sendMessage(ch, messages)
	return receiveMessage(ch)
}
