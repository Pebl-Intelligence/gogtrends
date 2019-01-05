package gogtrends

// TODO: export only needed structs, check validation and omitempty for every field

const (
	gAPI = "https://trends.google.com/trends/api"

	gDaily    = "/dailytrends"
	gRealtime = "/realtimetrends"

	gSExplore     = "/explore"
	gSCategories  = "/explore/pickers/category"
	gSGeo         = "/explore/pickers/geo"
	gSRelated     = "/widgetdata/relatedsearches"
	gSSuggestions = "/autocomplete"
	gSIntOverTime = "/widgetdata/multiline"
	gSIntOverReg  = "/widgetdata/comparedgeo"

	paramHl    = "hl"
	paramCat   = "cat"
	paramGeo   = "geo"
	paramReq   = "req"
	paramTZ    = "tz"
	paramToken = "token"

	intOverTimeWidgetID = "TIMESERIES"
	intOverRegionID     = "GEO_MAP"

	errParsing           = "failed to parse json"
	errRequestFailed     = "failed to perform http request to API"
	errReqDataF          = "request data: code = %d, status = %s, body = %s"
	errInvalidCategory   = "invalid category param"
	errInvalidLocation   = "invalid location param"
	errInvalidRequest    = "invalid request param"
	errInvalidWidgetType = "invalid widget type"

	timeLayoutFull = "2006-01-02T15:04:05Z07:00" // https://golang.org/src/time/format.go
)

var (
	defaultParams = map[string]string{
		paramTZ:  "0",
		paramCat: "all",
		"fi":     "0",
		"fs":     "0",
		paramGeo: "US",
		paramHl:  "EN",
		"ri":     "300",
		"rs":     "20",
	}
	availableLocations = map[string]string{
		"AU": "Australia",
		"AT": "Austria",
		"AR": "Argentina",
		"BE": "Belgium",
		"BR": "Brazil",
		"GB": "United Kingdom",
		"HU": "Hungary",
		"VN": "Vietnam",
		"DE": "Germany",
		"HK": "Hong Kong",
		"GR": "Greece",
		"DK": "Denmark",
		"EG": "Egypt",
		"IL": "Israel",
		"IN": "India",
		"IE": "Ireland",
		"IT": "Italy",
		"CA": "Canada",
		"KE": "Kenia",
		"CO": "Columbia",
		"MY": "Malaysia",
		"MX": "Mexico",
		"NG": "Nigeria",
		"NL": "Netherlands",
		"NZ": "New Zeland",
		"NO": "Norway",
		"PL": "Poland",
		"PT": "Portugal",
		"KR": "Korean Republic",
		"RU": "Russia",
		"RO": "Romania",
		"SA": "Saudi Arabia",
		"SG": "Singapore",
		"US": "United States",
		"TH": "Thailand",
		"TW": "Taiwan",
		"TR": "Turkey",
		"UA": "Ukraine",
		"PH": "Philippines",
		"FI": "Finland",
		"FR": "France",
		"CZ": "Czech Republic",
		"CL": "Chili",
		"CH": "Switzerland",
		"SE": "Sweden",
		"ZA": "Republic of South Africa",
		"JP": "Japan",
	}
	availableCategories = map[string]string{
		"all": "all",
		"b":   "business",
		"h":   "main news",
		"m":   "health",
		"t":   "science and technics",
		"e":   "entertainment",
		"s":   "sport",
	}
)

type dailyOut struct {
	Default *trendingSearchesDays `json:"default"`
}

type trendingSearchesDays struct {
	Searches []*trendingSearchDays `json:"trendingSearchesDays"`
}

type trendingSearchDays struct {
	FormattedDate string            `json:"formattedDate"`
	Searches      []*TrendingSearch `json:"trendingSearches"`
}

type TrendingSearch struct {
	Title            *SearchTitle     `json:"title"`
	FormattedTraffic string           `json:"formattedTraffic"`
	Image            *SearchImage     `json:"image"`
	Articles         []*SearchArticle `json:"articles"`
}

type SearchTitle struct {
	Query string `json:"query"`
}

type SearchImage struct {
	NewsURL  string `json:"newsUrl"`
	Source   string `json:"source"`
	ImageURL string `json:"imageUrl"`
}

type SearchArticle struct {
	Title   string       `json:"title"`
	TimeAgo string       `json:"timeAgo"`
	Source  string       `json:"source"`
	Image   *SearchImage `json:"image"`
	URL     string       `json:"url"`
	Snippet string       `json:"snippet"`
}

type realtimeOut struct {
	StorySummaries *storySummary `json:"storySummaries"`
}

type storySummary struct {
	TrendingStories []*TrendingStory `json:"trendingStories"`
}

type TrendingStory struct {
	Title    string             `json:"title"`
	Image    *SearchImage       `json:"image"`
	Articles []*TrendingArticle `json:"articles"`
}

type TrendingArticle struct {
	Title   string `json:"articleTitle"`
	URL     string `json:"url"`
	Source  string `json:"source"`
	Time    string `json:"time"`
	Snippet string `json:"snippet"`
}

type ExploreRequest struct {
	ComparisonItems []*ComparisonItem `json:"comparisonItem"`
	Category        int               `json:"category"`
	Property        string            `json:"property"`
}

type ComparisonItem struct {
	Keyword string `json:"keyword"`
	Geo     string `json:"geo"`
	Time    string `json:"time"`
}

type ExploreCatTree struct {
	Name     string            `json:"name"`
	ID       int               `json:"id"`
	Children []*ExploreCatTree `json:"children"`
}

type ExploreLocTree struct {
	Name     string            `json:"name"`
	ID       string            `json:"id"`
	Children []*ExploreLocTree `json:"children"`
}

type ExploreOut struct {
	Widgets []*ExploreWidget `json:"widgets"`
}

type ExploreWidget struct {
	Token   string          `json:"token"`
	Type    string          `json:"type"`
	Title   string          `json:"title"`
	ID      string          `json:"id"`
	Request *WidgetResponse `json:"request"`
}

type WidgetResponse struct {
	Geo         interface{}             `json:"geo,omitempty"`
	Time        string                  `json:"time"`
	Resolution  string                  `json:"resolution"`
	Locale      string                  `json:"locale"`
	RequestOpt  RequestOptions          `json:"requestOptions"`
	CompItem    []*WidgetComparisonItem `json:"comparisonItem"`
	KeywordType string                  `json:"keywordType"`
	Metric      []string                `json:"metric"`
	Restriction WidgetComparisonItem    `json:"restriction"`
	Language    string                  `json:"language"`
}

type WidgetComparisonItem struct {
	Geo                            map[string]string   `json:"geo,omitempty"`
	Time                           string              `json:"time,omitempty"`
	ComplexKeywordsRestriction     KeywordsRestriction `json:"complexKeywordsRestriction,omitempty"`
	OriginalTimeRangeForExploreUrl string              `json:"originalTimeRangeForExploreUrl,omitempty"`
}

type KeywordsRestriction struct {
	Keyword []*KeywordRestriction `json:"keyword"`
}

type KeywordRestriction struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type RequestOptions struct {
	Property string `json:"property"`
	Backend  string `json:"backend"`
	Category int    `json:"category"`
}

type MultilineOut struct {
	Default Multiline `json:"default"`
}

type Multiline struct {
	TimelineData []*Timeline `json:"timelineData"`
}

type Timeline struct {
	Time              string   `json:"time"`
	FormattedTime     string   `json:"formattedTime"`
	FormattedAxisTime string   `json:"formattedAxisTime"`
	Value             []int    `json:"value"`
	HasData           []bool   `json:"hasData"`
	FormattedValue    []string `json:"formattedValue"`
}

type GeoOut struct {
	Default Geo `json:"default"`
}

type Geo struct {
	GeoMapData []*GeoMap `json:"geoMapData"`
}

type GeoMap struct {
	GeoCode        string   `json:"geoCode"`
	GeoName        string   `json:"geoName"`
	Value          []int    `json:"value"`
	FormattedValue []string `json:"formattedValue"`
	MaxValueIndex  int      `json:"maxValueIndex"`
	HasData        []bool   `json:"hasData"`
}