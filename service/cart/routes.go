package cart

import (
	"fmt"
	"net/http"

	"github.com/Eslam-Amin/ecommerce/service/auth"
	"github.com/Eslam-Amin/ecommerce/types"
	"github.com/Eslam-Amin/ecommerce/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	store        types.OrderStore
	productStore types.ProductStore
	userStore    types.UserStore
}

func NewHandler(store types.OrderStore, productStore types.ProductStore) *Handler {
	return &Handler{store: store, productStore: productStore}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/cart/checkout", auth.WithJWTAut(h.handleCheckout, h.userStore)).Methods(http.MethodPost)

}
func (h *Handler) handleCheckout(w http.ResponseWriter, r *http.Request) {
	var cart types.CartCheckoutPayload
	userId := auth.GetUserIdFromContext(r.Context())
	if err := utils.ParseJSON(r, &cart); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(cart); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v:", errors))
		return
	}

	productIds, err := getCartItemsIds(cart.Items)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	products, err := h.productStore.GetProductsByIds(productIds)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	orderId, totalPrice, err := h.createOrder(products, cart.Items, userId)

	utils.WriteJSON(w, http.StatusCreated, types.ResponseBody{
		Success: true,
		Message: "Order Create successfully",
		Data: map[string]any{
			"orderId":    orderId,
			"totalPrice": totalPrice,
		},
	})
}
