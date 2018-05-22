package main

import (
	"log"
	"github.com/go-pg/pg"
	"sync"
)

func RunTransaction() {
	incrInTx := func(db *pg.DB) error {
		tx, err := db.Begin()
		if err != nil {
			return err
		}
		// Rollback tx on error.
		defer tx.Rollback()
	
		var counter int
		_, err = tx.QueryOne(
			pg.Scan(&counter), `SELECT counter FROM transactions FOR UPDATE`)
		if err != nil {
			return err
		}
	
		counter++
	
		_, err = tx.Exec(`UPDATE transactions SET counter = ?`, counter)
		if err != nil {
			return err
		}
	
		return tx.Commit()
	}
	
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := incrInTx(Db); err != nil {
				panic(err)
			}
		}()
	}
	wg.Wait()
	
	var counter int
	_, err := Db.QueryOne(pg.Scan(&counter), `SELECT counter FROM transactions`)
	if err != nil {
		panic(err)
	}
	log.Println(counter)
}