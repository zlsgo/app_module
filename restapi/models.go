package restapi

// type Models struct {
// 	db *zdb.DB
// 	ms *model.Models
// }

// func (p *Models) GetDB() *zdb.DB {
// 	return p.db
// }

// func (p *Models) GetModel(name string) (*model.Model, bool) {
// 	return p.ms.Get(name)
// }

// func (p *Models) RegModel(name string, data model.Define, force bool) error {
// 	_, err := p.ms.Reg(name, data, force)
// 	return err
// }

// func NewModels(models []model.Define, db *zdb.DB, async bool) (p *Models, err error) {
// 	p = &Models{
// 		db: db,
// 	}

// 	if p.db == nil {
// 		return nil, zerror.With(err, "init db error")
// 	}

// 	p.ms = model.New(model.NewSQL(p.db), func(o *model.ModelOptions) {
// 		o.Prefix = "model_"
// 	})

// 	run := func(name string, data model.Define) error {
// 		err = zerror.TryCatch(func() error {
// 			_, err := p.ms.Reg(name, data, false)
// 			return err
// 		})
// 		if err != nil {
// 			return zerror.With(err, "register "+name+" model error")
// 		}
// 		return nil
// 	}

// 	if async {
// 		size := len(models)
// 		pool := zpool.New(size)
// 		defer pool.Close()
// 		errChan := make(chan error, size)
// 		asyncRun := func(name string, data model.Define) {
// 			pool.Do(func() {
// 				errChan <- run(name, data)
// 			})
// 		}

// 		for _, data := range models {
// 			asyncRun(data.Name, data)
// 		}

// 		pool.Wait()

// 		for i := 0; i < size; i++ {
// 			err = <-errChan
// 			if err != nil {
// 				return
// 			}
// 		}
// 		return
// 	}

// 	for _, data := range models {
// 		err = run(data.Name, data)
// 		if err != nil {
// 			return
// 		}
// 	}

// 	return
// }
