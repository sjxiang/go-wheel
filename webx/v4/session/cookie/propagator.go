package cookie

import (
	"net/http"
)

type PropagatorOption func(p *Propagator)

type Propagator struct {
	cookieName   string
	cookieOption func(c *http.Cookie)
}

func NewPropagator(opts ...PropagatorOption) *Propagator {
	 res := &Propagator{
		cookieName: "sessid",
		cookieOption: func(c *http.Cookie) {

		},
	}

	for _, opt := range opts {
		opt(res)
	}

	return res
}

func WithCookieName(name string) PropagatorOption {
	return func(p *Propagator) {
		p.cookieName = name
	}
}

func (p *Propagator) Inject(id string, resp http.ResponseWriter) error {
	http.SetCookie(resp, &http.Cookie{
		Name: p.cookieName,
		Value: id,
	})

	return nil
}

func (p *Propagator) Extract(req *http.Request) (string, error) {
	c, err := req.Cookie(p.cookieName)
	if err != nil {
		return "", err
	}
	return c.Value, nil
}

func (p *Propagator) Remove(resp http.ResponseWriter) error {
	http.SetCookie(resp, &http.Cookie{
		Name: p.cookieName,
		MaxAge: -1,
	})
	return nil
}