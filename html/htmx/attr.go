package htmx

import (
	"strings"

	"github.com/zlsgo/app_module/html/el"
)

// HTMX 2.x attribute reference: https://htmx.org/reference/
//
// Constants intentionally mirror chasefleming/elem-go's htmx package so users
// can either call el.Attr(htmx.HXGet, "/url") directly or use the helper
// functions below.
const (
	// Core attributes.
	HXGet       = "hx-get"
	HXPost      = "hx-post"
	HXPushURL   = "hx-push-url"
	HXSelect    = "hx-select"
	HXSelectOOB = "hx-select-oob"
	HXSwap      = "hx-swap"
	HXSwapOOB   = "hx-swap-oob"
	HXTarget    = "hx-target"
	HXTrigger   = "hx-trigger"
	HXVals      = "hx-vals"

	// Additional attributes.
	HXBoost       = "hx-boost"
	HXConfirm     = "hx-confirm"
	HXDelete      = "hx-delete"
	HXDisable     = "hx-disable"
	HXDisabledElt = "hx-disabled-elt"
	HXDisinherit  = "hx-disinherit"
	HXEncoding    = "hx-encoding"
	HXExt         = "hx-ext"
	HXHeaders     = "hx-headers"
	HXHistory     = "hx-history"
	HXHistoryElt  = "hx-history-elt"
	HXInclude     = "hx-include"
	HXIndicator   = "hx-indicator"
	HXInherit     = "hx-inherit"
	HXParams      = "hx-params"
	HXPatch       = "hx-patch"
	HXPreserve    = "hx-preserve"
	HXPrompt      = "hx-prompt"
	HXPut         = "hx-put"
	HXReplaceURL  = "hx-replace-url"
	HXRequest     = "hx-request"
	HXSync        = "hx-sync"
	HXValidate    = "hx-validate"

	// Server Sent Events extension attributes.
	// Reference: https://htmx.org/extensions/sse/
	SSEConnect = "sse-connect"
	SSESwap    = "sse-swap"
	SSEClose   = "sse-close"

	// WebSockets extension attributes.
	// Reference: https://htmx.org/extensions/ws/
	WSConnect = "ws-connect"
	WSSend    = "ws-send"

	// HTMX event handler attributes.
	// Reference: https://htmx.org/reference/#events and https://htmx.org/attributes/hx-on/
	HXOnAbort                     = "hx-on--abort"
	HXOnAfterOnLoad               = "hx-on--after-on-load"
	HXOnAfterProcessNode          = "hx-on--after-process-node"
	HXOnAfterRequest              = "hx-on--after-request"
	HXOnAfterSettle               = "hx-on--after-settle"
	HXOnAfterSwap                 = "hx-on--after-swap"
	HXOnBeforeCleanupElement      = "hx-on--before-cleanup-element"
	HXOnBeforeOnLoad              = "hx-on--before-on-load"
	HXOnBeforeProcessNode         = "hx-on--before-process-node"
	HXOnBeforeRequest             = "hx-on--before-request"
	HXOnBeforeSwap                = "hx-on--before-swap"
	HXOnBeforeSend                = "hx-on--before-send"
	HXOnBeforeTransition          = "hx-on--before-transition"
	HXOnConfigRequest             = "hx-on--config-request"
	HXOnConfirm                   = "hx-on--confirm"
	HXOnHistoryCacheError         = "hx-on--history-cache-error"
	HXOnHistoryCacheHit           = "hx-on--history-cache-hit"
	HXOnHistoryCacheMiss          = "hx-on--history-cache-miss"
	HXOnHistoryCacheMissLoadError = "hx-on--history-cache-miss-load-error"
	HXOnHistoryCacheMissLoad      = "hx-on--history-cache-miss-load"
	HXOnHistoryRestore            = "hx-on--history-restore"
	HXOnBeforeHistorySave         = "hx-on--before-history-save"
	HXOnLoad                      = "hx-on--load"
	HXOnNoSSESourceError          = "hx-on--no-sse-source-error"
	HXOnOnLoadError               = "hx-on--on-load-error"
	HXOnOOBAfterSwap              = "hx-on--oob-after-swap"
	HXOnOOBBeforeSwap             = "hx-on--oob-before-swap"
	HXOnOOBErrorNoTarget          = "hx-on--oob-error-no-target"
	HXOnPrompt                    = "hx-on--prompt"
	HXOnPushedIntoHistory         = "hx-on--pushed-into-history"
	HXOnReplacedInHistory         = "hx-on--replaced-in-history"
	HXOnResponseError             = "hx-on--response-error"
	HXOnSendAbort                 = "hx-on--send-abort"
	HXOnSendError                 = "hx-on--send-error"
	HXOnSSEError                  = "hx-on--sse-error"
	HXOnSSEOpen                   = "hx-on--sse-open"
	HXOnSwapError                 = "hx-on--swap-error"
	HXOnTargetError               = "hx-on--target-error"
	HXOnTimeout                   = "hx-on--timeout"
	HXOnValidationValidate        = "hx-on--validation-validate"
	HXOnValidationFailed          = "hx-on--validation-failed"
	HXOnValidationHalted          = "hx-on--validation-halted"
	HXOnXHRAbort                  = "hx-on--xhr-abort"
	HXOnXHRLoadend                = "hx-on--xhr-loadend"
	HXOnXHRLoadstart              = "hx-on--xhr-loadstart"
	HXOnXHRProgress               = "hx-on--xhr-progress"
)

// Deprecated constants are kept for easier migration from older htmx examples.
const (
	// Deprecated: hx-vars is deprecated in htmx itself. Use HXVals instead.
	HXVars = "hx-vars"

	// Deprecated: hx-content is not an htmx attribute and has no effect.
	HXContent = "hx-content"

	// Deprecated: hx-values is not an htmx attribute and has no effect. Use HXVals instead.
	HXValues = "hx-values"

	// Deprecated: hx-timeout is not an htmx attribute and has no effect.
	// Set a timeout via HXRequest, e.g. hx-request='{"timeout": 1000}'.
	HXTimeout = "hx-timeout"

	// Deprecated: hx-retry is not an htmx attribute and has no effect.
	HXRetry = "hx-retry"

	// Deprecated: hx-retry-timeout is not an htmx attribute and has no effect.
	HXRetryTimeout = "hx-retry-timeout"

	// Deprecated: hx-triggering-element is not an htmx attribute and has no effect.
	HXTriggeringElement = "hx-triggering-element"

	// Deprecated: hx-triggering-event is not an htmx attribute and has no effect.
	HXTriggeringEvent = "hx-triggering-event"

	// Deprecated: hx-history-attr is not an htmx attribute and has no effect. Use HXHistory or HXHistoryElt instead.
	HXHistoryAttr = "hx-history-attr"

	// Deprecated: hx-error is not an htmx attribute and has no effect.
	HXError = "hx-error"

	// Deprecated: hx-cache is not an htmx attribute and has no effect.
	HXCache = "hx-cache"

	// Deprecated: hx-sse was removed in htmx 2.0. Use SSEConnect, SSESwap and SSEClose instead.
	HXSSE = "hx-sse"

	// Deprecated: hx-ws was removed in htmx 2.0. Use WSConnect and WSSend instead.
	HXWS = "hx-ws"

	// Deprecated: the bare hx-on attribute was removed in htmx 2.0. Use per-event HXOn* constants or On instead.
	HXOn = "hx-on"

	// Deprecated: use HXOnHistoryCacheMissLoadError instead.
	HXOnHistoryCacheMissError = "hx-on--history-cache-miss-load-error"
)

// Attr creates an htmx-related attribute with the same value semantics as el.Attr.
func Attr[T el.AttrValue](key string, value T) *el.Attribute { return el.Attr(key, value) }

// On creates a htmx event handler attribute.
//
// The event may be passed as a full attribute name ("hx-on--before-request"),
// htmx event name ("htmx:beforeRequest"), or shorthand event name
// ("before-request"). HTMX lifecycle events use htmx 2.x's
// hx-on--<event> form; ordinary DOM events use hx-on-<event>.
func On(event, script string) *el.Attribute { return el.Attr(OnName(event), script) }

// OnName normalizes a htmx event or hx-on attribute name to hx-on--<event>.
func OnName(event string) string {
	event = strings.TrimSpace(event)
	if event == "" {
		return "hx-on--"
	}
	if strings.HasPrefix(event, "hx-on--") {
		return "hx-on--" + kebabHTMXEvent(strings.TrimPrefix(event, "hx-on--"))
	}
	if strings.HasPrefix(event, "hx-on::") {
		event = strings.TrimPrefix(event, "hx-on::")
		return "hx-on--" + kebabHTMXEvent(event)
	}
	// The dashed long form is useful in markup where ':' is inconvenient:
	// hx-on-htmx-before-request corresponds to htmx:beforeRequest.
	if strings.HasPrefix(event, "hx-on-htmx-") {
		return "hx-on--" + kebabHTMXEvent(strings.TrimPrefix(event, "hx-on-htmx-"))
	}
	// Accept a complete ordinary DOM event attribute as well as an event name.
	// The double-dash HTMX form is handled above and must not be treated as
	// this single-dash form.
	if strings.HasPrefix(event, "hx-on-") {
		return "hx-on-" + kebabHTMXEvent(strings.TrimPrefix(event, "hx-on-"))
	}
	if strings.HasPrefix(event, "hx-on:") {
		event = strings.TrimPrefix(event, "hx-on:")
		if strings.HasPrefix(event, "htmx:") {
			return "hx-on--" + kebabHTMXEvent(strings.TrimPrefix(event, "htmx:"))
		}
		return "hx-on-" + event
	}
	if strings.HasPrefix(event, "htmx:") {
		return "hx-on--" + kebabHTMXEvent(strings.TrimPrefix(event, "htmx:"))
	}

	// Keep ordinary DOM events in the single-dash form. Bare HTMX event
	// shorthands such as "before-request" are recognized using the event
	// table below and use the double-dash form.
	if isHTMXEvent(kebabHTMXEvent(event)) {
		return "hx-on--" + kebabHTMXEvent(event)
	}
	return "hx-on-" + event
}

var htmxEvents = map[string]struct{}{
	"abort": {}, "after-on-load": {}, "after-process-node": {}, "after-request": {},
	"after-settle": {}, "after-swap": {}, "before-cleanup-element": {},
	"before-on-load": {}, "before-process-node": {}, "before-request": {},
	"before-swap": {}, "before-send": {}, "before-transition": {}, "config-request": {},
	"confirm": {}, "history-cache-error": {}, "history-cache-hit": {},
	"history-cache-miss": {}, "history-cache-miss-load-error": {},
	"history-cache-miss-load": {}, "history-restore": {}, "before-history-save": {},
	"load": {}, "no-sse-source-error": {}, "on-load-error": {}, "oob-after-swap": {},
	"oob-before-swap": {}, "oob-error-no-target": {}, "prompt": {},
	"pushed-into-history": {}, "replaced-in-history": {}, "response-error": {},
	"send-abort": {}, "send-error": {}, "sse-error": {}, "sse-open": {},
	"swap-error": {}, "target-error": {}, "timeout": {}, "validation-validate": {},
	"validation-failed": {}, "validation-halted": {}, "xhr-abort": {},
	"xhr-loadend": {}, "xhr-loadstart": {}, "xhr-progress": {},
}

func isHTMXEvent(event string) bool {
	_, ok := htmxEvents[event]
	return ok
}

func kebabHTMXEvent(event string) string {
	var b strings.Builder
	b.Grow(len(event) + 4)
	lastDash := false
	for i, r := range event {
		switch r {
		case ':', '_', ' ':
			if b.Len() > 0 && !lastDash {
				b.WriteByte('-')
				lastDash = true
			}
		case '-':
			if b.Len() > 0 && !lastDash {
				b.WriteByte('-')
				lastDash = true
			}
		default:
			if i > 0 && r >= 'A' && r <= 'Z' && b.Len() > 0 && !lastDash {
				b.WriteByte('-')
			}
			b.WriteRune(r)
			lastDash = false
		}
	}
	return strings.Trim(strings.ToLower(b.String()), "-")
}

// Core attribute helpers.
func Get(value string) *el.Attribute                { return el.Attr(HXGet, value) }
func Post(value string) *el.Attribute               { return el.Attr(HXPost, value) }
func PushURL[T el.AttrValue](value T) *el.Attribute { return el.Attr(HXPushURL, value) }
func Select(value string) *el.Attribute             { return el.Attr(HXSelect, value) }
func SelectOOB(value string) *el.Attribute          { return el.Attr(HXSelectOOB, value) }
func Swap(value string) *el.Attribute               { return el.Attr(HXSwap, value) }
func SwapOOB[T el.AttrValue](value T) *el.Attribute { return el.Attr(HXSwapOOB, value) }
func Target(value string) *el.Attribute             { return el.Attr(HXTarget, value) }
func Trigger(value string) *el.Attribute            { return el.Attr(HXTrigger, value) }
func Vals[T el.AttrValue](value T) *el.Attribute    { return el.Attr(HXVals, value) }

// Additional attribute helpers.
func Boost[T el.AttrValue](value T) *el.Attribute      { return el.Attr(HXBoost, value) }
func Confirm(value string) *el.Attribute               { return el.Attr(HXConfirm, value) }
func Delete(value string) *el.Attribute                { return el.Attr(HXDelete, value) }
func Disable() *el.Attribute                           { return el.BareAttr(HXDisable) }
func DisabledElt(value string) *el.Attribute           { return el.Attr(HXDisabledElt, value) }
func Disinherit(value string) *el.Attribute            { return el.Attr(HXDisinherit, value) }
func Encoding(value string) *el.Attribute              { return el.Attr(HXEncoding, value) }
func Ext(value string) *el.Attribute                   { return el.Attr(HXExt, value) }
func Headers[T el.AttrValue](value T) *el.Attribute    { return el.Attr(HXHeaders, value) }
func History[T el.AttrValue](value T) *el.Attribute    { return el.Attr(HXHistory, value) }
func HistoryElt(value string) *el.Attribute            { return el.Attr(HXHistoryElt, value) }
func Include(value string) *el.Attribute               { return el.Attr(HXInclude, value) }
func Indicator(value string) *el.Attribute             { return el.Attr(HXIndicator, value) }
func Inherit(value string) *el.Attribute               { return el.Attr(HXInherit, value) }
func Params(value string) *el.Attribute                { return el.Attr(HXParams, value) }
func Patch(value string) *el.Attribute                 { return el.Attr(HXPatch, value) }
func Preserve() *el.Attribute                          { return el.BareAttr(HXPreserve) }
func Prompt(value string) *el.Attribute                { return el.Attr(HXPrompt, value) }
func Put(value string) *el.Attribute                   { return el.Attr(HXPut, value) }
func ReplaceURL[T el.AttrValue](value T) *el.Attribute { return el.Attr(HXReplaceURL, value) }
func Request[T el.AttrValue](value T) *el.Attribute    { return el.Attr(HXRequest, value) }
func Sync(value string) *el.Attribute                  { return el.Attr(HXSync, value) }
func Validate[T el.AttrValue](value T) *el.Attribute   { return el.Attr(HXValidate, value) }

// Extension attribute helpers.
func SSEConnectAttr(value string) *el.Attribute { return el.Attr(SSEConnect, value) }
func SSESwapAttr(value string) *el.Attribute    { return el.Attr(SSESwap, value) }
func SSECloseAttr(value string) *el.Attribute   { return el.Attr(SSEClose, value) }
func WSConnectAttr(value string) *el.Attribute  { return el.Attr(WSConnect, value) }
func WSSendAttr() *el.Attribute                 { return el.BareAttr(WSSend) }
