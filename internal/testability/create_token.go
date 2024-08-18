package testability

import (
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/elct9620/clean-architecture-in-go-2025/internal/entity/tokens"
)

func (api *Api) CreateToken(w http.ResponseWriter, r *http.Request) {
	var req CreateTokenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data, err := base64.StdEncoding.DecodeString(req.Data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token := tokens.New(req.Id)
	token.SetData(data)

	err = api.TokenRepository.Save(r.Context(), token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := &CreatedResponse{
		Id: token.Id(),
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
