package db

import (
  "github.com/ALTSKUF/ALTSKUF.Back.SquadData/config"
  u "github.com/ALTSKUF/ALTSKUF.Back.SquadData/utils"
  e "github.com/ALTSKUF/ALTSKUF.Back.SquadData/apperror"
  "github.com/stretchr/testify/assert"

  "log"
  "testing"
)

// Init test database
func InitTestDb() Db {
  config := &config.Config {
    DbHost: "localhost",
    DbPort: "5432",
    DbUser: "postgres",
		DbPassword: "mypassword",
    DbName: "testdb",
    DbSSLMode: "disable",
  }

  db, err := Init(config)
  if err != nil {
    log.Fatal(err)
  }

	db.Migrate()

  return db
}

var (
	db = InitTestDb()
)

func TestGetSquadInfoSquadExists(t *testing.T) {
  u.LongTest(t)

  response := db.GetSquadInfo(1)
  
  assert.Condition(t, func () bool {
    return response.Name == "Test 1" && response.Description == "" && response.Error == nil
  })
}

func TestGetSquadInfoSquadNotExists(t *testing.T) {
	u.LongTest(t)

	response := db.GetSquadInfo(10)

	assert.Condition(t, func () bool {
		return response.Name == "" && response.Description == "" &&  response.Error == e.DbSquadNotFoundError
	})
}

func TestGetSquadMembersSquadExists(t *testing.T) {
	u.LongTest(t)

	_, err := db.GetSquadMembers(1)

	assert.Condition(t, func () bool {
		return err == nil
	})
}

func TestGetSquadMembersSquadNotExists(t *testing.T) {
	u.LongTest(t)

	uuids, err := db.GetSquadMembers(10)

	assert.Condition(t, func () bool {
		return len(uuids) == 0 && err != e.DbSquadNotFoundError
	})
}

func TestAllSquads(t *testing.T) {
	u.LongTest(t)

	squads, err := db.GetAllSquads()

	assert.Condition(t, func () bool {
		return len(squads) == 3 && err == nil
	})
}
