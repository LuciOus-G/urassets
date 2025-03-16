package service

import (
	"context"
	"database/sql"
	"github.com/lucious/urassets/UrAssetsCore/core/Iservice"
)

type Services struct {
	Ctx context.Context
	DB  *sql.DB
}

func Interfaces(init Services) Iservice.Interfaces {
	return Iservice.Interfaces(init)
}
