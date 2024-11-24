package response

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"time"
)

type setCookie struct {
	cookieName      string
	token           string
	cookieExpiresIn int
	path            string
	domain          string
	secure          bool
	httpOnly        bool
}

func (s setCookie) setCookie(c *gin.Context) {
	c.SetCookie(
		s.cookieName,
		s.token,
		s.cookieExpiresIn,
		"/",
		"localhost",
		false,
		true,
	)
}

// SetCookieHandler sets a cookie in the response header
// auto sets the path to "/" domain and secure depending on the environment
func SetCookieHandler(cookieName, token, path string, cookieExpiresIn time.Time, c *gin.Context) {
	log.Printf("Setting cookie with details - Name: %s, Path: %s, Expires In: %s",
		cookieName, path, cookieExpiresIn.String())
	cookie := setCookie{
		cookieName:      cookieName,
		token:           token,
		cookieExpiresIn: int(cookieExpiresIn.Sub(time.Now().UTC()).Seconds()),
		path:            path,
		domain:          os.Getenv("DOMAIN"),
		secure:          os.Getenv("APP_ENV") == "production",
		httpOnly:        os.Getenv("APP_ENV") != "production",
	}

	cookie.setCookie(c)
}
