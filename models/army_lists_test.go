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

func testArmyLists(t *testing.T) {
	t.Parallel()

	query := ArmyLists()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testArmyListsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ArmyList{}
	if err = randomize.Struct(seed, o, armyListDBTypes, true, armyListColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ArmyList struct: %s", err)
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

	count, err := ArmyLists().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testArmyListsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ArmyList{}
	if err = randomize.Struct(seed, o, armyListDBTypes, true, armyListColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ArmyList struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := ArmyLists().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ArmyLists().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testArmyListsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ArmyList{}
	if err = randomize.Struct(seed, o, armyListDBTypes, true, armyListColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ArmyList struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ArmyListSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ArmyLists().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testArmyListsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ArmyList{}
	if err = randomize.Struct(seed, o, armyListDBTypes, true, armyListColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ArmyList struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ArmyListExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if ArmyList exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ArmyListExists to return true, but got false.")
	}
}

func testArmyListsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ArmyList{}
	if err = randomize.Struct(seed, o, armyListDBTypes, true, armyListColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ArmyList struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	armyListFound, err := FindArmyList(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if armyListFound == nil {
		t.Error("want a record, got nil")
	}
}

func testArmyListsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ArmyList{}
	if err = randomize.Struct(seed, o, armyListDBTypes, true, armyListColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ArmyList struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = ArmyLists().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testArmyListsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ArmyList{}
	if err = randomize.Struct(seed, o, armyListDBTypes, true, armyListColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ArmyList struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := ArmyLists().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testArmyListsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	armyListOne := &ArmyList{}
	armyListTwo := &ArmyList{}
	if err = randomize.Struct(seed, armyListOne, armyListDBTypes, false, armyListColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ArmyList struct: %s", err)
	}
	if err = randomize.Struct(seed, armyListTwo, armyListDBTypes, false, armyListColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ArmyList struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = armyListOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = armyListTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ArmyLists().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testArmyListsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	armyListOne := &ArmyList{}
	armyListTwo := &ArmyList{}
	if err = randomize.Struct(seed, armyListOne, armyListDBTypes, false, armyListColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ArmyList struct: %s", err)
	}
	if err = randomize.Struct(seed, armyListTwo, armyListDBTypes, false, armyListColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ArmyList struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = armyListOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = armyListTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ArmyLists().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func armyListBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *ArmyList) error {
	*o = ArmyList{}
	return nil
}

func armyListAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *ArmyList) error {
	*o = ArmyList{}
	return nil
}

func armyListAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *ArmyList) error {
	*o = ArmyList{}
	return nil
}

func armyListBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ArmyList) error {
	*o = ArmyList{}
	return nil
}

func armyListAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ArmyList) error {
	*o = ArmyList{}
	return nil
}

func armyListBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ArmyList) error {
	*o = ArmyList{}
	return nil
}

func armyListAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ArmyList) error {
	*o = ArmyList{}
	return nil
}

func armyListBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ArmyList) error {
	*o = ArmyList{}
	return nil
}

func armyListAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ArmyList) error {
	*o = ArmyList{}
	return nil
}

func testArmyListsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &ArmyList{}
	o := &ArmyList{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, armyListDBTypes, false); err != nil {
		t.Errorf("Unable to randomize ArmyList object: %s", err)
	}

	AddArmyListHook(boil.BeforeInsertHook, armyListBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	armyListBeforeInsertHooks = []ArmyListHook{}

	AddArmyListHook(boil.AfterInsertHook, armyListAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	armyListAfterInsertHooks = []ArmyListHook{}

	AddArmyListHook(boil.AfterSelectHook, armyListAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	armyListAfterSelectHooks = []ArmyListHook{}

	AddArmyListHook(boil.BeforeUpdateHook, armyListBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	armyListBeforeUpdateHooks = []ArmyListHook{}

	AddArmyListHook(boil.AfterUpdateHook, armyListAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	armyListAfterUpdateHooks = []ArmyListHook{}

	AddArmyListHook(boil.BeforeDeleteHook, armyListBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	armyListBeforeDeleteHooks = []ArmyListHook{}

	AddArmyListHook(boil.AfterDeleteHook, armyListAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	armyListAfterDeleteHooks = []ArmyListHook{}

	AddArmyListHook(boil.BeforeUpsertHook, armyListBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	armyListBeforeUpsertHooks = []ArmyListHook{}

	AddArmyListHook(boil.AfterUpsertHook, armyListAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	armyListAfterUpsertHooks = []ArmyListHook{}
}

func testArmyListsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ArmyList{}
	if err = randomize.Struct(seed, o, armyListDBTypes, true, armyListColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ArmyList struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ArmyLists().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testArmyListsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ArmyList{}
	if err = randomize.Struct(seed, o, armyListDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ArmyList struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(armyListColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := ArmyLists().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testArmyListToManyListBattlereports(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ArmyList
	var b, c Battlereport

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, armyListDBTypes, true, armyListColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ArmyList struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, battlereportDBTypes, false, battlereportColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, battlereportDBTypes, false, battlereportColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&b.ListID, a.ID)
	queries.Assign(&c.ListID, a.ID)
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.ListBattlereports().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if queries.Equal(v.ListID, b.ListID) {
			bFound = true
		}
		if queries.Equal(v.ListID, c.ListID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := ArmyListSlice{&a}
	if err = a.L.LoadListBattlereports(ctx, tx, false, (*[]*ArmyList)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ListBattlereports); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.ListBattlereports = nil
	if err = a.L.LoadListBattlereports(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ListBattlereports); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testArmyListToManyAddOpListBattlereports(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ArmyList
	var b, c, d, e Battlereport

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, armyListDBTypes, false, strmangle.SetComplement(armyListPrimaryKeyColumns, armyListColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Battlereport{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, battlereportDBTypes, false, strmangle.SetComplement(battlereportPrimaryKeyColumns, battlereportColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Battlereport{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddListBattlereports(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if !queries.Equal(a.ID, first.ListID) {
			t.Error("foreign key was wrong value", a.ID, first.ListID)
		}
		if !queries.Equal(a.ID, second.ListID) {
			t.Error("foreign key was wrong value", a.ID, second.ListID)
		}

		if first.R.List != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.List != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.ListBattlereports[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.ListBattlereports[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.ListBattlereports().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testArmyListToManySetOpListBattlereports(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ArmyList
	var b, c, d, e Battlereport

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, armyListDBTypes, false, strmangle.SetComplement(armyListPrimaryKeyColumns, armyListColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Battlereport{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, battlereportDBTypes, false, strmangle.SetComplement(battlereportPrimaryKeyColumns, battlereportColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.SetListBattlereports(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.ListBattlereports().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetListBattlereports(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.ListBattlereports().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.ListID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.ListID) {
		t.Error("want c's foreign key value to be nil")
	}
	if !queries.Equal(a.ID, d.ListID) {
		t.Error("foreign key was wrong value", a.ID, d.ListID)
	}
	if !queries.Equal(a.ID, e.ListID) {
		t.Error("foreign key was wrong value", a.ID, e.ListID)
	}

	if b.R.List != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.List != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.List != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.List != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.ListBattlereports[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.ListBattlereports[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testArmyListToManyRemoveOpListBattlereports(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ArmyList
	var b, c, d, e Battlereport

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, armyListDBTypes, false, strmangle.SetComplement(armyListPrimaryKeyColumns, armyListColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Battlereport{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, battlereportDBTypes, false, strmangle.SetComplement(battlereportPrimaryKeyColumns, battlereportColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddListBattlereports(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.ListBattlereports().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveListBattlereports(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.ListBattlereports().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.ListID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.ListID) {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.List != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.List != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.List != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.List != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.ListBattlereports) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.ListBattlereports[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.ListBattlereports[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testArmyListsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ArmyList{}
	if err = randomize.Struct(seed, o, armyListDBTypes, true, armyListColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ArmyList struct: %s", err)
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

func testArmyListsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ArmyList{}
	if err = randomize.Struct(seed, o, armyListDBTypes, true, armyListColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ArmyList struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ArmyListSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testArmyListsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ArmyList{}
	if err = randomize.Struct(seed, o, armyListDBTypes, true, armyListColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ArmyList struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ArmyLists().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	armyListDBTypes = map[string]string{`ID`: `integer`, `ListName`: `character varying`, `Faction`: `character varying`, `List`: `text`, `UserID`: `character varying`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`, `PL`: `integer`, `PTS`: `integer`, `CP`: `integer`}
	_               = bytes.MinRead
)

func testArmyListsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(armyListPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(armyListColumns) == len(armyListPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ArmyList{}
	if err = randomize.Struct(seed, o, armyListDBTypes, true, armyListColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ArmyList struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ArmyLists().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, armyListDBTypes, true, armyListPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ArmyList struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testArmyListsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(armyListColumns) == len(armyListPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ArmyList{}
	if err = randomize.Struct(seed, o, armyListDBTypes, true, armyListColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ArmyList struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ArmyLists().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, armyListDBTypes, true, armyListPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ArmyList struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(armyListColumns, armyListPrimaryKeyColumns) {
		fields = armyListColumns
	} else {
		fields = strmangle.SetComplement(
			armyListColumns,
			armyListPrimaryKeyColumns,
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

	slice := ArmyListSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testArmyListsUpsert(t *testing.T) {
	t.Parallel()

	if len(armyListColumns) == len(armyListPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := ArmyList{}
	if err = randomize.Struct(seed, &o, armyListDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ArmyList struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ArmyList: %s", err)
	}

	count, err := ArmyLists().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, armyListDBTypes, false, armyListPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ArmyList struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ArmyList: %s", err)
	}

	count, err = ArmyLists().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
