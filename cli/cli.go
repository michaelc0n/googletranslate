package cli

import (
	"log"
	"net/http"
	"sync"

	"github.com/Jeffail/gabs"
)

type RequestBody struct {
	SourceLang string
	TargetLang string
	SourceText string
}

const translateUrl = "https://translate.googleapis.com/translate_a/single"

func RequestTranslate(body *RequestBody, str chan string, wg *sync.WaitGroup) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", translateUrl, nil)

	query := req.URL.Query()

	query.Add("client", "gtx")

	query.Add("sl", body.SourceLang)
	query.Add("tl", body.TargetLang)
	query.Add("dt", "t")
	query.Add("q", body.SourceText)

	req.URL.RawQuery = query.Encode()

	if err != nil {
		log.Fatal("1. error:", err)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal("2. error:", err)
	}

	defer res.Body.Close()

	if res.StatusCode == http.StatusTooManyRequests {
		str <- "You have been rate limited, try again later."
		wg.Done()
		return
	}

	parsedJson, err := gabs.ParseJSONBuffer(res.Body)

	if err != nil {
		log.Fatal("3. error:", err)
	}

	nestOne, err := parsedJson.ArrayElement(0)

	if err != nil {
		log.Fatal("4. error:", err)
	}

	nestTwo, err := nestOne.ArrayElement(0)
	if err != nil {
		log.Fatal("5. error:", err)
	}

	translateStr, err := nestTwo.ArrayElement(0)

	if err != nil {
		log.Fatal("6. error:", err)
	}

	str <- translateStr.Data().(string)
	wg.Done()
}
