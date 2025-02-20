package casbin

import (
	"fmt"
	"github.com/casbin/casbin/v2"
)

func InitCasbin() (*casbin.Enforcer, error) {
	e, err := casbin.NewEnforcer("model.conf", "policy.csv")
	if err != nil {
		return nil, fmt.Errorf("unable to create Casbin enforcer: %w", err)
	}

	return e, nil
}
