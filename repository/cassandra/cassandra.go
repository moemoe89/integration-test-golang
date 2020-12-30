//
//  integration-test-golang
//
//  Copyright © 2020. All rights reserved.
//

package cassandra

import (
	repo "github.com/moemoe89/integration-test-golang/repository"
	"github.com/gocql/gocql"
)

// repository represent the repository model
type repository struct {
	session *gocql.Session
}

// NewRepository will create a variable that represent the Repository struct
func NewRepository(addr string, port int, username, password string) (repo.Repository, error) {
	cluster := gocql.NewCluster(addr)
	cluster.Port = port
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: username,
		Password: password,
	}
	sess, err := cluster.CreateSession()
	if err != nil {
		return nil, err
	}

	return &repository{sess}, nil
}

// Close attaches the provider and close the connection
func (r *repository) Close() {
	r.session.Close()
}

// Up attaches the provider and create the table
func (r *repository) Up() error {
	err := r.session.Query("CREATE KEYSPACE IF NOT EXISTS example WITH replication = {'class':'SimpleStrategy','replication_factor':'1'};").Exec()
	if err != nil {
		return err
	}

	query :=
		"CREATE TABLE IF NOT EXISTS example.users (" +
			"id UUID," +
			"name text," +
			"email text," +
			"phone text," +
			"PRIMARY KEY(id)" +
		");"

	return r.session.Query(query).Exec()
}

// Drop attaches the provider and drop the table
func (r *repository) Drop() error {
	err := r.session.Query("DROP KEYSPACE IF EXISTS example;").Exec()
	if err != nil {
		return err
	}

	return r.session.Query("DROP TABLE IF EXISTS example.users;").Exec()
}

// FindByID attaches the user repository and find data based on id
func (r *repository) FindByID(id string) (*repo.UserModel, error) {
	uuid, _ := gocql.ParseUUID(id)
	user := &repo.UserModel{}
	err := r.session.Query(`SELECT id, name, email, phone FROM example.users WHERE id = ? LIMIT 1`,
		uuid).Consistency(gocql.One).Scan(&user.ID, &user.Name, &user.Email, &user.Phone)
	return user, err
}

// Find attaches the user repository and find all data
func (r *repository) Find() ([]*repo.UserModel, error) {
	users := []*repo.UserModel{}
	user := &repo.UserModel{}
	iter := r.session.Query("SELECT id, name, email, phone FROM example.users").Iter()
	for iter.Scan(&user.ID, &user.Name, &user.Email, &user.Phone) {
		users = append(users, user)
	}

	if err := iter.Close(); err != nil {
		return nil, err
	}
	return users, nil
}

// Create attaches the user repository and creating the data
func (r *repository) Create(user *repo.UserModel) error {
	err := r.session.Query(`INSERT INTO example.users (id, name, email, phone) VALUES (?, ?, ?, ?)`, user.ID, user.Name, user.Email, user.Phone).Exec()
	return err
}

// Update attaches the user repository and update data based on id
func (r *repository) Update(user *repo.UserModel) error {
	uuid, _ := gocql.ParseUUID(user.ID)
	err := r.session.Query(`UPDATE example.users SET name = ?, email = ?, phone = ? WHERE id = ?`, user.Name, user.Email, user.Phone, uuid).Exec()
	return err
}

// Delete attaches the user repository and delete data based on id
func (r *repository) Delete(id string) error {
	uuid, _ := gocql.ParseUUID(id)
	err := r.session.Query(`DELETE FROM example.users WHERE id = ?`, uuid).Exec()
	return err
}
