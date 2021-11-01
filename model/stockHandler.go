package model

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3" //암시적
)

type stockHandler struct {
	db3 *sql.DB // 멤버변수로 가진다
}

//----------------------------------------Done
func (s *stockHandler) GetStocks() []*Stock {
	stocks := []*Stock{}                                                                        //list를 만든다
	rows, err := s.db3.Query("SELECT id, pswd, name, birth, gender, email, mobile FROM stocks") //데이터를 가져오는 쿼리는 SELECT
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() { //rows 행이다. Next() 다음 레코드로 간다, true가 계속될 때까지 돌면서 레코드를 읽어온다.
		var stock Stock                                                                                          //받아온 데이터를 담을 공간을 만든다
		rows.Scan(&stock.ID, &stock.PSWD, &stock.Name, &stock.Birth, &stock.Gender, &stock.Email, &stock.Mobile) // 첫 번째부터 네 번째까지 컬럼을 쿼리에서 받아(가져)온다.
		stocks = append(stocks, &stock)
	}
	return stocks
}

func (s *stockHandler) AddStock(id string, pswd string, name string, birth string, gender string, email string, mobile string) *Stock { //VALUES는 각 항목, (?,?)어떤 VALUES? (?,?) 첫 번째는 name 두 번째는 completed
	stmt, err := s.db3.Prepare("INSERT INTO stock (id, pswd, name, birth, gender, email, mobile) VALUES (?, ?, ?, ?, ?, ?, ?)") //datetime은 내장함수
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec(id, pswd, name, birth, gender, email, mobile)
	if err != nil {
		panic(err)
	}
	var stock Stock
	stock.ID = id
	stock.PSWD = pswd
	stock.Name = name
	stock.Birth = birth
	stock.Gender = gender
	stock.Email = email
	stock.Mobile = mobile
	return &stock
}
func (s *stockHandler) RemoveStock(id string) bool { //WHERE 구문 특정값만 특정 id=?
	stmt, err := s.db3.Prepare("DELETE FROM dones WHERE id=?")
	if err != nil {
		panic(err)
	}
	rst, err := stmt.Exec(id)
	if err != nil {
		panic(err)
	}
	cnt, _ := rst.RowsAffected()
	return cnt > 0
}

// 함수추가, 프로그램 종료전에 함수를 사용할 수 있도록 해준다.
func (s *stockHandler) Close() {
	s.db3.Close()
}

func newStockHandler(filepath string) DBHandler3 {
	database, err := sql.Open("sqlite3", filepath)
	if err != nil {
		panic(err)
	}
	statement, _ := database.Prepare( //아래 Table에서 sql 쿼리문을 만들어준다
		`CREATE TABLE IF NOT EXISTS Stock (
			id			TEXT PRIMARY KEY,
			pswd		TEXT NOT NULL,
			name		TEXT NOT NULL,
			birth		DATE NOT NULL,
			gender		TEXT NOT NULL,
			email		TEXT,
			mobile		TEXT NOT NULL
		);`)
	statement.Exec()
	return &stockHandler{db3: database} // &sqliteHandler{}를 반환
}
