> **⚠️ Note:** This project is still in progress. It's not feature-complete yet, but you're welcome to explore, contribute, or use it as a base for your own project.

# Tiny Express (TE)

Tiny Express (TE) is a lightweight and flexible web framework for Go developers, heavily inspired by [Express.js](https://expressjs.com/). It enables you to build modular and scalable web applications with a clean and intuitive API, while embracing Go’s performance and concurrency.

---

## ✨ Features

- 🚀 ExpressJS-like routing (`Get`, `Post`, `Put`, `Delete`, etc.)
- 🧱 Middleware support
- ⚙️ Custom route parameters with type parsing (`:id{int}`)
- 📦 Chainable responses (e.g., `SendJSON`, `SendText`, `SendFile`)
- 🔁 Route resolver chaining with `Next()`
- 🔒 Request and response lifecycle hooks (e.g., `OnBeforeSend`)
- 🔍 Flexible logger middleware support
- 🛠 Utility-rich request and response wrappers
- 🌐 Zero-dependency, minimal server setup

---

## 📦 Installation

To get started, clone this repository and include it in your Go project:

```bash
git clone https://github.com/your-username/tiny-express.git
cd tiny-express
```

---

## 🧪 Example Usage

```go
package main

import te "tinyExpress"

func main() {
	app := te.App()

	app.Get("/hello", func(req *te.Request, res *te.Response) {
		res.SendText("Hello, World!")
	})

	app.Listen()
}
```

---

## 🧩 Advanced Routing

Support for dynamic parameters with types:

```go
app.Get("/user/:id{int}", func(req *te.Request, res *te.Response) {
	id := req.GetParam("id")
	res.SendText("User ID: " + id)
})
```

---

## 🧱 Middleware

You can use middleware to extend the behavior of requests:

```go
app.Use(te.Handler(func(req *te.Request, res *te.Response) {
	fmt.Println("Middleware activated!")
	if req.Next != nil {
		(*req.Next)()
	}
}))
```

---

## 🔧 Configuration

You can pass optional configurations like `Port`, `Host`, and `AllowedMethod`:

```go
config := te.AppConfig{
	Port: 8080,
	Host: "0.0.0.0",
}
app := te.App(config)
```

---

## 📤 Response Methods

- `SendText(string)`
- `SendHTML(string)`
- `SendJSON(interface{})`
- `SendBytes([]byte)`
- `SendFile(path string)`
- `Redirect(url string, statusCode int)`

---

## 🧾 Logger Middleware Example

Add structured logging with configurable options:

```go
app.Use(middleware.TeLog(middleware.LogOption{
	TimeStamp: true,
	Method:    true,
	Path:      true,
	FullUrl:   true,
}))
```

---

## 🧠 Core Components

### `Router`
The `Router` in Tiny Express is responsible for handling routes and middleware. Like ExpressJS, you can mount route handlers with methods like `Get`, `Post`, `Put`, `Delete`, `Patch`, and more.

```go
handler:= Router()

handler.Get("", te.Handler(func(req *te.Request, res *te.Response) {
		fmt.Println("Middleware activated!")
		if req.Next != nil {
			(*req.Next)()
		}
	}) 
)

app.Get("/", handler)
app.Post("/submit", handler)
```

You can also chain middleware and mount path-based routers:

```go
app.Use(middleware1, middleware2)
app.UsePath("/api", apiRouter)
```

### `Request`
The custom `Request` object extends Go's `http.Request` and adds:
- `Params`: URL parameters
- `Query`: Query string values
- `Body`: Raw request body as bytes
- `GetParam`, `GetQuery`, `GetHeader`, `GetCookie` helpers

### `Response`
The `Response` object wraps `http.ResponseWriter` and provides methods like:
- `SendText`, `SendHTML`, `SendJSON`, `SendFile`
- `SetHeader`, `SetStatusCode`, `SetCookie`
- `Redirect`, `IsResolved`

### `Route`
Routes in Tiny Express are powered by regex and allow expressive matching like:

```go
/user/:id{int}  →  matches /user/123 and parses id
```

You can define routes using:

```go
Route("/path", te.RequestMethodGet, handler)
```

### `App`
The `Application` is the entry point. You can create an app with optional config:

```go
app := te.App(te.AppConfig{Port: 3000, Host: "127.0.0.1"})
app.Listen()
```

It also supports a `BeforeSend` hook to execute logic before the response is finalized.

---

## 📂 Project Structure

- `app.go`: Entry point and app/server config
- `request.go`: Enhanced request wrapper
- `response.go`: Response helper with send methods
- `route.go`, `router.go`: Routing engine and method handling
- `types.go`: Core types and interfaces
- `utils.go`: Utility functions (e.g., route matching)
- `logger.go`: Custom middleware example

---

## 🛣 Roadmap

- [ ] File upload handling
- [ ] Error middleware support
- [ ] Static file serving
- [ ] Unit tests
- [ ] Performance benchmarks

---

## 👨‍💻 Contributing

Want to contribute? Awesome! Feel free to submit pull requests or open issues to help improve this project.

---

## 📄 License

MIT License. Feel free to use this in personal and commercial projects.