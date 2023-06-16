package morningpost

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

type guardianResponse struct {
	Response struct {
		Status      string
		UserTier    string
		Total       int
		StartIndex  int
		PageSize    int
		CurrentPage int
		Pages       int
		OrderBy     string
		Results     []struct {
			ID                 string
			Type               string
			SectionID          string
			SectionName        string
			WebPublicationDate string
			WebTitle           string
			WebURL             string
			APIURL             string
			IsHosted           bool
			PillarID           string
			PillarName         string
		} `json:"results"`
	} `json:"response"`
}

type News struct {
	Title string `json:"title"`
	Link  string `json:"link"`
	Date  string `json:"date"`
}

type Client struct {
	HTTPClient *http.Client
	BaseURL    string
	APIKey     string
}

func NewClient(apikey string) *Client {
	c := Client{
		HTTPClient: http.DefaultClient,
		BaseURL:    "https://content.guardianapis.com",
		APIKey:     apikey,
	}
	return &c
}

func (c *Client) GetNews() ([]News, error) {
	url := fmt.Sprintf("%s/search?sectionId=news&format=json&showFields=headline,short-url", c.BaseURL)
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
	if err != nil {
		return []News{}, err
	}
	req.Header.Set("api-key", c.APIKey)
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return []News{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return []News{}, fmt.Errorf("response status error: %d", resp.StatusCode)
	}

	var gr guardianResponse
	if err := json.NewDecoder(resp.Body).Decode(&gr); err != nil {
		return []News{}, err
	}

	news := make([]News, 0, len(gr.Response.Results))
	for _, v := range gr.Response.Results {
		n := News{
			Title: v.WebTitle,
			Link:  v.WebURL,
			Date:  v.WebPublicationDate,
		}
		news = append(news, n)
	}
	return news, nil
}

func FetchNews() ([]News, error) {
	apiKey, ok := os.LookupEnv("API_KEY_GUARDIAN")
	if !ok {
		return nil, errors.New("api key `API_KEY_GUARDIAN` not set")
	}
	c := NewClient(apiKey)
	news, err := c.GetNews()
	if err != nil {
		return nil, fmt.Errorf("fetching news, %w", err)
	}
	return news, err
}

func Main() int {
	news, err := FetchNews()
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		return 1
	}
	n, err := json.MarshalIndent(news, "", "  ")
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		return 1
	}

	fmt.Fprintf(os.Stdout, "%s\n", string(n))
	return 0
}
