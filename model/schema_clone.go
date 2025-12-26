package model

import "sync"

func cloneSchemaWithStorage(root *Schema, storage Storageer) *Schema {
	if root == nil || storage == nil {
		return root
	}

	baseGet := root.getSchema
	if baseGet == nil {
		clone := *root
		clone.Storage = storage
		clone.model = nil
		return &clone
	}

	cache := make(map[string]*Schema)
	var mu sync.Mutex

	var get func(alias string) (*Schema, bool)
	get = func(alias string) (*Schema, bool) {
		mu.Lock()
		if s, ok := cache[alias]; ok {
			mu.Unlock()
			return s, true
		}
		mu.Unlock()

		base, ok := baseGet(alias)
		if !ok || base == nil {
			return nil, false
		}

		clone := *base
		clone.Storage = storage
		clone.getSchema = get
		clone.model = nil

		mu.Lock()
		cache[alias] = &clone
		mu.Unlock()

		return &clone, true
	}

	alias := root.GetAlias()
	if alias != "" {
		if s, ok := get(alias); ok {
			return s
		}
	}

	clone := *root
	clone.Storage = storage
	clone.getSchema = get
	clone.model = nil

	if alias != "" {
		mu.Lock()
		cache[alias] = &clone
		mu.Unlock()
	}

	return &clone
}
