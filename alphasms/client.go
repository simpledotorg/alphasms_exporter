package alphasms

type Client struct {
	APIKey string
}

type BalanceResponse struct {
	Balance float64 `json:"balance"`
	Error   int     `json:"error"`
	Date    string  `json:"date"`
}

func (c *Client) GetUserBalance() (*BalanceResponse, error) {
	// Implement AlphaSMS API call here
	return nil, nil
}
