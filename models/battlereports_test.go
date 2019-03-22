// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testBattlereports(t *testing.T) {
	t.Parallel()

	query := Battlereports()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testBattlereportsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Battlereport{}
	if err = randomize.Struct(seed, o, battlereportDBTypes, true, battlereportColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Battlereport struct: %s", err)
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

	count, err := Battlereports().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBattlereportsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Battlereport{}
	if err = randomize.Struct(seed, o, battlereportDBTypes, true, battlereportColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Battlereport struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Battlereports().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Battlereports().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBattlereportsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Battlereport{}
	if err = randomize.Struct(seed, o, battlereportDBTypes, true, battlereportColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Battlereport struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BattlereportSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Battlereports().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBattlereportsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Battlereport{}
	if err = randomize.Struct(seed, o, battlereportDBTypes, true, battlereportColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Battlereport struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := BattlereportExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Battlereport exists: %s", err)
	}
	if !e {
		t.Errorf("Expected BattlereportExists to return true, but got false.")
	}
}

func testBattlereportsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Battlereport{}
	if err = randomize.Struct(seed, o, battlereportDBTypes, true, battlereportColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Battlereport struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	battlereportFound, err := FindBattlereport(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if battlereportFound == nil {
		t.Error("want a record, got nil")
	}
}

func testBattlereportsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Battlereport{}
	if err = randomize.Struct(seed, o, battlereportDBTypes, true, battlereportColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Battlereport struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Battlereports().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testBattlereportsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Battlereport{}
	if err = randomize.Struct(seed, o, battlereportDBTypes, true, battlereportColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Battlereport struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Battlereports().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testBattlereportsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	battlereportOne := &Battlereport{}
	battlereportTwo := &Battlereport{}
	if err = randomize.Struct(seed, battlereportOne, battlereportDBTypes, false, battlereportColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Battlereport struct: %s", err)
	}
	if err = randomize.Struct(seed, battlereportTwo, battlereportDBTypes, false, battlereportColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Battlereport struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = battlereportOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = battlereportTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Battlereports().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testBattlereportsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	battlereportOne := &Battlereport{}
	battlereportTwo := &Battlereport{}
	if err = randomize.Struct(seed, battlereportOne, battlereportDBTypes, false, battlereportColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Battlereport struct: %s", err)
	}
	if err = randomize.Struct(seed, battlereportTwo, battlereportDBTypes, false, battlereportColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Battlereport struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = battlereportOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = battlereportTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Battlereports().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func battlereportBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Battlereport) error {
	*o = Battlereport{}
	return nil
}

func battlereportAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Battlereport) error {
	*o = Battlereport{}
	return nil
}

func battlereportAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Battlereport) error {
	*o = Battlereport{}
	return nil
}

func battlereportBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Battlereport) error {
	*o = Battlereport{}
	return nil
}

func battlereportAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Battlereport) error {
	*o = Battlereport{}
	return nil
}

func battlereportBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Battlereport) error {
	*o = Battlereport{}
	return nil
}

func battlereportAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Battlereport) error {
	*o = Battlereport{}
	return nil
}

func battlereportBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Battlereport) error {
	*o = Battlereport{}
	return nil
}

func battlereportAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Battlereport) error {
	*o = Battlereport{}
	return nil
}

func testBattlereportsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Battlereport{}
	o := &Battlereport{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, battlereportDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Battlereport object: %s", err)
	}

	AddBattlereportHook(boil.BeforeInsertHook, battlereportBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	battlereportBeforeInsertHooks = []BattlereportHook{}

	AddBattlereportHook(boil.AfterInsertHook, battlereportAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	battlereportAfterInsertHooks = []BattlereportHook{}

	AddBattlereportHook(boil.AfterSelectHook, battlereportAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	battlereportAfterSelectHooks = []BattlereportHook{}

	AddBattlereportHook(boil.BeforeUpdateHook, battlereportBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	battlereportBeforeUpdateHooks = []BattlereportHook{}

	AddBattlereportHook(boil.AfterUpdateHook, battlereportAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	battlereportAfterUpdateHooks = []BattlereportHook{}

	AddBattlereportHook(boil.BeforeDeleteHook, battlereportBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	battlereportBeforeDeleteHooks = []BattlereportHook{}

	AddBattlereportHook(boil.AfterDeleteHook, battlereportAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	battlereportAfterDeleteHooks = []BattlereportHook{}

	AddBattlereportHook(boil.BeforeUpsertHook, battlereportBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	battlereportBeforeUpsertHooks = []BattlereportHook{}

	AddBattlereportHook(boil.AfterUpsertHook, battlereportAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	battlereportAfterUpsertHooks = []BattlereportHook{}
}

func testBattlereportsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Battlereport{}
	if err = randomize.Struct(seed, o, battlereportDBTypes, true, battlereportColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Battlereport struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Battlereports().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBattlereportsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Battlereport{}
	if err = randomize.Struct(seed, o, battlereportDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Battlereport struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(battlereportColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Battlereports().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBattlereportToOneArmyListUsingList(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Battlereport
	var foreign ArmyList

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, battlereportDBTypes, true, battlereportColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Battlereport struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, armyListDBTypes, false, armyListColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ArmyList struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.ListID, foreign.ID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.List().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := BattlereportSlice{&local}
	if err = local.L.LoadList(ctx, tx, false, (*[]*Battlereport)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.List == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.List = nil
	if err = local.L.LoadList(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.List == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testBattlereportToOneSetOpArmyListUsingList(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Battlereport
	var b, c ArmyList

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, battlereportDBTypes, false, strmangle.SetComplement(battlereportPrimaryKeyColumns, battlereportColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, armyListDBTypes, false, strmangle.SetComplement(armyListPrimaryKeyColumns, armyListColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, armyListDBTypes, false, strmangle.SetComplement(armyListPrimaryKeyColumns, armyListColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*ArmyList{&b, &c} {
		err = a.SetList(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.List != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ListBattlereports[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.ListID, x.ID) {
			t.Error("foreign key was wrong value", a.ListID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.ListID))
		reflect.Indirect(reflect.ValueOf(&a.ListID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.ListID, x.ID) {
			t.Error("foreign key was wrong value", a.ListID, x.ID)
		}
	}
}

func testBattlereportToOneRemoveOpArmyListUsingList(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Battlereport
	var b ArmyList

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, battlereportDBTypes, false, strmangle.SetComplement(battlereportPrimaryKeyColumns, battlereportColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, armyListDBTypes, false, strmangle.SetComplement(armyListPrimaryKeyColumns, armyListColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetList(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveList(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.List().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.List != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.ListID) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.ListBattlereports) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testBattlereportsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Battlereport{}
	if err = randomize.Struct(seed, o, battlereportDBTypes, true, battlereportColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Battlereport struct: %s", err)
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

func testBattlereportsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Battlereport{}
	if err = randomize.Struct(seed, o, battlereportDBTypes, true, battlereportColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Battlereport struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BattlereportSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testBattlereportsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Battlereport{}
	if err = randomize.Struct(seed, o, battlereportDBTypes, true, battlereportColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Battlereport struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Battlereports().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	battlereportDBTypes = map[string]string{`ID`: `integer`, `UserID`: `character varying`, `UserFaction`: `character varying`, `ListID`: `integer`, `EnemyFaction`: `character varying`, `EnemyList`: `text`, `GameMode`: `integer`, `Win`: `boolean`, `PlayerScore`: `integer`, `EnemyScore`: `integer`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`}
	_                   = bytes.MinRead
)

func testBattlereportsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(battlereportPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(battlereportColumns) == len(battlereportPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Battlereport{}
	if err = randomize.Struct(seed, o, battlereportDBTypes, true, battlereportColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Battlereport struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Battlereports().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, battlereportDBTypes, true, battlereportPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Battlereport struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testBattlereportsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(battlereportColumns) == len(battlereportPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Battlereport{}
	if err = randomize.Struct(seed, o, battlereportDBTypes, true, battlereportColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Battlereport struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Battlereports().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, battlereportDBTypes, true, battlereportPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Battlereport struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(battlereportColumns, battlereportPrimaryKeyColumns) {
		fields = battlereportColumns
	} else {
		fields = strmangle.SetComplement(
			battlereportColumns,
			battlereportPrimaryKeyColumns,
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

	slice := BattlereportSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testBattlereportsUpsert(t *testing.T) {
	t.Parallel()

	if len(battlereportColumns) == len(battlereportPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Battlereport{}
	if err = randomize.Struct(seed, &o, battlereportDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Battlereport struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Battlereport: %s", err)
	}

	count, err := Battlereports().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, battlereportDBTypes, false, battlereportPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Battlereport struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Battlereport: %s", err)
	}

	count, err = Battlereports().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
