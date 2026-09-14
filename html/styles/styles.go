// Package styles provides deterministic CSS values, inline style helpers and
// an optional in-memory stylesheet manager for html/el.
//
// It deliberately has no dependency on the html module or frontend zcss
// runtime. Use it for server-rendered CSS; use html/static for zcss assets.
package styles

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/zlsgo/app_module/html/el"
)

// Props is a CSS declaration set. Property names and values are emitted as
// provided; callers should only use trusted CSS values when using CSS.
type Props map[string]string

// ToInline returns a deterministic inline CSS declaration string.
func (p Props) ToInline() string {
	keys := sortedKeys(p)
	var b strings.Builder
	for _, key := range keys {
		if strings.TrimSpace(key) == "" {
			continue
		}
		b.WriteString(key)
		b.WriteString(": ")
		b.WriteString(p[key])
		b.WriteString("; ")
	}
	return strings.TrimSuffix(b.String(), " ")
}

// Attr converts declarations into the HTML style attribute.
func (p Props) Attr() *el.Attribute { return el.Style(p.ToInline()) }

// Merge combines declarations; later sets override earlier sets.
func Merge(sets ...Props) Props {
	result := make(Props)
	for _, set := range sets {
		for key, value := range set {
			result[key] = value
		}
	}
	return result
}

// CSS returns trusted raw CSS content for use inside el.STYLE.
func CSS(content string) el.Raw { return el.Raw(content) }

// StyleTag creates a style element containing trusted raw CSS.
func StyleTag(content string) *el.Element { return el.STYLE(CSS(content)) }

// Keyframes maps keyframe selectors such as "from", "50%" and "to" to declarations.
type Keyframes map[string]Props

// CompositeStyle describes a class and its pseudo/media variants.
type CompositeStyle struct {
	Default        Props
	PseudoClasses  map[string]Props
	PseudoElements map[string]Props
	MediaQueries   map[string]Props
}

// StyleManager deduplicates declarations and generates deterministic CSS.
type StyleManager struct {
	mu              sync.RWMutex
	styles          map[string]Props
	compositeStyles map[string]CompositeStyle
	animations      map[string]Keyframes
}

// NewStyleManager creates an empty stylesheet manager.
func NewStyleManager() *StyleManager {
	return &StyleManager{
		styles:          make(map[string]Props),
		compositeStyles: make(map[string]CompositeStyle),
		animations:      make(map[string]Keyframes),
	}
}

// AddStyle registers declarations and returns a stable generated class name.
func (m *StyleManager) AddStyle(style Props) string {
	name := "cls_" + hashValue(style)
	m.mu.Lock()
	defer m.mu.Unlock()
	m.initLocked()
	if _, ok := m.styles[name]; !ok {
		m.styles[name] = cloneProps(style)
	}
	return name
}

// Class registers declarations and returns a class attribute ready for an
// element constructor.
func (m *StyleManager) Class(style Props) *el.Attribute {
	return el.Class(m.AddStyle(style))
}

// AddCompositeStyle registers a class with pseudo-elements/classes and media queries.
func (m *StyleManager) AddCompositeStyle(style CompositeStyle) string {
	name := "cls_" + hashValue(style)
	m.mu.Lock()
	defer m.mu.Unlock()
	m.initLocked()
	if _, ok := m.compositeStyles[name]; !ok {
		m.compositeStyles[name] = cloneComposite(style)
	}
	return name
}

// CompositeClass registers a composite style and returns a class attribute.
func (m *StyleManager) CompositeClass(style CompositeStyle) *el.Attribute {
	return el.Class(m.AddCompositeStyle(style))
}

// AddAnimation registers keyframes and returns a stable animation name.
func (m *StyleManager) AddAnimation(keyframes Keyframes) string {
	name := "anim_" + hashValue(keyframes)
	m.mu.Lock()
	defer m.mu.Unlock()
	m.initLocked()
	if _, ok := m.animations[name]; !ok {
		m.animations[name] = cloneKeyframes(keyframes)
	}
	return name
}

// Animation registers keyframes and returns a CSS animation-name value.
func (m *StyleManager) Animation(keyframes Keyframes) string {
	return m.AddAnimation(keyframes)
}

// GenerateCSS emits all managed rules in deterministic order.
func (m *StyleManager) GenerateCSS() string {
	m.mu.RLock()
	defer m.mu.RUnlock()
	var b strings.Builder
	styleNames := sortedKeys(m.styles)
	for _, name := range styleNames {
		writeRule(&b, "."+name, m.styles[name])
	}
	animationNames := sortedKeys(m.animations)
	for _, name := range animationNames {
		b.WriteString("@keyframes ")
		b.WriteString(name)
		b.WriteString(" { ")
		for _, frame := range sortedKeys(m.animations[name]) {
			writeRuleBody(&b, frame, m.animations[name][frame])
		}
		b.WriteString("} ")
	}
	compositeNames := sortedKeys(m.compositeStyles)
	for _, name := range compositeNames {
		composite := m.compositeStyles[name]
		writeRule(&b, "."+name, composite.Default)
		for _, pseudo := range sortedKeys(composite.PseudoClasses) {
			writeRule(&b, "."+name+ensurePrefix(pseudo, ":"), composite.PseudoClasses[pseudo])
		}
		for _, pseudo := range sortedKeys(composite.PseudoElements) {
			writeRule(&b, "."+name+ensurePrefix(pseudo, "::"), composite.PseudoElements[pseudo])
		}
		for _, media := range sortedKeys(composite.MediaQueries) {
			query := media
			if !strings.HasPrefix(query, "@media") {
				query = "@media " + query
			}
			b.WriteString(query)
			b.WriteString(" { ")
			writeRuleBody(&b, "."+name, composite.MediaQueries[media])
			b.WriteString("} ")
		}
	}
	return b.String()
}

// StyleTag returns a style element containing the generated stylesheet.
func (m *StyleManager) StyleTag() *el.Element { return StyleTag(m.GenerateCSS()) }

func (m *StyleManager) initLocked() {
	if m.styles == nil {
		m.styles = make(map[string]Props)
	}
	if m.compositeStyles == nil {
		m.compositeStyles = make(map[string]CompositeStyle)
	}
	if m.animations == nil {
		m.animations = make(map[string]Keyframes)
	}
}

func writeRule(b *strings.Builder, selector string, props Props) {
	b.WriteString(selector)
	b.WriteString(" { ")
	writeDeclarations(b, props)
	b.WriteString("} ")
}

func writeRuleBody(b *strings.Builder, selector string, props Props) {
	b.WriteString(selector)
	b.WriteString(" { ")
	writeDeclarations(b, props)
	b.WriteString("} ")
}

func writeDeclarations(b *strings.Builder, props Props) {
	for _, key := range sortedKeys(props) {
		if strings.TrimSpace(key) == "" {
			continue
		}
		b.WriteString(key)
		b.WriteString(": ")
		b.WriteString(props[key])
		b.WriteString("; ")
	}
}

func ensurePrefix(value, prefix string) string {
	value = strings.TrimSpace(value)
	value = strings.TrimLeft(value, ":")
	return prefix + value
}

func sortedKeys[T any](values map[string]T) []string {
	keys := make([]string, 0, len(values))
	for key := range values {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

func cloneProps(props Props) Props { return Merge(props) }

func cloneComposite(value CompositeStyle) CompositeStyle {
	result := CompositeStyle{Default: cloneProps(value.Default), PseudoClasses: make(map[string]Props), PseudoElements: make(map[string]Props), MediaQueries: make(map[string]Props)}
	for key, props := range value.PseudoClasses {
		result.PseudoClasses[key] = cloneProps(props)
	}
	for key, props := range value.PseudoElements {
		result.PseudoElements[key] = cloneProps(props)
	}
	for key, props := range value.MediaQueries {
		result.MediaQueries[key] = cloneProps(props)
	}
	return result
}

func cloneKeyframes(value Keyframes) Keyframes {
	result := make(Keyframes)
	for key, props := range value {
		result[key] = cloneProps(props)
	}
	return result
}

func hashValue(value any) string {
	data, _ := json.Marshal(value)
	sum := sha256.Sum256(data)
	return fmt.Sprintf("%x", sum[:5])
}

func Em(value float64) string  { return strconv.FormatFloat(value, 'f', 2, 64) + "em" }
func Rem(value float64) string { return strconv.FormatFloat(value, 'f', 2, 64) + "rem" }
func Ch(value float64) string  { return strconv.FormatFloat(value, 'f', 2, 64) + "ch" }
func Ex(value float64) string  { return strconv.FormatFloat(value, 'f', 2, 64) + "ex" }
func Pixels(value int) string {
	if value == 0 {
		return "0"
	}
	return strconv.Itoa(value) + "px"
}
func Percent(value int) string        { return strconv.Itoa(value) + "%" }
func ViewportHeight(value int) string { return strconv.Itoa(value) + "vh" }
func ViewportWidth(value int) string  { return strconv.Itoa(value) + "vw" }
func ViewportMin(value int) string    { return strconv.Itoa(value) + "vmin" }
func ViewportMax(value int) string    { return strconv.Itoa(value) + "vmax" }
func Seconds(value float64) string    { return strconv.FormatFloat(value, 'f', 2, 64) + "s" }
func Milliseconds(value int) string   { return strconv.Itoa(value) + "ms" }
func Int(value int) string            { return strconv.Itoa(value) }
func Float(value float64) string      { return fmt.Sprintf("%.2f", value) }
func RGB(r, g, b int) string          { return fmt.Sprintf("rgb(%d,%d,%d)", r, g, b) }
func RGBA(r, g, b int, a float64) string {
	return fmt.Sprintf("rgba(%d,%d,%d,%.1f)", r, g, b, a)
}
func HSL(h, s, l int) string { return fmt.Sprintf("hsl(%d,%d%%,%d%%)", h, s, l) }
func HSLA(h, s, l int, a float64) string {
	return fmt.Sprintf("hsla(%d,%d%%,%d%%,%.2f)", h, s, l, a)
}
func URL(value string) string { return "url('" + strings.ReplaceAll(value, "'", "\\'") + "')" }
func Var(name string) string  { return "var(--" + strings.TrimPrefix(name, "--") + ")" }
