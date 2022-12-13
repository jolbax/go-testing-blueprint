package fruits_client

import "github.com/go-resty/resty/v2"

const (
	MockAPIGetURL  = "https://run.mocky.io/v3/d959e861-a8b3-44be-b3a6-ee58fb42f05e"
	MockAPIPostURL = "https://run.mocky.io/v3/8090b334-31b1-45eb-bb8e-6b821b0341ee"
)

type FruitsClient struct {
	client *resty.Client
}

func (f FruitsClient) Get() (*FruitsResponse, error) {
	resp, err := f.client.R().SetResult(Fruits{}).Get(MockAPIGetURL)
	if err != nil {
		return &FruitsResponse{nil, 504}, err
	}

	fruits := resp.Result().(*Fruits)
	return &FruitsResponse{fruits, resp.StatusCode()}, nil
}

func (f FruitsClient) Post() (*FruitsResponse, error) {
	resp, err := f.client.R().SetResult(Fruits{}).Post(MockAPIPostURL)
	if err != nil {
		return &FruitsResponse{nil, 504}, err
	}

	fruits := resp.Result().(*Fruits)

	return &FruitsResponse{fruits, resp.StatusCode()}, nil
}

func NewFruitsClient() *FruitsClient {
	client := resty.New()
	return &FruitsClient{client: client}
}
