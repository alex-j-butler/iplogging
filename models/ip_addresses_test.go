package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testIPAddresses(t *testing.T) {
	t.Parallel()

	query := IPAddresses(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testIPAddressesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	ipAddress := &IPAddress{}
	if err = randomize.Struct(seed, ipAddress, ipAddressDBTypes, true); err != nil {
		t.Errorf("Unable to randomize IPAddress struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = ipAddress.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = ipAddress.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := IPAddresses(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testIPAddressesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	ipAddress := &IPAddress{}
	if err = randomize.Struct(seed, ipAddress, ipAddressDBTypes, true); err != nil {
		t.Errorf("Unable to randomize IPAddress struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = ipAddress.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = IPAddresses(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := IPAddresses(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testIPAddressesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	ipAddress := &IPAddress{}
	if err = randomize.Struct(seed, ipAddress, ipAddressDBTypes, true); err != nil {
		t.Errorf("Unable to randomize IPAddress struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = ipAddress.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := IPAddressSlice{ipAddress}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := IPAddresses(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testIPAddressesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	ipAddress := &IPAddress{}
	if err = randomize.Struct(seed, ipAddress, ipAddressDBTypes, true, ipAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IPAddress struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = ipAddress.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := IPAddressExists(tx, ipAddress.IpaddressID)
	if err != nil {
		t.Errorf("Unable to check if IPAddress exists: %s", err)
	}
	if !e {
		t.Errorf("Expected IPAddressExistsG to return true, but got false.")
	}
}
func testIPAddressesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	ipAddress := &IPAddress{}
	if err = randomize.Struct(seed, ipAddress, ipAddressDBTypes, true, ipAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IPAddress struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = ipAddress.Insert(tx); err != nil {
		t.Error(err)
	}

	ipAddressFound, err := FindIPAddress(tx, ipAddress.IpaddressID)
	if err != nil {
		t.Error(err)
	}

	if ipAddressFound == nil {
		t.Error("want a record, got nil")
	}
}
func testIPAddressesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	ipAddress := &IPAddress{}
	if err = randomize.Struct(seed, ipAddress, ipAddressDBTypes, true, ipAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IPAddress struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = ipAddress.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = IPAddresses(tx).Bind(ipAddress); err != nil {
		t.Error(err)
	}
}

func testIPAddressesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	ipAddress := &IPAddress{}
	if err = randomize.Struct(seed, ipAddress, ipAddressDBTypes, true, ipAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IPAddress struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = ipAddress.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := IPAddresses(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testIPAddressesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	ipAddressOne := &IPAddress{}
	ipAddressTwo := &IPAddress{}
	if err = randomize.Struct(seed, ipAddressOne, ipAddressDBTypes, false, ipAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IPAddress struct: %s", err)
	}
	if err = randomize.Struct(seed, ipAddressTwo, ipAddressDBTypes, false, ipAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IPAddress struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = ipAddressOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = ipAddressTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := IPAddresses(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testIPAddressesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	ipAddressOne := &IPAddress{}
	ipAddressTwo := &IPAddress{}
	if err = randomize.Struct(seed, ipAddressOne, ipAddressDBTypes, false, ipAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IPAddress struct: %s", err)
	}
	if err = randomize.Struct(seed, ipAddressTwo, ipAddressDBTypes, false, ipAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IPAddress struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = ipAddressOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = ipAddressTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := IPAddresses(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func ipAddressBeforeInsertHook(e boil.Executor, o *IPAddress) error {
	*o = IPAddress{}
	return nil
}

func ipAddressAfterInsertHook(e boil.Executor, o *IPAddress) error {
	*o = IPAddress{}
	return nil
}

func ipAddressAfterSelectHook(e boil.Executor, o *IPAddress) error {
	*o = IPAddress{}
	return nil
}

func ipAddressBeforeUpdateHook(e boil.Executor, o *IPAddress) error {
	*o = IPAddress{}
	return nil
}

func ipAddressAfterUpdateHook(e boil.Executor, o *IPAddress) error {
	*o = IPAddress{}
	return nil
}

func ipAddressBeforeDeleteHook(e boil.Executor, o *IPAddress) error {
	*o = IPAddress{}
	return nil
}

func ipAddressAfterDeleteHook(e boil.Executor, o *IPAddress) error {
	*o = IPAddress{}
	return nil
}

func ipAddressBeforeUpsertHook(e boil.Executor, o *IPAddress) error {
	*o = IPAddress{}
	return nil
}

func ipAddressAfterUpsertHook(e boil.Executor, o *IPAddress) error {
	*o = IPAddress{}
	return nil
}

func testIPAddressesHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &IPAddress{}
	o := &IPAddress{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, ipAddressDBTypes, false); err != nil {
		t.Errorf("Unable to randomize IPAddress object: %s", err)
	}

	AddIPAddressHook(boil.BeforeInsertHook, ipAddressBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	ipAddressBeforeInsertHooks = []IPAddressHook{}

	AddIPAddressHook(boil.AfterInsertHook, ipAddressAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	ipAddressAfterInsertHooks = []IPAddressHook{}

	AddIPAddressHook(boil.AfterSelectHook, ipAddressAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	ipAddressAfterSelectHooks = []IPAddressHook{}

	AddIPAddressHook(boil.BeforeUpdateHook, ipAddressBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	ipAddressBeforeUpdateHooks = []IPAddressHook{}

	AddIPAddressHook(boil.AfterUpdateHook, ipAddressAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	ipAddressAfterUpdateHooks = []IPAddressHook{}

	AddIPAddressHook(boil.BeforeDeleteHook, ipAddressBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	ipAddressBeforeDeleteHooks = []IPAddressHook{}

	AddIPAddressHook(boil.AfterDeleteHook, ipAddressAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	ipAddressAfterDeleteHooks = []IPAddressHook{}

	AddIPAddressHook(boil.BeforeUpsertHook, ipAddressBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	ipAddressBeforeUpsertHooks = []IPAddressHook{}

	AddIPAddressHook(boil.AfterUpsertHook, ipAddressAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	ipAddressAfterUpsertHooks = []IPAddressHook{}
}
func testIPAddressesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	ipAddress := &IPAddress{}
	if err = randomize.Struct(seed, ipAddress, ipAddressDBTypes, true, ipAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IPAddress struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = ipAddress.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := IPAddresses(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testIPAddressesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	ipAddress := &IPAddress{}
	if err = randomize.Struct(seed, ipAddress, ipAddressDBTypes, true); err != nil {
		t.Errorf("Unable to randomize IPAddress struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = ipAddress.Insert(tx, ipAddressColumns...); err != nil {
		t.Error(err)
	}

	count, err := IPAddresses(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testIPAddressToManyIpaddresIPAddressUsers(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a IPAddress
	var b, c IPAddressUser

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, ipAddressDBTypes, true, ipAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IPAddress struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, ipAddressUserDBTypes, false, ipAddressUserColumnsWithDefault...)
	randomize.Struct(seed, &c, ipAddressUserDBTypes, false, ipAddressUserColumnsWithDefault...)
	b.IpaddressID.Valid = true
	c.IpaddressID.Valid = true
	b.IpaddressID.Int = a.IpaddressID
	c.IpaddressID.Int = a.IpaddressID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	ipAddressUser, err := a.IpaddresIPAddressUsers(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range ipAddressUser {
		if v.IpaddressID.Int == b.IpaddressID.Int {
			bFound = true
		}
		if v.IpaddressID.Int == c.IpaddressID.Int {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := IPAddressSlice{&a}
	if err = a.L.LoadIpaddresIPAddressUsers(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.IpaddresIPAddressUsers); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.IpaddresIPAddressUsers = nil
	if err = a.L.LoadIpaddresIPAddressUsers(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.IpaddresIPAddressUsers); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", ipAddressUser)
	}
}

func testIPAddressToManyAddOpIpaddresIPAddressUsers(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a IPAddress
	var b, c, d, e IPAddressUser

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, ipAddressDBTypes, false, strmangle.SetComplement(ipAddressPrimaryKeyColumns, ipAddressColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*IPAddressUser{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, ipAddressUserDBTypes, false, strmangle.SetComplement(ipAddressUserPrimaryKeyColumns, ipAddressUserColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*IPAddressUser{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddIpaddresIPAddressUsers(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.IpaddressID != first.IpaddressID.Int {
			t.Error("foreign key was wrong value", a.IpaddressID, first.IpaddressID.Int)
		}
		if a.IpaddressID != second.IpaddressID.Int {
			t.Error("foreign key was wrong value", a.IpaddressID, second.IpaddressID.Int)
		}

		if first.R.Ipaddre != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Ipaddre != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.IpaddresIPAddressUsers[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.IpaddresIPAddressUsers[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.IpaddresIPAddressUsers(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testIPAddressToManySetOpIpaddresIPAddressUsers(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a IPAddress
	var b, c, d, e IPAddressUser

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, ipAddressDBTypes, false, strmangle.SetComplement(ipAddressPrimaryKeyColumns, ipAddressColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*IPAddressUser{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, ipAddressUserDBTypes, false, strmangle.SetComplement(ipAddressUserPrimaryKeyColumns, ipAddressUserColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	err = a.SetIpaddresIPAddressUsers(tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.IpaddresIPAddressUsers(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetIpaddresIPAddressUsers(tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.IpaddresIPAddressUsers(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if b.IpaddressID.Valid {
		t.Error("want b's foreign key value to be nil")
	}
	if c.IpaddressID.Valid {
		t.Error("want c's foreign key value to be nil")
	}
	if a.IpaddressID != d.IpaddressID.Int {
		t.Error("foreign key was wrong value", a.IpaddressID, d.IpaddressID.Int)
	}
	if a.IpaddressID != e.IpaddressID.Int {
		t.Error("foreign key was wrong value", a.IpaddressID, e.IpaddressID.Int)
	}

	if b.R.Ipaddre != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Ipaddre != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Ipaddre != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Ipaddre != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.IpaddresIPAddressUsers[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.IpaddresIPAddressUsers[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testIPAddressToManyRemoveOpIpaddresIPAddressUsers(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a IPAddress
	var b, c, d, e IPAddressUser

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, ipAddressDBTypes, false, strmangle.SetComplement(ipAddressPrimaryKeyColumns, ipAddressColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*IPAddressUser{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, ipAddressUserDBTypes, false, strmangle.SetComplement(ipAddressUserPrimaryKeyColumns, ipAddressUserColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	err = a.AddIpaddresIPAddressUsers(tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.IpaddresIPAddressUsers(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveIpaddresIPAddressUsers(tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.IpaddresIPAddressUsers(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if b.IpaddressID.Valid {
		t.Error("want b's foreign key value to be nil")
	}
	if c.IpaddressID.Valid {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.Ipaddre != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Ipaddre != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Ipaddre != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.Ipaddre != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.IpaddresIPAddressUsers) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.IpaddresIPAddressUsers[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.IpaddresIPAddressUsers[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testIPAddressesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	ipAddress := &IPAddress{}
	if err = randomize.Struct(seed, ipAddress, ipAddressDBTypes, true, ipAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IPAddress struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = ipAddress.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = ipAddress.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testIPAddressesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	ipAddress := &IPAddress{}
	if err = randomize.Struct(seed, ipAddress, ipAddressDBTypes, true, ipAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IPAddress struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = ipAddress.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := IPAddressSlice{ipAddress}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testIPAddressesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	ipAddress := &IPAddress{}
	if err = randomize.Struct(seed, ipAddress, ipAddressDBTypes, true, ipAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IPAddress struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = ipAddress.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := IPAddresses(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	ipAddressDBTypes = map[string]string{`Ipaddress`: `integer`, `IpaddressID`: `integer`}
	_                = bytes.MinRead
)

func testIPAddressesUpdate(t *testing.T) {
	t.Parallel()

	if len(ipAddressColumns) == len(ipAddressPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	ipAddress := &IPAddress{}
	if err = randomize.Struct(seed, ipAddress, ipAddressDBTypes, true); err != nil {
		t.Errorf("Unable to randomize IPAddress struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = ipAddress.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := IPAddresses(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, ipAddress, ipAddressDBTypes, true, ipAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IPAddress struct: %s", err)
	}

	if err = ipAddress.Update(tx); err != nil {
		t.Error(err)
	}
}

func testIPAddressesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(ipAddressColumns) == len(ipAddressPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	ipAddress := &IPAddress{}
	if err = randomize.Struct(seed, ipAddress, ipAddressDBTypes, true); err != nil {
		t.Errorf("Unable to randomize IPAddress struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = ipAddress.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := IPAddresses(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, ipAddress, ipAddressDBTypes, true, ipAddressPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize IPAddress struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(ipAddressColumns, ipAddressPrimaryKeyColumns) {
		fields = ipAddressColumns
	} else {
		fields = strmangle.SetComplement(
			ipAddressColumns,
			ipAddressPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(ipAddress))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := IPAddressSlice{ipAddress}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testIPAddressesUpsert(t *testing.T) {
	t.Parallel()

	if len(ipAddressColumns) == len(ipAddressPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	ipAddress := IPAddress{}
	if err = randomize.Struct(seed, &ipAddress, ipAddressDBTypes, true); err != nil {
		t.Errorf("Unable to randomize IPAddress struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = ipAddress.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert IPAddress: %s", err)
	}

	count, err := IPAddresses(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &ipAddress, ipAddressDBTypes, false, ipAddressPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize IPAddress struct: %s", err)
	}

	if err = ipAddress.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert IPAddress: %s", err)
	}

	count, err = IPAddresses(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
