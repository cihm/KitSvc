package eventutil

import "github.com/gin-gonic/gin"

type Engine struct {
	Gin       *gin.Engine
	Listeners []Listener
}

type Listener struct {
	Method  string
	Path    string
	Stream  string
	Handler func(*gin.Context)
}

func New(e *gin.Engine) *Engine {
	return &Engine{Gin: e}
}

func (e *Engine) POST(relativePath string, stream string, handler func(*gin.Context)) {
	e.Gin.POST(relativePath, handler)
	e.Listeners = append(e.Listeners, Listener{"POST", relativePath, stream, handler})
}

func (e *Engine) GET(relativePath string, stream string, handler func(*gin.Context)) {
	e.Gin.GET(relativePath, handler)
	e.Listeners = append(e.Listeners, Listener{"GET", relativePath, stream, handler})
}

func (e *Engine) DELETE(relativePath string, stream string, handler func(*gin.Context)) {
	e.Gin.DELETE(relativePath, handler)
	e.Listeners = append(e.Listeners, Listener{"DELETE", relativePath, stream, handler})
}

func (e *Engine) PUT(relativePath string, stream string, handler func(*gin.Context)) {
	e.Gin.PUT(relativePath, handler)
	e.Listeners = append(e.Listeners, Listener{"PUT", relativePath, stream, handler})
}

func (e *Engine) PATCH(relativePath string, stream string, handler func(*gin.Context)) {
	e.Gin.PATCH(relativePath, handler)
	e.Listeners = append(e.Listeners, Listener{"PATCH", relativePath, stream, handler})
}