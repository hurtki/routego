# routego

> A lightweight, zero-dependency HTTP router for Go — clean, fast, and straightforward.

```
go get github.com/hurtki/routego
```

---

## Features

- **Typed URL parameters** — capture `{num}` or `{string}` segments directly from the path
- **All standard HTTP methods** — GET, POST, PUT, PATCH, DELETE, HEAD, OPTIONS
- **`http.Handler` & `http.HandlerFunc`** — works with anything from the standard library
- **Custom 404 handler** — bring your own not-found response
- **Zero dependencies** — pure Go, nothing else

---

## Quick Start

```go
package main

import (
    "fmt"
    "net/http"

    "github.com/hurtki/routego"
)

func main() {
    router := routego.NewRouter(nil)

    router.GetFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Hello, World!")
    })

    router.GetFunc("/users/{num}", func(w http.ResponseWriter, r *http.Request) {
        id := r.Context().Value("urlParameter").(int)
        fmt.Fprintf(w, "User ID: %d\n", id)
    })

    router.PostFunc("/posts/{string}", func(w http.ResponseWriter, r *http.Request) {
        slug := r.Context().Value("urlParameter").(string)
        fmt.Fprintf(w, "Post: %s\n", slug)
    })

    http.ListenAndServe(":8080", &router)
}
```

---

## URL Parameters

Route parameters are declared inline in the path pattern using curly braces.

| Pattern           | Matches              | Parameter value                 |
| ----------------- | -------------------- | ------------------------------- |
| `/users/{num}`    | `/users/42`          | `42` (type `int`)               |
| `/posts/{string}` | `/posts/hello-world` | `"hello-world"` (type `string`) |
| `/api/health`     | `/api/health`        | — (strict, no parameter)        |

> **Note:** a route may contain at most **one** parameter segment.

Parameters are injected into the request context under the key `"urlParameter"`:

```go
router.GetFunc("/items/{num}", func(w http.ResponseWriter, r *http.Request) {
    id := r.Context().Value("urlParameter").(int)
    // ...
})
```

---

## Router Configuration

```go
router := routego.NewRouter(&routego.RouterConfig{
    NotFoundHandler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        http.Error(w, "nothing here", http.StatusNotFound)
    }),
})
```

If `RouterConfig` is `nil` or `NotFoundHandler` is unset, the standard `http.NotFoundHandler()` is used.

---

## All Available Methods

```go
// http.Handler variants
router.Get(path, handler)
router.Post(path, handler)
router.Put(path, handler)
router.Patch(path, handler)
router.Delete(path, handler)
router.Head(path, handler)
router.Options(path, handler)

// http.HandlerFunc variants
router.GetFunc(path, handlerFunc)
router.PostFunc(path, handlerFunc)
router.PutFunc(path, handlerFunc)
router.PatchFunc(path, handlerFunc)
router.DeleteFunc(path, handlerFunc)
router.HeadFunc(path, handlerFunc)
router.OptionsFunc(path, handlerFunc)
```

---

## How It Works

`routego` splits each registered pattern and incoming URL into path segments and compares them one by one. Strict segments must match exactly; parameter segments accept any value of the declared type. The first matching route wins.

```
Pattern:  /api/users/{num}
Request:  /api/users/99

  "api"   == "api"   ✓  strict
  "users" == "users" ✓  strict
  {num}   <- 99      ✓  parameter → int(99)
```

---

## License

MIT
