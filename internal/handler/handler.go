package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/vanyovan/test-product.git/internal/entity"
	"github.com/vanyovan/test-product.git/internal/usecase"
)

type Handler struct {
	ProductUsecase usecase.ProductService
}

func NewHandler(ProductUsecase usecase.ProductService) *Handler {
	return &Handler{
		ProductUsecase: ProductUsecase,
	}
}

func (h *Handler) HandleCreateProduct(w http.ResponseWriter, r *http.Request) {

	request := CreateProductRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "fail",
			"data": map[string]interface{}{
				"error": err.Error(),
			},
		})
		return
	}

	param := entity.Product{
		ProductName:        request.ProductName,
		ProductDescription: request.ProductDescription,
		ProductPrice:       request.ProductPrice,
		ProductVariety:     request.ProductVariety,
		ProductRating:      request.ProductRating,
		ProductStock:       request.ProductStock,
	}

	result, err := h.ProductUsecase.CreateProduct(r.Context(), param)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "fail",
			"data": map[string]interface{}{
				"error": err.Error(),
			},
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
		"data":   result,
	})
}

func (h *Handler) HandleViewProduct(w http.ResponseWriter, r *http.Request) {
	result, err := h.ProductUsecase.ViewProduct(r.Context())
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "fail",
			"data": map[string]interface{}{
				"error": err.Error(),
			},
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
		"data":   result,
	})
}

func (h *Handler) HandleDeleteProduct(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	segments := strings.Split(path, "/")

	param := segments[4]

	if param == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "fail",
			"data": map[string]interface{}{
				"error": "param not found",
			},
		})
		return
	}

	productId, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "fail",
			"data": map[string]interface{}{
				"error": err.Error(),
			},
		})
		return
	}

	err = h.ProductUsecase.DeleteProduct(r.Context(), productId)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "fail",
			"data": map[string]interface{}{
				"error": err.Error(),
			},
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
	})
}

func (h *Handler) HandleUpdateProduct(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	segments := strings.Split(path, "/")

	param := segments[4]

	if param == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "fail",
			"data": map[string]interface{}{
				"error": "param not found",
			},
		})
		return
	}

	productId, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "fail",
			"data": map[string]interface{}{
				"error": err.Error(),
			},
		})
		return
	}

	request := UpdateProductRequest{}
	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "fail",
			"data": map[string]interface{}{
				"error": err.Error(),
			},
		})
		return
	}

	product := entity.Product{
		ProductName:        request.ProductName,
		ProductDescription: request.ProductDescription,
		ProductPrice:       request.ProductPrice,
		ProductVariety:     request.ProductVariety,
		ProductRating:      request.ProductRating,
		ProductStock:       request.ProductStock,
	}

	err = h.ProductUsecase.UpdateProduct(r.Context(), productId, product)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "fail",
			"data": map[string]interface{}{
				"error": err.Error(),
			},
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
	})
}
