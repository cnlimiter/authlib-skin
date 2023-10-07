package handle

import (
	"encoding/json"
	"net/http"

	"github.com/xmdhs/authlib-skin/model"
)

func (h *Handel) GetConfig() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		c := h.webService.GetConfig(ctx)
		m := model.API[model.Config]{
			Code: 0,
			Data: c,
		}
		json.NewEncoder(w).Encode(m)
	}
}
