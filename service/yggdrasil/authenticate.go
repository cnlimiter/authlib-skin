package yggdrasil

import (
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/xmdhs/authlib-skin/db/cache"
	"github.com/xmdhs/authlib-skin/db/ent"
	"github.com/xmdhs/authlib-skin/db/ent/user"
	"github.com/xmdhs/authlib-skin/model"
	"github.com/xmdhs/authlib-skin/model/yggdrasil"
	"github.com/xmdhs/authlib-skin/utils"
)

var (
	ErrRate     = errors.New("频率限制")
	ErrPassWord = errors.New("错误的密码或邮箱")
)

func (y *Yggdrasil) Authenticate(cxt context.Context, auth yggdrasil.Authenticate) (yggdrasil.Token, error) {
	err := rate("Authenticate"+auth.Username, y.cache, 10*time.Second, 3)
	if err != nil {
		return yggdrasil.Token{}, fmt.Errorf("Authenticate: %w", err)
	}

	u, err := y.client.User.Query().Where(user.EmailEQ(auth.Username)).WithProfile().First(cxt)
	if err != nil {
		var nf *ent.NotFoundError
		if errors.As(err, &nf) {
			return yggdrasil.Token{}, fmt.Errorf("Authenticate: %w", ErrPassWord)
		}
		return yggdrasil.Token{}, fmt.Errorf("Authenticate: %w", err)
	}
	if !utils.Argon2Compare(auth.Password, u.Password, u.Salt) {
		return yggdrasil.Token{}, fmt.Errorf("Authenticate: %w", ErrPassWord)
	}
	clientToken := auth.ClientToken
	if clientToken == "" {
		clientToken = strings.ReplaceAll(uuid.New().String(), "-", "")
	}

	var utoken *ent.UserToken
	err = utils.WithTx(cxt, y.client, func(tx *ent.Tx) error {
		utoken, err = tx.User.QueryToken(u).ForUpdate().First(cxt)
		if err != nil {
			var nf *ent.NotFoundError
			if !errors.As(err, &nf) {
				return err
			}
		}
		if utoken == nil {
			ut, err := tx.UserToken.Create().SetTokenID(1).SetUUID(u.Edges.Profile.UUID).Save(cxt)
			if err != nil {
				return err
			}
			err = tx.User.UpdateOne(u).SetToken(ut).Exec(cxt)
			if err != nil {
				return err
			}
			utoken = ut
		}
		return nil
	})
	if err != nil {
		return yggdrasil.Token{}, fmt.Errorf("Authenticate: %w", err)
	}

	claims := model.TokenClaims{
		Tid: strconv.FormatUint(utoken.TokenID, 10),
		CID: clientToken,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * 24 * time.Hour)),
			Issuer:    "authlib-skin",
			Subject:   u.Edges.Profile.UUID,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwts, err := token.SignedString([]byte(y.config.JwtKey))
	if err != nil {
		return yggdrasil.Token{}, fmt.Errorf("Authenticate: %w", err)
	}
	p := yggdrasil.TokenProfile{
		ID:   u.Edges.Profile.UUID,
		Name: u.Edges.Profile.Name,
	}

	return yggdrasil.Token{
		AccessToken:       jwts,
		AvailableProfiles: []yggdrasil.TokenProfile{p},
		ClientToken:       clientToken,
		SelectedProfile:   p,
		User: yggdrasil.TokenUser{
			ID:         u.Edges.Profile.UUID,
			Properties: []any{},
		},
	}, nil
}

func rate(k string, c cache.Cache, d time.Duration, count uint) error {
	key := []byte(k)
	v, err := c.Get([]byte(key))
	if err != nil {
		return fmt.Errorf("rate: %w", err)
	}
	if v == nil {
		err := putUint(1, c, key, d)
		if err != nil {
			return fmt.Errorf("rate: %w", err)
		}
		return nil
	}
	n := binary.BigEndian.Uint64(v)
	if n > uint64(count) {
		return fmt.Errorf("rate: %w", ErrRate)
	}
	err = putUint(n+1, c, key, d)
	if err != nil {
		return fmt.Errorf("rate: %w", err)
	}
	return nil
}

func putUint(n uint64, c cache.Cache, key []byte, d time.Duration) error {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, n)
	err := c.Put(key, b, time.Now().Add(d))
	if err != nil {
		return fmt.Errorf("rate: %w", err)
	}
	return nil
}
