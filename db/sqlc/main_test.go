package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/simplebankapp/db/util"
)

var testQueries *Queries
var testDBConnPool *pgxpool.Pool

func TestMain(m *testing.M) {
	var err error 
	ctx := context.Background()

	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	testDBConnPool, err = pgxpool.New(ctx, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to DB:", err)
	}
	defer testDBConnPool.Close()
	
	testQueries = New(testDBConnPool)

	os.Exit(m.Run())
}