#Food2Fork API Client 

[![GoDoc](https://godoc.org/github.com/Rompei/food2fork?status.png)](https://godoc.org/github.com/Rompei/food2fork)

Food2Fork API client implemented with Golang.  
API document: [http://food2fork.com/about/api](http://food2fork.com/about/api)

##Usage

```
$ go get github.com/Rompei/food2fork
```

##Example

```go
package main

import (
    "fmt"
    "github.com/Rompei/food2fork"
    "net/url"
    "os"
)

func main() {

    // Initialize api structure.
    api := food2fork.New(os.Getenv("FOOD_2_FORK_API_KEY"))

    // Prepare query strings.
    v := url.Values{}
    v.Add("q", "chicken")

    // SearchRecipe returns recipes searched with keywords.
    recipes, err := api.SearchRecipe(v)
    if err != nil {
        panic(err)
    }

    if len(recipes) > 0 {

    // GetRecipe returns a recipe specialized by recipe id.
    recipe, err := api.GetRecipe(recipes[0].RecipeID)
        if err != nil {
           panic(err)
        }
        fmt.Printf("Title: %v\n", recipe.Title)
        for _, ing := range recipe.Ingredients {
            fmt.Printf("Ingredients: %v\n", ing)
        }
    }
}


/*
Title: Bacon Wrapped Jalapeno Popper Stuffed Chicken
Ingredients: 4 small chicken breasts, pounded thin
Ingredients: salt and pepper to taste
Ingredients: 4 jalapenos, diced
Ingredients: 4 ounces cream cheese, room temperature
Ingredients: 1 cup cheddar cheese, shredded
Ingredients: 8 slices bacon
*/


```

##Referenced software

The software design of this OSS references to  [ChimeraCoder/anaconda](https://github.com/ChimeraCoder/anaconda)

##License

[MIT License](https://opensource.org/licenses/MIT)
