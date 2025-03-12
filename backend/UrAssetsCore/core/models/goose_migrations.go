// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// GooseMigration is an object representing the database table.
type GooseMigration struct {
	ID        int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	VersionID int64     `boil:"version_id" json:"version_id" toml:"version_id" yaml:"version_id"`
	IsApplied bool      `boil:"is_applied" json:"is_applied" toml:"is_applied" yaml:"is_applied"`
	Tstamp    time.Time `boil:"tstamp" json:"tstamp" toml:"tstamp" yaml:"tstamp"`

	R *gooseMigrationR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L gooseMigrationL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var GooseMigrationColumns = struct {
	ID        string
	VersionID string
	IsApplied string
	Tstamp    string
}{
	ID:        "id",
	VersionID: "version_id",
	IsApplied: "is_applied",
	Tstamp:    "tstamp",
}

var GooseMigrationTableColumns = struct {
	ID        string
	VersionID string
	IsApplied string
	Tstamp    string
}{
	ID:        "goose_migrations.id",
	VersionID: "goose_migrations.version_id",
	IsApplied: "goose_migrations.is_applied",
	Tstamp:    "goose_migrations.tstamp",
}

// Generated where

type whereHelperint struct{ field string }

func (w whereHelperint) EQ(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint) NEQ(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint) LT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint) LTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint) GT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint) GTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint) IN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint) NIN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperint64 struct{ field string }

func (w whereHelperint64) EQ(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint64) NEQ(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint64) LT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint64) LTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint64) GT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint64) GTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint64) IN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint64) NIN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperbool struct{ field string }

func (w whereHelperbool) EQ(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperbool) NEQ(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperbool) LT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperbool) LTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperbool) GT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperbool) GTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var GooseMigrationWhere = struct {
	ID        whereHelperint
	VersionID whereHelperint64
	IsApplied whereHelperbool
	Tstamp    whereHelpertime_Time
}{
	ID:        whereHelperint{field: "\"goose_migrations\".\"id\""},
	VersionID: whereHelperint64{field: "\"goose_migrations\".\"version_id\""},
	IsApplied: whereHelperbool{field: "\"goose_migrations\".\"is_applied\""},
	Tstamp:    whereHelpertime_Time{field: "\"goose_migrations\".\"tstamp\""},
}

// GooseMigrationRels is where relationship names are stored.
var GooseMigrationRels = struct {
}{}

// gooseMigrationR is where relationships are stored.
type gooseMigrationR struct {
}

// NewStruct creates a new relationship struct
func (*gooseMigrationR) NewStruct() *gooseMigrationR {
	return &gooseMigrationR{}
}

// gooseMigrationL is where Load methods for each relationship are stored.
type gooseMigrationL struct{}

var (
	gooseMigrationAllColumns            = []string{"id", "version_id", "is_applied", "tstamp"}
	gooseMigrationColumnsWithoutDefault = []string{"version_id", "is_applied"}
	gooseMigrationColumnsWithDefault    = []string{"id", "tstamp"}
	gooseMigrationPrimaryKeyColumns     = []string{"id"}
	gooseMigrationGeneratedColumns      = []string{}
)

type (
	// GooseMigrationSlice is an alias for a slice of pointers to GooseMigration.
	// This should almost always be used instead of []GooseMigration.
	GooseMigrationSlice []*GooseMigration
	// GooseMigrationHook is the signature for custom GooseMigration hook methods
	GooseMigrationHook func(context.Context, boil.ContextExecutor, *GooseMigration) error

	gooseMigrationQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	gooseMigrationType                 = reflect.TypeOf(&GooseMigration{})
	gooseMigrationMapping              = queries.MakeStructMapping(gooseMigrationType)
	gooseMigrationPrimaryKeyMapping, _ = queries.BindMapping(gooseMigrationType, gooseMigrationMapping, gooseMigrationPrimaryKeyColumns)
	gooseMigrationInsertCacheMut       sync.RWMutex
	gooseMigrationInsertCache          = make(map[string]insertCache)
	gooseMigrationUpdateCacheMut       sync.RWMutex
	gooseMigrationUpdateCache          = make(map[string]updateCache)
	gooseMigrationUpsertCacheMut       sync.RWMutex
	gooseMigrationUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var gooseMigrationAfterSelectMu sync.Mutex
var gooseMigrationAfterSelectHooks []GooseMigrationHook

var gooseMigrationBeforeInsertMu sync.Mutex
var gooseMigrationBeforeInsertHooks []GooseMigrationHook
var gooseMigrationAfterInsertMu sync.Mutex
var gooseMigrationAfterInsertHooks []GooseMigrationHook

var gooseMigrationBeforeUpdateMu sync.Mutex
var gooseMigrationBeforeUpdateHooks []GooseMigrationHook
var gooseMigrationAfterUpdateMu sync.Mutex
var gooseMigrationAfterUpdateHooks []GooseMigrationHook

var gooseMigrationBeforeDeleteMu sync.Mutex
var gooseMigrationBeforeDeleteHooks []GooseMigrationHook
var gooseMigrationAfterDeleteMu sync.Mutex
var gooseMigrationAfterDeleteHooks []GooseMigrationHook

var gooseMigrationBeforeUpsertMu sync.Mutex
var gooseMigrationBeforeUpsertHooks []GooseMigrationHook
var gooseMigrationAfterUpsertMu sync.Mutex
var gooseMigrationAfterUpsertHooks []GooseMigrationHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *GooseMigration) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gooseMigrationAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *GooseMigration) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gooseMigrationBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *GooseMigration) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gooseMigrationAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *GooseMigration) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gooseMigrationBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *GooseMigration) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gooseMigrationAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *GooseMigration) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gooseMigrationBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *GooseMigration) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gooseMigrationAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *GooseMigration) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gooseMigrationBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *GooseMigration) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gooseMigrationAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddGooseMigrationHook registers your hook function for all future operations.
func AddGooseMigrationHook(hookPoint boil.HookPoint, gooseMigrationHook GooseMigrationHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		gooseMigrationAfterSelectMu.Lock()
		gooseMigrationAfterSelectHooks = append(gooseMigrationAfterSelectHooks, gooseMigrationHook)
		gooseMigrationAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		gooseMigrationBeforeInsertMu.Lock()
		gooseMigrationBeforeInsertHooks = append(gooseMigrationBeforeInsertHooks, gooseMigrationHook)
		gooseMigrationBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		gooseMigrationAfterInsertMu.Lock()
		gooseMigrationAfterInsertHooks = append(gooseMigrationAfterInsertHooks, gooseMigrationHook)
		gooseMigrationAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		gooseMigrationBeforeUpdateMu.Lock()
		gooseMigrationBeforeUpdateHooks = append(gooseMigrationBeforeUpdateHooks, gooseMigrationHook)
		gooseMigrationBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		gooseMigrationAfterUpdateMu.Lock()
		gooseMigrationAfterUpdateHooks = append(gooseMigrationAfterUpdateHooks, gooseMigrationHook)
		gooseMigrationAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		gooseMigrationBeforeDeleteMu.Lock()
		gooseMigrationBeforeDeleteHooks = append(gooseMigrationBeforeDeleteHooks, gooseMigrationHook)
		gooseMigrationBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		gooseMigrationAfterDeleteMu.Lock()
		gooseMigrationAfterDeleteHooks = append(gooseMigrationAfterDeleteHooks, gooseMigrationHook)
		gooseMigrationAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		gooseMigrationBeforeUpsertMu.Lock()
		gooseMigrationBeforeUpsertHooks = append(gooseMigrationBeforeUpsertHooks, gooseMigrationHook)
		gooseMigrationBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		gooseMigrationAfterUpsertMu.Lock()
		gooseMigrationAfterUpsertHooks = append(gooseMigrationAfterUpsertHooks, gooseMigrationHook)
		gooseMigrationAfterUpsertMu.Unlock()
	}
}

// One returns a single gooseMigration record from the query.
func (q gooseMigrationQuery) One(ctx context.Context, exec boil.ContextExecutor) (*GooseMigration, error) {
	o := &GooseMigration{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for goose_migrations")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all GooseMigration records from the query.
func (q gooseMigrationQuery) All(ctx context.Context, exec boil.ContextExecutor) (GooseMigrationSlice, error) {
	var o []*GooseMigration

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to GooseMigration slice")
	}

	if len(gooseMigrationAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all GooseMigration records in the query.
func (q gooseMigrationQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count goose_migrations rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q gooseMigrationQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if goose_migrations exists")
	}

	return count > 0, nil
}

// GooseMigrations retrieves all the records using an executor.
func GooseMigrations(mods ...qm.QueryMod) gooseMigrationQuery {
	mods = append(mods, qm.From("\"goose_migrations\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"goose_migrations\".*"})
	}

	return gooseMigrationQuery{q}
}

// FindGooseMigration retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindGooseMigration(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*GooseMigration, error) {
	gooseMigrationObj := &GooseMigration{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"goose_migrations\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, gooseMigrationObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from goose_migrations")
	}

	if err = gooseMigrationObj.doAfterSelectHooks(ctx, exec); err != nil {
		return gooseMigrationObj, err
	}

	return gooseMigrationObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *GooseMigration) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no goose_migrations provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(gooseMigrationColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	gooseMigrationInsertCacheMut.RLock()
	cache, cached := gooseMigrationInsertCache[key]
	gooseMigrationInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			gooseMigrationAllColumns,
			gooseMigrationColumnsWithDefault,
			gooseMigrationColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(gooseMigrationType, gooseMigrationMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(gooseMigrationType, gooseMigrationMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"goose_migrations\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"goose_migrations\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into goose_migrations")
	}

	if !cached {
		gooseMigrationInsertCacheMut.Lock()
		gooseMigrationInsertCache[key] = cache
		gooseMigrationInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the GooseMigration.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *GooseMigration) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	gooseMigrationUpdateCacheMut.RLock()
	cache, cached := gooseMigrationUpdateCache[key]
	gooseMigrationUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			gooseMigrationAllColumns,
			gooseMigrationPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update goose_migrations, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"goose_migrations\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, gooseMigrationPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(gooseMigrationType, gooseMigrationMapping, append(wl, gooseMigrationPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update goose_migrations row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for goose_migrations")
	}

	if !cached {
		gooseMigrationUpdateCacheMut.Lock()
		gooseMigrationUpdateCache[key] = cache
		gooseMigrationUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q gooseMigrationQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for goose_migrations")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for goose_migrations")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o GooseMigrationSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), gooseMigrationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"goose_migrations\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, gooseMigrationPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in gooseMigration slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all gooseMigration")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *GooseMigration) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no goose_migrations provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(gooseMigrationColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	gooseMigrationUpsertCacheMut.RLock()
	cache, cached := gooseMigrationUpsertCache[key]
	gooseMigrationUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			gooseMigrationAllColumns,
			gooseMigrationColumnsWithDefault,
			gooseMigrationColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			gooseMigrationAllColumns,
			gooseMigrationPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert goose_migrations, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(gooseMigrationPrimaryKeyColumns))
			copy(conflict, gooseMigrationPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"goose_migrations\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(gooseMigrationType, gooseMigrationMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(gooseMigrationType, gooseMigrationMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert goose_migrations")
	}

	if !cached {
		gooseMigrationUpsertCacheMut.Lock()
		gooseMigrationUpsertCache[key] = cache
		gooseMigrationUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single GooseMigration record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *GooseMigration) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no GooseMigration provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), gooseMigrationPrimaryKeyMapping)
	sql := "DELETE FROM \"goose_migrations\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from goose_migrations")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for goose_migrations")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q gooseMigrationQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no gooseMigrationQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from goose_migrations")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for goose_migrations")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o GooseMigrationSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(gooseMigrationBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), gooseMigrationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"goose_migrations\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, gooseMigrationPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from gooseMigration slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for goose_migrations")
	}

	if len(gooseMigrationAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *GooseMigration) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindGooseMigration(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *GooseMigrationSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := GooseMigrationSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), gooseMigrationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"goose_migrations\".* FROM \"goose_migrations\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, gooseMigrationPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in GooseMigrationSlice")
	}

	*o = slice

	return nil
}

// GooseMigrationExists checks if the GooseMigration row exists.
func GooseMigrationExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"goose_migrations\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if goose_migrations exists")
	}

	return exists, nil
}

// Exists checks if the GooseMigration row exists.
func (o *GooseMigration) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return GooseMigrationExists(ctx, exec, o.ID)
}
