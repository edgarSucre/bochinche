package bot

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/edgarSucre/bochinche/domain"
)

type Bot struct {
	client   domain.Broker
	fetchUrl string
}

func New(c domain.Broker) *Bot {
	return &Bot{
		client:   c,
		fetchUrl: "https://stooq.com/q/l/",
	}
}

func (b *Bot) ListenForRequest() error {
	consumer, err := b.client.GetQuoteConsummer()
	if err != nil {
		return err
	}

	go func() {
		for request := range consumer {
			go b.GetQuoteFile(request.Body)
		}
	}()

	return nil
}

func (b *Bot) GetQuoteFile(msg []byte) {
	var req domain.QuoteMessage

	err := json.Unmarshal(msg, &req)

	if err != nil {
		// can't do anything
		log.Printf("Can't parse: %s\n", msg)
		return
	}

	response := domain.QuoteMessage{
		Room: req.Room,
	}

	if req.Message == "" {
		response.Message = "I can't search empty codes"
		b.client.PublishResponseMessage(response)
		return
	}

	if strings.Contains(req.Message, "&") {
		response.Message = "Extra query detected!"
		b.client.PublishResponseMessage(response)
		return
	}

	url := fmt.Sprintf("%s?s=%s&f=%s&e=%s", b.fetchUrl, req.Message, "sd2t2ohlcv", "csv")
	resp, err := http.Get(url)
	if err != nil {
		response.Message = fmt.Sprintf("Error querying %s", req.Message)
		b.client.PublishResponseMessage(response)
		return
	}

	defer resp.Body.Close()

	csvReader := csv.NewReader(resp.Body)
	content, err := csvReader.Read()

	if err != nil {
		response.Message = fmt.Sprintf("Error reading response for  %s", req.Message)
		b.client.PublishResponseMessage(response)
		return
	}

	//check if content != "N/D"
	if content[3] == "N/D" {
		response.Message = fmt.Sprintf("Can't find info on %s", req.Message)
		b.client.PublishResponseMessage(response)
		return
	}
	response.Message = fmt.Sprintf("%s quote is $%s per share", req.Message, content[3])

	//open = 3
	b.client.PublishResponseMessage(response)
}
