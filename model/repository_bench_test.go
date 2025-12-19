package model

import (
	"testing"

	"github.com/sohaha/zlsgo/ztype"
)

func BenchmarkRepositoryInsert(b *testing.B) {
	_, m := newTestDB(&testing.T{}, "bench_insert")
	repo := m.Model().Repository()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = repo.Insert(ztype.Map{
			"name":   "BenchUser",
			"email":  "bench@test.com",
			"age":    25,
			"status": 1,
		})
	}
}

func BenchmarkRepositoryFind(b *testing.B) {
	_, m := newTestDB(&testing.T{}, "bench_find")
	repo := m.Model().Repository()

	for i := 0; i < 100; i++ {
		_, _ = repo.Insert(ztype.Map{
			"name":   "BenchUser" + ztype.ToString(i),
			"email":  "bench" + ztype.ToString(i) + "@test.com",
			"age":    20 + i%30,
			"status": 1,
		})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = repo.Find(Eq("status", 1))
	}
}

func BenchmarkRepositoryFindByID(b *testing.B) {
	_, m := newTestDB(&testing.T{}, "bench_find_by_id")
	repo := m.Model().Repository()

	id, _ := repo.Insert(ztype.Map{
		"name":   "BenchUser",
		"email":  "bench@test.com",
		"age":    25,
		"status": 1,
	})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = repo.FindByID(id)
	}
}

func BenchmarkRepositoryUpdate(b *testing.B) {
	_, m := newTestDB(&testing.T{}, "bench_update")
	repo := m.Model().Repository()

	id, _ := repo.Insert(ztype.Map{
		"name":   "BenchUser",
		"email":  "bench@test.com",
		"age":    25,
		"status": 1,
	})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = repo.UpdateByID(id, ztype.Map{"age": 26})
	}
}

func BenchmarkRepositoryDelete(b *testing.B) {
	_, m := newTestDB(&testing.T{}, "bench_delete")
	repo := m.Model().Repository()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		id, _ := repo.Insert(ztype.Map{
			"name":   "BenchUser",
			"email":  "bench@test.com",
			"age":    25,
			"status": 1,
		})
		b.StartTimer()

		_, _ = repo.DeleteByID(id)
	}
}

func BenchmarkRepositoryQueryBuilder(b *testing.B) {
	_, m := newTestDB(&testing.T{}, "bench_query")
	repo := m.Model().Repository()

	for i := 0; i < 100; i++ {
		_, _ = repo.Insert(ztype.Map{
			"name":   "QueryUser" + ztype.ToString(i),
			"email":  "query" + ztype.ToString(i) + "@test.com",
			"age":    20 + i%30,
			"status": i % 2,
		})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = repo.Query().
			Where("status", 1).
			WhereGt("age", 25).
			OrderByDesc("age").
			Limit(10).
			Find()
	}
}

func BenchmarkRepositoryStructMapper(b *testing.B) {
	_, m := newTestDB(&testing.T{}, "bench_struct")
	repo := Repo[TestUser](m.Model())

	for i := 0; i < 100; i++ {
		_, _ = repo.Insert(ztype.Map{
			"name":   "StructUser" + ztype.ToString(i),
			"email":  "struct" + ztype.ToString(i) + "@test.com",
			"age":    20 + i%30,
			"status": 1,
		})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = repo.Find(Eq("status", 1))
	}
}

func BenchmarkRepositoryBatchInsert(b *testing.B) {
	_, m := newTestDB(&testing.T{}, "bench_batch")
	repo := m.Model().Repository()

	data := make(ztype.Maps, 100)
	for i := 0; i < 100; i++ {
		data[i] = ztype.Map{
			"name":   "BatchUser" + ztype.ToString(i),
			"email":  "batch" + ztype.ToString(i) + "@test.com",
			"age":    20 + i%30,
			"status": 1,
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = repo.BatchInsert(data, BatchSize(20))
	}
}

func BenchmarkQueryFilterBuild(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = And(
			Eq("status", 1),
			Gt("age", 18),
			In("role", []string{"admin", "user"}),
			Or(
				Like("name", "%test%"),
				IsNotNull("email"),
			),
		).ToMap()
	}
}

func BenchmarkCondOptionsPool(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		opts := acquireCondOptions()
		opts.Fields = append(opts.Fields, "id", "name", "email")
		opts.Limit = 10
		releaseCondOptions(opts)
	}
}

func BenchmarkCondOptionsNoPool(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		opts := &CondOptions{
			Fields:  make([]string, 0, 8),
			OrderBy: make([]OrderByItem, 0, 4),
			GroupBy: make([]string, 0, 4),
		}
		opts.Fields = append(opts.Fields, "id", "name", "email")
		opts.Limit = 10
		_ = opts
	}
}
