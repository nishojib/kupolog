package repository_test

import (
	"context"
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	repoErrors "github.com/nishojib/ffxivdailies/internal/errors"
	"github.com/nishojib/ffxivdailies/internal/repository"
	"github.com/nishojib/ffxivdailies/internal/task"
	"github.com/nishojib/ffxivdailies/migrations"
	"github.com/pressly/goose/v3"
	"github.com/pressly/goose/v3/database"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
)

func TestAddUserTask_Success(t *testing.T) {
	db, tearDown := setupTest(t)
	defer tearDown(t)

	bunDB := bun.NewDB(db, sqlitedialect.New())

	repo := repository.New(bunDB)

	tsk := task.Task{
		UserID:    task.ID("test"),
		TaskID:    task.ID("test"),
		Completed: true,
		Hidden:    false,
		Kind:      "daily",
	}

	err := repo.AddUserTask(context.Background(), &tsk)
	require.NoError(t, err)

	var got task.Task
	err = bunDB.NewSelect().
		Model(&got).
		Where("ut.task_id = ?", tsk.TaskID).
		Where("ut.user_id = ?", tsk.UserID).
		Scan(context.Background())
	require.NoError(t, err)

	assert.Equal(t, tsk, got)

	_, err = bunDB.NewDelete().
		Model(&got).
		Where("ut.task_id = ?", got.TaskID).
		Where("ut.user_id = ?", got.UserID).
		Exec(context.Background())
	require.NoError(t, err)
}

func TestAddUserTask_Error(t *testing.T) {
	db, tearDown := setupTest(t)
	defer tearDown(t)

	bunDB := bun.NewDB(db, sqlitedialect.New())

	repo := repository.New(bunDB)

	tsk1 := task.Task{
		UserID:    task.ID("test"),
		TaskID:    task.ID("test"),
		Completed: true,
		Hidden:    false,
		Kind:      "daily",
	}

	err := repo.AddUserTask(context.Background(), &tsk1)
	require.NoError(t, err)

	tsk2 := task.Task{
		UserID:    task.ID("test"),
		TaskID:    task.ID("test"),
		Completed: true,
		Hidden:    false,
		Kind:      "weekly",
	}

	err = repo.AddUserTask(context.Background(), &tsk2)
	require.Error(t, err)

	_, err = bunDB.NewDelete().
		Model(&tsk1).
		Where("ut.task_id = ?", tsk1.TaskID).
		Where("ut.user_id = ?", tsk1.UserID).
		Exec(context.Background())
	require.NoError(t, err)
}

func TestUpdateUserTask(t *testing.T) {
	db, tearDown := setupTest(t)
	defer tearDown(t)

	bunDB := bun.NewDB(db, sqlitedialect.New())

	repo := repository.New(bunDB)

	tsk := task.Task{
		UserID:    task.ID("test"),
		TaskID:    task.ID("test"),
		Completed: true,
		Hidden:    false,
		Kind:      "daily",
	}
	err := repo.AddUserTask(context.Background(), &tsk)
	require.NoError(t, err)

	tsk.Completed = true
	tsk.Hidden = true
	err = repo.UpdateUserTask(context.Background(), &tsk)
	require.NoError(t, err)

	var got task.Task
	err = bunDB.NewSelect().
		Model(&got).
		Where("ut.task_id = ?", tsk.TaskID).
		Where("ut.user_id = ?", tsk.UserID).
		Scan(context.Background())
	require.NoError(t, err)

	assert.Equal(t, tsk.Completed, got.Completed)
	assert.Equal(t, tsk.Hidden, got.Hidden)
	assert.Equal(t, tsk.Version+1, got.Version)

	_, err = bunDB.NewDelete().
		Model(&got).
		Where("ut.task_id = ?", got.TaskID).
		Where("ut.user_id = ?", got.UserID).
		Exec(context.Background())
	require.NoError(t, err)
}

func TestUpdateUserTask_Error(t *testing.T) {
	db, tearDown := setupTest(t)
	defer tearDown(t)

	bunDB := bun.NewDB(db, sqlitedialect.New())

	repo := repository.New(bunDB)

	err := repo.UpdateUserTask(context.Background(), &task.Task{})
	require.ErrorIs(t, err, repoErrors.ErrEditConflict)

	tsk := task.Task{
		UserID:    task.ID("test"),
		TaskID:    task.ID("test"),
		Completed: true,
		Hidden:    false,
		Kind:      "daily",
	}
	err = repo.AddUserTask(context.Background(), &tsk)
	require.NoError(t, err)

	tsk.Version = 3

	err = repo.UpdateUserTask(context.Background(), &tsk)
	require.Error(t, err)
}

func setupTest(t *testing.T) (*sql.DB, func(t *testing.T)) {
	db, err := sql.Open("libsql", "file::memory:?cache=shared")
	if err != nil {
		t.Fatal(err)
	}

	goose.SetBaseFS(migrations.EmbedMigrations)

	if err := goose.SetDialect(string(database.DialectTurso)); err != nil {
		t.Fatal(err)
	}

	if err := goose.Up(db, "."); err != nil {
		t.Fatal(err)
	}

	return db, func(t *testing.T) {
		if err := goose.Down(db, "."); err != nil {
			t.Fatal(err)
		}
	}
}
