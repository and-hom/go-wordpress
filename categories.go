package wordpress

import (
	"fmt"
	"net/http"
)

type Category struct {
	Description  string     `json:"description,omitempty"`
	Hierarchical bool       `json:"hierarchical,omitempty"`
	Name         string     `json:"name,omitempty"`
	Slug         string     `json:"slug,omitempty"`
}

type CategoriesCollection struct {
	client *Client
	url    string
}

func (col *CategoriesCollection) List(params interface{}) ([]Category, *http.Response, []byte, error) {
	var categories []Category
	resp, body, err := col.client.List(col.url, params, &categories)
	return categories, resp, body, err
}

func (col *CategoriesCollection) Get(slug string, params interface{}) (*Category, *http.Response, []byte, error) {
	var entity Category
	entityURL := fmt.Sprintf("%v/%v", col.url, slug)
	resp, body, err := col.client.Get(entityURL, params, &entity)
	return &entity, resp, body, err
}
