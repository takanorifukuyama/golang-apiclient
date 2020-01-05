package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"time"
)

func main() {
	endpoint := "https://api-dev.i-kasa.com/db/v1/transaction/deleted"
	ikasa_id := "206cc453-2df0-11ea-ab85-f268b5f830dc"
	u, err := url.Parse(endpoint)
	if err != nil {
		panic(err)
	}
	u.Path = path.Join(u.Path, ikasa_id)
	fmt.Println(u.String())
	req, _ := http.NewRequest("GET", u.String(), nil)

	client := new(http.Client)
	rsp, err := client.Do(req)
	if err != nil {
		fmt.Printf("client do err : %s", err)
		panic(err)
	}

	var transactions []TTransaction
	if err := decodeBody(rsp, &transactions); err != nil {
		fmt.Printf("decode err : %s", err)
		panic(err)
	}
	for _, v := range transactions {
		fmt.Println(v.TransactionID)
	}
}

func decodeBody(resp *http.Response, out interface{}) error {
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(out)
	return nil
}

type TTransaction struct {
	TransactionID string `json:"transaction_id"`
	IkasaID       string
	Ksid          int
	RentedAt      time.Time
	RentedScd     int
	ReturnedScd   int
	DeletedAt     time.Time
}
