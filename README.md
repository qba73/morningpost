# The Morning Post

`morningpost` is a cli tool that curates a little “morning newspaper” for you, by scraping news websites and selecting stories.

## Using `morningpost`

- Register and get an api key from the [Open Platform](https://open-platform.theguardian.com/access/)
- Export the env var `API_KEY_GUARDIAN`
- Install `morningpost`
- Get latest 10 news:

```bash
./morningpost
[
  {
    "title": "Australia news live: Liberal women’s committee backs Deeming; Thorpe urges more support for assault survivors",
    "link": "https://www.theguardian.com/australia-news/live/2023/jun/16/australia-news-david-van-lidia-thorpe-amanda-stoker-indigenous-voice-cost-of-living-interest-rates-james-hardy",
    "date": "2023-06-16T07:29:43Z"
  },
  {
    "title": "The Ashes 2023: England v Australia, first Test, day one – live",
    "link": "https://www.theguardian.com/sport/live/2023/jun/16/the-ashes-2023-england-vs-australia-live-updates-first-test-cricket-eng-v-aus-latest-score-day-one-egdbaston",
    "date": "2023-06-16T07:27:52Z"
  },
  ...
]
```

## Installation

### Using `brew` package manager

```bash
brew tap qba73/tap
brew install morningpost
```

### From source

- Clone the repo.
- Build the `morningpost` binary.

```bash
go build -o morningpost ./cmd/morningpost/main.go
```

## Description

Full project description and instructions [link](./INSTRUCTIONS.md).
