package model

import (
	"github.com/sohaha/zlsgo/ztype"
)

const DefaultBatchSize = 1000

type BatchOptions struct {
	Size int
}

type BatchOption func(*BatchOptions)

func BatchSize(size int) BatchOption {
	return func(o *BatchOptions) {
		if size > 0 {
			o.Size = size
		}
	}
}

// BatchInsert inserts data in batches and returns the IDs of inserted records.
// Note: The returned IDs count may be less than input data count if some records
// are skipped due to validation or empty data.
func (r *Repository[T]) BatchInsert(data ztype.Maps, opts ...BatchOption) ([]any, error) {
	if len(data) == 0 {
		return nil, nil
	}

	options := &BatchOptions{Size: DefaultBatchSize}
	for _, opt := range opts {
		opt(options)
	}

	allIDs := make([]any, 0, len(data))

	for i := 0; i < len(data); i += options.Size {
		end := i + options.Size
		if end > len(data) {
			end = len(data)
		}

		batch := data[i:end]
		ids, err := r.store.InsertMany(batch)
		if err != nil {
			return allIDs, err
		}

		if idSlice, ok := ids.([]interface{}); ok {
			allIDs = append(allIDs, idSlice...)
		} else if ids != nil {
			allIDs = append(allIDs, ids)
		}
	}

	return allIDs, nil
}

func (r *Repository[T]) BatchInsertTx(data ztype.Maps, opts ...BatchOption) ([]any, error) {
	if len(data) == 0 {
		return nil, nil
	}

	options := &BatchOptions{Size: DefaultBatchSize}
	for _, opt := range opts {
		opt(options)
	}

	var allIDs []any

	err := r.Tx(func(txRepo *Repository[T]) error {
		for i := 0; i < len(data); i += options.Size {
			end := i + options.Size
			if end > len(data) {
				end = len(data)
			}

			batch := data[i:end]
			ids, err := txRepo.store.InsertMany(batch)
			if err != nil {
				return err
			}

			if idSlice, ok := ids.([]interface{}); ok {
				allIDs = append(allIDs, idSlice...)
			} else if ids != nil {
				allIDs = append(allIDs, ids)
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return allIDs, nil
}

func (r *Repository[T]) BatchUpdate(filter QueryFilter, data ztype.Map, opts ...BatchOption) (int64, error) {
	options := &BatchOptions{Size: DefaultBatchSize}
	for _, opt := range opts {
		opt(options)
	}

	rows, err := r.store.Find(Filter(filter.ToMap()), func(co *CondOptions) {
		co.Fields = []string{idKey}
	})
	if err != nil {
		return 0, err
	}

	if len(rows) == 0 {
		return 0, nil
	}

	var total int64

	for i := 0; i < len(rows); i += options.Size {
		end := i + options.Size
		if end > len(rows) {
			end = len(rows)
		}

		ids := make([]any, end-i)
		for j, row := range rows[i:end] {
			ids[j] = row.Get(idKey).Value()
		}

		count, err := r.store.UpdateMany(Filter(In(idKey, ids).ToMap()), data)
		if err != nil {
			return total, err
		}
		total += count
	}

	return total, nil
}

func (r *Repository[T]) BatchDelete(filter QueryFilter, opts ...BatchOption) (int64, error) {
	options := &BatchOptions{Size: DefaultBatchSize}
	for _, opt := range opts {
		opt(options)
	}

	var total int64

	for {
		count, err := r.store.DeleteMany(Filter(filter.ToMap()), func(co *CondOptions) {
			co.Limit = options.Size
		})
		if err != nil {
			return total, err
		}

		total += count

		if count < int64(options.Size) {
			break
		}
	}

	return total, nil
}
