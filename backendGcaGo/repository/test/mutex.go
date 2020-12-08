package repository

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"
)

type TestRepo struct{}

var (
	mutex   sync.Mutex
	balance int
	wg      sync.WaitGroup
)

//**************this section is an example of using mutex inside two concurrent go routines**************

//TestMutexIsolationLevels this is to test trying to lock the database/test isolation levels to make sure the current transaction will not be overwritten by another goroutines changing the same information
// func (ts TestRepo) TestMutexIsolationLevels(db *sql.DB, h map[string]string, id string) error {
// 	balance = 1000

// 	wg.Add(2)
// 	go withdraw(700, &wg)
// 	go deposit(500, &wg)
// 	wg.Wait()

// 	fmt.Printf("New Balance %d\n", balance)

// 	return nil
// }

// func deposit(value int, wg *sync.WaitGroup) {
// 	mutex.Lock()
// 	fmt.Printf("Depositing %d to account with balance: %d\n", value, balance)
// 	balance += value
// 	mutex.Unlock()
// 	wg.Done()
// }

// func withdraw(value int, wg *sync.WaitGroup) {
// 	mutex.Lock()
// 	fmt.Printf("Withdrawing %d from account with balance: %d\n", value, balance)
// 	balance -= value
// 	mutex.Unlock()
// 	wg.Done()
// }

//**************this section is an example of using mutex inside two concurrent go routines inside a loop**************

// func (ts TestRepo) TestMutexIsolationLevels(db *sql.DB, h map[string]string, id string) error {
// 	balance = 1000

// 	wg.Add(2)
// 	go withdraw(700, &wg)
// 	go deposit(500, &wg)
// 	wg.Wait()

// 	fmt.Printf("New Balance %d\n", balance)

// 	return nil
// }

// func deposit(value int, wg *sync.WaitGroup) {
// 	for i := 0; i < 10; i++ {
// 		mutex.Lock()
// 		balance += value
// 		fmt.Printf("Depositing %d to account with balance: %d\n", value, balance)
// 		mutex.Unlock()
// 	}
// 	wg.Done()
// }

// func withdraw(value int, wg *sync.WaitGroup) {

// 	for i := 0; i < 10; i++ {
// 		mutex.Lock()
// 		balance -= value
// 		fmt.Printf("Withdrawing %d from account with balance: %d\n", value, balance)
// 		mutex.Unlock()
// 	}
// 	wg.Done()
// }

//**************isolation level example (wont get row until update is committed)**************

type TxOptions struct {
	// Isolation is the transaction isolation level.
	// If zero, the driver or database's default level is used.
	Isolation sql.IsolationLevel
	ReadOnly  bool
}

func (ts TestRepo) TestMutexIsolationLevels(db *sql.DB, h map[string]string, id string) error {

	errorCh := make(chan error, 3)
	// go getPrice(db, id, errorCh)
	go changePrice(db, id, 4, errorCh)
	time.Sleep(1 * time.Second)
	go getPrice(db, id, errorCh)
	time.Sleep(1 * time.Second)
	go changePrice(db, id, 5, errorCh)

	for i := 0; i < 3; i++ {
		if flag := <-errorCh; flag != nil {
			return flag
		}
	}

	return nil
}

func getPrice(db *sql.DB, id string, errorCh chan error) {
	var price string

	tx, err := db.BeginTx(context.Background(), &sql.TxOptions{
		Isolation: sql.LevelRepeatableRead,
	})

	if err != nil {
		errorCh <- err
	}

	stmt := "SELECT `product_price` FROM `prices` WHERE `id` = ?"
	_ = tx.QueryRow(stmt, id).Scan(&price)

	fmt.Println(price)
	err = tx.Commit()
	if err != nil {
		errorCh <- err
	}
	errorCh <- nil
}

func changePrice(db *sql.DB, id string, val int, errorCh chan error) {
	tx, err := db.BeginTx(context.Background(), &sql.TxOptions{
		Isolation: sql.LevelRepeatableRead,
	})

	if err != nil {
		errorCh <- err
	}

	valuesArr := make([]interface{}, 0)
	valuesArr = append(valuesArr, val, id)

	stmt := "UPDATE `prices` SET `product_price` = ? WHERE `id` = ?"
	_, _ = tx.ExecContext(context.Background(), stmt, valuesArr...)

	fmt.Println("updated: ", val)
	err = tx.Commit()
	fmt.Println(err)

	if err != nil {
		errorCh <- err
	}
	errorCh <- nil

}
