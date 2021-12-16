package controllers

import (
	"net/http"

	"github.com/leslesnoa/bookstore_items-api/domain/items"
	"github.com/leslesnoa/bookstore_items-api/services"
	"github.com/leslesnoa/bookstore_items-api/utils/http_utils"
	"github.com/leslesnoa/bookstore_oauth-go/oauth"
)

var (
	ItemsController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type itemsController struct{}

func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		http_utils.RespondJson(w, err.Status, err)
		// w.Header().Set("Content-Type", "application/json")
		// w.WriteHeader(err.Status)
		// json.NewEncoder(w).Encode(err)
		return
	}
	item := items.Item{
		Seller: oauth.GetCallerId(r),
	}
	result, err := services.ItemsService.Create(item)
	if err != nil {
		http_utils.RespondJson(w, err.Status, err)
		// w.Header().Set("Content-Type", "application/json")
		// w.WriteHeader(err.Status)
		// json.NewEncoder(w).Encode(err)
		return
	}
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusCreated)
	// json.NewEncoder(w).Encode(result)
	http_utils.RespondJson(w, http.StatusCreated, result)
	return
}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {

}
