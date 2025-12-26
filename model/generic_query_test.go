package model

import (
	"errors"
	"testing"

	"github.com/sohaha/zlsgo"
	"github.com/sohaha/zlsgo/ztype"
)

func TestFindOne(t *testing.T) {
	tt := zlsgo.NewTest(t)
	_, m := newTestDB(t, "find_one_as")
	store := m.Model()

	_, err := store.Insert(TestUserCreate{
		Name:   "User1",
		Email:  "u1@test.com",
		Age:    20,
		Status: 1,
	})
	tt.NoError(err)

	user, err := FindOne[TestUser](store, Q(TestUserFilter{Name: "User1"}))
	tt.NoError(err)
	tt.Equal("User1", user.Name)

	_, err = FindOne[TestUser](store, Q(TestUserFilter{Name: "Missing"}))
	tt.Equal(true, errors.Is(err, ErrNoRecord))
}

func TestPages(t *testing.T) {
	tt := zlsgo.NewTest(t)
	_, m := newTestDB(t, "pages_as")
	store := m.Model()

	for i := 0; i < 3; i++ {
		_, err := store.Insert(TestUserCreate{
			Name:   "User" + ztype.ToString(i+1),
			Email:  "u" + ztype.ToString(i+1) + "@test.com",
			Age:    20 + i,
			Status: 1,
		})
		tt.NoError(err)
	}

	pageData, err := Pages[TestUser](m, 1, 2, Q(TestUserFilter{Status: 1}))
	tt.NoError(err)
	tt.Equal(2, len(pageData.Items))
	tt.Equal(uint(3), pageData.Page.Total)
}

func TestFindCols(t *testing.T) {
	tt := zlsgo.NewTest(t)
	_, m := newTestDB(t, "find_cols_as")
	store := m.Model()

	_, err := store.Insert(TestUserCreate{Name: "User1", Email: "u1@test.com", Age: 22, Status: 1})
	tt.NoError(err)
	_, err = store.Insert(TestUserCreate{Name: "User2", Email: "u2@test.com", Age: 25, Status: 1})
	tt.NoError(err)

	ages, err := FindCols[int](store, "age", Q(TestUserFilter{Status: 1}))
	tt.NoError(err)
	tt.Equal(2, len(ages))

	age, ok, err := FindCol[int](store, "age", Q(TestUserFilter{Name: "User1"}))
	tt.NoError(err)
	tt.Equal(true, ok)
	tt.Equal(22, age)
}
