package scyllauser

import (
	"github.com/rezaAmiri123/ormus/logger"
	"github.com/rezaAmiri123/ormus/manager/entity"
	"github.com/rezaAmiri123/ormus/manager/repository/scyllarepo"
	"github.com/scylladb/gocqlx/v2/qb"
)

func init() {
	statements["GetUserByEmail"] = scyllarepo.Statement{
		Query:  "SELECT * FROM users WHERE email = ? AND is_active = true",
		Values: []string{"email"},
	}
}

func (r Repository) GetUserByEmail(email string) (entity.User, error) {
	// TODO : check if this is the right way to do it
	query, err := r.db.GetStatement(statements["GetUserByEmail"])
	if err != nil {
		logger.L().Debug("GetUserByEmail GetStatement", err)

		return entity.User{}, err
	}
	query.BindMap(qb.M{
		"email": email,
	})

	var user entity.User
	if err = query.Get(&user); err != nil {
		logger.L().Debug("GetUserByEmail Get", err)

		return entity.User{}, err
	}

	return user, nil
}
