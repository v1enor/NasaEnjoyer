package domain

type APOD struct {
	Date           string `json:"date"`
	CopyRight      string `json:"copyright"`
	Explanation    string `json:"explanation"`
	HDURL          string `json:"hdurl"`
	MediaType      string `json:"media_type"`
	ServiceVersion string `json:"service_version"`
	Title          string `json:"title"`
	URL            string `json:"url"`
	ImagePath      string `json:"imagpath"`
}
