//Package app el paquete application contiene todos los métodos y recursos comunes para usar por una api
package app

import (
	"github.com/go-chi/jwtauth"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"time"
)

const (
	LogLevelDB   = 1
	LogLevelAuth = 2
	LogLevelHttp = 4
)

var (
	//TokenAuth Estructura para crear y validar token JWT
	TokenAuth *jwtauth.JWTAuth
	//Exp Tiempo de expiracion del JWT
	Exp time.Duration
	//Secret utilizado para la creacion de JWT
	Secret string
	//Logger inicializado
	Logger zerolog.Logger
)

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(-1)
	Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	Exp = 600 * time.Minute
	Secret = os.Getenv("API_JWT_SECRET")
	TokenAuth = jwtauth.New("HS256", []byte(Secret), nil)

}
