package model

type Member struct {
	ID     string `json:"id"`
	PSWD   string `json:"pswd"`
	Name   string `json:"name"`
	Birth  string `json:"birth"`
	Gender string `json:"gender"`
	Email  string `json:"email"`
	Mobile string `json:"mobile"`
}

type Done struct {
	ID     string `json:"id"`
	PSWD   string `json:"pswd"`
	Name   string `json:"name"`
	Birth  string `json:"birth"`
	Gender string `json:"gender"`
	Email  string `json:"email"`
	Mobile string `json:"mobile"`
}

type Account struct {
	ID     string `json:"id"`
	PSWD   string `json:"pswd"`
	Name   string `json:"name"`
	Birth  string `json:"birth"`
	Gender string `json:"gender"`
	Email  string `json:"email"`
	Mobile string `json:"mobile"`
}

type Stock struct {
	ID     string `json:"id"`
	PSWD   string `json:"pswd"`
	Name   string `json:"name"`
	Birth  string `json:"birth"`
	Gender string `json:"gender"`
	Email  string `json:"email"`
	Mobile string `json:"mobile"`
}

type DBHandler interface {
	GetMembers() []*Member
	AddMember(id string, pswd string, name string, birth string, gender string, email string, mobile string) *Member
	RemoveMember(id string) bool
	//인스턴스를 사용하는 측에 대문자로 인터페이스를 추가하고 외부 공개
	Close() //인스턴스를 사용하는 측에 대문자로 인터페이스를 추가하고 외부 공개
}

type DBHandler1 interface {
	GetDones() []*Done
	AddDone(id string, pswd string, name string, birth string, gender string, email string, mobile string) *Done
	RemoveDone(id string) bool
	Close()
}

type DBHandler2 interface {
	GetAccounts() []*Account
	AddAccount(id string, pswd string, name string, birth string, gender string, email string, mobile string) *Account
	RemoveAccount(id string) bool
	Close()
}

type DBHandler3 interface {
	GetStocks() []*Stock
	AddStock(id string, pswd string, name string, birth string, gender string, email string, mobile string) *Stock
	RemoveStock(id string) bool
	Close()
}

func NewDBHandler(filepath string) DBHandler { //DBHandler를 사용하다가 필요없을 때 Close()를 호출한다.
	return newSqliteHandler(filepath)
}

func NewDBHandler1(filepath string) DBHandler1 { //DBHandler를 사용하다가 필요없을 때 Close()를 호출한다.
	return newDoneHandler(filepath)
}

func NewDBHandler2(filepath string) DBHandler2 { //DBHandler를 사용하다가 필요없을 때 Close()를 호출한다.
	return newAccountHandler(filepath)
}

func NewDBHandler3(filepath string) DBHandler3 { //DBHandler를 사용하다가 필요없을 때 Close()를 호출한다.
	return newStockHandler(filepath)
}
