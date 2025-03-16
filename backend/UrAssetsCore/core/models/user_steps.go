// Code generated by SQLBoiler 4.18.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// UserStep is an object representing the database table.
type UserStep struct {
	UserID     string    `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	Step1      null.Bool `boil:"step_1" json:"step_1,omitempty" toml:"step_1" yaml:"step_1,omitempty"`
	Step2      null.Bool `boil:"step_2" json:"step_2,omitempty" toml:"step_2" yaml:"step_2,omitempty"`
	Step3      null.Bool `boil:"step_3" json:"step_3,omitempty" toml:"step_3" yaml:"step_3,omitempty"`
	IsComplete null.Bool `boil:"is_complete" json:"is_complete,omitempty" toml:"is_complete" yaml:"is_complete,omitempty"`

	R *userStepR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L userStepL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var UserStepColumns = struct {
	UserID     string
	Step1      string
	Step2      string
	Step3      string
	IsComplete string
}{
	UserID:     "user_id",
	Step1:      "step_1",
	Step2:      "step_2",
	Step3:      "step_3",
	IsComplete: "is_complete",
}

var UserStepTableColumns = struct {
	UserID     string
	Step1      string
	Step2      string
	Step3      string
	IsComplete string
}{
	UserID:     "user_steps.user_id",
	Step1:      "user_steps.step_1",
	Step2:      "user_steps.step_2",
	Step3:      "user_steps.step_3",
	IsComplete: "user_steps.is_complete",
}

// Generated where

var UserStepWhere = struct {
	UserID     whereHelperstring
	Step1      whereHelpernull_Bool
	Step2      whereHelpernull_Bool
	Step3      whereHelpernull_Bool
	IsComplete whereHelpernull_Bool
}{
	UserID:     whereHelperstring{field: "\"user_steps\".\"user_id\""},
	Step1:      whereHelpernull_Bool{field: "\"user_steps\".\"step_1\""},
	Step2:      whereHelpernull_Bool{field: "\"user_steps\".\"step_2\""},
	Step3:      whereHelpernull_Bool{field: "\"user_steps\".\"step_3\""},
	IsComplete: whereHelpernull_Bool{field: "\"user_steps\".\"is_complete\""},
}

// UserStepRels is where relationship names are stored.
var UserStepRels = struct {
	User string
}{
	User: "User",
}

// userStepR is where relationships are stored.
type userStepR struct {
	User *User `boil:"User" json:"User" toml:"User" yaml:"User"`
}

// NewStruct creates a new relationship struct
func (*userStepR) NewStruct() *userStepR {
	return &userStepR{}
}

func (r *userStepR) GetUser() *User {
	if r == nil {
		return nil
	}
	return r.User
}

// userStepL is where Load methods for each relationship are stored.
type userStepL struct{}

var (
	userStepAllColumns            = []string{"user_id", "step_1", "step_2", "step_3", "is_complete"}
	userStepColumnsWithoutDefault = []string{"user_id"}
	userStepColumnsWithDefault    = []string{"step_1", "step_2", "step_3", "is_complete"}
	userStepPrimaryKeyColumns     = []string{"user_id"}
	userStepGeneratedColumns      = []string{}
)

type (
	// UserStepSlice is an alias for a slice of pointers to UserStep.
	// This should almost always be used instead of []UserStep.
	UserStepSlice []*UserStep
	// UserStepHook is the signature for custom UserStep hook methods
	UserStepHook func(context.Context, boil.ContextExecutor, *UserStep) error

	userStepQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	userStepType                 = reflect.TypeOf(&UserStep{})
	userStepMapping              = queries.MakeStructMapping(userStepType)
	userStepPrimaryKeyMapping, _ = queries.BindMapping(userStepType, userStepMapping, userStepPrimaryKeyColumns)
	userStepInsertCacheMut       sync.RWMutex
	userStepInsertCache          = make(map[string]insertCache)
	userStepUpdateCacheMut       sync.RWMutex
	userStepUpdateCache          = make(map[string]updateCache)
	userStepUpsertCacheMut       sync.RWMutex
	userStepUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var userStepAfterSelectMu sync.Mutex
var userStepAfterSelectHooks []UserStepHook

var userStepBeforeInsertMu sync.Mutex
var userStepBeforeInsertHooks []UserStepHook
var userStepAfterInsertMu sync.Mutex
var userStepAfterInsertHooks []UserStepHook

var userStepBeforeUpdateMu sync.Mutex
var userStepBeforeUpdateHooks []UserStepHook
var userStepAfterUpdateMu sync.Mutex
var userStepAfterUpdateHooks []UserStepHook

var userStepBeforeDeleteMu sync.Mutex
var userStepBeforeDeleteHooks []UserStepHook
var userStepAfterDeleteMu sync.Mutex
var userStepAfterDeleteHooks []UserStepHook

var userStepBeforeUpsertMu sync.Mutex
var userStepBeforeUpsertHooks []UserStepHook
var userStepAfterUpsertMu sync.Mutex
var userStepAfterUpsertHooks []UserStepHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *UserStep) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userStepAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *UserStep) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userStepBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *UserStep) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userStepAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *UserStep) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userStepBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *UserStep) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userStepAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *UserStep) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userStepBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *UserStep) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userStepAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *UserStep) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userStepBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *UserStep) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userStepAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddUserStepHook registers your hook function for all future operations.
func AddUserStepHook(hookPoint boil.HookPoint, userStepHook UserStepHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		userStepAfterSelectMu.Lock()
		userStepAfterSelectHooks = append(userStepAfterSelectHooks, userStepHook)
		userStepAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		userStepBeforeInsertMu.Lock()
		userStepBeforeInsertHooks = append(userStepBeforeInsertHooks, userStepHook)
		userStepBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		userStepAfterInsertMu.Lock()
		userStepAfterInsertHooks = append(userStepAfterInsertHooks, userStepHook)
		userStepAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		userStepBeforeUpdateMu.Lock()
		userStepBeforeUpdateHooks = append(userStepBeforeUpdateHooks, userStepHook)
		userStepBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		userStepAfterUpdateMu.Lock()
		userStepAfterUpdateHooks = append(userStepAfterUpdateHooks, userStepHook)
		userStepAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		userStepBeforeDeleteMu.Lock()
		userStepBeforeDeleteHooks = append(userStepBeforeDeleteHooks, userStepHook)
		userStepBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		userStepAfterDeleteMu.Lock()
		userStepAfterDeleteHooks = append(userStepAfterDeleteHooks, userStepHook)
		userStepAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		userStepBeforeUpsertMu.Lock()
		userStepBeforeUpsertHooks = append(userStepBeforeUpsertHooks, userStepHook)
		userStepBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		userStepAfterUpsertMu.Lock()
		userStepAfterUpsertHooks = append(userStepAfterUpsertHooks, userStepHook)
		userStepAfterUpsertMu.Unlock()
	}
}

// One returns a single userStep record from the query.
func (q userStepQuery) One(ctx context.Context, exec boil.ContextExecutor) (*UserStep, error) {
	o := &UserStep{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for user_steps")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all UserStep records from the query.
func (q userStepQuery) All(ctx context.Context, exec boil.ContextExecutor) (UserStepSlice, error) {
	var o []*UserStep

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to UserStep slice")
	}

	if len(userStepAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all UserStep records in the query.
func (q userStepQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count user_steps rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q userStepQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if user_steps exists")
	}

	return count > 0, nil
}

// User pointed to by the foreign key.
func (o *UserStep) User(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	return Users(queryMods...)
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (userStepL) LoadUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybeUserStep interface{}, mods queries.Applicator) error {
	var slice []*UserStep
	var object *UserStep

	if singular {
		var ok bool
		object, ok = maybeUserStep.(*UserStep)
		if !ok {
			object = new(UserStep)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeUserStep)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeUserStep))
			}
		}
	} else {
		s, ok := maybeUserStep.(*[]*UserStep)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeUserStep)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeUserStep))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &userStepR{}
		}
		args[object.UserID] = struct{}{}

	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &userStepR{}
			}

			args[obj.UserID] = struct{}{}

		}
	}

	if len(args) == 0 {
		return nil
	}

	argsSlice := make([]interface{}, len(args))
	i := 0
	for arg := range args {
		argsSlice[i] = arg
		i++
	}

	query := NewQuery(
		qm.From(`users`),
		qm.WhereIn(`users.id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load User")
	}

	var resultSlice []*User
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice User")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for users")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for users")
	}

	if len(userAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.User = foreign
		if foreign.R == nil {
			foreign.R = &userR{}
		}
		foreign.R.UserStep = object
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserID == foreign.ID {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.UserStep = local
				break
			}
		}
	}

	return nil
}

// SetUser of the userStep to the related item.
// Sets o.R.User to related.
// Adds o to related.R.UserStep.
func (o *UserStep) SetUser(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"user_steps\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"user_id"}),
		strmangle.WhereClause("\"", "\"", 2, userStepPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.UserID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.UserID = related.ID
	if o.R == nil {
		o.R = &userStepR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userR{
			UserStep: o,
		}
	} else {
		related.R.UserStep = o
	}

	return nil
}

// UserSteps retrieves all the records using an executor.
func UserSteps(mods ...qm.QueryMod) userStepQuery {
	mods = append(mods, qm.From("\"user_steps\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"user_steps\".*"})
	}

	return userStepQuery{q}
}

// FindUserStep retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindUserStep(ctx context.Context, exec boil.ContextExecutor, userID string, selectCols ...string) (*UserStep, error) {
	userStepObj := &UserStep{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"user_steps\" where \"user_id\"=$1", sel,
	)

	q := queries.Raw(query, userID)

	err := q.Bind(ctx, exec, userStepObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from user_steps")
	}

	if err = userStepObj.doAfterSelectHooks(ctx, exec); err != nil {
		return userStepObj, err
	}

	return userStepObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *UserStep) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no user_steps provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(userStepColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	userStepInsertCacheMut.RLock()
	cache, cached := userStepInsertCache[key]
	userStepInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			userStepAllColumns,
			userStepColumnsWithDefault,
			userStepColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(userStepType, userStepMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(userStepType, userStepMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"user_steps\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"user_steps\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into user_steps")
	}

	if !cached {
		userStepInsertCacheMut.Lock()
		userStepInsertCache[key] = cache
		userStepInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the UserStep.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *UserStep) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	userStepUpdateCacheMut.RLock()
	cache, cached := userStepUpdateCache[key]
	userStepUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			userStepAllColumns,
			userStepPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update user_steps, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"user_steps\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, userStepPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(userStepType, userStepMapping, append(wl, userStepPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update user_steps row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for user_steps")
	}

	if !cached {
		userStepUpdateCacheMut.Lock()
		userStepUpdateCache[key] = cache
		userStepUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q userStepQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for user_steps")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for user_steps")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o UserStepSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userStepPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"user_steps\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, userStepPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in userStep slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all userStep")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *UserStep) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("models: no user_steps provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(userStepColumnsWithDefault, o)

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

	userStepUpsertCacheMut.RLock()
	cache, cached := userStepUpsertCache[key]
	userStepUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			userStepAllColumns,
			userStepColumnsWithDefault,
			userStepColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			userStepAllColumns,
			userStepPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert user_steps, could not build update column list")
		}

		ret := strmangle.SetComplement(userStepAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(userStepPrimaryKeyColumns) == 0 {
				return errors.New("models: unable to upsert user_steps, could not build conflict column list")
			}

			conflict = make([]string, len(userStepPrimaryKeyColumns))
			copy(conflict, userStepPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"user_steps\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(userStepType, userStepMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(userStepType, userStepMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert user_steps")
	}

	if !cached {
		userStepUpsertCacheMut.Lock()
		userStepUpsertCache[key] = cache
		userStepUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single UserStep record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *UserStep) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no UserStep provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), userStepPrimaryKeyMapping)
	sql := "DELETE FROM \"user_steps\" WHERE \"user_id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from user_steps")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for user_steps")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q userStepQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no userStepQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from user_steps")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for user_steps")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o UserStepSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(userStepBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userStepPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"user_steps\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, userStepPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from userStep slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for user_steps")
	}

	if len(userStepAfterDeleteHooks) != 0 {
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
func (o *UserStep) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindUserStep(ctx, exec, o.UserID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *UserStepSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := UserStepSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userStepPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"user_steps\".* FROM \"user_steps\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, userStepPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in UserStepSlice")
	}

	*o = slice

	return nil
}

// UserStepExists checks if the UserStep row exists.
func UserStepExists(ctx context.Context, exec boil.ContextExecutor, userID string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"user_steps\" where \"user_id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, userID)
	}
	row := exec.QueryRowContext(ctx, sql, userID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if user_steps exists")
	}

	return exists, nil
}

// Exists checks if the UserStep row exists.
func (o *UserStep) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return UserStepExists(ctx, exec, o.UserID)
}
