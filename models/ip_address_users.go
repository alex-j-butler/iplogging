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

// IPAddressUser is an object representing the database table.
type IPAddressUser struct {
	RefID          int       `boil:"ref_id" json:"ref_id" toml:"ref_id" yaml:"ref_id"`
	IpaddressID    null.Int  `boil:"ipaddress_id" json:"ipaddress_id,omitempty" toml:"ipaddress_id" yaml:"ipaddress_id,omitempty"`
	UserID         null.Int  `boil:"user_id" json:"user_id,omitempty" toml:"user_id" yaml:"user_id,omitempty"`
	ConnectionTime null.Time `boil:"connection_time" json:"connection_time,omitempty" toml:"connection_time" yaml:"connection_time,omitempty"`

	R *ipAddressUserR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L ipAddressUserL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// ipAddressUserR is where relationships are stored.
type ipAddressUserR struct {
	Ipaddres *IPAddress
	User     *User
}

// ipAddressUserL is where Load methods for each relationship are stored.
type ipAddressUserL struct{}

var (
	ipAddressUserColumns               = []string{"ref_id", "ipaddress_id", "user_id", "connection_time"}
	ipAddressUserColumnsWithoutDefault = []string{"ipaddress_id", "user_id", "connection_time"}
	ipAddressUserColumnsWithDefault    = []string{"ref_id"}
	ipAddressUserPrimaryKeyColumns     = []string{"ref_id"}
)

type (
	// IPAddressUserSlice is an alias for a slice of pointers to IPAddressUser.
	// This should generally be used opposed to []IPAddressUser.
	IPAddressUserSlice []*IPAddressUser
	// IPAddressUserHook is the signature for custom IPAddressUser hook methods
	IPAddressUserHook func(boil.Executor, *IPAddressUser) error

	ipAddressUserQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	ipAddressUserType                 = reflect.TypeOf(&IPAddressUser{})
	ipAddressUserMapping              = queries.MakeStructMapping(ipAddressUserType)
	ipAddressUserPrimaryKeyMapping, _ = queries.BindMapping(ipAddressUserType, ipAddressUserMapping, ipAddressUserPrimaryKeyColumns)
	ipAddressUserInsertCacheMut       sync.RWMutex
	ipAddressUserInsertCache          = make(map[string]insertCache)
	ipAddressUserUpdateCacheMut       sync.RWMutex
	ipAddressUserUpdateCache          = make(map[string]updateCache)
	ipAddressUserUpsertCacheMut       sync.RWMutex
	ipAddressUserUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var ipAddressUserBeforeInsertHooks []IPAddressUserHook
var ipAddressUserBeforeUpdateHooks []IPAddressUserHook
var ipAddressUserBeforeDeleteHooks []IPAddressUserHook
var ipAddressUserBeforeUpsertHooks []IPAddressUserHook

var ipAddressUserAfterInsertHooks []IPAddressUserHook
var ipAddressUserAfterSelectHooks []IPAddressUserHook
var ipAddressUserAfterUpdateHooks []IPAddressUserHook
var ipAddressUserAfterDeleteHooks []IPAddressUserHook
var ipAddressUserAfterUpsertHooks []IPAddressUserHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *IPAddressUser) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range ipAddressUserBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *IPAddressUser) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range ipAddressUserBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *IPAddressUser) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range ipAddressUserBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *IPAddressUser) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range ipAddressUserBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *IPAddressUser) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range ipAddressUserAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *IPAddressUser) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range ipAddressUserAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *IPAddressUser) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range ipAddressUserAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *IPAddressUser) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range ipAddressUserAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *IPAddressUser) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range ipAddressUserAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddIPAddressUserHook registers your hook function for all future operations.
func AddIPAddressUserHook(hookPoint boil.HookPoint, ipAddressUserHook IPAddressUserHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		ipAddressUserBeforeInsertHooks = append(ipAddressUserBeforeInsertHooks, ipAddressUserHook)
	case boil.BeforeUpdateHook:
		ipAddressUserBeforeUpdateHooks = append(ipAddressUserBeforeUpdateHooks, ipAddressUserHook)
	case boil.BeforeDeleteHook:
		ipAddressUserBeforeDeleteHooks = append(ipAddressUserBeforeDeleteHooks, ipAddressUserHook)
	case boil.BeforeUpsertHook:
		ipAddressUserBeforeUpsertHooks = append(ipAddressUserBeforeUpsertHooks, ipAddressUserHook)
	case boil.AfterInsertHook:
		ipAddressUserAfterInsertHooks = append(ipAddressUserAfterInsertHooks, ipAddressUserHook)
	case boil.AfterSelectHook:
		ipAddressUserAfterSelectHooks = append(ipAddressUserAfterSelectHooks, ipAddressUserHook)
	case boil.AfterUpdateHook:
		ipAddressUserAfterUpdateHooks = append(ipAddressUserAfterUpdateHooks, ipAddressUserHook)
	case boil.AfterDeleteHook:
		ipAddressUserAfterDeleteHooks = append(ipAddressUserAfterDeleteHooks, ipAddressUserHook)
	case boil.AfterUpsertHook:
		ipAddressUserAfterUpsertHooks = append(ipAddressUserAfterUpsertHooks, ipAddressUserHook)
	}
}

// OneP returns a single ipAddressUser record from the query, and panics on error.
func (q ipAddressUserQuery) OneP() *IPAddressUser {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single ipAddressUser record from the query.
func (q ipAddressUserQuery) One() (*IPAddressUser, error) {
	o := &IPAddressUser{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for ip_address_users")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all IPAddressUser records from the query, and panics on error.
func (q ipAddressUserQuery) AllP() IPAddressUserSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all IPAddressUser records from the query.
func (q ipAddressUserQuery) All() (IPAddressUserSlice, error) {
	var o IPAddressUserSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to IPAddressUser slice")
	}

	if len(ipAddressUserAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all IPAddressUser records in the query, and panics on error.
func (q ipAddressUserQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all IPAddressUser records in the query.
func (q ipAddressUserQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count ip_address_users rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q ipAddressUserQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q ipAddressUserQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if ip_address_users exists")
	}

	return count > 0, nil
}

// IpaddresG pointed to by the foreign key.
func (o *IPAddressUser) IpaddresG(mods ...qm.QueryMod) ipAddressQuery {
	return o.Ipaddres(boil.GetDB(), mods...)
}

// Ipaddres pointed to by the foreign key.
func (o *IPAddressUser) Ipaddres(exec boil.Executor, mods ...qm.QueryMod) ipAddressQuery {
	queryMods := []qm.QueryMod{
		qm.Where("ipaddress_id=?", o.IpaddressID),
	}

	queryMods = append(queryMods, mods...)

	query := IPAddresses(exec, queryMods...)
	queries.SetFrom(query.Query, "\"ip_addresses\"")

	return query
}

// UserG pointed to by the foreign key.
func (o *IPAddressUser) UserG(mods ...qm.QueryMod) userQuery {
	return o.User(boil.GetDB(), mods...)
}

// User pointed to by the foreign key.
func (o *IPAddressUser) User(exec boil.Executor, mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("user_id=?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	query := Users(exec, queryMods...)
	queries.SetFrom(query.Query, "\"users\"")

	return query
}

// LoadIpaddres allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (ipAddressUserL) LoadIpaddres(e boil.Executor, singular bool, maybeIPAddressUser interface{}) error {
	var slice []*IPAddressUser
	var object *IPAddressUser

	count := 1
	if singular {
		object = maybeIPAddressUser.(*IPAddressUser)
	} else {
		slice = *maybeIPAddressUser.(*IPAddressUserSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &ipAddressUserR{}
		}
		args[0] = object.IpaddressID
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &ipAddressUserR{}
			}
			args[i] = obj.IpaddressID
		}
	}

	query := fmt.Sprintf(
		"select * from \"ip_addresses\" where \"ipaddress_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load IPAddress")
	}
	defer results.Close()

	var resultSlice []*IPAddress
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice IPAddress")
	}

	if len(ipAddressUserAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Ipaddres = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.IpaddressID.Int == foreign.IpaddressID {
				local.R.Ipaddres = foreign
				break
			}
		}
	}

	return nil
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (ipAddressUserL) LoadUser(e boil.Executor, singular bool, maybeIPAddressUser interface{}) error {
	var slice []*IPAddressUser
	var object *IPAddressUser

	count := 1
	if singular {
		object = maybeIPAddressUser.(*IPAddressUser)
	} else {
		slice = *maybeIPAddressUser.(*IPAddressUserSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &ipAddressUserR{}
		}
		args[0] = object.UserID
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &ipAddressUserR{}
			}
			args[i] = obj.UserID
		}
	}

	query := fmt.Sprintf(
		"select * from \"users\" where \"user_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load User")
	}
	defer results.Close()

	var resultSlice []*User
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice User")
	}

	if len(ipAddressUserAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.User = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.UserID.Int == foreign.UserID {
				local.R.User = foreign
				break
			}
		}
	}

	return nil
}

// SetIpaddres of the ip_address_user to the related item.
// Sets o.R.Ipaddres to related.
// Adds o to related.R.IpaddresIPAddressUsers.
func (o *IPAddressUser) SetIpaddres(exec boil.Executor, insert bool, related *IPAddress) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"ip_address_users\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"ipaddress_id"}),
		strmangle.WhereClause("\"", "\"", 2, ipAddressUserPrimaryKeyColumns),
	)
	values := []interface{}{related.IpaddressID, o.RefID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.IpaddressID.Int = related.IpaddressID
	o.IpaddressID.Valid = true

	if o.R == nil {
		o.R = &ipAddressUserR{
			Ipaddres: related,
		}
	} else {
		o.R.Ipaddres = related
	}

	if related.R == nil {
		related.R = &ipAddressR{
			IpaddresIPAddressUsers: IPAddressUserSlice{o},
		}
	} else {
		related.R.IpaddresIPAddressUsers = append(related.R.IpaddresIPAddressUsers, o)
	}

	return nil
}

// RemoveIpaddres relationship.
// Sets o.R.Ipaddres to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *IPAddressUser) RemoveIpaddres(exec boil.Executor, related *IPAddress) error {
	var err error

	o.IpaddressID.Valid = false
	if err = o.Update(exec, "ipaddress_id"); err != nil {
		o.IpaddressID.Valid = true
		return errors.Wrap(err, "failed to update local table")
	}

	o.R.Ipaddres = nil
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.IpaddresIPAddressUsers {
		if o.IpaddressID.Int != ri.IpaddressID.Int {
			continue
		}

		ln := len(related.R.IpaddresIPAddressUsers)
		if ln > 1 && i < ln-1 {
			related.R.IpaddresIPAddressUsers[i] = related.R.IpaddresIPAddressUsers[ln-1]
		}
		related.R.IpaddresIPAddressUsers = related.R.IpaddresIPAddressUsers[:ln-1]
		break
	}
	return nil
}

// SetUser of the ip_address_user to the related item.
// Sets o.R.User to related.
// Adds o to related.R.IPAddressUsers.
func (o *IPAddressUser) SetUser(exec boil.Executor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"ip_address_users\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"user_id"}),
		strmangle.WhereClause("\"", "\"", 2, ipAddressUserPrimaryKeyColumns),
	)
	values := []interface{}{related.UserID, o.RefID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.UserID.Int = related.UserID
	o.UserID.Valid = true

	if o.R == nil {
		o.R = &ipAddressUserR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userR{
			IPAddressUsers: IPAddressUserSlice{o},
		}
	} else {
		related.R.IPAddressUsers = append(related.R.IPAddressUsers, o)
	}

	return nil
}

// RemoveUser relationship.
// Sets o.R.User to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *IPAddressUser) RemoveUser(exec boil.Executor, related *User) error {
	var err error

	o.UserID.Valid = false
	if err = o.Update(exec, "user_id"); err != nil {
		o.UserID.Valid = true
		return errors.Wrap(err, "failed to update local table")
	}

	o.R.User = nil
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.IPAddressUsers {
		if o.UserID.Int != ri.UserID.Int {
			continue
		}

		ln := len(related.R.IPAddressUsers)
		if ln > 1 && i < ln-1 {
			related.R.IPAddressUsers[i] = related.R.IPAddressUsers[ln-1]
		}
		related.R.IPAddressUsers = related.R.IPAddressUsers[:ln-1]
		break
	}
	return nil
}

// IPAddressUsersG retrieves all records.
func IPAddressUsersG(mods ...qm.QueryMod) ipAddressUserQuery {
	return IPAddressUsers(boil.GetDB(), mods...)
}

// IPAddressUsers retrieves all the records using an executor.
func IPAddressUsers(exec boil.Executor, mods ...qm.QueryMod) ipAddressUserQuery {
	mods = append(mods, qm.From("\"ip_address_users\""))
	return ipAddressUserQuery{NewQuery(exec, mods...)}
}

// FindIPAddressUserG retrieves a single record by ID.
func FindIPAddressUserG(refID int, selectCols ...string) (*IPAddressUser, error) {
	return FindIPAddressUser(boil.GetDB(), refID, selectCols...)
}

// FindIPAddressUserGP retrieves a single record by ID, and panics on error.
func FindIPAddressUserGP(refID int, selectCols ...string) *IPAddressUser {
	retobj, err := FindIPAddressUser(boil.GetDB(), refID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindIPAddressUser retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindIPAddressUser(exec boil.Executor, refID int, selectCols ...string) (*IPAddressUser, error) {
	ipAddressUserObj := &IPAddressUser{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"ip_address_users\" where \"ref_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, refID)

	err := q.Bind(ipAddressUserObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from ip_address_users")
	}

	return ipAddressUserObj, nil
}

// FindIPAddressUserP retrieves a single record by ID with an executor, and panics on error.
func FindIPAddressUserP(exec boil.Executor, refID int, selectCols ...string) *IPAddressUser {
	retobj, err := FindIPAddressUser(exec, refID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *IPAddressUser) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *IPAddressUser) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *IPAddressUser) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *IPAddressUser) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no ip_address_users provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(ipAddressUserColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	ipAddressUserInsertCacheMut.RLock()
	cache, cached := ipAddressUserInsertCache[key]
	ipAddressUserInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			ipAddressUserColumns,
			ipAddressUserColumnsWithDefault,
			ipAddressUserColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(ipAddressUserType, ipAddressUserMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(ipAddressUserType, ipAddressUserMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"ip_address_users\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "models: unable to insert into ip_address_users")
	}

	if !cached {
		ipAddressUserInsertCacheMut.Lock()
		ipAddressUserInsertCache[key] = cache
		ipAddressUserInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single IPAddressUser record. See Update for
// whitelist behavior description.
func (o *IPAddressUser) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single IPAddressUser record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *IPAddressUser) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the IPAddressUser, and panics on error.
// See Update for whitelist behavior description.
func (o *IPAddressUser) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the IPAddressUser.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *IPAddressUser) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	ipAddressUserUpdateCacheMut.RLock()
	cache, cached := ipAddressUserUpdateCache[key]
	ipAddressUserUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(ipAddressUserColumns, ipAddressUserPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update ip_address_users, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"ip_address_users\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, ipAddressUserPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(ipAddressUserType, ipAddressUserMapping, append(wl, ipAddressUserPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update ip_address_users row")
	}

	if !cached {
		ipAddressUserUpdateCacheMut.Lock()
		ipAddressUserUpdateCache[key] = cache
		ipAddressUserUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q ipAddressUserQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q ipAddressUserQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for ip_address_users")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o IPAddressUserSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o IPAddressUserSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o IPAddressUserSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o IPAddressUserSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), ipAddressUserPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"ip_address_users\" SET %s WHERE (\"ref_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(ipAddressUserPrimaryKeyColumns), len(colNames)+1, len(ipAddressUserPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in ipAddressUser slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *IPAddressUser) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *IPAddressUser) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *IPAddressUser) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *IPAddressUser) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no ip_address_users provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(ipAddressUserColumnsWithDefault, o)

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

	ipAddressUserUpsertCacheMut.RLock()
	cache, cached := ipAddressUserUpsertCache[key]
	ipAddressUserUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			ipAddressUserColumns,
			ipAddressUserColumnsWithDefault,
			ipAddressUserColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			ipAddressUserColumns,
			ipAddressUserPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert ip_address_users, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(ipAddressUserPrimaryKeyColumns))
			copy(conflict, ipAddressUserPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"ip_address_users\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(ipAddressUserType, ipAddressUserMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(ipAddressUserType, ipAddressUserMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for ip_address_users")
	}

	if !cached {
		ipAddressUserUpsertCacheMut.Lock()
		ipAddressUserUpsertCache[key] = cache
		ipAddressUserUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single IPAddressUser record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *IPAddressUser) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single IPAddressUser record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *IPAddressUser) DeleteG() error {
	if o == nil {
		return errors.New("models: no IPAddressUser provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single IPAddressUser record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *IPAddressUser) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single IPAddressUser record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *IPAddressUser) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no IPAddressUser provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), ipAddressUserPrimaryKeyMapping)
	sql := "DELETE FROM \"ip_address_users\" WHERE \"ref_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from ip_address_users")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q ipAddressUserQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q ipAddressUserQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no ipAddressUserQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from ip_address_users")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o IPAddressUserSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o IPAddressUserSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no IPAddressUser slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o IPAddressUserSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o IPAddressUserSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no IPAddressUser slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(ipAddressUserBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), ipAddressUserPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"ip_address_users\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ipAddressUserPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(ipAddressUserPrimaryKeyColumns), 1, len(ipAddressUserPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from ipAddressUser slice")
	}

	if len(ipAddressUserAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *IPAddressUser) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *IPAddressUser) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *IPAddressUser) ReloadG() error {
	if o == nil {
		return errors.New("models: no IPAddressUser provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *IPAddressUser) Reload(exec boil.Executor) error {
	ret, err := FindIPAddressUser(exec, o.RefID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *IPAddressUserSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *IPAddressUserSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *IPAddressUserSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty IPAddressUserSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *IPAddressUserSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	ipAddressUsers := IPAddressUserSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), ipAddressUserPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"ip_address_users\".* FROM \"ip_address_users\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ipAddressUserPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(ipAddressUserPrimaryKeyColumns), 1, len(ipAddressUserPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&ipAddressUsers)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in IPAddressUserSlice")
	}

	*o = ipAddressUsers

	return nil
}

// IPAddressUserExists checks if the IPAddressUser row exists.
func IPAddressUserExists(exec boil.Executor, refID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"ip_address_users\" where \"ref_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, refID)
	}

	row := exec.QueryRow(sql, refID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if ip_address_users exists")
	}

	return exists, nil
}

// IPAddressUserExistsG checks if the IPAddressUser row exists.
func IPAddressUserExistsG(refID int) (bool, error) {
	return IPAddressUserExists(boil.GetDB(), refID)
}

// IPAddressUserExistsGP checks if the IPAddressUser row exists. Panics on error.
func IPAddressUserExistsGP(refID int) bool {
	e, err := IPAddressUserExists(boil.GetDB(), refID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// IPAddressUserExistsP checks if the IPAddressUser row exists. Panics on error.
func IPAddressUserExistsP(exec boil.Executor, refID int) bool {
	e, err := IPAddressUserExists(exec, refID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
