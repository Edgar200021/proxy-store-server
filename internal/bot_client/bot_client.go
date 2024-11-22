package bot_client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"proxyStoreServer/internal/config"
)

type BotClient struct {
	telegramUrl    string
	webhookUrl     string
	apiSecretToken string
	httpClient     *http.Client
}

func New(botConfig *config.BotConfig, applicationConfig *config.ApplicationConfig) *BotClient {
	httpClient := &http.Client{
		Timeout: botConfig.RequestTimeout,
	}

	return &BotClient{
		telegramUrl:    fmt.Sprintf("%s%s", botConfig.TelegramUrl, botConfig.Token),
		webhookUrl:     fmt.Sprintf("%s:%d/bot", applicationConfig.Host, applicationConfig.Port),
		apiSecretToken: botConfig.BotApiSecretToken,
		httpClient:     httpClient,
	}
}

func (b *BotClient) SetWebHook() error {
	req, err := http.NewRequest("GET", b.telegramUrl+"/setWebhook", nil)
	if err != nil {
		return err
	}

	req.Header.Add("Accept", "application/json")

	q := req.URL.Query()
	q.Add("url", b.webhookUrl)
	q.Add("secret_token", b.apiSecretToken)

	req.URL.RawQuery = q.Encode()

	resp, err := b.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to set webhook, status code: %d", resp.StatusCode)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	type Data struct {
		Ok          bool   `json:"ok"`
		ErrorCode   int    `json:"error_code"`
		Description string `json:"description"`
	}

	var data Data

	if err := json.Unmarshal(bytes, &data); err != nil {
		return err
	}

	if !data.Ok {
		return fmt.Errorf("%s", data.Description)
	}

	return nil
}
