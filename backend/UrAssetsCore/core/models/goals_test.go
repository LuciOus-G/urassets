// Code generated by SQLBoiler 4.18.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testGoals(t *testing.T) {
	t.Parallel()

	query := Goals()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testGoalsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Goal{}
	if err = randomize.Struct(seed, o, goalDBTypes, true, goalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Goal struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Goals().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testGoalsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Goal{}
	if err = randomize.Struct(seed, o, goalDBTypes, true, goalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Goal struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Goals().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Goals().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testGoalsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Goal{}
	if err = randomize.Struct(seed, o, goalDBTypes, true, goalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Goal struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := GoalSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Goals().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testGoalsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Goal{}
	if err = randomize.Struct(seed, o, goalDBTypes, true, goalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Goal struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := GoalExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Goal exists: %s", err)
	}
	if !e {
		t.Errorf("Expected GoalExists to return true, but got false.")
	}
}

func testGoalsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Goal{}
	if err = randomize.Struct(seed, o, goalDBTypes, true, goalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Goal struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	goalFound, err := FindGoal(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if goalFound == nil {
		t.Error("want a record, got nil")
	}
}

func testGoalsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Goal{}
	if err = randomize.Struct(seed, o, goalDBTypes, true, goalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Goal struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Goals().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testGoalsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Goal{}
	if err = randomize.Struct(seed, o, goalDBTypes, true, goalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Goal struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Goals().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testGoalsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	goalOne := &Goal{}
	goalTwo := &Goal{}
	if err = randomize.Struct(seed, goalOne, goalDBTypes, false, goalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Goal struct: %s", err)
	}
	if err = randomize.Struct(seed, goalTwo, goalDBTypes, false, goalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Goal struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = goalOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = goalTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Goals().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testGoalsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	goalOne := &Goal{}
	goalTwo := &Goal{}
	if err = randomize.Struct(seed, goalOne, goalDBTypes, false, goalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Goal struct: %s", err)
	}
	if err = randomize.Struct(seed, goalTwo, goalDBTypes, false, goalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Goal struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = goalOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = goalTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Goals().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func goalBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Goal) error {
	*o = Goal{}
	return nil
}

func goalAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Goal) error {
	*o = Goal{}
	return nil
}

func goalAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Goal) error {
	*o = Goal{}
	return nil
}

func goalBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Goal) error {
	*o = Goal{}
	return nil
}

func goalAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Goal) error {
	*o = Goal{}
	return nil
}

func goalBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Goal) error {
	*o = Goal{}
	return nil
}

func goalAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Goal) error {
	*o = Goal{}
	return nil
}

func goalBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Goal) error {
	*o = Goal{}
	return nil
}

func goalAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Goal) error {
	*o = Goal{}
	return nil
}

func testGoalsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Goal{}
	o := &Goal{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, goalDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Goal object: %s", err)
	}

	AddGoalHook(boil.BeforeInsertHook, goalBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	goalBeforeInsertHooks = []GoalHook{}

	AddGoalHook(boil.AfterInsertHook, goalAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	goalAfterInsertHooks = []GoalHook{}

	AddGoalHook(boil.AfterSelectHook, goalAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	goalAfterSelectHooks = []GoalHook{}

	AddGoalHook(boil.BeforeUpdateHook, goalBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	goalBeforeUpdateHooks = []GoalHook{}

	AddGoalHook(boil.AfterUpdateHook, goalAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	goalAfterUpdateHooks = []GoalHook{}

	AddGoalHook(boil.BeforeDeleteHook, goalBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	goalBeforeDeleteHooks = []GoalHook{}

	AddGoalHook(boil.AfterDeleteHook, goalAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	goalAfterDeleteHooks = []GoalHook{}

	AddGoalHook(boil.BeforeUpsertHook, goalBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	goalBeforeUpsertHooks = []GoalHook{}

	AddGoalHook(boil.AfterUpsertHook, goalAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	goalAfterUpsertHooks = []GoalHook{}
}

func testGoalsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Goal{}
	if err = randomize.Struct(seed, o, goalDBTypes, true, goalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Goal struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Goals().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testGoalsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Goal{}
	if err = randomize.Struct(seed, o, goalDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Goal struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(goalColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Goals().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testGoalToOneUserUsingUser(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Goal
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, goalDBTypes, false, goalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Goal struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.UserID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.User().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	ranAfterSelectHook := false
	AddUserHook(boil.AfterSelectHook, func(ctx context.Context, e boil.ContextExecutor, o *User) error {
		ranAfterSelectHook = true
		return nil
	})

	slice := GoalSlice{&local}
	if err = local.L.LoadUser(ctx, tx, false, (*[]*Goal)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.User = nil
	if err = local.L.LoadUser(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}

	if !ranAfterSelectHook {
		t.Error("failed to run AfterSelect hook for relationship")
	}
}

func testGoalToOneSetOpUserUsingUser(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Goal
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, goalDBTypes, false, strmangle.SetComplement(goalPrimaryKeyColumns, goalColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*User{&b, &c} {
		err = a.SetUser(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.User != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Goals[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.UserID != x.ID {
			t.Error("foreign key was wrong value", a.UserID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.UserID))
		reflect.Indirect(reflect.ValueOf(&a.UserID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.UserID != x.ID {
			t.Error("foreign key was wrong value", a.UserID, x.ID)
		}
	}
}

func testGoalsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Goal{}
	if err = randomize.Struct(seed, o, goalDBTypes, true, goalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Goal struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testGoalsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Goal{}
	if err = randomize.Struct(seed, o, goalDBTypes, true, goalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Goal struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := GoalSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testGoalsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Goal{}
	if err = randomize.Struct(seed, o, goalDBTypes, true, goalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Goal struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Goals().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	goalDBTypes = map[string]string{`ID`: `uuid`, `UserID`: `uuid`, `GoalName`: `character varying`, `GoalMotivation`: `character varying`, `GoalDescriptions`: `text`, `GoalAmount`: `numeric`, `Priority`: `integer`}
	_           = bytes.MinRead
)

func testGoalsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(goalPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(goalAllColumns) == len(goalPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Goal{}
	if err = randomize.Struct(seed, o, goalDBTypes, true, goalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Goal struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Goals().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, goalDBTypes, true, goalPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Goal struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testGoalsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(goalAllColumns) == len(goalPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Goal{}
	if err = randomize.Struct(seed, o, goalDBTypes, true, goalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Goal struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Goals().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, goalDBTypes, true, goalPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Goal struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(goalAllColumns, goalPrimaryKeyColumns) {
		fields = goalAllColumns
	} else {
		fields = strmangle.SetComplement(
			goalAllColumns,
			goalPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := GoalSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testGoalsUpsert(t *testing.T) {
	t.Parallel()

	if len(goalAllColumns) == len(goalPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Goal{}
	if err = randomize.Struct(seed, &o, goalDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Goal struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Goal: %s", err)
	}

	count, err := Goals().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, goalDBTypes, false, goalPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Goal struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Goal: %s", err)
	}

	count, err = Goals().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
