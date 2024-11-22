package cryptomus_client

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"io"
	"net/http"
	"proxyStoreServer/internal/config"
	"proxyStoreServer/internal/dto"
)

type CryptomusClient struct {
	config     *config.CryptomusConfig
	httpClient *http.Client
}

func New(config *config.CryptomusConfig) *CryptomusClient {
	httpClient := http.Client{
		Timeout: config.RequestTimeout,
	}

	return &CryptomusClient{
		config,
		&httpClient,
	}
}

func (c *CryptomusClient) CreateInvoice(payload *dto.CreateCryptomusInvoiceRequest) (*dto.CreateCryptomusInvoiceResponse, error) {
	req, err := http.NewRequest("POST", c.config.Url+"v1/payment", nil)
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req.Header = c.generateHeader(data)

	response, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var responseData dto.CreateCryptomusInvoiceResponse

	if err := json.Unmarshal(bytes, &responseData); err != nil {
		return nil, err
	}

	return &responseData, nil
}

func (c *CryptomusClient) generateHeader(data []byte) http.Header {
	sign := md5.Sum([]byte(base64.StdEncoding.EncodeToString(data) + c.config.ApiKey))

	header := http.Header{
		"Content-Type": []string{"application/json"},
		"sign":         []string{hex.EncodeToString(sign[:])},
		"merchant":     []string{c.config.MerchantId},
	}

	return header
}
