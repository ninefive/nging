package defaults

import (
	"github.com/webx-top/echo"
	"github.com/webx-top/echo/engine"
	"github.com/webx-top/echo/logger"
)

var Default = echo.New()

func ParseHeaderAccept(on bool) *echo.Echo {
	return Default.ParseHeaderAccept(on)
}

func SetAcceptFormats(acceptFormats map[string]string) *echo.Echo {
	return Default.SetAcceptFormats(acceptFormats)
}

func AddAcceptFormat(mime, format string) *echo.Echo {
	return Default.AddAcceptFormat(mime, format)
}

func SetFormatRenderers(formatRenderers map[string]func(c echo.Context, data interface{}) error) *echo.Echo {
	return Default.SetFormatRenderers(formatRenderers)
}

func AddFormatRenderer(format string, renderer func(c echo.Context, data interface{}) error) *echo.Echo {
	return Default.AddFormatRenderer(format, renderer)
}

func RemoveFormatRenderer(formats ...string) *echo.Echo {
	return Default.RemoveFormatRenderer(formats...)
}

// Router returns router.
func Router() *echo.Router {
	return Default.Router()
}

// SetLogger sets the logger instance.
func SetLogger(l logger.Logger) {
	Default.SetLogger(l)
}

// Logger returns the logger instance.
func Logger() logger.Logger {
	return Default.Logger()
}

// DefaultHTTPErrorHandler invokes the default HTTP error handler.
func DefaultHTTPErrorHandler(err error, c echo.Context) {
	Default.DefaultHTTPErrorHandler(err, c)
}

// SetHTTPErrorHandler registers a custom Echo.HTTPErrorHandler.
func SetHTTPErrorHandler(h echo.HTTPErrorHandler) {
	Default.SetHTTPErrorHandler(h)
}

// HTTPErrorHandler returns the HTTPErrorHandler
func HTTPErrorHandler() echo.HTTPErrorHandler {
	return Default.HTTPErrorHandler()
}

// SetBinder registers a custom binder. It's invoked by Context.Bind().
func SetBinder(b echo.Binder) {
	Default.SetBinder(b)
}

// Binder returns the binder instance.
func Binder() echo.Binder {
	return Default.Binder()
}

// SetRenderer registers an HTML template renderer. It's invoked by Context.Render().
func SetRenderer(r echo.Renderer) {
	Default.SetRenderer(r)
}

// Renderer returns the renderer instance.
func Renderer() echo.Renderer {
	return Default.Renderer()
}

// SetDebug enable/disable debug mode.
func SetDebug(on bool) {
	Default.SetDebug(on)
}

// Debug returns debug mode (enabled or disabled).
func Debug() bool {
	return Default.Debug()
}

// Use adds handler to the middleware chain.
func Use(middleware ...interface{}) {
	Default.Use(middleware...)
}

// Pre is alias
func Pre(middleware ...interface{}) {
	Default.Pre(middleware...)
}

// PreUse adds handler to the middleware chain.
func PreUse(middleware ...interface{}) {
	Default.PreUse(middleware...)
}

// Clear middleware
func Clear(middleware ...interface{}) {
	Default.Clear(middleware...)
}

// Connect adds a CONNECT route > handler to the router.
func Connect(path string, h interface{}, m ...interface{}) {
	Default.Connect(path, h, m...)
}

// Delete adds a DELETE route > handler to the router.
func Delete(path string, h interface{}, m ...interface{}) {
	Default.Delete(path, h, m...)
}

// Get adds a GET route > handler to the router.
func Get(path string, h interface{}, m ...interface{}) {
	Default.Get(path, h, m...)
}

// Head adds a HEAD route > handler to the router.
func Head(path string, h interface{}, m ...interface{}) {
	Default.Head(path, h, m...)
}

// Options adds an OPTIONS route > handler to the router.
func Options(path string, h interface{}, m ...interface{}) {
	Default.Options(path, h, m...)
}

// Patch adds a PATCH route > handler to the router.
func Patch(path string, h interface{}, m ...interface{}) {
	Default.Patch(path, h, m...)
}

// Post adds a POST route > handler to the router.
func Post(path string, h interface{}, m ...interface{}) {
	Default.Post(path, h, m...)
}

// Put adds a PUT route > handler to the router.
func Put(path string, h interface{}, m ...interface{}) {
	Default.Put(path, h, m...)
}

// Trace adds a TRACE route > handler to the router.
func Trace(path string, h interface{}, m ...interface{}) {
	Default.Trace(path, h, m...)
}

// Any adds a route > handler to the router for all HTTP methods.
func Any(path string, h interface{}, m ...interface{}) {
	Default.Any(path, h, m...)
}

func Route(methods string, path string, h interface{}, m ...interface{}) {
	Default.Route(methods, path, h, m...)
}

// Match adds a route > handler to the router for multiple HTTP methods provided.
func Match(methods []string, path string, h interface{}, m ...interface{}) {
	Default.Match(methods, path, h, m...)
}

func SetHandlerWrapper(funcs ...func(interface{}) echo.Handler) {
	Default.SetHandlerWrapper(funcs...)
}

func SetMiddlewareWrapper(funcs ...func(interface{}) echo.Middleware) {
	Default.SetMiddlewareWrapper(funcs...)
}

func AddHandlerWrapper(funcs ...func(interface{}) echo.Handler) {
	Default.AddHandlerWrapper(funcs...)
}

func AddMiddlewareWrapper(funcs ...func(interface{}) echo.Middleware) {
	Default.AddMiddlewareWrapper(funcs...)
}

func Prefix() string {
	return Default.Prefix()
}

func SetPrefix(prefix string) *echo.Echo {
	return Default.SetPrefix(prefix)
}

// MetaHandler Add meta information about endpoint
func MetaHandler(m echo.H, handler interface{}) interface{} {
	return Default.MetaHandler(m, handler)
}

// RebuildRouter rebuild router
func RebuildRouter(args ...[]*echo.Route) {
	Default.RebuildRouter(args...)
}

// Group creates a new sub-router with prefix.
func Group(prefix string, m ...interface{}) *echo.Group {
	return Default.Group(prefix, m...)
}

// URI generates a URI from handler.
func URI(handler interface{}, params ...interface{}) string {
	return Default.URI(handler, params...)
}

// URL is an alias for `URI` function.
func URL(h interface{}, params ...interface{}) string {
	return Default.URL(h, params...)
}

// Routes returns the registered routes.
func Routes() []*echo.Route {
	return Default.Routes()
}

// NamedRoutes returns the registered handler name.
func NamedRoutes() map[string][]int {
	return Default.NamedRoutes()
}

func ServeHTTP(req engine.Request, res engine.Response) {
	Default.ServeHTTP(req, res)
}

// Run starts the HTTP engine.
func Run(eng engine.Engine, handler ...engine.Handler) error {
	return Default.Run(eng, handler...)
}

func Engine() engine.Engine {
	return Default.Engine()
}

// Stop stops the HTTP server.
func Stop() error {
	return Default.Stop()
}
