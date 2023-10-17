package botClient

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// HTTPClient is the type needed for the bot to perform HTTP requests.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// BotAPI allows you to interact with the Telegram Bot API.
type BotAPI struct {
	Token  string `json:"token"`
	Debug  bool   `json:"debug"`
	Buffer int    `json:"buffer"`

	Self            User       `json:"-"`
	Client          HTTPClient `json:"-"`
	shutdownChannel chan interface{}

	apiEndpoint string
}

// NewBotAPI creates a new BotAPI instance.
//
// It requires a token, provided by @BotFather on Telegram.
func NewBotAPI(token string) (*BotAPI, error) {
	return NewBotAPIWithClient(token, APIEndpoint, &http.Client{})
}

// NewBotAPIWithClient creates a new BotAPI instance
// and allows you to pass a http.Client.
//
// It requires a token, provided by @BotFather on Telegram and API endpoint.
func NewBotAPIWithClient(token, apiEndpoint string, client HTTPClient) (*BotAPI, error) {
	bot := &BotAPI{
		Token:           token,
		Client:          client,
		Buffer:          100,
		shutdownChannel: make(chan interface{}),

		apiEndpoint: apiEndpoint,
	}

	self, err := bot.GetMe()
	if err != nil {
		return nil, err
	}

	bot.Self = self

	return bot, nil
}

// GetMe fetches the currently authenticated bot.
//
// This method is called upon creation to validate the token,
// and so you may get this data from BotAPI.Self without the need for
// another request.
func (bot *BotAPI) GetMe() (User, error) {
	resp, err := bot.MakeRequest("getMe", nil)
	if err != nil {
		return User{}, err
	}

	var user User
	err = json.Unmarshal(resp.Result, &user)

	return user, err
}

// MakeRequest makes a request to a specific endpoint with our token.
func (bot *BotAPI) MakeRequest(endpoint string, params Params) (*APIResponse, error) {
	if bot.Debug {
		log.Printf("Endpoint: %s, params: %v\n", endpoint, params)
	}

	method := fmt.Sprintf(bot.apiEndpoint, bot.Token, endpoint)

	values := buildParams(params)

	req, err := http.NewRequest("POST", method, strings.NewReader(values.Encode()))
	if err != nil {
		return &APIResponse{}, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := bot.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var apiResp APIResponse
	bytes, err := bot.decodeAPIResponse(resp.Body, &apiResp)
	if err != nil {
		return &apiResp, err
	}

	if bot.Debug {
		log.Printf("Endpoint: %s, response: %s\n", endpoint, string(bytes))
	}

	if !apiResp.Ok {
		var parameters ResponseParameters

		if apiResp.Parameters != nil {
			parameters = *apiResp.Parameters
		}

		return &apiResp, &Error{
			Code:               apiResp.ErrorCode,
			Message:            apiResp.Description,
			ResponseParameters: parameters,
		}
	}

	return &apiResp, nil
}

func buildParams(in Params) url.Values {
	if in == nil {
		return url.Values{}
	}

	out := url.Values{}

	for key, value := range in {
		out.Set(key, value)
	}

	return out
}

// decodeAPIResponse decode response and return slice of bytes if debug enabled.
// If debug disabled, just decode http.Response.Body stream to APIResponse struct
// for efficient memory usage
func (bot *BotAPI) decodeAPIResponse(responseBody io.Reader, resp *APIResponse) ([]byte, error) {
	if !bot.Debug {
		dec := json.NewDecoder(responseBody)
		err := dec.Decode(resp)
		return nil, err
	}

	// if debug, read response body
	data, err := ioutil.ReadAll(responseBody)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, resp)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// GetUpdatesChan starts and returns a channel for getting updates.
func (bot *BotAPI) GetUpdatesChan(config UpdateConfig) UpdatesChannel {
	ch := make(chan Update, bot.Buffer)

	go func() {
		for {
			select {
			case <-bot.shutdownChannel:
				close(ch)
				return
			default:
			}

			updates, err := bot.GetUpdates(config)
			if err != nil {
				log.Println(err)
				log.Println("Failed to get updates, retrying in 3 seconds...")
				time.Sleep(time.Second * 3)

				continue
			}

			for _, update := range updates {
				if update.UpdateID >= config.Offset {
					config.Offset = update.UpdateID + 1
					ch <- update
				}
			}
		}
	}()

	return ch
}

// GetUpdates fetches updates.
// If a WebHook is set, this will not return any data!
//
// Offset, Limit, Timeout, and AllowedUpdates are optional.
// To avoid stale items, set Offset to one higher than the previous item.
// Set Timeout to a large number to reduce requests, so you can get updates
// instantly instead of having to wait between requests.
func (bot *BotAPI) GetUpdates(config UpdateConfig) ([]Update, error) {
	resp, err := bot.Request(config)
	if err != nil {
		return []Update{}, err
	}

	var updates []Update
	err = json.Unmarshal(resp.Result, &updates)

	return updates, err
}

// Request sends a Chattable to Telegram, and returns the APIResponse.
func (bot *BotAPI) Request(c Chattable) (*APIResponse, error) {
	params, err := c.params()
	if err != nil {
		return nil, err
	}

	return bot.MakeRequest(c.method(), params)
}

// Send will send a Chattable item to Telegram and provides the
// returned Message.
func (bot *BotAPI) Send(c Chattable) (Message, error) {
	resp, err := bot.Request(c)
	if err != nil {
		return Message{}, err
	}

	var message Message
	err = json.Unmarshal(resp.Result, &message)

	return message, err
}
