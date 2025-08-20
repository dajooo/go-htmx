package htmx

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
	Reswap             string
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

func (h *Htmx) Refresh(url string) *Htmx {
	h.Response.Refresh = url
	return h
}

func (h *Htmx) ReplaceUrl(url string) *Htmx {
	h.Response.ReplaceUrl = url
	return h
}

func (h *Htmx) Reswap(url string) *Htmx {
	h.Response.Reswap = url
	return h
}

func (h *Htmx) Retarget(url string) *Htmx {
	h.Response.Retarget = url
	return h
}

func (h *Htmx) Reselect(url string) *Htmx {
	h.Response.Reselect = url
	return h
}

func (h *Htmx) Trigger(name string) *Htmx {
	h.Response.Trigger = name
	return h
}

func (h *Htmx) TriggerAfterSettle(name string) *Htmx {
	h.Response.TriggerAfterSettle = name
	return h
}

func (h *Htmx) TriggerAfterSwap(name string) *Htmx {
	h.Response.TriggerAfterSwap = name
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
		header.Set("HX-Reswap", h.Response.Reswap)
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
