package onetimesecret

import (
	"github.com/mholt/caddy"
	"github.com/mholt/caddy/caddyhttp/httpserver"
)

func init() {
	caddy.RegisterPlugin("onetimesecret", caddy.Plugin{
		ServerType: "http",
		Action:     setup,
	})
}

func setup(c *caddy.Controller) error {
	cfg := httpserver.GetConfig(c)
	root := cfg.Root

	dirs, err := oneTimeSecretParse(c)
	if err != nil {
		return err
	}

	onetimesecret := OneTimeSecret{Dirs: dirs}

	cfg.AddMiddleware(func(next httpserver.Handler) httpserver.Handler {
		onetimesecret.Next = next
		onetimesecret.SiteRoot = root
		return onetimesecret
	})

	return nil
}

func oneTimeSecretParse(c *caddy.Controller) ([]string, error) {
	var dirs []string

	for c.Next() {
		args := c.RemainingArgs()
		if len(args) == 0 {
			return dirs, c.Err("expected at least one directory name")
		}
		dirs = append(dirs, args...)
	}

	return dirs, nil
}
