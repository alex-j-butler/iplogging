package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("SchemaMigrations", testSchemaMigrations)
	t.Run("IPAddresses", testIPAddresses)
	t.Run("Users", testUsers)
	t.Run("IPAddressUsers", testIPAddressUsers)
}

func TestDelete(t *testing.T) {
	t.Run("SchemaMigrations", testSchemaMigrationsDelete)
	t.Run("IPAddresses", testIPAddressesDelete)
	t.Run("Users", testUsersDelete)
	t.Run("IPAddressUsers", testIPAddressUsersDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("SchemaMigrations", testSchemaMigrationsQueryDeleteAll)
	t.Run("IPAddresses", testIPAddressesQueryDeleteAll)
	t.Run("Users", testUsersQueryDeleteAll)
	t.Run("IPAddressUsers", testIPAddressUsersQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("SchemaMigrations", testSchemaMigrationsSliceDeleteAll)
	t.Run("IPAddresses", testIPAddressesSliceDeleteAll)
	t.Run("Users", testUsersSliceDeleteAll)
	t.Run("IPAddressUsers", testIPAddressUsersSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("SchemaMigrations", testSchemaMigrationsExists)
	t.Run("IPAddresses", testIPAddressesExists)
	t.Run("Users", testUsersExists)
	t.Run("IPAddressUsers", testIPAddressUsersExists)
}

func TestFind(t *testing.T) {
	t.Run("SchemaMigrations", testSchemaMigrationsFind)
	t.Run("IPAddresses", testIPAddressesFind)
	t.Run("Users", testUsersFind)
	t.Run("IPAddressUsers", testIPAddressUsersFind)
}

func TestBind(t *testing.T) {
	t.Run("SchemaMigrations", testSchemaMigrationsBind)
	t.Run("IPAddresses", testIPAddressesBind)
	t.Run("Users", testUsersBind)
	t.Run("IPAddressUsers", testIPAddressUsersBind)
}

func TestOne(t *testing.T) {
	t.Run("SchemaMigrations", testSchemaMigrationsOne)
	t.Run("IPAddresses", testIPAddressesOne)
	t.Run("Users", testUsersOne)
	t.Run("IPAddressUsers", testIPAddressUsersOne)
}

func TestAll(t *testing.T) {
	t.Run("SchemaMigrations", testSchemaMigrationsAll)
	t.Run("IPAddresses", testIPAddressesAll)
	t.Run("Users", testUsersAll)
	t.Run("IPAddressUsers", testIPAddressUsersAll)
}

func TestCount(t *testing.T) {
	t.Run("SchemaMigrations", testSchemaMigrationsCount)
	t.Run("IPAddresses", testIPAddressesCount)
	t.Run("Users", testUsersCount)
	t.Run("IPAddressUsers", testIPAddressUsersCount)
}

func TestHooks(t *testing.T) {
	t.Run("SchemaMigrations", testSchemaMigrationsHooks)
	t.Run("IPAddresses", testIPAddressesHooks)
	t.Run("Users", testUsersHooks)
	t.Run("IPAddressUsers", testIPAddressUsersHooks)
}

func TestInsert(t *testing.T) {
	t.Run("SchemaMigrations", testSchemaMigrationsInsert)
	t.Run("SchemaMigrations", testSchemaMigrationsInsertWhitelist)
	t.Run("IPAddresses", testIPAddressesInsert)
	t.Run("IPAddresses", testIPAddressesInsertWhitelist)
	t.Run("Users", testUsersInsert)
	t.Run("Users", testUsersInsertWhitelist)
	t.Run("IPAddressUsers", testIPAddressUsersInsert)
	t.Run("IPAddressUsers", testIPAddressUsersInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("IPAddressUserToIPAddressUsingIpaddres", testIPAddressUserToOneIPAddressUsingIpaddres)
	t.Run("IPAddressUserToUserUsingUser", testIPAddressUserToOneUserUsingUser)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("IPAddressToIpaddresIPAddressUsers", testIPAddressToManyIpaddresIPAddressUsers)
	t.Run("UserToIPAddressUsers", testUserToManyIPAddressUsers)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("IPAddressUserToIPAddressUsingIpaddres", testIPAddressUserToOneSetOpIPAddressUsingIpaddres)
	t.Run("IPAddressUserToUserUsingUser", testIPAddressUserToOneSetOpUserUsingUser)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {
	t.Run("IPAddressUserToIPAddressUsingIpaddres", testIPAddressUserToOneRemoveOpIPAddressUsingIpaddres)
	t.Run("IPAddressUserToUserUsingUser", testIPAddressUserToOneRemoveOpUserUsingUser)
}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("IPAddressToIpaddresIPAddressUsers", testIPAddressToManyAddOpIpaddresIPAddressUsers)
	t.Run("UserToIPAddressUsers", testUserToManyAddOpIPAddressUsers)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {
	t.Run("IPAddressToIpaddresIPAddressUsers", testIPAddressToManySetOpIpaddresIPAddressUsers)
	t.Run("UserToIPAddressUsers", testUserToManySetOpIPAddressUsers)
}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {
	t.Run("IPAddressToIpaddresIPAddressUsers", testIPAddressToManyRemoveOpIpaddresIPAddressUsers)
	t.Run("UserToIPAddressUsers", testUserToManyRemoveOpIPAddressUsers)
}

func TestReload(t *testing.T) {
	t.Run("SchemaMigrations", testSchemaMigrationsReload)
	t.Run("IPAddresses", testIPAddressesReload)
	t.Run("Users", testUsersReload)
	t.Run("IPAddressUsers", testIPAddressUsersReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("SchemaMigrations", testSchemaMigrationsReloadAll)
	t.Run("IPAddresses", testIPAddressesReloadAll)
	t.Run("Users", testUsersReloadAll)
	t.Run("IPAddressUsers", testIPAddressUsersReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("SchemaMigrations", testSchemaMigrationsSelect)
	t.Run("IPAddresses", testIPAddressesSelect)
	t.Run("Users", testUsersSelect)
	t.Run("IPAddressUsers", testIPAddressUsersSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("SchemaMigrations", testSchemaMigrationsUpdate)
	t.Run("IPAddresses", testIPAddressesUpdate)
	t.Run("Users", testUsersUpdate)
	t.Run("IPAddressUsers", testIPAddressUsersUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("SchemaMigrations", testSchemaMigrationsSliceUpdateAll)
	t.Run("IPAddresses", testIPAddressesSliceUpdateAll)
	t.Run("Users", testUsersSliceUpdateAll)
	t.Run("IPAddressUsers", testIPAddressUsersSliceUpdateAll)
}

func TestUpsert(t *testing.T) {
	t.Run("SchemaMigrations", testSchemaMigrationsUpsert)
	t.Run("IPAddresses", testIPAddressesUpsert)
	t.Run("Users", testUsersUpsert)
	t.Run("IPAddressUsers", testIPAddressUsersUpsert)
}
