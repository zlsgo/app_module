package model

// DefaultBatchSize 默认批次大小
const DefaultBatchSize = 1000

// BatchOptions 批次选项
type BatchOptions struct {
	Size int
}

// BatchOption 批次选项函数
type BatchOption func(*BatchOptions)

// BatchSize 设置批次大小
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
// BatchInsert 批量插入数据
func (r *Repository[T, F, C, U]) BatchInsert(data []C, opts ...BatchOption) ([]any, error) {
	dataMaps, err := dataToMaps(data)
	if err != nil {
		return nil, err
	}
	if len(dataMaps) == 0 {
		return nil, nil
	}

	options := &BatchOptions{Size: DefaultBatchSize}
	for _, opt := range opts {
		opt(options)
	}

	allIDs := make([]any, 0, len(dataMaps))

	for i := 0; i < len(dataMaps); i += options.Size {
		end := i + options.Size
		if end > len(dataMaps) {
			end = len(dataMaps)
		}

		batch := dataMaps[i:end]
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

// BatchInsertTx 在事务中批量插入数据
func (r *Repository[T, F, C, U]) BatchInsertTx(data []C, opts ...BatchOption) ([]any, error) {
	dataMaps, err := dataToMaps(data)
	if err != nil {
		return nil, err
	}
	if len(dataMaps) == 0 {
		return nil, nil
	}

	options := &BatchOptions{Size: DefaultBatchSize}
	for _, opt := range opts {
		opt(options)
	}

	var allIDs []any

	err = r.Tx(func(txRepo *Repository[T, F, C, U]) error {
		for i := 0; i < len(dataMaps); i += options.Size {
			end := i + options.Size
			if end > len(dataMaps) {
				end = len(dataMaps)
			}

			batch := dataMaps[i:end]
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

// BatchUpdate 批量更新数据
func (r *Repository[T, F, C, U]) BatchUpdate(filter F, data U, opts ...BatchOption) (int64, error) {
	options := &BatchOptions{Size: DefaultBatchSize}
	for _, opt := range opts {
		opt(options)
	}

	rows, err := r.store.Find(Q(filter), func(co *CondOptions) {
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

		count, err := r.store.UpdateMany(In(idKey, ids), data)
		if err != nil {
			return total, err
		}
		total += count
	}

	return total, nil
}

// BatchDelete 批量删除数据
func (r *Repository[T, F, C, U]) BatchDelete(filter F, opts ...BatchOption) (int64, error) {
	options := &BatchOptions{Size: DefaultBatchSize}
	for _, opt := range opts {
		opt(options)
	}

	var total int64

	for {
		count, err := r.store.DeleteMany(Q(filter), func(co *CondOptions) {
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
