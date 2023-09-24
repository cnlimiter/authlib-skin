package handle

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/netip"

	"github.com/julienschmidt/httprouter"
	"github.com/samber/lo"
	"github.com/xmdhs/authlib-skin/model"
	"github.com/xmdhs/authlib-skin/service"
	"github.com/xmdhs/authlib-skin/utils"
)

func (h *Handel) Reg() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		ctx := r.Context()

		ip, err := utils.GetIP(r, h.config.RaelIP)
		if err != nil {
			h.logger.InfoContext(ctx, err.Error())
			handleError(ctx, w, err.Error(), model.ErrInput, 400)
			return
		}

		u, err := utils.DeCodeBody[model.User](r.Body, h.validate)
		if err != nil {
			h.logger.InfoContext(ctx, err.Error())
			handleError(ctx, w, err.Error(), model.ErrInput, 400)
			return
		}
		rip, err := getPrefix(ip)
		if err != nil {
			h.logger.WarnContext(ctx, err.Error())
			handleError(ctx, w, err.Error(), model.ErrUnknown, 500)
			return
		}
		err = h.webService.Reg(ctx, u, rip, ip)
		if err != nil {
			if errors.Is(err, service.ErrExistUser) {
				h.logger.DebugContext(ctx, err.Error())
				handleError(ctx, w, err.Error(), model.ErrExistUser, 400)
				return
			}
			if errors.Is(err, service.ErrRegLimit) {
				h.logger.DebugContext(ctx, err.Error())
				handleError(ctx, w, err.Error(), model.ErrRegLimit, 400)
				return
			}
			h.logger.WarnContext(ctx, err.Error())
			handleError(ctx, w, err.Error(), model.ErrService, 500)
			return
		}
		json.NewEncoder(w).Encode(model.API[any]{
			Code: 0,
			Data: nil,
			Msg:  "",
		})

	}
}

func getPrefix(ip string) (string, error) {
	ipa, err := netip.ParseAddr(ip)
	if err != nil {
		return "", fmt.Errorf("getPrefix: %w", err)
	}
	if ipa.Is6() {
		return lo.Must1(ipa.Prefix(48)).String(), nil
	}
	return lo.Must1(ipa.Prefix(24)).String(), nil
}
