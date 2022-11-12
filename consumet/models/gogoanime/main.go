package gogoanime

type Response interface {
	Info | Search | Watch
}

type Episode struct {
	Number uint32 `json:"number"`
	Id     string `json:"id"`
}

// GET https://api.consumet.org/anime/gogoanime/info/{id}
type Info struct {
	Episodes []Episode
}

type Source struct {
	Quality string `json:"quality"`
	Url     string `json:"url"`
}

// GET https://api.consumet.org/anime/gogoanime/watch/{id}
type Watch struct {
	Sources []Source `json:"sources"`
}

type Anime struct {
	EpisodeNumber int `json:"episodeNumber"`

	Title string `json:"title"`
	Id    string `json:"id"`
}

// GET https://api.consumet.org/anime/gogoanime/recent-episodes
// GET https://api.consumet.org/anime/gogoanime/search/{query}
type Search struct {
	Results []Anime `json:"results"`
}
