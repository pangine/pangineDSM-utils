package general

import (
	"database/sql"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	// For sqlite3 sql plugin init
	_ "github.com/mattn/go-sqlite3"
)

// InsnRst represents a row stored in the insn table
type InsnRst struct {
	Offset int
	ID     int
}

const maxSQLVals = 100
const maxSQLQuery = 50000

// CreateSqliteRst creates sqlite for storing result on disk
func CreateSqliteRst(sqlPath string) {
	os.Remove(sqlPath)
	db, err := sql.Open("sqlite3", sqlPath)
	if err != nil {
		fmt.Printf("FATAL: sqlite file %s create failed\n", sqlPath)
		panic(err)
	}
	defer db.Close()
	// create tables
	stm, err := db.Prepare("CREATE TABLE IF NOT EXISTS insn (" +
		"offset INTEGER PRIMARY KEY, " +
		"iter INTEGER" +
		")")
	if err != nil {
		fmt.Println("FATAL: sqlite statement error")
		panic(err)
	}
	stm.Exec()
	stm.Close()

	stm, err = db.Prepare("CREATE TABLE IF NOT EXISTS func (" +
		"start INTEGER PRIMARY KEY" +
		")")
	if err != nil {
		fmt.Println("FATAL: sqlite statement error")
		panic(err)
	}
	stm.Exec()
	stm.Close()

	stm, err = db.Prepare("CREATE TABLE IF NOT EXISTS backtrack (" +
		"iter INTEGER PRIMARY KEY, " +
		"source TEXT" +
		")")
	if err != nil {
		fmt.Println("FATAL: sqlite statement error")
		panic(err)
	}
	stm.Exec()
	stm.Close()
}

// InsertSqliteRstInsn insert input offsets into insn table with iter=id.
// Will ignore if the offset alreay exists
func InsertSqliteRstInsn(sqlPath string, insns []int, id int) {
	db, err := sql.Open("sqlite3", sqlPath)
	if err != nil {
		fmt.Printf("FATAL: sqlite file %s open failed\n", sqlPath)
		panic(err)
	}
	defer db.Close()
	// Insert in groups of maxSQLVals
	insertStr := "INSERT OR IGNORE INTO insn (offset, iter) VALUES "
	value := "(?, ?)"
	insertFormation := make([]string, 0)
	vals := make([]interface{}, 0)
	counter := 0
	for _, insn := range insns {
		if counter++; counter >= maxSQLVals {
			// sqlite3 plugin cannot support too many vals insertion at once
			counter = 0
			insertStr += strings.Join(insertFormation, ",")
			stm, err := db.Prepare(insertStr)
			if err != nil {
				fmt.Println("FATAL: sqlite insertion statement incorrect")
				panic(err)
			}
			_, err = stm.Exec(vals...)
			if err != nil {
				fmt.Println("FATAL: sqlite insert values failed")
				panic(err)
			}
			stm.Close()
			insertStr = "INSERT OR IGNORE INTO insn (offset, iter) VALUES "
			insertFormation = make([]string, 0)
			vals = make([]interface{}, 0)
		}
		insertFormation = append(insertFormation, value)
		vals = append(vals, insn, id)
	}
	insertStr += strings.Join(insertFormation, ",")
	if len(vals) > 0 {
		stm, err := db.Prepare(insertStr)
		if err != nil {
			fmt.Println("FATAL: sqlite insertion statement incorrect")
			panic(err)
		}
		_, err = stm.Exec(vals...)
		stm.Close()
		if err != nil {
			fmt.Println("FATAL: sqlite insert values failed")
			panic(err)
		}
	}
}

// InsertSqliteRstBT insert input (id, source) pair into backtrack table
func InsertSqliteRstBT(sqlPath string, id int, source string) {
	db, err := sql.Open("sqlite3", sqlPath)
	if err != nil {
		fmt.Printf("FATAL: sqlite file %s open failed\n", sqlPath)
		panic(err)
	}
	defer db.Close()
	// Insert one by one
	insertStr := "INSERT INTO backtrack (iter, source) VALUES (?, ?)"
	stm, err := db.Prepare(insertStr)
	if err != nil {
		fmt.Println("FATAL: sqlite insertion statement incorrect")
		panic(err)
	}
	_, err = stm.Exec(id, source)
	if err != nil {
		fmt.Println("FATAL: sqlite insert values failed")
		panic(err)
	}
	stm.Close()
}

// InsertSqliteRstFunc insert input offsets into func table
func InsertSqliteRstFunc(sqlPath string, offsets []int) {
	db, err := sql.Open("sqlite3", sqlPath)
	if err != nil {
		fmt.Printf("FATAL: sqlite file %s open failed\n", sqlPath)
		panic(err)
	}
	defer db.Close()
	// Insert in groups of maxSQLVals
	insertStr := "INSERT OR IGNORE INTO func (start) VALUES "
	value := "(?)"
	insertFormation := make([]string, 0)
	vals := make([]interface{}, 0)
	counter := 0
	for _, o := range offsets {
		if counter++; counter >= maxSQLVals {
			// sqlite3 plugin cannot support too many vals insertion at once
			counter = 0
			insertStr += strings.Join(insertFormation, ",")
			stm, err := db.Prepare(insertStr)
			if err != nil {
				fmt.Println("FATAL: sqlite insertion statement incorrect")
				panic(err)
			}
			_, err = stm.Exec(vals...)
			if err != nil {
				fmt.Println("FATAL: sqlite insert values failed")
				panic(err)
			}
			stm.Close()
			insertStr = "INSERT OR IGNORE INTO func (start) VALUES "
			insertFormation = make([]string, 0)
			vals = make([]interface{}, 0)
		}
		insertFormation = append(insertFormation, value)
		vals = append(vals, o)
	}
	insertStr += strings.Join(insertFormation, ",")
	if len(vals) > 0 {
		stm, err := db.Prepare(insertStr)
		if err != nil {
			fmt.Println("FATAL: sqlite insertion statement incorrect")
			panic(err)
		}
		_, err = stm.Exec(vals...)
		stm.Close()
		if err != nil {
			fmt.Println("FATAL: sqlite insert values failed")
			panic(err)
		}
	}
}

// SumSqliteRstInsnOffset sum up rows from insn table where iter in [from, to)
// If to <= from, there is no upbound restriction
func SumSqliteRstInsnOffset(sqlPath string, from, to int) (count int) {
	db, err := sql.Open("sqlite3", sqlPath)
	if err != nil {
		fmt.Printf("FATAL: sqlite file %s open failed\n", sqlPath)
		panic(err)
	}
	defer db.Close()
	fromStr := "FROM insn WHERE iter >= " + strconv.Itoa(from)
	if to > from {
		fromStr += " AND iter < " + strconv.Itoa(to)
	}

	sum, err := db.Query("SELECT COUNT(*) " + fromStr)
	if err != nil {
		fmt.Println("FATAL: sqlite selection count from insn failed")
	}
	sum.Next()
	sum.Scan(&count)
	sum.Close()
	return
}

// ReadSqliteRstInsnOffset select offsets from insn table where iter in [from, to)
// If to <= from, there is no upbound restriction
func ReadSqliteRstInsnOffset(sqlPath string, from, to int) (offsets []int) {
	offsets = make([]int, 0)
	db, err := sql.Open("sqlite3", sqlPath)
	if err != nil {
		fmt.Printf("FATAL: sqlite file %s open failed\n", sqlPath)
		panic(err)
	}
	defer db.Close()
	fromStr := "FROM insn WHERE iter >= " + strconv.Itoa(from)
	if to > from {
		fromStr += " AND iter < " + strconv.Itoa(to)
	}

	sum, err := db.Query("SELECT COUNT(*) " + fromStr)
	if err != nil {
		fmt.Println("FATAL: sqlite selection count from insn failed")
	}
	var count int
	sum.Next()
	sum.Scan(&count)
	sum.Close()

	for i := 0; i < count; i += maxSQLQuery {
		rows, err := db.Query("SELECT offset " + fromStr + " LIMIT " +
			strconv.Itoa(maxSQLQuery) + " OFFSET " + strconv.Itoa(i))
		if err != nil {
			fmt.Println("FATAL: sqlite select from insn failed")
			panic(err)
		}
		var offset int
		for rows.Next() {
			rows.Scan(&offset)
			offsets = append(offsets, offset)
		}
		rows.Close()
	}
	return
}

// ReadSqliteRstInsnInFunc select offsets from insn table in a function defined by [from, to)
// If to <= from, there is no upbound restriction
func ReadSqliteRstInsnInFunc(sqlPath string, from, to int) (offsets []int) {
	offsets = make([]int, 0)
	db, err := sql.Open("sqlite3", sqlPath)
	if err != nil {
		fmt.Printf("FATAL: sqlite file %s open failed\n", sqlPath)
		panic(err)
	}
	defer db.Close()
	fromStr := "FROM insn WHERE offset >= " + strconv.Itoa(from)
	if to > from {
		fromStr += " AND offset < " + strconv.Itoa(to)
	}

	sum, err := db.Query("SELECT COUNT(*) " + fromStr)
	if err != nil {
		fmt.Println("FATAL: sqlite selection count from insn failed")
	}
	var count int
	sum.Next()
	sum.Scan(&count)
	sum.Close()

	for i := 0; i < count; i += maxSQLQuery {
		rows, err := db.Query("SELECT offset " + fromStr + " LIMIT " +
			strconv.Itoa(maxSQLQuery) + " OFFSET " + strconv.Itoa(i))
		if err != nil {
			fmt.Println("FATAL: sqlite select from insn failed")
			panic(err)
		}
		var offset int
		for rows.Next() {
			rows.Scan(&offset)
			offsets = append(offsets, offset)
		}
		rows.Close()
	}
	return
}

// ReadSqliteRstInsnAll select all elements from insn table
func ReadSqliteRstInsnAll(sqlPath string) (insns []InsnRst) {
	insns = make([]InsnRst, 0)
	db, err := sql.Open("sqlite3", sqlPath)
	if err != nil {
		fmt.Printf("FATAL: sqlite file %s open failed\n", sqlPath)
		panic(err)
	}
	defer db.Close()

	sum, err := db.Query("SELECT COUNT(*) FROM insn")
	if err != nil {
		fmt.Println("FATAL: sqlite selection count from insn failed")
	}
	var count int
	sum.Next()
	sum.Scan(&count)
	sum.Close()

	for i := 0; i < count; i += maxSQLQuery {
		rows, err := db.Query("SELECT * From insn LIMIT " +
			strconv.Itoa(maxSQLQuery) + " OFFSET " + strconv.Itoa(i))
		if err != nil {
			fmt.Println("FATAL: sqlite select from insn failed")
			panic(err)
		}
		var offset int
		var id int
		for rows.Next() {
			rows.Scan(&offset, &id)
			insns = append(insns, InsnRst{Offset: offset, ID: id})
		}
		rows.Close()
	}
	return
}

// ReadSqliteRstFunc select start from func table
func ReadSqliteRstFunc(sqlPath string) (funcs []int) {
	funcs = make([]int, 0)
	db, err := sql.Open("sqlite3", sqlPath)
	if err != nil {
		fmt.Printf("FATAL: sqlite file %s open failed\n", sqlPath)
		panic(err)
	}
	defer db.Close()

	sum, err := db.Query("SELECT COUNT(*) FROM func")
	if err != nil {
		fmt.Println("FATAL: sqlite selection count from func failed")
	}
	var count int
	sum.Next()
	sum.Scan(&count)
	sum.Close()

	for i := 0; i < count; i += maxSQLQuery {
		rows, err := db.Query("SELECT start From func LIMIT " +
			strconv.Itoa(maxSQLQuery) + " OFFSET " + strconv.Itoa(i))
		if err != nil {
			fmt.Println("FATAL: sqlite select from func failed")
			panic(err)
		}
		var start int
		for rows.Next() {
			rows.Scan(&start)
			funcs = append(funcs, start)
		}
		rows.Close()
	}
	return
}

// ReadSqliteRstBT select all from backtrack table
func ReadSqliteRstBT(sqlPath string) (bt map[int]string) {
	bt = make(map[int]string)
	db, err := sql.Open("sqlite3", sqlPath)
	if err != nil {
		fmt.Printf("FATAL: sqlite file %s open failed\n", sqlPath)
		panic(err)
	}
	defer db.Close()

	sum, err := db.Query("SELECT COUNT(*) FROM backtrack")
	if err != nil {
		fmt.Println("FATAL: sqlite selection count from backtrack failed")
	}
	var count int
	sum.Next()
	sum.Scan(&count)
	sum.Close()

	for i := 0; i < count; i += maxSQLQuery {
		rows, err := db.Query("SELECT iter, source From backtrack LIMIT " +
			strconv.Itoa(maxSQLQuery) + " OFFSET " + strconv.Itoa(i))
		if err != nil {
			fmt.Println("FATAL: sqlite select from backtrack failed")
			panic(err)
		}
		var id int
		var source string
		for rows.Next() {
			rows.Scan(&id, &source)
			bt[id] = source
		}
		rows.Close()
	}
	return
}

// FindSqliteEarliestIter returns the iter that
// belongs to source and contains the earlist offset in the input
func FindSqliteEarliestIter(sqlPath string, offsets []int, source string) (iter int) {
	iter = math.MaxInt64
	db, err := sql.Open("sqlite3", sqlPath)
	if err != nil {
		fmt.Printf("FATAL: sqlite file %s open failed\n", sqlPath)
		panic(err)
	}
	defer db.Close()

	itersMap := make(map[int]bool)
	itersRows, err := db.Query("SELECT iter FROM backtrack WHERE source = '" +
		source + "'")
	if err != nil {
		fmt.Println("FATAL: sqlite select from backtrack failed")
		panic(err)
	}
	var it int
	for itersRows.Next() {
		itersRows.Scan(&it)
		itersMap[it] = true
	}
	itersRows.Close()

	for _, offset := range offsets {
		rows, err := db.Query("SELECT iter FROM insn WHERE offset = " +
			strconv.Itoa(offset))
		if err != nil {
			fmt.Println("FATAL: sqlite select from insn failed")
			panic(err)
		}
		it = math.MaxInt64
		if rows.Next() {
			rows.Scan(&it)
		}
		if itersMap[it] && it < iter {
			iter = it
		}
		rows.Close()
	}
	return
}

// DeleteSqliteRstInsnInIters deletes offsets within iters [from, to)
// If to <= from, there is no upbound restriction
func DeleteSqliteRstInsnInIters(sqlPath string, from, to int) {
	db, err := sql.Open("sqlite3", sqlPath)
	if err != nil {
		fmt.Printf("FATAL: sqlite file %s open failed\n", sqlPath)
		panic(err)
	}
	defer db.Close()
	delStr := "DELETE FROM insn WHERE iter >= " + strconv.Itoa(from)
	if to > from {
		delStr += " AND iter < " + strconv.Itoa(to)
	}

	_, err = db.Exec(delStr)
	if err != nil {
		fmt.Println("FATAL: sqlite delete from insn failed")
	}
}
