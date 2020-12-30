//
//  integration-test-golang
//
//  Copyright © 2020. All rights reserved.
//

package cassandra_test

import (
	"log"
	"os"
	"strconv"
	"testing"

	r "github.com/moemoe89/integration-test-golang/repository"
	"github.com/moemoe89/integration-test-golang/repository/cassandra"

	"github.com/gocql/gocql"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/stretchr/testify/assert"
)

var (
	repo r.Repository
)

var u = &r.UserModel{
	ID:    gocql.TimeUUID().String(),
	Name:  "Momo",
	Email: "momo@mail.com",
	Phone: "08123456789",
}

var (
	addr     = "localhost"
	port     = 9042
	username = "cassandra"
	password = "cassandra"
)

func TestMain(m *testing.M) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	opts := dockertest.RunOptions{
		Repository:   "bitnami/cassandra",
		Tag:          "latest",
		ExposedPorts: []string{"9042"},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"9042": {
				{HostIP: "0.0.0.0", HostPort: strconv.Itoa(port)},
			},
		},
	}

	resource, err := pool.RunWithOptions(&opts)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err.Error())
	}

	if err = pool.Retry(func() error {
		repo, err = cassandra.NewRepository(addr, port, username, password)
		return err
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err.Error())
	}

	defer func() {
		repo.Close()
	}()

	err = repo.Drop()
	if err != nil {
		panic(err)
	}

	err = repo.Up()
	if err != nil {
		panic(err)
	}

	code := m.Run()

	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}

func TestCreate(t *testing.T) {
	err := repo.Create(u)
	assert.NoError(t, err)
}

func TestFindByID(t *testing.T) {
	res, err := repo.FindByID(u.ID)
	assert.NotEmpty(t, res)
	assert.NoError(t, err)
	assert.Equal(t, u.ID, res.ID)
	assert.Equal(t, u.Name, res.Name)
	assert.Equal(t, u.Email, res.Email)
	assert.Equal(t, u.Phone, res.Phone)
}

func TestFind(t *testing.T) {
	users, err := repo.Find()
	assert.NotEmpty(t, users)
	assert.NoError(t, err)
	assert.NotEqual(t, 0, len(users))
}

func TestUpdate(t *testing.T) {
	err := repo.Update(u)
	assert.NoError(t, err)
}

func TestDelete(t *testing.T) {
	err := repo.Delete(u.ID)
	assert.NoError(t, err)
}

