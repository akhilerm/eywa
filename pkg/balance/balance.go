package balance

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/anvari1313/splitwise.go"

	"eywa/pkg/shortener"
	"eywa/pkg/upi"
	"eywa/pkg/urlserver"
)

const SPLITWISEAPIKEYENV = "SPLIT_KEY"

type Config struct {
	PayeeAddress string
	PayeeName    string
	Port         string
	StartServer  bool
	// currently not used
	removeExpiredLinks bool
	sendWAMessage      bool
}

func Run(config Config) error {
	splitWiseAPIKey := os.Getenv(SPLITWISEAPIKEYENV)
	if len(splitWiseAPIKey) == 0 {
		return fmt.Errorf("splitwise API key not provided")
	}

	paymentRequest := upi.NewPaymentRequest().
		WithPayeeAddress(config.PayeeAddress).
		WithPayeeName(config.PayeeName)

	auth := splitwise.NewAPIKeyAuth("gs0I5pOnk2YXD7P5xBfaKcd2seYPC07Lz7SuU4rd")
	client := splitwise.NewClient(auth)

	friends, err := client.Friends(context.Background())
	if err != nil {
		return fmt.Errorf("error getting friends list from splitwise. Err: %v", err)
	}

	for _, friend := range friends {
		name := friend.FirstName + " " + friend.LastName
		for _, balance := range friend.Balance {
			if balance.CurrencyCode != upi.INR {
				continue
			}
			fbalance, err := strconv.ParseFloat(balance.Amount, 64)
			if err != nil {
				log.Printf("error parsing balance for %s. Err: %v", name, err)
			}
			if fbalance <= 0 {
				continue
			}
			paymentRequest.
				WithAmount(balance.Amount).
				WithTransactionNote(
					fmt.Sprintf(
						"Balance settlement for the month of %s",
						time.Now().AddDate(0, -1, 0).Month(),
					))
			paymentURL, err := paymentRequest.GenerateUPIPaymentURL()
			if err != nil {
				log.Printf("error creating payment URL for %s for payment request: %v. Err: %v",
					name, paymentRequest, err)
			}

			hash, err := shortener.Short.Add(paymentURL)
			if err != nil {
				log.Printf("failed to add shortened data for %s for payment request: %v. Err: %v",
					name, paymentRequest, err)
			}
			if config.StartServer && err == nil {
				log.Printf("%s: http://0.0.0.0%s/%s", name, config.Port, hash)
			} else {
				// if server is not running log the payment URL
				log.Printf("%s: %s", name, paymentURL)
			}
		}
	}

	if !config.StartServer {
		return nil
	}

	c := urlserver.Config{
		Port: config.Port,
	}

	s := urlserver.NewServer(c)
	return s.Start()
}
