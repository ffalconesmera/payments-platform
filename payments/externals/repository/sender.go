package repository

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ffalconesmera/payments-platform/payments/externals/dto"
)

func SendRequestApiExternal[V dto.JSONMerchant | dto.BankPayment | dto.BankRefund](endpoint string, method string, bodyString string, receiver *V) error {
	body := []byte(bodyString)
	r, err := http.NewRequest(method, endpoint, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	r.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return errors.New("provider not available")
	}

	defer res.Body.Close()
	json.NewDecoder(res.Body).Decode(&receiver)
	return nil
}
