package food2fork

import (
	f2f "."
	"net/url"
	"os"
	"testing"
)

// APIKey is api key from envitonment valuables.
var APIKey = os.Getenv("FOOD_2_FORK_API_KEY")

var api *f2f.Food2Fork

func init() {
	api = f2f.New(APIKey)
}

func TestSearchRecipe(t *testing.T) {
	v := url.Values{}
	r1, err := api.SearchRecipe(v)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", r1)
}

func TestGetRecipe(t *testing.T) {
	r, err := api.GetRecipe("35382")
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", r)

}
