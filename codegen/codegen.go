package codegen

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"math"
	"os"
	"path"
)

type CodeGen struct {
}

type RecipesPage struct {
	NumCols int        `json:"num-cols"`
	Output  string     `json:"output"`
	Recipes []*Recipes `json:"recipes"`
}

type Recipes struct {
	Name       string `json:"name"`
	RecipeLink string `json:"recipe-link"`
	InstaLink  string `json:"insta-link"`
	ImgSrc     string `json:"img-src"`
	Desc       string `json:"desc"`
}

func New() *CodeGen {
	return &CodeGen{}
}

// recipes points to links to
func (c *CodeGen) Recipes() error {
	// open and readd recipes config - contains the list of recipes each
	// containing name, links, img, desc, etc..
	fileReader, err := os.Open(RecipesConfigFile)
	if err != nil {
		return err
	}
	recipesPage := &RecipesPage{}
	byteValue, _ := ioutil.ReadAll(fileReader)
	if err := json.Unmarshal(byteValue, recipesPage); err != nil {
		return err
	}

	// create output file
	outputFile, err := os.Create(recipesPage.Output)
	if err != nil {
		return err
	}

	tmpl, err := template.New(path.Base(RecipesTemplateFile)).Funcs(
		template.FuncMap{
			"getTotalRows": func(numRecipes, numCols int) int {
				if (numRecipes % numCols) == 0 {
					return numRecipes / numCols
				}
				return (numRecipes / numCols) + 1
			},
			"N": func(start, end int) (stream chan int) {
				stream = make(chan int)
				go func() {
					for i := start; i < end; i++ {
						stream <- i
					}
					close(stream)
				}()
				return
			},
			"min": func(x, y int) int {
				return int(math.Min(float64(x), float64(y)))
			},
			"sub": func(x, y int) int {
				return x - y
			},
			"add": func(x, y int) int {
				return x + y
			},
			"div": func(x, y int) int {
				return x / y
			}}).ParseFiles(RecipesTemplateFile)
	if err != nil {
		return err
	}

	if err := tmpl.Execute(outputFile, recipesPage); err != nil {
		return err
	}

	return nil
}

// articles are the content that is home grown from this website
func (c *CodeGen) Artices() {
	// TODO
}
