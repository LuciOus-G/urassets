package controllers

import (
	"context"
	"database/sql"
	"github.com/lucious/urassets/UrAssetsCore/core/Iservice"
	"github.com/lucious/urassets/Utilities"
)

type UrAssetsHandler struct {
	Ctx      context.Context
	Service  Iservice.Interfaces
	DB       *sql.DB
	Response *Utilities.BaseResponse
}
