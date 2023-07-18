package server

import (
	"net/http"
	"userservice/pkg/server/gen"
	"userservice/pkg/storage"
	"userservice/pkg/token"

	"github.com/labstack/echo/v4"
	"go.elastic.co/apm/v2"
	"go.uber.org/zap"
)

type Handler struct {
	logger         *zap.Logger
	storage        storage.Storage
	tokenGenerator token.TokenGenerator
}

func (h Handler) Health(ctx echo.Context) error {
	return ctx.NoContent(http.StatusOK)
}

func (h Handler) GetUserDetail(ctx echo.Context, uuid string) error {
	span, apmCtx := apm.StartSpan(ctx.Request().Context(), "GetUserDetail", "request")
	defer span.End()

	user := h.storage.GetByUUID(apmCtx, uuid)
	if user == nil {
		h.logger.Debug("user with uuid cannot be found", zap.String("uuid", uuid))
		return ctx.NoContent(http.StatusNotFound)
	}

	userDetail := gen.UserDetail{
		Email:    user.Email,
		Fullname: user.Fullname,
		Username: user.Username,
		Uuid:     user.Uuid,
	}
	return ctx.JSON(http.StatusOK, userDetail)
}

func (h Handler) Authenticate(ctx echo.Context) error {
	span, apmCtx := apm.StartSpan(ctx.Request().Context(), "Authenticate", "request")
	defer span.End()

	creds := new(gen.UserCredentials)
	if ctx.Bind(creds) != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}

	user := h.storage.GetByUsername(apmCtx, creds.Username)
	if user == nil {
		h.logger.Debug("user with username cannot be found", zap.String("username", creds.Username))
		return ctx.NoContent(http.StatusNotFound)
	}

	if !user.ValidatePassword(creds.Password) {
		h.logger.Debug("authentication failed for user", zap.String("username", creds.Username))
		return ctx.NoContent(http.StatusUnauthorized)
	}

	tokenData := token.Data{
		UserId:   user.Uuid,
		Username: user.Username,
		Roles:    user.Roles,
	}

	tokenString, err := h.tokenGenerator.Generate(apmCtx, tokenData)
	if err != nil {
		h.logger.Error("error while generating token for user", zap.Error(err), zap.String("username", user.Username))
		return ctx.NoContent(http.StatusInternalServerError)
	}

	token := gen.Token{
		Token: tokenString,
	}

	h.logger.Debug("user authenticated with roles", zap.String("username", user.Username), zap.Strings("roles", user.Roles))
	return ctx.JSON(http.StatusOK, token)
}

func NewHandler(logger *zap.Logger, storage storage.Storage, tokenGenerator token.TokenGenerator) Handler {
	return Handler{
		logger:         logger,
		storage:        storage,
		tokenGenerator: tokenGenerator,
	}
}
