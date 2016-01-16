package food2fork

import (
	//"fmt"
	"net/url"
)

// SearchRecipe searches recipes.
func (f *Food2Fork) SearchRecipe(params url.Values) (r []Recipe, err error) {
	params = f.cleanValue(params)
	responseCh := make(chan response)
	var sr searchResponse
	f.queryQueue <- query{f.baseURL + "/search", params, &sr, responseCh}
	return sr.Recipes, (<-responseCh).err
}

// GetRecipe gets recipe of recipe id.
func (f *Food2Fork) GetRecipe(rID string) (r Recipe, err error) {
	params := url.Values{}
	params.Add("rId", rID)
	params = f.cleanValue(params)
	responseCh := make(chan response)
	var gr getResponse
	f.queryQueue <- query{f.baseURL + "/get", params, &gr, responseCh}
	return gr.Recipe, (<-responseCh).err
}
