package weread

import (
	"fmt"
	"io"
	"net/http"
)

const BookURL = "https://i.weread.qq.com/user/notebooks"

type Client struct {
	Cookie string
}

func NewClient(cookie string) *Client {
	return &Client{Cookie: cookie}
}

func (c *Client) GetBooks() ([]Book, error) {
	req, err := http.NewRequest("GET", BookURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error while creating request: %w", err)
	}
	req.Header.Set("Cookie", c.Cookie)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error while getting books: %w", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error while reading books body: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error while getting books: %s, body: %s", resp.Status, body)
	}

	books, err := ParseBooks(body)
	if err != nil {
		return nil, fmt.Errorf("error while parsing books: %w", err)
	}

	return books, nil
}
