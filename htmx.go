package htmx

import "strconv"

type Swap string

const (
  SwapInnerHTML   Swap = "innerHTML"
  SwapOuterHTML   Swap = "outerHTML"
  SwapTextContent Swap = "textContent"
  SwapBeforebegin Swap = "beforebegin"
  SwapAfterbegin  Swap = "afterbegin"
  SwapBeforeend   Swap = "beforeend"
  SwapAfterend    Swap = "afterend"
  SwapDelete      Swap = "delete"
  SwapNone        Swap = "none"
)

type HtmxRequest struct {
  Request               bool
  Boosted               bool
  CurrentUrl            string
  HistoryRestoreRequest bool
  Prompt                string
  Target                string
  TriggerName           string
  Trigger               string
}

type HtmxResponse struct {
  Location           string
  PushUrl            string
  Redirect           string
  Refresh            string
  ReplaceUrl         string
  Reswap             Swap
  Retarget           string
  Reselect           string
  Trigger            string
  TriggerAfterSettle string
  TriggerAfterSwap   string
}

type RequestHeader interface {
  Get(key string) string
}

type FastHttpRequestHeader interface {
  Get(key string, defaultValue ...string) string
}

type ResponseHeader interface {
  Set(key, value string)
}

type Htmx struct {
  Request  HtmxRequest
  Response HtmxResponse
}

func New(header RequestHeader) *Htmx {
  return NewUniversal(header)
}

func NewFastHttp(requestHeader FastHttpRequestHeader) *Htmx {
  return NewUniversal(requestHeader)
}

func NewUniversal(header any) *Htmx {
  return &Htmx{
    Request: HtmxRequest{
      Request:               getHeader(header, "HX-Request") == "true",
      Boosted:               getHeader(header, "HX-Boosted") == "true",
      CurrentUrl:            getHeader(header, "HX-Current-URL"),
      HistoryRestoreRequest: getHeader(header, "HX-History-Restore-Request") == "true",
      Prompt:                getHeader(header, "HX-Prompt"),
      Target:                getHeader(header, "HX-Target"),
      TriggerName:           getHeader(header, "HX-Trigger-Name"),
      Trigger:               getHeader(header, "HX-Trigger"),
    },
    Response: HtmxResponse{},
  }
}

func IsHtmxRequest(header RequestHeader) bool {
  return getHeader(header, "HX-Request") == "true"
}

func IsFastHttpHtmxRequest(header FastHttpRequestHeader) bool {
  return getHeader(header, "HX-Request") == "true"
}

func (h *Htmx) IsRequest() bool {
  return h.Request.Request
}

func (h *Htmx) IsBoosted() bool {
  return h.Request.Boosted
}

func (h *Htmx) IsHistoryRestoreRequest() bool {
  return h.Request.HistoryRestoreRequest
}

func (h *Htmx) GetPrompt() string {
  return h.Request.Prompt
}

func (h *Htmx) GetTarget() string {
  return h.Request.Target
}

func (h *Htmx) GetTriggerName() string {
  return h.Request.TriggerName
}

func (h *Htmx) GetTrigger() string {
  return h.Request.Trigger
}

func (h *Htmx) Location(location string) *Htmx {
  h.Response.Location = location
  return h
}

func (h *Htmx) PushUrl(url string) *Htmx {
  h.Response.PushUrl = url
  return h
}

func (h *Htmx) Redirect(url string) *Htmx {
  h.Response.Redirect = url
  return h
}

func (h *Htmx) Refresh(refresh bool) *Htmx {
  h.Response.Refresh = strconv.FormatBool(refresh)
  return h
}

func (h *Htmx) ReplaceUrl(url string) *Htmx {
  h.Response.ReplaceUrl = url
  return h
}

func (h *Htmx) Reswap(swap Swap) *Htmx {
  h.Response.Reswap = swap
  return h
}

func (h *Htmx) Retarget(selector string) *Htmx {
  h.Response.Retarget = selector
  return h
}

func (h *Htmx) Reselect(selector string) *Htmx {
  h.Response.Reselect = selector
  return h
}

func (h *Htmx) Trigger(trigger string) *Htmx {
  h.Response.Trigger = trigger
  return h
}

func (h *Htmx) TriggerAfterSettle(trigger string) *Htmx {
  h.Response.TriggerAfterSettle = trigger
  return h
}

func (h *Htmx) TriggerAfterSwap(trigger string) *Htmx {
  h.Response.TriggerAfterSwap = trigger
  return h
}

func (h *Htmx) Apply(header ResponseHeader) *Htmx {
  if h.Response.Location != "" {
    header.Set("HX-Location", h.Response.Location)
  }
  if h.Response.PushUrl != "" {
    header.Set("HX-Push-URL", h.Response.PushUrl)
  }
  if h.Response.Redirect != "" {
    header.Set("HX-Redirect", h.Response.Redirect)
  }
  if h.Response.Refresh != "" {
    header.Set("HX-Refresh", h.Response.Refresh)
  }
  if h.Response.ReplaceUrl != "" {
    header.Set("HX-Replace-URL", h.Response.ReplaceUrl)
  }
  if h.Response.Reswap != "" {
    header.Set("HX-Reswap", string(h.Response.Reswap))
  }
  if h.Response.Retarget != "" {
    header.Set("HX-Retarget", h.Response.Retarget)
  }
  if h.Response.Reselect != "" {
    header.Set("HX-Reselect", h.Response.Reselect)
  }
  if h.Response.Trigger != "" {
    header.Set("HX-Trigger", h.Response.Trigger)
  }
  if h.Response.TriggerAfterSettle != "" {
    header.Set("HX-Trigger-After-Settle", h.Response.TriggerAfterSettle)
  }
  if h.Response.TriggerAfterSwap != "" {
    header.Set("HX-Trigger-After-Swap", h.Response.TriggerAfterSwap)
  }
  return h
}

func getHeader(header any, key string) string {
  netHttpHeader, ok := header.(RequestHeader)
  if ok {
    return netHttpHeader.Get(key)
  }
  fastHttpHeader, ok := header.(FastHttpRequestHeader)
  if ok {
    return fastHttpHeader.Get(key)
  }
  panic("unsupported header type")
}
