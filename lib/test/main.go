package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

var requestVolume int = 5000

var availableChannels []string = []string{"Main", "Secondary", "Example", "Testing", "Beta"}

func main() {

	for i := 0; i < requestVolume; i++ {
		sendRandomRequest()
		time.Sleep(time.Second * time.Duration(.25+rand.Float32()))
	}

}

func sendRandomRequest() {
	// Pick random element
	ch := availableChannels[rand.Intn(len(availableChannels))]
	// Send the test request
	sendTestRequest(ch, gofakeit.HackerPhrase())
}

func sendTestRequest(channel string, content string) {
	hc := http.Client{}

	form := url.Values{}
	form.Add("channel", channel)
	form.Add("content", content)

	req, err := http.NewRequest("POST", "http://localhost:8080/api/log", strings.NewReader(form.Encode()))
	if err != nil {
		fmt.Println("Error:", err)
		// Quit
		return
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Send the info
	fmt.Printf("Form: %v\n", form)

	// Send the request
	resp, err := hc.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Printf("Form: %v\n", resp)
}
