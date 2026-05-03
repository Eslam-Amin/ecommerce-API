package product

import (
	"fmt"
	"net/http"

	"github.com/Eslam-Amin/ecommerce/types"
	"github.com/Eslam-Amin/ecommerce/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.ProductStore
}

func NewHandler(store types.ProductStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/products", h.handleCreateProduct).Methods(http.MethodPost)
	router.HandleFunc("/products", h.handleGetProducts).Methods(http.MethodGet)
}

func (h *Handler) handleCreateProduct(w http.ResponseWriter, r *http.Request) {
	var payload types.CreateProductPayload
	err := utils.ParseJSON(r, &payload)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	errors := utils.Validate.Struct(payload)
	if errors != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}
	product := types.Product{
		Name:        payload.Name,
		Description: payload.Description,
		Image:       payload.Image,
		Price:       payload.Price,
		Quantity:    payload.Quantity,
	}
	err = h.store.CreateProduct(product)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, types.ResponseBody{
		Success: true,
		Message: "Product Created Successfully",
		Data:    product,
	})
}

func (h *Handler) handleGetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.store.GetProducts()

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, types.ResponseBody{
		Success: true,
		Data:    products,
	})
}
