// models/coin.go

package models

// Localization representa informações de localização para uma criptomoeda
type Localization struct {
	EN string `json:"en"`
	DE string `json:"de"`
	ES string `json:"es"`
	FI string `json:"fi"`
	FR string `json:"fr"`
	IT string `json:"it"`
	KO string `json:"ko"`
	NL string `json:"nl"`
	PL string `json:"pl"`
	PT string `json:"pt"`
	RU string `json:"ru"`
	ZH string `json:"zh"`
}

// Image representa URLs de imagem para uma criptomoeda
type Image struct {
	Thumb string `json:"thumb"`
	Small string `json:"small"`
	Large string `json:"large"`
}

// CurrentPrice representa o preço atual de uma criptomoeda
type CurrentPrice struct {
	USD float64 `json:"usd"`
}

// MarketCap representa a capitalização de mercado de uma criptomoeda
type MarketCap struct {
	USD float64 `json:"usd"`
}

// TotalVolume representa o volume total de uma criptomoeda
type TotalVolume struct {
	USD float64 `json:"usd"`
}

// ATH representa o valor máximo histórico de uma criptomoeda
type ATH struct {
	USD float64 `json:"usd"`
}

// ATL representa o valor mínimo histórico de uma criptomoeda
type ATL struct {
	USD float64 `json:"usd"`
}

// CoinInfo representa informações sobre uma criptomoeda
type CoinInfo struct {
	ID                 string       `json:"id"`
	Symbol             string       `json:"symbol"`
	Name               string       `json:"name"`
	Localization       Localization `json:"localization"`
	Image              Image        `json:"image"`
	MarketData         MarketData   `json:"market_data"`
	CommunityData      Community    `json:"community_data"`
	DeveloperData      Developer    `json:"developer_data"`
	PublicInterestData interface{}  `json:"public_interest_data"`
}

// MarketData representa dados de mercado para uma criptomoeda
type MarketData struct {
	CurrentPrice                 CurrentPrice `json:"current_price"`
	TotalValueLocked             float64      `json:"total_value_locked"`
	MarketCap                    MarketCap    `json:"market_cap"`
	MarketCapRank                int          `json:"market_cap_rank"`
	TotalVolume                  TotalVolume  `json:"total_volume"`
	High24h                      CurrentPrice `json:"high_24h"`
	Low24h                       CurrentPrice `json:"low_24h"`
	PriceChange24h               float64      `json:"price_change_24h"`
	PriceChangePercentage24h     float64      `json:"price_change_percentage_24h"`
	PriceChangePercentage7d      float64      `json:"price_change_percentage_7d"`
	PriceChangePercentage14d     float64      `json:"price_change_percentage_14d"`
	PriceChangePercentage30d     float64      `json:"price_change_percentage_30d"`
	PriceChangePercentage60d     float64      `json:"price_change_percentage_60d"`
	PriceChangePercentage200d    float64      `json:"price_change_percentage_200d"`
	PriceChangePercentage1y      float64      `json:"price_change_percentage_1y"`
	MarketCapChange24h           float64      `json:"market_cap_change_24h"`
	MarketCapChangePercentage24h float64      `json:"market_cap_change_percentage_24h"`
	CirculatingSupply            float64      `json:"circulating_supply"`
	TotalSupply                  float64      `json:"total_supply"`
	MaxSupply                    float64      `json:"max_supply"`
	ATH                          ATH          `json:"ath"`
	ATL                          ATL          `json:"atl"`
}

// Community representa dados da comunidade para uma criptomoeda
type Community struct {
	FacebookLikes         int     `json:"facebook_likes"`
	TwitterFollowers      int     `json:"twitter_followers"`
	RedditAveragePosts    float64 `json:"reddit_average_posts_48h"`
	RedditAverageComments float64 `json:"reddit_average_comments_48h"`
	RedditSubscribers     int     `json:"reddit_subscribers"`
	RedditAccountsActive  int     `json:"reddit_accounts_active_48h"`
}

// Developer representa dados do desenvolvedor para uma criptomoeda
type Developer struct {
	Forks                   int         `json:"forks"`
	Stars                   int         `json:"stars"`
	Subscribers             int         `json:"subscribers"`
	TotalIssues             int         `json:"total_issues"`
	ClosedIssues            int         `json:"closed_issues"`
	PullRequestsMerged      int         `json:"pull_requests_merged"`
	PullRequestContributors int         `json:"pull_request_contributors"`
	CodeAdditionsDeletions  CodeChanges `json:"code_additions_deletions_4_weeks"`
	CommitCount4Weeks       int         `json:"commit_count_4_weeks"`
}

// CodeChanges representa mudanças no código para uma criptomoeda
type CodeChanges struct {
	Additions int `json:"additions"`
	Deletions int `json:"deletions"`
}

// models/coin.go
