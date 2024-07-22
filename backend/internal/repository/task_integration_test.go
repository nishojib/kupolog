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
	db, tearDown := setupIntegrationTest(t)
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
}

func TestAddUserTask_Error(t *testing.T) {
	db, tearDown := setupIntegrationTest(t)
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
}

func TestUpdateUserTask_Success(t *testing.T) {
	db, tearDown := setupIntegrationTest(t)
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
}

func TestUpdateUserTask_Error(t *testing.T) {
	db, tearDown := setupIntegrationTest(t)
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
	require.ErrorIs(t, err, repoErrors.ErrEditConflict)

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	err = repo.UpdateUserTask(ctx, &tsk)
	require.Error(t, err)
}

func TestGetUserTask_Success(t *testing.T) {
	db, tearDown := setupIntegrationTest(t)
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

	got, err := repo.GetUserTask(context.Background(), "test", "test")
	require.NoError(t, err)

	assert.Equal(t, tsk, got)
}

func TestGetUserTask_Error(t *testing.T) {
	db, tearDown := setupIntegrationTest(t)
	defer tearDown(t)

	bunDB := bun.NewDB(db, sqlitedialect.New())

	repo := repository.New(bunDB)

	_, err := repo.GetUserTask(context.Background(), "test", "test")
	require.ErrorIs(t, err, repoErrors.ErrRecordNotFound)

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err = repo.GetUserTask(ctx, "test", "test")
	require.Error(t, err)
}

func TestGetTasksForUser_Success(t *testing.T) {
	db, tearDown := setupIntegrationTest(t)
	defer tearDown(t)

	bunDB := bun.NewDB(db, sqlitedialect.New())

	repo := repository.New(bunDB)

	tsk1 := task.Task{
		UserID:    task.ID("test"),
		TaskID:    task.ID("test1"),
		Completed: true,
		Hidden:    false,
		Kind:      "daily",
	}

	err := repo.AddUserTask(context.Background(), &tsk1)
	require.NoError(t, err)

	tsk2 := task.Task{
		UserID:    task.ID("test"),
		TaskID:    task.ID("test2"),
		Completed: true,
		Hidden:    false,
		Kind:      "weekly",
	}

	err = repo.AddUserTask(context.Background(), &tsk2)
	require.NoError(t, err)

	got, err := repo.GetTasksForUser(context.Background(), "test")
	require.NoError(t, err)

	assert.Equal(t, []task.Task{tsk1, tsk2}, got)
}

func TestGetTasksForUser_Error(t *testing.T) {
	db, tearDown := setupIntegrationTest(t)
	defer tearDown(t)

	bunDB := bun.NewDB(db, sqlitedialect.New())

	repo := repository.New(bunDB)

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := repo.GetTasksForUser(ctx, "test")
	require.Error(t, err)
}

func TestUpdateTaskForKind_Success(t *testing.T) {
	db, tearDown := setupIntegrationTest(t)
	defer tearDown(t)

	bunDB := bun.NewDB(db, sqlitedialect.New())

	repo := repository.New(bunDB)

	tsk1 := task.Task{
		UserID:    task.ID("test"),
		TaskID:    task.ID("test1"),
		Completed: true,
		Hidden:    false,
		Kind:      "daily",
	}

	err := repo.AddUserTask(context.Background(), &tsk1)
	require.NoError(t, err)

	tsk2 := task.Task{
		UserID:    task.ID("test"),
		TaskID:    task.ID("test2"),
		Completed: true,
		Hidden:    false,
		Kind:      "daily",
	}

	err = repo.AddUserTask(context.Background(), &tsk2)
	require.NoError(t, err)

	tsk3 := task.Task{
		UserID:    task.ID("test"),
		TaskID:    task.ID("test3"),
		Completed: false,
		Hidden:    false,
		Kind:      "daily",
	}

	err = repo.AddUserTask(context.Background(), &tsk3)
	require.NoError(t, err)

	err = repo.UpdateTaskForKind(context.Background(), "daily")
	require.NoError(t, err)

	got, err := repo.GetTasksForUser(context.Background(), "test")
	require.NoError(t, err)

	for _, tsk := range got {
		assert.Equal(t, string(tsk.Kind), "daily")
		assert.Equal(t, bool(tsk.Completed), false)
		assert.Equal(t, bool(tsk.Hidden), false)
	}
}

func TestUpdateTaskForKind_Error(t *testing.T) {
	db, tearDown := setupIntegrationTest(t)
	defer tearDown(t)

	bunDB := bun.NewDB(db, sqlitedialect.New())

	repo := repository.New(bunDB)

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	err := repo.UpdateTaskForKind(ctx, "daily")
	require.Error(t, err)
}

func setupIntegrationTest(t *testing.T) (*sql.DB, func(t *testing.T)) {
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
		if err := goose.Reset(db, "."); err != nil {
			t.Fatal(err)
		}
	}
}
