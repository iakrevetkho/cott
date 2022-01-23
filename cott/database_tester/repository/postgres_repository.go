package repository

import (
	"bytes"
	"database/sql"
	"strconv"

	"github.com/iakrevetkho/components-tests/cott/domain"

	_ "github.com/lib/pq"
)

type postgresDatabaseTesterRepository struct {
	db       *sql.DB
	port     uint16
	host     string
	user     string
	password string
	dbname   string
}

func NewPostgresDatabaseTesterRepository(port uint16, host, user, password, dbname string) DatabaseTesterRepository {
	r := new(postgresDatabaseTesterRepository)
	r.port = port
	r.host = host
	r.user = user
	r.password = password
	r.dbname = dbname
	return r
}

func (r *postgresDatabaseTesterRepository) Open() error {
	var err error
	r.db, err = sql.Open("postgres", r.createConnString(r.port, r.host, r.user, r.password, r.dbname))
	if err != nil {
		return err
	}

	return nil
}

func (r *postgresDatabaseTesterRepository) CreateDatabase(name string) error {
	if r.db == nil {
		return domain.CONNECTION_WAS_NOT_ESTABLISHED
	}

	var buf bytes.Buffer
	buf.WriteString("CREATE DATABASE ")
	buf.WriteString(name)

	_, err := r.db.Exec(buf.String())
	if err != nil {
		return err
	}

	return nil
}

func (r *postgresDatabaseTesterRepository) DropDatabase(name string) error {
	if r.db == nil {
		return domain.CONNECTION_WAS_NOT_ESTABLISHED
	}

	var buf bytes.Buffer
	buf.WriteString("DROP DATABASE ")
	buf.WriteString(name)

	_, err := r.db.Exec(buf.String())
	if err != nil {
		return err
	}

	return nil
}

func (r *postgresDatabaseTesterRepository) SwitchDatabase(name string) error {
	if r.db == nil {
		return domain.CONNECTION_WAS_NOT_ESTABLISHED
	}

	r.Close()

	r.dbname = name

	r.Open()

	return nil
}

func (r *postgresDatabaseTesterRepository) CreateTable(name string) error {
	if r.db == nil {
		return domain.CONNECTION_WAS_NOT_ESTABLISHED
	}

	var buf bytes.Buffer
	buf.WriteString("CREATE TABLE ")
	buf.WriteString(name)
	buf.WriteString(" (")
	buf.WriteString("id SERIAL PRIMARY KEY")
	buf.WriteString(");")

	_, err := r.db.Exec(buf.String())
	if err != nil {
		return err
	}

	return nil
}

func (r *postgresDatabaseTesterRepository) DropTable(name string) error {
	if r.db == nil {
		return domain.CONNECTION_WAS_NOT_ESTABLISHED
	}

	var buf bytes.Buffer
	buf.WriteString("DROP TABLE ")
	buf.WriteString(name)

	_, err := r.db.Exec(buf.String())
	if err != nil {
		return err
	}

	return nil
}

func (r *postgresDatabaseTesterRepository) Close() error {
	if r.db == nil {
		return domain.CONNECTION_WAS_NOT_ESTABLISHED
	}

	if err := r.db.Close(); err != nil {
		return err
	}

	r.db = nil

	return nil
}

func (r *postgresDatabaseTesterRepository) createConnString(port uint16, host, user, password, dbname string) string {
	var buf bytes.Buffer

	buf.WriteString("host=")
	buf.WriteString(host)
	buf.WriteString(" port=")
	buf.WriteString(strconv.FormatUint(uint64(port), 10))
	buf.WriteString(" user=")
	buf.WriteString(user)
	buf.WriteString(" password=")
	buf.WriteString(password)
	if dbname != "" {
		buf.WriteString(" dbname=")
		buf.WriteString(dbname)
	}
	buf.WriteString(" sslmode=disable")

	return buf.String()
}
