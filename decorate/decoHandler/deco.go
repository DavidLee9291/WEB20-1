package decoHansdler

import "net/http"

type DecoratorFunc func(http.ResponseWriter, *http.Request, http.Handler)

type DecoHansdler struct {
	fn DecoratorFunc
	h  http.Handler
}

func (self *DecoHansdler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	self.fn(w, r, self.h)
}

func NewDecoHandler(h http.Handler, fn DecoHansdler) http.Handler {
	return &DecoHansdler{
		fn: fn,
		h:  h,
	}
}
