package morningpost

import (
	"net/http"
	"time"
)

type GuardianResponse struct {
	Status      string `json:"status"`
	UserTier    string `json:"userTier"`
	Total       int    `json:"total"`
	StartIndex  int    `json:"startIndex"`
	PageSize    int    `json:"pageSize"`
	CurrentPage int    `json:"currentPage"`
	Pages       int    `json:"pages"`
	OrderBy     string `json:"orderBy"`
	Results     []struct {
		ID                 string    `json:"id"`
		Type               string    `json:"type"`
		SectionID          string    `json:"sectionId"`
		SectionName        string    `json:"sectionName"`
		WebPublicationDate time.Time `json:"webPublicationDate"`
		WebTitle           string    `json:"webTitle"`
		WebURL             string    `json:"webUrl"`
		APIURL             string    `json:"apiUrl"`
		IsHosted           bool      `json:"isHosted"`
		PillarID           string    `json:"pillarId"`
		PillarName         string    `json:"pillarName"`
	} `json:"results"`
}

type News struct {
	Title string
	Link  string
	Date  string
}

type Client struct {
	HTTPClient *http.Client
	BaseURL    string
}

func NewClient() *Client {
	c := Client{
		HTTPClient: http.DefaultClient,
	}
	return &c
}

func (c *Client) GetNews() ([]News, error) {
	return nil, nil
}
