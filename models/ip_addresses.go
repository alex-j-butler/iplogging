package models

import (
	"bytes"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/queries"
	"github.com/vattle/sqlboiler/queries/qm"
	"github.com/vattle/sqlboiler/strmangle"
	"gopkg.in/nullbio/null.v6"
)

// IPAddress is an object representing the database table.
type IPAddress struct {
	IpaddressID int      `boil:"ipaddress_id" json:"ipaddress_id" toml:"ipaddress_id" yaml:"ipaddress_id"`
	Ipaddress   null.Int `boil:"ipaddress" json:"ipaddress,omitempty" toml:"ipaddress" yaml:"ipaddress,omitempty"`

	R *ipAddressR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L ipAddressL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// ipAddressR is where relationships are stored.
type ipAddressR struct {
	IpaddresIPAddressUsers IPAddressUserSlice
}

// ipAddressL is where Load methods for each relationship are stored.
type ipAddressL struct{}

var (
	ipAddressColumns               = []string{"ipaddress_id", "ipaddress"}
	ipAddressColumnsWithoutDefault = []string{"ipaddress"}
	ipAddressColumnsWithDefault    = []string{"ipaddress_id"}
	ipAddressPrimaryKeyColumns     = []string{"ipaddress_id"}
)

type (
	// IPAddressSlice is an alias for a slice of pointers to IPAddress.
	// This should generally be used opposed to []IPAddress.
	IPAddressSlice []*IPAddress
	// IPAddressHook is the signature for custom IPAddress hook methods
	IPAddressHook func(boil.Executor, *IPAddress) error

	ipAddressQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	ipAddressType                 = reflect.TypeOf(&IPAddress{})
	ipAddressMapping              = queries.MakeStructMapping(ipAddressType)
	ipAddressPrimaryKeyMapping, _ = queries.BindMapping(ipAddressType, ipAddressMapping, ipAddressPrimaryKeyColumns)
	ipAddressInsertCacheMut       sync.RWMutex
	ipAddressInsertCache          = make(map[string]insertCache)
	ipAddressUpdateCacheMut       sync.RWMutex
	ipAddressUpdateCache          = make(map[string]updateCache)
	ipAddressUpsertCacheMut       sync.RWMutex
	ipAddressUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var ipAddressBeforeInsertHooks []IPAddressHook
var ipAddressBeforeUpdateHooks []IPAddressHook
var ipAddressBeforeDeleteHooks []IPAddressHook
var ipAddressBeforeUpsertHooks []IPAddressHook

var ipAddressAfterInsertHooks []IPAddressHook
var ipAddressAfterSelectHooks []IPAddressHook
var ipAddressAfterUpdateHooks []IPAddressHook
var ipAddressAfterDeleteHooks []IPAddressHook
var ipAddressAfterUpsertHooks []IPAddressHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *IPAddress) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range ipAddressBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *IPAddress) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range ipAddressBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *IPAddress) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range ipAddressBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *IPAddress) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range ipAddressBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *IPAddress) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range ipAddressAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *IPAddress) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range ipAddressAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *IPAddress) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range ipAddressAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *IPAddress) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range ipAddressAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *IPAddress) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range ipAddressAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddIPAddressHook registers your hook function for all future operations.
func AddIPAddressHook(hookPoint boil.HookPoint, ipAddressHook IPAddressHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		ipAddressBeforeInsertHooks = append(ipAddressBeforeInsertHooks, ipAddressHook)
	case boil.BeforeUpdateHook:
		ipAddressBeforeUpdateHooks = append(ipAddressBeforeUpdateHooks, ipAddressHook)
	case boil.BeforeDeleteHook:
		ipAddressBeforeDeleteHooks = append(ipAddressBeforeDeleteHooks, ipAddressHook)
	case boil.BeforeUpsertHook:
		ipAddressBeforeUpsertHooks = append(ipAddressBeforeUpsertHooks, ipAddressHook)
	case boil.AfterInsertHook:
		ipAddressAfterInsertHooks = append(ipAddressAfterInsertHooks, ipAddressHook)
	case boil.AfterSelectHook:
		ipAddressAfterSelectHooks = append(ipAddressAfterSelectHooks, ipAddressHook)
	case boil.AfterUpdateHook:
		ipAddressAfterUpdateHooks = append(ipAddressAfterUpdateHooks, ipAddressHook)
	case boil.AfterDeleteHook:
		ipAddressAfterDeleteHooks = append(ipAddressAfterDeleteHooks, ipAddressHook)
	case boil.AfterUpsertHook:
		ipAddressAfterUpsertHooks = append(ipAddressAfterUpsertHooks, ipAddressHook)
	}
}

// OneP returns a single ipAddress record from the query, and panics on error.
func (q ipAddressQuery) OneP() *IPAddress {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single ipAddress record from the query.
func (q ipAddressQuery) One() (*IPAddress, error) {
	o := &IPAddress{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for ip_addresses")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all IPAddress records from the query, and panics on error.
func (q ipAddressQuery) AllP() IPAddressSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all IPAddress records from the query.
func (q ipAddressQuery) All() (IPAddressSlice, error) {
	var o IPAddressSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to IPAddress slice")
	}

	if len(ipAddressAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all IPAddress records in the query, and panics on error.
func (q ipAddressQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all IPAddress records in the query.
func (q ipAddressQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count ip_addresses rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q ipAddressQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q ipAddressQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if ip_addresses exists")
	}

	return count > 0, nil
}

// IpaddresIPAddressUsersG retrieves all the ip_address_user's ip address users via ipaddress_id column.
func (o *IPAddress) IpaddresIPAddressUsersG(mods ...qm.QueryMod) ipAddressUserQuery {
	return o.IpaddresIPAddressUsers(boil.GetDB(), mods...)
}

// IpaddresIPAddressUsers retrieves all the ip_address_user's ip address users with an executor via ipaddress_id column.
func (o *IPAddress) IpaddresIPAddressUsers(exec boil.Executor, mods ...qm.QueryMod) ipAddressUserQuery {
	queryMods := []qm.QueryMod{
		qm.Select("\"a\".*"),
	}

	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"a\".\"ipaddress_id\"=?", o.IpaddressID),
	)

	query := IPAddressUsers(exec, queryMods...)
	queries.SetFrom(query.Query, "\"ip_address_users\" as \"a\"")
	return query
}

// LoadIpaddresIPAddressUsers allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (ipAddressL) LoadIpaddresIPAddressUsers(e boil.Executor, singular bool, maybeIPAddress interface{}) error {
	var slice []*IPAddress
	var object *IPAddress

	count := 1
	if singular {
		object = maybeIPAddress.(*IPAddress)
	} else {
		slice = *maybeIPAddress.(*IPAddressSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &ipAddressR{}
		}
		args[0] = object.IpaddressID
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &ipAddressR{}
			}
			args[i] = obj.IpaddressID
		}
	}

	query := fmt.Sprintf(
		"select * from \"ip_address_users\" where \"ipaddress_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)
	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load ip_address_users")
	}
	defer results.Close()

	var resultSlice []*IPAddressUser
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice ip_address_users")
	}

	if len(ipAddressUserAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.IpaddresIPAddressUsers = resultSlice
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.IpaddressID == foreign.IpaddressID.Int {
				local.R.IpaddresIPAddressUsers = append(local.R.IpaddresIPAddressUsers, foreign)
				break
			}
		}
	}

	return nil
}

// AddIpaddresIPAddressUsers adds the given related objects to the existing relationships
// of the ip_address, optionally inserting them as new records.
// Appends related to o.R.IpaddresIPAddressUsers.
// Sets related.R.Ipaddre appropriately.
func (o *IPAddress) AddIpaddresIPAddressUsers(exec boil.Executor, insert bool, related ...*IPAddressUser) error {
	var err error
	for _, rel := range related {
		rel.IpaddressID.Int = o.IpaddressID
		rel.IpaddressID.Valid = true
		if insert {
			if err = rel.Insert(exec); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			if err = rel.Update(exec, "ipaddress_id"); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}
		}
	}

	if o.R == nil {
		o.R = &ipAddressR{
			IpaddresIPAddressUsers: related,
		}
	} else {
		o.R.IpaddresIPAddressUsers = append(o.R.IpaddresIPAddressUsers, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &ipAddressUserR{
				Ipaddre: o,
			}
		} else {
			rel.R.Ipaddre = o
		}
	}
	return nil
}

// SetIpaddresIPAddressUsers removes all previously related items of the
// ip_address replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Ipaddre's IpaddresIPAddressUsers accordingly.
// Replaces o.R.IpaddresIPAddressUsers with related.
// Sets related.R.Ipaddre's IpaddresIPAddressUsers accordingly.
func (o *IPAddress) SetIpaddresIPAddressUsers(exec boil.Executor, insert bool, related ...*IPAddressUser) error {
	query := "update \"ip_address_users\" set \"ipaddress_id\" = null where \"ipaddress_id\" = $1"
	values := []interface{}{o.IpaddressID}
	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err := exec.Exec(query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	if o.R != nil {
		for _, rel := range o.R.IpaddresIPAddressUsers {
			rel.IpaddressID.Valid = false
			if rel.R == nil {
				continue
			}

			rel.R.Ipaddre = nil
		}

		o.R.IpaddresIPAddressUsers = nil
	}
	return o.AddIpaddresIPAddressUsers(exec, insert, related...)
}

// RemoveIpaddresIPAddressUsers relationships from objects passed in.
// Removes related items from R.IpaddresIPAddressUsers (uses pointer comparison, removal does not keep order)
// Sets related.R.Ipaddre.
func (o *IPAddress) RemoveIpaddresIPAddressUsers(exec boil.Executor, related ...*IPAddressUser) error {
	var err error
	for _, rel := range related {
		rel.IpaddressID.Valid = false
		if rel.R != nil {
			rel.R.Ipaddre = nil
		}
		if err = rel.Update(exec, "ipaddress_id"); err != nil {
			return err
		}
	}
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.IpaddresIPAddressUsers {
			if rel != ri {
				continue
			}

			ln := len(o.R.IpaddresIPAddressUsers)
			if ln > 1 && i < ln-1 {
				o.R.IpaddresIPAddressUsers[i] = o.R.IpaddresIPAddressUsers[ln-1]
			}
			o.R.IpaddresIPAddressUsers = o.R.IpaddresIPAddressUsers[:ln-1]
			break
		}
	}

	return nil
}

// IPAddressesG retrieves all records.
func IPAddressesG(mods ...qm.QueryMod) ipAddressQuery {
	return IPAddresses(boil.GetDB(), mods...)
}

// IPAddresses retrieves all the records using an executor.
func IPAddresses(exec boil.Executor, mods ...qm.QueryMod) ipAddressQuery {
	mods = append(mods, qm.From("\"ip_addresses\""))
	return ipAddressQuery{NewQuery(exec, mods...)}
}

// FindIPAddressG retrieves a single record by ID.
func FindIPAddressG(ipaddressID int, selectCols ...string) (*IPAddress, error) {
	return FindIPAddress(boil.GetDB(), ipaddressID, selectCols...)
}

// FindIPAddressGP retrieves a single record by ID, and panics on error.
func FindIPAddressGP(ipaddressID int, selectCols ...string) *IPAddress {
	retobj, err := FindIPAddress(boil.GetDB(), ipaddressID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindIPAddress retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindIPAddress(exec boil.Executor, ipaddressID int, selectCols ...string) (*IPAddress, error) {
	ipAddressObj := &IPAddress{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"ip_addresses\" where \"ipaddress_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, ipaddressID)

	err := q.Bind(ipAddressObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from ip_addresses")
	}

	return ipAddressObj, nil
}

// FindIPAddressP retrieves a single record by ID with an executor, and panics on error.
func FindIPAddressP(exec boil.Executor, ipaddressID int, selectCols ...string) *IPAddress {
	retobj, err := FindIPAddress(exec, ipaddressID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *IPAddress) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *IPAddress) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *IPAddress) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *IPAddress) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no ip_addresses provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(ipAddressColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	ipAddressInsertCacheMut.RLock()
	cache, cached := ipAddressInsertCache[key]
	ipAddressInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			ipAddressColumns,
			ipAddressColumnsWithDefault,
			ipAddressColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(ipAddressType, ipAddressMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(ipAddressType, ipAddressMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"ip_addresses\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

		if len(cache.retMapping) != 0 {
			cache.query += fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into ip_addresses")
	}

	if !cached {
		ipAddressInsertCacheMut.Lock()
		ipAddressInsertCache[key] = cache
		ipAddressInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single IPAddress record. See Update for
// whitelist behavior description.
func (o *IPAddress) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single IPAddress record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *IPAddress) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the IPAddress, and panics on error.
// See Update for whitelist behavior description.
func (o *IPAddress) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the IPAddress.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *IPAddress) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	ipAddressUpdateCacheMut.RLock()
	cache, cached := ipAddressUpdateCache[key]
	ipAddressUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(ipAddressColumns, ipAddressPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update ip_addresses, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"ip_addresses\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, ipAddressPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(ipAddressType, ipAddressMapping, append(wl, ipAddressPrimaryKeyColumns...))
		if err != nil {
			return err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err = exec.Exec(cache.query, values...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update ip_addresses row")
	}

	if !cached {
		ipAddressUpdateCacheMut.Lock()
		ipAddressUpdateCache[key] = cache
		ipAddressUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q ipAddressQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q ipAddressQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for ip_addresses")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o IPAddressSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o IPAddressSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o IPAddressSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o IPAddressSlice) UpdateAll(exec boil.Executor, cols M) error {
	ln := int64(len(o))
	if ln == 0 {
		return nil
	}

	if len(cols) == 0 {
		return errors.New("models: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), ipAddressPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"ip_addresses\" SET %s WHERE (\"ipaddress_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(ipAddressPrimaryKeyColumns), len(colNames)+1, len(ipAddressPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in ipAddress slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *IPAddress) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *IPAddress) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *IPAddress) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *IPAddress) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no ip_addresses provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(ipAddressColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs postgres problems
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
	for _, c := range updateColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range whitelist {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	ipAddressUpsertCacheMut.RLock()
	cache, cached := ipAddressUpsertCache[key]
	ipAddressUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			ipAddressColumns,
			ipAddressColumnsWithDefault,
			ipAddressColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			ipAddressColumns,
			ipAddressPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert ip_addresses, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(ipAddressPrimaryKeyColumns))
			copy(conflict, ipAddressPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"ip_addresses\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(ipAddressType, ipAddressMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(ipAddressType, ipAddressMapping, ret)
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

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(returns...)
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for ip_addresses")
	}

	if !cached {
		ipAddressUpsertCacheMut.Lock()
		ipAddressUpsertCache[key] = cache
		ipAddressUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single IPAddress record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *IPAddress) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single IPAddress record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *IPAddress) DeleteG() error {
	if o == nil {
		return errors.New("models: no IPAddress provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single IPAddress record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *IPAddress) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single IPAddress record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *IPAddress) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no IPAddress provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), ipAddressPrimaryKeyMapping)
	sql := "DELETE FROM \"ip_addresses\" WHERE \"ipaddress_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from ip_addresses")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q ipAddressQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q ipAddressQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no ipAddressQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from ip_addresses")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o IPAddressSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o IPAddressSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no IPAddress slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o IPAddressSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o IPAddressSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no IPAddress slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(ipAddressBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), ipAddressPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"ip_addresses\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ipAddressPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(ipAddressPrimaryKeyColumns), 1, len(ipAddressPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from ipAddress slice")
	}

	if len(ipAddressAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *IPAddress) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *IPAddress) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *IPAddress) ReloadG() error {
	if o == nil {
		return errors.New("models: no IPAddress provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *IPAddress) Reload(exec boil.Executor) error {
	ret, err := FindIPAddress(exec, o.IpaddressID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *IPAddressSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *IPAddressSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *IPAddressSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty IPAddressSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *IPAddressSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	ipAddresses := IPAddressSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), ipAddressPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"ip_addresses\".* FROM \"ip_addresses\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ipAddressPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(ipAddressPrimaryKeyColumns), 1, len(ipAddressPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&ipAddresses)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in IPAddressSlice")
	}

	*o = ipAddresses

	return nil
}

// IPAddressExists checks if the IPAddress row exists.
func IPAddressExists(exec boil.Executor, ipaddressID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"ip_addresses\" where \"ipaddress_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, ipaddressID)
	}

	row := exec.QueryRow(sql, ipaddressID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if ip_addresses exists")
	}

	return exists, nil
}

// IPAddressExistsG checks if the IPAddress row exists.
func IPAddressExistsG(ipaddressID int) (bool, error) {
	return IPAddressExists(boil.GetDB(), ipaddressID)
}

// IPAddressExistsGP checks if the IPAddress row exists. Panics on error.
func IPAddressExistsGP(ipaddressID int) bool {
	e, err := IPAddressExists(boil.GetDB(), ipaddressID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// IPAddressExistsP checks if the IPAddress row exists. Panics on error.
func IPAddressExistsP(exec boil.Executor, ipaddressID int) bool {
	e, err := IPAddressExists(exec, ipaddressID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
