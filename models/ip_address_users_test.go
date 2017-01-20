package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testIPAddressUsers(t *testing.T) {
	t.Parallel()

	query := IPAddressUsers(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testIPAddressUsersDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	ipAddressUser := &IPAddressUser{}
	if err = randomize.Struct(seed, ipAddressUser, ipAddressUserDBTypes, true); err != nil {
		t.Errorf("Unable to randomize IPAddressUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = ipAddressUser.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = ipAddressUser.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := IPAddressUsers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testIPAddressUsersQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	ipAddressUser := &IPAddressUser{}
	if err = randomize.Struct(seed, ipAddressUser, ipAddressUserDBTypes, true); err != nil {
		t.Errorf("Unable to randomize IPAddressUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = ipAddressUser.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = IPAddressUsers(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := IPAddressUsers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testIPAddressUsersSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	ipAddressUser := &IPAddressUser{}
	if err = randomize.Struct(seed, ipAddressUser, ipAddressUserDBTypes, true); err != nil {
		t.Errorf("Unable to randomize IPAddressUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = ipAddressUser.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := IPAddressUserSlice{ipAddressUser}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := IPAddressUsers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testIPAddressUsersExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	ipAddressUser := &IPAddressUser{}
	if err = randomize.Struct(seed, ipAddressUser, ipAddressUserDBTypes, true, ipAddressUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IPAddressUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = ipAddressUser.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := IPAddressUserExists(tx, ipAddressUser.RefID)
	if err != nil {
		t.Errorf("Unable to check if IPAddressUser exists: %s", err)
	}
	if !e {
		t.Errorf("Expected IPAddressUserExistsG to return true, but got false.")
	}
}
func testIPAddressUsersFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	ipAddressUser := &IPAddressUser{}
	if err = randomize.Struct(seed, ipAddressUser, ipAddressUserDBTypes, true, ipAddressUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IPAddressUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = ipAddressUser.Insert(tx); err != nil {
		t.Error(err)
	}

	ipAddressUserFound, err := FindIPAddressUser(tx, ipAddressUser.RefID)
	if err != nil {
		t.Error(err)
	}

	if ipAddressUserFound == nil {
		t.Error("want a record, got nil")
	}
}
func testIPAddressUsersBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	ipAddressUser := &IPAddressUser{}
	if err = randomize.Struct(seed, ipAddressUser, ipAddressUserDBTypes, true, ipAddressUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IPAddressUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = ipAddressUser.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = IPAddressUsers(tx).Bind(ipAddressUser); err != nil {
		t.Error(err)
	}
}

func testIPAddressUsersOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	ipAddressUser := &IPAddressUser{}
	if err = randomize.Struct(seed, ipAddressUser, ipAddressUserDBTypes, true, ipAddressUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IPAddressUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = ipAddressUser.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := IPAddressUsers(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testIPAddressUsersAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	ipAddressUserOne := &IPAddressUser{}
	ipAddressUserTwo := &IPAddressUser{}
	if err = randomize.Struct(seed, ipAddressUserOne, ipAddressUserDBTypes, false, ipAddressUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IPAddressUser struct: %s", err)
	}
	if err = randomize.Struct(seed, ipAddressUserTwo, ipAddressUserDBTypes, false, ipAddressUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IPAddressUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = ipAddressUserOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = ipAddressUserTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := IPAddressUsers(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testIPAddressUsersCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	ipAddressUserOne := &IPAddressUser{}
	ipAddressUserTwo := &IPAddressUser{}
	if err = randomize.Struct(seed, ipAddressUserOne, ipAddressUserDBTypes, false, ipAddressUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IPAddressUser struct: %s", err)
	}
	if err = randomize.Struct(seed, ipAddressUserTwo, ipAddressUserDBTypes, false, ipAddressUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IPAddressUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = ipAddressUserOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = ipAddressUserTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := IPAddressUsers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func ipAddressUserBeforeInsertHook(e boil.Executor, o *IPAddressUser) error {
	*o = IPAddressUser{}
	return nil
}

func ipAddressUserAfterInsertHook(e boil.Executor, o *IPAddressUser) error {
	*o = IPAddressUser{}
	return nil
}

func ipAddressUserAfterSelectHook(e boil.Executor, o *IPAddressUser) error {
	*o = IPAddressUser{}
	return nil
}

func ipAddressUserBeforeUpdateHook(e boil.Executor, o *IPAddressUser) error {
	*o = IPAddressUser{}
	return nil
}

func ipAddressUserAfterUpdateHook(e boil.Executor, o *IPAddressUser) error {
	*o = IPAddressUser{}
	return nil
}

func ipAddressUserBeforeDeleteHook(e boil.Executor, o *IPAddressUser) error {
	*o = IPAddressUser{}
	return nil
}

func ipAddressUserAfterDeleteHook(e boil.Executor, o *IPAddressUser) error {
	*o = IPAddressUser{}
	return nil
}

func ipAddressUserBeforeUpsertHook(e boil.Executor, o *IPAddressUser) error {
	*o = IPAddressUser{}
	return nil
}

func ipAddressUserAfterUpsertHook(e boil.Executor, o *IPAddressUser) error {
	*o = IPAddressUser{}
	return nil
}

func testIPAddressUsersHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &IPAddressUser{}
	o := &IPAddressUser{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, ipAddressUserDBTypes, false); err != nil {
		t.Errorf("Unable to randomize IPAddressUser object: %s", err)
	}

	AddIPAddressUserHook(boil.BeforeInsertHook, ipAddressUserBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	ipAddressUserBeforeInsertHooks = []IPAddressUserHook{}

	AddIPAddressUserHook(boil.AfterInsertHook, ipAddressUserAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	ipAddressUserAfterInsertHooks = []IPAddressUserHook{}

	AddIPAddressUserHook(boil.AfterSelectHook, ipAddressUserAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	ipAddressUserAfterSelectHooks = []IPAddressUserHook{}

	AddIPAddressUserHook(boil.BeforeUpdateHook, ipAddressUserBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	ipAddressUserBeforeUpdateHooks = []IPAddressUserHook{}

	AddIPAddressUserHook(boil.AfterUpdateHook, ipAddressUserAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	ipAddressUserAfterUpdateHooks = []IPAddressUserHook{}

	AddIPAddressUserHook(boil.BeforeDeleteHook, ipAddressUserBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	ipAddressUserBeforeDeleteHooks = []IPAddressUserHook{}

	AddIPAddressUserHook(boil.AfterDeleteHook, ipAddressUserAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	ipAddressUserAfterDeleteHooks = []IPAddressUserHook{}

	AddIPAddressUserHook(boil.BeforeUpsertHook, ipAddressUserBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	ipAddressUserBeforeUpsertHooks = []IPAddressUserHook{}

	AddIPAddressUserHook(boil.AfterUpsertHook, ipAddressUserAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	ipAddressUserAfterUpsertHooks = []IPAddressUserHook{}
}
func testIPAddressUsersInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	ipAddressUser := &IPAddressUser{}
	if err = randomize.Struct(seed, ipAddressUser, ipAddressUserDBTypes, true, ipAddressUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IPAddressUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = ipAddressUser.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := IPAddressUsers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testIPAddressUsersInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	ipAddressUser := &IPAddressUser{}
	if err = randomize.Struct(seed, ipAddressUser, ipAddressUserDBTypes, true); err != nil {
		t.Errorf("Unable to randomize IPAddressUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = ipAddressUser.Insert(tx, ipAddressUserColumns...); err != nil {
		t.Error(err)
	}

	count, err := IPAddressUsers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testIPAddressUserToOneIPAddressUsingIpaddres(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local IPAddressUser
	var foreign IPAddress

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, ipAddressUserDBTypes, true, ipAddressUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IPAddressUser struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, ipAddressDBTypes, true, ipAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IPAddress struct: %s", err)
	}

	local.IpaddressID.Valid = true

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.IpaddressID.Int = foreign.IpaddressID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Ipaddres(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.IpaddressID != foreign.IpaddressID {
		t.Errorf("want: %v, got %v", foreign.IpaddressID, check.IpaddressID)
	}

	slice := IPAddressUserSlice{&local}
	if err = local.L.LoadIpaddres(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Ipaddres == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Ipaddres = nil
	if err = local.L.LoadIpaddres(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Ipaddres == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testIPAddressUserToOneUserUsingUser(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local IPAddressUser
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, ipAddressUserDBTypes, true, ipAddressUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IPAddressUser struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, userDBTypes, true, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	local.UserID.Valid = true

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.UserID.Int = foreign.UserID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.User(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.UserID != foreign.UserID {
		t.Errorf("want: %v, got %v", foreign.UserID, check.UserID)
	}

	slice := IPAddressUserSlice{&local}
	if err = local.L.LoadUser(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.User = nil
	if err = local.L.LoadUser(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testIPAddressUserToOneSetOpIPAddressUsingIpaddres(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a IPAddressUser
	var b, c IPAddress

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, ipAddressUserDBTypes, false, strmangle.SetComplement(ipAddressUserPrimaryKeyColumns, ipAddressUserColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, ipAddressDBTypes, false, strmangle.SetComplement(ipAddressPrimaryKeyColumns, ipAddressColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, ipAddressDBTypes, false, strmangle.SetComplement(ipAddressPrimaryKeyColumns, ipAddressColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*IPAddress{&b, &c} {
		err = a.SetIpaddres(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Ipaddres != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.IpaddresIPAddressUsers[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.IpaddressID.Int != x.IpaddressID {
			t.Error("foreign key was wrong value", a.IpaddressID.Int)
		}

		zero := reflect.Zero(reflect.TypeOf(a.IpaddressID.Int))
		reflect.Indirect(reflect.ValueOf(&a.IpaddressID.Int)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.IpaddressID.Int != x.IpaddressID {
			t.Error("foreign key was wrong value", a.IpaddressID.Int, x.IpaddressID)
		}
	}
}

func testIPAddressUserToOneRemoveOpIPAddressUsingIpaddres(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a IPAddressUser
	var b IPAddress

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, ipAddressUserDBTypes, false, strmangle.SetComplement(ipAddressUserPrimaryKeyColumns, ipAddressUserColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, ipAddressDBTypes, false, strmangle.SetComplement(ipAddressPrimaryKeyColumns, ipAddressColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	if err = a.SetIpaddres(tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveIpaddres(tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.Ipaddres(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.Ipaddres != nil {
		t.Error("R struct entry should be nil")
	}

	if a.IpaddressID.Valid {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.IpaddresIPAddressUsers) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testIPAddressUserToOneSetOpUserUsingUser(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a IPAddressUser
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, ipAddressUserDBTypes, false, strmangle.SetComplement(ipAddressUserPrimaryKeyColumns, ipAddressUserColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*User{&b, &c} {
		err = a.SetUser(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.User != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.IPAddressUsers[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.UserID.Int != x.UserID {
			t.Error("foreign key was wrong value", a.UserID.Int)
		}

		zero := reflect.Zero(reflect.TypeOf(a.UserID.Int))
		reflect.Indirect(reflect.ValueOf(&a.UserID.Int)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.UserID.Int != x.UserID {
			t.Error("foreign key was wrong value", a.UserID.Int, x.UserID)
		}
	}
}

func testIPAddressUserToOneRemoveOpUserUsingUser(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a IPAddressUser
	var b User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, ipAddressUserDBTypes, false, strmangle.SetComplement(ipAddressUserPrimaryKeyColumns, ipAddressUserColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	if err = a.SetUser(tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveUser(tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.User(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.User != nil {
		t.Error("R struct entry should be nil")
	}

	if a.UserID.Valid {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.IPAddressUsers) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testIPAddressUsersReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	ipAddressUser := &IPAddressUser{}
	if err = randomize.Struct(seed, ipAddressUser, ipAddressUserDBTypes, true, ipAddressUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IPAddressUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = ipAddressUser.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = ipAddressUser.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testIPAddressUsersReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	ipAddressUser := &IPAddressUser{}
	if err = randomize.Struct(seed, ipAddressUser, ipAddressUserDBTypes, true, ipAddressUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IPAddressUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = ipAddressUser.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := IPAddressUserSlice{ipAddressUser}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testIPAddressUsersSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	ipAddressUser := &IPAddressUser{}
	if err = randomize.Struct(seed, ipAddressUser, ipAddressUserDBTypes, true, ipAddressUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IPAddressUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = ipAddressUser.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := IPAddressUsers(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	ipAddressUserDBTypes = map[string]string{`ConnectionTime`: `timestamp without time zone`, `IpaddressID`: `integer`, `RefID`: `integer`, `UserID`: `integer`}
	_                    = bytes.MinRead
)

func testIPAddressUsersUpdate(t *testing.T) {
	t.Parallel()

	if len(ipAddressUserColumns) == len(ipAddressUserPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	ipAddressUser := &IPAddressUser{}
	if err = randomize.Struct(seed, ipAddressUser, ipAddressUserDBTypes, true); err != nil {
		t.Errorf("Unable to randomize IPAddressUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = ipAddressUser.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := IPAddressUsers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, ipAddressUser, ipAddressUserDBTypes, true, ipAddressUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IPAddressUser struct: %s", err)
	}

	if err = ipAddressUser.Update(tx); err != nil {
		t.Error(err)
	}
}

func testIPAddressUsersSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(ipAddressUserColumns) == len(ipAddressUserPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	ipAddressUser := &IPAddressUser{}
	if err = randomize.Struct(seed, ipAddressUser, ipAddressUserDBTypes, true); err != nil {
		t.Errorf("Unable to randomize IPAddressUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = ipAddressUser.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := IPAddressUsers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, ipAddressUser, ipAddressUserDBTypes, true, ipAddressUserPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize IPAddressUser struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(ipAddressUserColumns, ipAddressUserPrimaryKeyColumns) {
		fields = ipAddressUserColumns
	} else {
		fields = strmangle.SetComplement(
			ipAddressUserColumns,
			ipAddressUserPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(ipAddressUser))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := IPAddressUserSlice{ipAddressUser}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testIPAddressUsersUpsert(t *testing.T) {
	t.Parallel()

	if len(ipAddressUserColumns) == len(ipAddressUserPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	ipAddressUser := IPAddressUser{}
	if err = randomize.Struct(seed, &ipAddressUser, ipAddressUserDBTypes, true); err != nil {
		t.Errorf("Unable to randomize IPAddressUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = ipAddressUser.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert IPAddressUser: %s", err)
	}

	count, err := IPAddressUsers(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &ipAddressUser, ipAddressUserDBTypes, false, ipAddressUserPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize IPAddressUser struct: %s", err)
	}

	if err = ipAddressUser.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert IPAddressUser: %s", err)
	}

	count, err = IPAddressUsers(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
