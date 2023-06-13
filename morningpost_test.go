package morningpost_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/qba73/morningpost"
)

func TestGetNews(t *testing.T) {
	t.Parallel()

	ts := newTestServerWithPathValidator(
		res_guardian,
		"/search?sectionId=news&format=json&showFields=headline,short-url",
		t,
	)
	client := newClient(ts.URL, t)

	_, err := client.GetNews()
	if err != nil {
		t.Fatal(err)
	}

}

func newClient(baseURL string, t *testing.T) *morningpost.Client {
	t.Helper()
	c := morningpost.NewClient()
	c.BaseURL = baseURL
	return c
}

func newTestServerWithPathValidator(respBody string, wantURL string, t *testing.T) *httptest.Server {
	t.Helper()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotReqURI := r.RequestURI
		verifyURL(wantURL, gotReqURI, t)
		_, err := w.Write([]byte(respBody))
		if err != nil {
			t.Fatal(err)
		}
	}))
	return ts
}

// verifyURL verifies if URLs are equal.
func verifyURL(wantURL, gotURL string, t *testing.T) {
	wantU, err := url.Parse(wantURL)
	if err != nil {
		t.Fatalf("error parsing URL %q, %v", wantURL, err)
	}
	gotU, err := url.Parse(gotURL)
	if err != nil {
		t.Fatalf("error parsing URL %q, %v", wantURL, err)
	}
	// Verify if paths of both URLs are the same.
	if wantU.Path != gotU.Path {
		t.Fatalf("want %q, got %q", wantU.Path, gotU.Path)
	}

	wantQuery, err := url.ParseQuery(wantU.RawQuery)
	if err != nil {
		t.Fatal(err)
	}
	gotQuery, err := url.ParseQuery(gotU.RawQuery)
	if err != nil {
		t.Fatal(err)
	}

	if !cmp.Equal(wantQuery, gotQuery) {
		t.Fatalf("query params do not match, \n%s", cmp.Diff(wantQuery, gotQuery))
	}
}

var (
	res_guardian = `{
		"response": {
			"status": "ok",
			"userTier": "developer",
			"total": 2430672,
			"startIndex": 1,
			"pageSize": 10,
			"currentPage": 1,
			"pages": 243068,
			"orderBy": "newest",
			"results": [
				{
					"id": "politics/2023/jun/11/senior-tories-tell-boris-johnson-and-allies-to-shut-up-and-go-away",
					"type": "article",
					"sectionId": "politics",
					"sectionName": "Politics",
					"webPublicationDate": "2023-06-11T19:07:10Z",
					"webTitle": "Senior Tories tell Boris Johnson and allies to ‘shut up and go away’",
					"webUrl": "https://www.theguardian.com/politics/2023/jun/11/senior-tories-tell-boris-johnson-and-allies-to-shut-up-and-go-away",
					"apiUrl": "https://content.guardianapis.com/politics/2023/jun/11/senior-tories-tell-boris-johnson-and-allies-to-shut-up-and-go-away",
					"isHosted": false,
					"pillarId": "pillar/news",
					"pillarName": "News"
				},
				{
					"id": "environment/2023/jun/11/yorkshire-water-bosss-decision-to-forgo-bonus-labelled-hollow-by-union",
					"type": "article",
					"sectionId": "environment",
					"sectionName": "Environment",
					"webPublicationDate": "2023-06-11T19:00:00Z",
					"webTitle": "Yorkshire Water boss’s decision to forgo bonus labelled ‘hollow’ by union",
					"webUrl": "https://www.theguardian.com/environment/2023/jun/11/yorkshire-water-bosss-decision-to-forgo-bonus-labelled-hollow-by-union",
					"apiUrl": "https://content.guardianapis.com/environment/2023/jun/11/yorkshire-water-bosss-decision-to-forgo-bonus-labelled-hollow-by-union",
					"isHosted": false,
					"pillarId": "pillar/news",
					"pillarName": "News"
				},
				{
					"id": "world/2023/jun/11/ukraine-claims-to-have-liberated-two-frontline-villages-in-donetsk",
					"type": "article",
					"sectionId": "world",
					"sectionName": "World news",
					"webPublicationDate": "2023-06-11T17:56:28Z",
					"webTitle": "Ukraine claims to have liberated three frontline villages in Donetsk",
					"webUrl": "https://www.theguardian.com/world/2023/jun/11/ukraine-claims-to-have-liberated-two-frontline-villages-in-donetsk",
					"apiUrl": "https://content.guardianapis.com/world/2023/jun/11/ukraine-claims-to-have-liberated-two-frontline-villages-in-donetsk",
					"isHosted": false,
					"pillarId": "pillar/news",
					"pillarName": "News"
				}
			]
		}
	}`
)
