package auth_db

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"math/big"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"

	errormsgs "zk_auth_service/pkg/domain/errors"
	"zk_auth_service/pkg/infra/env"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	fs       = flag.NewFlagSet("database", flag.ExitOnError)
	host     = fs.String("postgres-host", env.WithDefaultString("PGHOST", "localhost"), "postgres connection string")
	port     = fs.String("postgres-port", env.WithDefaultString("PGPORT", "5432"), "postgres port")
	dbname   = fs.String("postgres-db", env.WithDefaultString("ZK_AUTH_DB_NAME", "zk_auth"), "postgres database name")
	user     = fs.String("postgres-user", env.WithDefaultString("ZK_AUTH_DB_USER", "zk_auth_user"), "postgres user")
	password = fs.String("postgres-password", env.WithDefaultString("ZK_AUTH_DB_PASSWORD", ""), "postgres user password")
	ssl      = fs.String("postgres-ssl", env.WithDefaultString("PGSSLMODE", "disable"), "postgres ssl mode")
	debug    = fs.String("postgres-debug", env.WithDefaultString("POSTGRES_DEBUG", "true"), "postgres debug mode")
)

const (
	maxConnections int = 20
)

type authDB struct {
	pool *pgxpool.Pool
}

func NewAuthDB() *authDB {
	return &authDB{}
}

func (db *authDB) Connect() error {
	log.Trace("AuthDB connect")

	conf, err := pgxpool.ParseConfig(db.getConnectionString())
	if err != nil {
		log.Error(fmt.Printf("Unable to parse database connection string: %v \n", err))
		return err
	}

	// tell the driver to sanitize statement strings
	conf.ConnConfig.RuntimeParams = map[string]string{
		"standard_conforming_strings": "on",
		"timezone":                    "UTC",
	}
	// no statement preparation call, whole queries within a single network call
	conf.ConnConfig.PreferSimpleProtocol = true
	if *debug == "true" {
		conf.ConnConfig.LogLevel = pgx.LogLevelDebug
	} else {
		conf.ConnConfig.LogLevel = pgx.LogLevelError
	}

	ctx := context.Background()
	db.pool, err = pgxpool.ConnectConfig(ctx, conf)
	if err != nil {
		log.Error(fmt.Printf("Unable to create database connection pool: %v \n", err))
	}
	// reconnect is handled by the driver
	return err
}

func (db *authDB) Close() {
	log.Trace("AuthDB Close")
	db.pool.Close()
}

func (db *authDB) Save(userID string, y1, y2 *big.Int) (uuid.UUID, error) {
	log.Trace("AuthDB Save")

	ctx := context.Background()

	var authID string
	// bigInt in postgres is an int64 in Go
	// big.Int is stored as a numeric in postgres
	ny1 := pgtype.Numeric{Int: y1, Exp: 0, Status: pgtype.Present}
	ny2 := pgtype.Numeric{Int: y2, Exp: 0, Status: pgtype.Present}

	var sqlStatementAuth = `
	INSERT INTO ` + *dbname + `.auth (id, user_id, y1, y2)
	VALUES (uuid_generate_v4(), $1, $2, $3)
	RETURNING id
	`
	err := db.pool.QueryRow(ctx,
		sqlStatementAuth,
		userID,
		&ny1,
		&ny2,
	).Scan(&authID)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" { // `pgerrcode.UniqueViolation`
				return uuid.Nil, errors.New(errormsgs.ErrExists)
			}
		}
		wrappedErr := fmt.Errorf("datatabase error. Failed to insert auth: %v", err)
		return uuid.Nil, wrappedErr
	}
	authUID := uuid.MustParse(authID)
	return authUID, nil
}

func (db *authDB) Update(userID string, y1, y2 *big.Int) (uuid.UUID, error) {
	log.Trace("AuthDB Update")

	var authID string
	ny1 := pgtype.Numeric{Int: y1, Exp: 0, Status: pgtype.Present}
	ny2 := pgtype.Numeric{Int: y2, Exp: 0, Status: pgtype.Present}

	ctx := context.Background()
	var sqlStatementAuth = `
	UPDATE ` + *dbname + `.auth SET
		y1 = $1,
		y2 = $2
		WHERE user_id = $3
		RETURNING id
		`
	err := db.pool.QueryRow(ctx,
		sqlStatementAuth,
		&ny1,
		&ny2,
		userID,
	).Scan(&authID)
	if err == pgx.ErrNoRows {
		return uuid.Nil, errors.New(errormsgs.ErrNotFound)
	}
	if err != nil {
		wrappedErr := fmt.Errorf("datatabase error. Failed to update auth: %v", err)
		return uuid.Nil, wrappedErr
	}
	authUID := uuid.MustParse(authID)
	return authUID, nil
}

func (db *authDB) Read(authID uuid.UUID) (*big.Int, *big.Int, error) {
	log.Trace("AuthDB Read")

	ctx := context.Background()
	var sqlStatement = `
	SELECT y1, y2 FROM ` + *dbname + `.auth WHERE id=$1;`
	row := db.pool.QueryRow(ctx, sqlStatement, authID)

	ny1 := pgtype.Numeric{}
	ny2 := pgtype.Numeric{}

	switch err := row.Scan(&ny1, &ny2); err {
	case pgx.ErrNoRows:
		return nil, nil, errors.New("not found")
	case nil:
		var y1, y2 *big.Int
		y1 = ny1.Int
		y2 = ny2.Int
		return y1, y2, nil
	default:
		wrappedErr := fmt.Errorf("datatabase error. Failed to retrieve user for authID %v: %v", authID, err)
		return nil, nil, wrappedErr
	}
}

func (db *authDB) ReadAuthID(userID string) (uuid.UUID, error) {
	log.Trace("AuthDB Read")

	ctx := context.Background()
	var sqlStatement = `
	SELECT id FROM ` + *dbname + `.auth WHERE user_id=$1;`
	row := db.pool.QueryRow(ctx, sqlStatement, userID)

	var authID string
	switch err := row.Scan(&authID); err {
	case pgx.ErrNoRows:
		return uuid.Nil, errors.New("not found")
	case nil:
		authUID := uuid.MustParse(authID)
		return authUID, nil
	default:
		wrappedErr := fmt.Errorf("datatabase error. Failed to retrieve user for authID %v: %v", authID, err)
		return uuid.Nil, wrappedErr
	}
}

func (db *authDB) getConnectionString() string {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s&pool_max_conns=%d", *user, *password, *host, *port, *dbname, *ssl, maxConnections)
	return connString
}
