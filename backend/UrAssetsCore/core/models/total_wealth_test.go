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

func testTotalWealths(t *testing.T) {
	t.Parallel()

	query := TotalWealths()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testTotalWealthsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TotalWealth{}
	if err = randomize.Struct(seed, o, totalWealthDBTypes, true, totalWealthColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TotalWealth struct: %s", err)
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

	count, err := TotalWealths().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTotalWealthsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TotalWealth{}
	if err = randomize.Struct(seed, o, totalWealthDBTypes, true, totalWealthColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TotalWealth struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := TotalWealths().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := TotalWealths().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTotalWealthsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TotalWealth{}
	if err = randomize.Struct(seed, o, totalWealthDBTypes, true, totalWealthColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TotalWealth struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := TotalWealthSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := TotalWealths().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTotalWealthsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TotalWealth{}
	if err = randomize.Struct(seed, o, totalWealthDBTypes, true, totalWealthColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TotalWealth struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := TotalWealthExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if TotalWealth exists: %s", err)
	}
	if !e {
		t.Errorf("Expected TotalWealthExists to return true, but got false.")
	}
}

func testTotalWealthsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TotalWealth{}
	if err = randomize.Struct(seed, o, totalWealthDBTypes, true, totalWealthColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TotalWealth struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	totalWealthFound, err := FindTotalWealth(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if totalWealthFound == nil {
		t.Error("want a record, got nil")
	}
}

func testTotalWealthsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TotalWealth{}
	if err = randomize.Struct(seed, o, totalWealthDBTypes, true, totalWealthColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TotalWealth struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = TotalWealths().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testTotalWealthsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TotalWealth{}
	if err = randomize.Struct(seed, o, totalWealthDBTypes, true, totalWealthColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TotalWealth struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := TotalWealths().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testTotalWealthsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	totalWealthOne := &TotalWealth{}
	totalWealthTwo := &TotalWealth{}
	if err = randomize.Struct(seed, totalWealthOne, totalWealthDBTypes, false, totalWealthColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TotalWealth struct: %s", err)
	}
	if err = randomize.Struct(seed, totalWealthTwo, totalWealthDBTypes, false, totalWealthColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TotalWealth struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = totalWealthOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = totalWealthTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := TotalWealths().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testTotalWealthsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	totalWealthOne := &TotalWealth{}
	totalWealthTwo := &TotalWealth{}
	if err = randomize.Struct(seed, totalWealthOne, totalWealthDBTypes, false, totalWealthColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TotalWealth struct: %s", err)
	}
	if err = randomize.Struct(seed, totalWealthTwo, totalWealthDBTypes, false, totalWealthColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TotalWealth struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = totalWealthOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = totalWealthTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := TotalWealths().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func totalWealthBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *TotalWealth) error {
	*o = TotalWealth{}
	return nil
}

func totalWealthAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *TotalWealth) error {
	*o = TotalWealth{}
	return nil
}

func totalWealthAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *TotalWealth) error {
	*o = TotalWealth{}
	return nil
}

func totalWealthBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *TotalWealth) error {
	*o = TotalWealth{}
	return nil
}

func totalWealthAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *TotalWealth) error {
	*o = TotalWealth{}
	return nil
}

func totalWealthBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *TotalWealth) error {
	*o = TotalWealth{}
	return nil
}

func totalWealthAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *TotalWealth) error {
	*o = TotalWealth{}
	return nil
}

func totalWealthBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *TotalWealth) error {
	*o = TotalWealth{}
	return nil
}

func totalWealthAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *TotalWealth) error {
	*o = TotalWealth{}
	return nil
}

func testTotalWealthsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &TotalWealth{}
	o := &TotalWealth{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, totalWealthDBTypes, false); err != nil {
		t.Errorf("Unable to randomize TotalWealth object: %s", err)
	}

	AddTotalWealthHook(boil.BeforeInsertHook, totalWealthBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	totalWealthBeforeInsertHooks = []TotalWealthHook{}

	AddTotalWealthHook(boil.AfterInsertHook, totalWealthAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	totalWealthAfterInsertHooks = []TotalWealthHook{}

	AddTotalWealthHook(boil.AfterSelectHook, totalWealthAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	totalWealthAfterSelectHooks = []TotalWealthHook{}

	AddTotalWealthHook(boil.BeforeUpdateHook, totalWealthBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	totalWealthBeforeUpdateHooks = []TotalWealthHook{}

	AddTotalWealthHook(boil.AfterUpdateHook, totalWealthAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	totalWealthAfterUpdateHooks = []TotalWealthHook{}

	AddTotalWealthHook(boil.BeforeDeleteHook, totalWealthBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	totalWealthBeforeDeleteHooks = []TotalWealthHook{}

	AddTotalWealthHook(boil.AfterDeleteHook, totalWealthAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	totalWealthAfterDeleteHooks = []TotalWealthHook{}

	AddTotalWealthHook(boil.BeforeUpsertHook, totalWealthBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	totalWealthBeforeUpsertHooks = []TotalWealthHook{}

	AddTotalWealthHook(boil.AfterUpsertHook, totalWealthAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	totalWealthAfterUpsertHooks = []TotalWealthHook{}
}

func testTotalWealthsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TotalWealth{}
	if err = randomize.Struct(seed, o, totalWealthDBTypes, true, totalWealthColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TotalWealth struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := TotalWealths().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTotalWealthsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TotalWealth{}
	if err = randomize.Struct(seed, o, totalWealthDBTypes, true); err != nil {
		t.Errorf("Unable to randomize TotalWealth struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(totalWealthColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := TotalWealths().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTotalWealthToOneUserUsingUser(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local TotalWealth
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, totalWealthDBTypes, false, totalWealthColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TotalWealth struct: %s", err)
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

	slice := TotalWealthSlice{&local}
	if err = local.L.LoadUser(ctx, tx, false, (*[]*TotalWealth)(&slice), nil); err != nil {
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

func testTotalWealthToOneSetOpUserUsingUser(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a TotalWealth
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, totalWealthDBTypes, false, strmangle.SetComplement(totalWealthPrimaryKeyColumns, totalWealthColumnsWithoutDefault)...); err != nil {
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

		if x.R.TotalWealths[0] != &a {
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

func testTotalWealthsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TotalWealth{}
	if err = randomize.Struct(seed, o, totalWealthDBTypes, true, totalWealthColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TotalWealth struct: %s", err)
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

func testTotalWealthsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TotalWealth{}
	if err = randomize.Struct(seed, o, totalWealthDBTypes, true, totalWealthColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TotalWealth struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := TotalWealthSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testTotalWealthsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TotalWealth{}
	if err = randomize.Struct(seed, o, totalWealthDBTypes, true, totalWealthColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TotalWealth struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := TotalWealths().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	totalWealthDBTypes = map[string]string{`ID`: `uuid`, `UserID`: `uuid`, `Balance`: `numeric`}
	_                  = bytes.MinRead
)

func testTotalWealthsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(totalWealthPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(totalWealthAllColumns) == len(totalWealthPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &TotalWealth{}
	if err = randomize.Struct(seed, o, totalWealthDBTypes, true, totalWealthColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TotalWealth struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := TotalWealths().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, totalWealthDBTypes, true, totalWealthPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize TotalWealth struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testTotalWealthsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(totalWealthAllColumns) == len(totalWealthPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &TotalWealth{}
	if err = randomize.Struct(seed, o, totalWealthDBTypes, true, totalWealthColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TotalWealth struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := TotalWealths().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, totalWealthDBTypes, true, totalWealthPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize TotalWealth struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(totalWealthAllColumns, totalWealthPrimaryKeyColumns) {
		fields = totalWealthAllColumns
	} else {
		fields = strmangle.SetComplement(
			totalWealthAllColumns,
			totalWealthPrimaryKeyColumns,
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

	slice := TotalWealthSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testTotalWealthsUpsert(t *testing.T) {
	t.Parallel()

	if len(totalWealthAllColumns) == len(totalWealthPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := TotalWealth{}
	if err = randomize.Struct(seed, &o, totalWealthDBTypes, true); err != nil {
		t.Errorf("Unable to randomize TotalWealth struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert TotalWealth: %s", err)
	}

	count, err := TotalWealths().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, totalWealthDBTypes, false, totalWealthPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize TotalWealth struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert TotalWealth: %s", err)
	}

	count, err = TotalWealths().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
