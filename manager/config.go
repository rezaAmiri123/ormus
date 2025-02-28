package manager

import "github.com/rezaAmiri123/ormus/manager/service/authservice"

type Config struct {
	JWTConfig authservice.JwtConfig `koanf:"jwt_config"`
}
