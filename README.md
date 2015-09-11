# MWare

Collection of Go middleware patterns and handlers using the pattern:

```
type Middleware func(h http.Handler) http.Handler
```
