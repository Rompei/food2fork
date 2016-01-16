package food2fork

type searchResponse struct {
	Count   int      `json:"count"`
	Recipes []Recipe `json:"recipes"`
}

type getResponse struct {
	Recipe Recipe `json:"recipe"`
}

// Recipe is recipe object.
type Recipe struct {
	RecipeID     string   `json:"recipe_id"`
	ImageURL     string   `json:"image_url"`
	SourceURL    string   `json:"source_url"`
	F2FURL       string   `json:"f2f_url"`
	Title        string   `json:"title"`
	Publisher    string   `josn:"publisher"`
	PublisherURL string   `json:"publisher_url"`
	SocialRank   float64  `json:"social_rank"`
	Page         int      `json:"page"`
	Ingredients  []string `json:"ingredients"`
}
