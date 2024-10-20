package weatherapi

import "fmt"

type Client struct {
	APIKey string
}

type NewAPIClientOpts struct {
	APIKey string
}

func NewAPIClient(opts NewAPIClientOpts) (*Client, error) {
	if err := opts.Validate(); err != nil {
		return nil, err
	}

	return &Client{
		APIKey: opts.APIKey,
	}, nil
}

func (opts NewAPIClientOpts) Validate() error {
	if opts.APIKey == "" {
		return fmt.Errorf("APIKey is required")
	}
	return nil
}
