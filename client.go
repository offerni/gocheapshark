package gocheapshark

type Client struct {
	BaseURL string
}

type NewClientOpts struct {
	BaseUrl string
}

func (opts *NewClientOpts) validate() error {
	if opts.BaseUrl == "" {
		return ErrNoBaseURL
	}

	return nil
}

func NewClient(opts NewClientOpts) (*Client, error) {
	if err := opts.validate(); err != nil {
		return nil, err
	}

	return &Client{
		BaseURL: opts.BaseUrl,
	}, nil
}
