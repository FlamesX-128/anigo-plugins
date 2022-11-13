package fembed

type Response struct {
	Data []struct {
		Quality string `json:"label"`
		Url     string `json:"file"`
	} `json:"data"`
}
