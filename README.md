# go-htmx

A lightweight Go library for handling HTMX requests and responses with support for both `net/http` and `fasthttp` frameworks.

## Installation

```bash
go get github.com/dajooo/go-htmx
```

## Features

- **Universal compatibility**: Works with both `net/http` and `fasthttp`
- **Type-safe**: Structured approach to HTMX request/response handling
- **Fluent API**: Chainable methods for building responses
- **Complete HTMX support**: All HTMX request headers and response headers are supported

## Quick Start

### Basic Usage with net/http

```go
package main

import (
    "net/http"
    "github.com/dajooo/go-htmx"
)

func handler(w http.ResponseWriter, r *http.Request) {
    h := htmx.New(r.Header)
    
    // Check if it's an HTMX request
    if h.IsRequest() {
        // Handle HTMX-specific logic
        h.PushUrl("/new-url").Trigger("myEvent").Apply(w.Header())
        w.Write([]byte("<div>Updated content</div>"))
    } else {
        // Handle regular HTTP request
        w.Write([]byte("<html>Full page</html>"))
    }
}
```

### Usage with fasthttp

```go
package main

import (
    "github.com/valyala/fasthttp"
    "github.com/dajooo/go-htmx"
)

func handler(ctx *fasthttp.RequestCtx) {
    h := htmx.NewFastHttp(&ctx.Request.Header)
    
    if h.IsRequest() {
        h.Redirect("/dashboard").Apply(&ctx.Response.Header)
        ctx.SetBody([]byte("<div>Redirecting...</div>"))
    }
}
```

## API Reference

### Creating HTMX Instance

```go
// For net/http
h := htmx.New(request.Header)

// For fasthttp  
h := htmx.NewFastHttp(&request.Header)

// Universal (auto-detects header type)
h := htmx.NewUniversal(header)
```

### Checking HTMX Request Status

```go
// Check if request is from HTMX
if htmx.IsHtmxRequest(request.Header) { ... }
if htmx.IsFastHttpHtmxRequest(&request.Header) { ... }

// Or using instance methods
if h.IsRequest() { ... }
if h.IsBoosted() { ... }
if h.IsHistoryRestoreRequest() { ... }
```

### Reading Request Information

```go
// Get request details
currentURL := h.Request.CurrentUrl
target := h.GetTarget()
trigger := h.GetTrigger()
triggerName := h.GetTriggerName() 
prompt := h.GetPrompt()
```

### Building Responses (Fluent API)

All response methods are chainable and return `*Htmx`:

```go
h.Location("/users/123").
  PushUrl("/users/123").
  Trigger("userUpdated").
  TriggerAfterSwap("highlight").
  Apply(responseHeader)
```

#### Response Methods

| Method | HTMX Header | Description |
|--------|-------------|-------------|
| `Location(url)` | `HX-Location` | Navigate to URL |
| `PushUrl(url)` | `HX-Push-URL` | Push URL to browser history |
| `Redirect(url)` | `HX-Redirect` | Client-side redirect |
| `Refresh(url)` | `HX-Refresh` | Refresh the page |
| `ReplaceUrl(url)` | `HX-Replace-URL` | Replace current URL in history |
| `Reswap(value)` | `HX-Reswap` | Change swap method |
| `Retarget(selector)` | `HX-Retarget` | Change target element |
| `Reselect(selector)` | `HX-Reselect` | Change content selection |
| `Trigger(event)` | `HX-Trigger` | Trigger event immediately |
| `TriggerAfterSettle(event)` | `HX-Trigger-After-Settle` | Trigger after settle |
| `TriggerAfterSwap(event)` | `HX-Trigger-After-Swap` | Trigger after swap |

## Complete Example

```go
package main

import (
    "encoding/json"
    "net/http"
    "github.com/dajooo/go-htmx"
)

type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

func updateUserHandler(w http.ResponseWriter, r *http.Request) {
    h := htmx.New(r.Header)
    
    if !h.IsRequest() {
        http.Error(w, "HTMX required", http.StatusBadRequest)
        return
    }
    
    // Get the target element that triggered the request
    target := h.GetTarget()
    trigger := h.GetTrigger()
    
    // Process the update
    user := User{ID: 1, Name: "Updated User"}
    
    // Build response with multiple HTMX directives
    h.PushUrl("/users/1").
      Trigger("userUpdated").
      TriggerAfterSwap("highlight").
      Apply(w.Header())
    
    // Send JSON data to trigger client-side event
    triggerData := map[string]interface{}{
        "user": user,
        "message": "User updated successfully",
    }
    triggerJSON, _ := json.Marshal(triggerData)
    h.Response.Trigger = string(triggerJSON)
    h.Apply(w.Header())
    
    w.Header().Set("Content-Type", "text/html")
    w.Write([]byte(`<div id="user-1">` + user.Name + `</div>`))
}

func main() {
    http.HandleFunc("/users/update", updateUserHandler)
    http.ListenAndServe(":8080", nil)
}
```

## Request Headers Supported

The library automatically parses these HTMX request headers:

- `HX-Request` - Always true for HTMX requests
- `HX-Boosted` - True if request is via hx-boost
- `HX-Current-URL` - Current page URL
- `HX-History-Restore-Request` - True if from history restoration
- `HX-Prompt` - User response from hx-prompt
- `HX-Target` - ID of target element
- `HX-Trigger` - ID of element that triggered request
- `HX-Trigger-Name` - Name of element that triggered request

## Response Headers Supported

The library can set these HTMX response headers:

- `HX-Location` - Navigate to URL
- `HX-Push-URL` - Push URL to history
- `HX-Redirect` - Client-side redirect
- `HX-Refresh` - Refresh page
- `HX-Replace-URL` - Replace URL in history
- `HX-Reswap` - Change swap behavior
- `HX-Retarget` - Change target element
- `HX-Reselect` - Change selection
- `HX-Trigger` - Trigger events
- `HX-Trigger-After-Settle` - Trigger after settle
- `HX-Trigger-After-Swap` - Trigger after swap

## License

MIT License

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.