package model

type memoryHandler struct {
	memberMap map[string]*Member
}
type memory1Handler struct {
	doneMap map[string]*Done
}
type memory2Handler struct {
	accountMap map[string]*Account
}
type memory3Handler struct {
	stockMap map[string]*Stock
}

//4개 func을 만든다
func (m *memoryHandler) GetMembers() []*Member {
	list := []*Member{}
	for _, v := range m.memberMap {
		list = append(list, v)
	}
	return list
}

func (m *memoryHandler) AddMember(id string, pswd string, name string, birth string, gender string, email string, mobile string) *Member {
	member := &Member{id, pswd, name, birth, gender, email, mobile}
	m.memberMap[id] = member
	return member
}

func (m *memoryHandler) RemoveMember(id string) bool {
	if _, ok := m.memberMap[id]; ok { // memberMap id 값이 있으면
		delete(m.memberMap, id) //지우고
		return true
	}
	return false
}

func (d *memoryHandler) Close() {

}

func newMemoryHandler() DBHandler {
	m := &memoryHandler{}
	m.memberMap = make(map[string]*Member)
	return m
}

//-------------------------------------------Done

func (d *memory1Handler) GetDones() []*Done {
	list := []*Done{}
	for _, v := range d.doneMap {
		list = append(list, v)
	}
	return list
}

func (d *memory1Handler) AddDone(id string, pswd string, name string, birth string, gender string, email string, mobile string) *Done {
	done := &Done{id, pswd, name, birth, gender, email, mobile}
	d.doneMap[id] = done
	return done
}

func (d *memory1Handler) RemoveDone(id string) bool {
	if _, ok := d.doneMap[id]; ok { // memberMap id 값이 있으면
		delete(d.doneMap, id) //지우고
		return true
	}
	return false
}

func (d *memory1Handler) Close() {

}

func newMemory1Handler() DBHandler1 {
	d := &memory1Handler{}
	d.doneMap = make(map[string]*Done)
	return d
}

//-------------------------------------------Account

func (ac *memory2Handler) GetAccounts() []*Account {
	list := []*Account{}
	for _, v := range ac.accountMap {
		list = append(list, v)
	}
	return list
}

func (ac *memory2Handler) AddAccount(id string, pswd string, name string, birth string, gender string, email string, mobile string) *Account {
	account := &Account{id, pswd, name, birth, gender, email, mobile}
	ac.accountMap[id] = account
	return account
}

func (ac *memory2Handler) RemoveAccount(id string) bool {
	if _, ok := ac.accountMap[id]; ok { // memberMap id 값이 있으면
		delete(ac.accountMap, id) //지우고
		return true
	}
	return false
}

func (ac *memory2Handler) Close() {

}

func newMemory2Handler() DBHandler2 {
	ac := &memory2Handler{}
	ac.accountMap = make(map[string]*Account)
	return ac
}

//-------------------------------------------Stock

func (st *memory3Handler) GetStocks() []*Stock {
	list := []*Stock{}
	for _, v := range st.stockMap {
		list = append(list, v)
	}
	return list
}

func (st *memory3Handler) AddStock(id string, pswd string, name string, birth string, gender string, email string, mobile string) *Stock {
	stock := &Stock{id, pswd, name, birth, gender, email, mobile}
	st.stockMap[id] = stock
	return stock
}

func (st *memory3Handler) RemoveStock(id string) bool {
	if _, ok := st.stockMap[id]; ok { // memberMap id 값이 있으면
		delete(st.stockMap, id) //지우고
		return true
	}
	return false
}

func (st *memory3Handler) Close() {

}

func newMemory3Handler() DBHandler3 {
	st := &memory3Handler{}
	st.stockMap = make(map[string]*Stock)
	return st
}
