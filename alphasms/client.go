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
	println("hello, getting a call...")
	return &BalanceResponse{
		Balance: 7945.6100,
		Error:   0,
		Date:    "2024-08-16 00:00:00",
	}, nil
}
