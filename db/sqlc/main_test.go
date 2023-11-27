package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	// dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries
var testDBConnPool *pgxpool.Pool

func TestMain(m *testing.M) {
	var err error 
	ctx := context.Background()
	testDBConnPool, err = pgxpool.New(ctx, dbSource)
	if err != nil {
		log.Fatal("cannot connect to DB:", err)
	}
	defer testDBConnPool.Close()
	
	testQueries = New(testDBConnPool)

	os.Exit(m.Run())
}