package app

import (
	"log"
	"net/http"
	"os"
	"strings"

	"GOWEB/WEB20-1/model"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
var rd *render.Render = render.New() //전역변수 render.New() 초기화

type AppHandler struct {
	http.Handler //handler http.Handler인데 handler를 생략, 암시적으로 인터페이스를 포함한 멤버 변수를 포함한 상태
	db           model.DBHandler
	db1          model.DBHandler1
	db2          model.DBHandler2
	db3          model.DBHandler3
}

func getSesssionID(r *http.Request) string {
	session, err := store.Get(r, "session")
	if err != nil {
		return ""
	}

	// Set some session values.
	val := session.Values["id"]
	if val == nil {
		return ""
	}
	return val.(string)
}

//핸들러들을 (a *AppHandler)메소드로 바꾼다
func (a *AppHandler) indexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/index.html", http.StatusTemporaryRedirect)
	log.Println(w, r)
}

//핸들러들을 (a *AppHandler)메소드로 바꾼다
func (a *AppHandler) getMemberListHandler(w http.ResponseWriter, r *http.Request) {
	list := a.db.GetMembers() //model -> a.db로 바꾼다
	rd.JSON(w, http.StatusOK, list)
}

//핸들러들을 (a *AppHandler)메소드로 바꾼다
func (a *AppHandler) addMemberHandler(w http.ResponseWriter, r *http.Request) { //member list add 해주는 핸들러

	id := r.FormValue("id")
	pswd := r.FormValue("pswd")
	name := r.FormValue("name")
	birth := r.FormValue("birth")
	gender := r.FormValue("gender")
	email := r.FormValue("email")
	mobile := r.FormValue("mobile")                                        // js에서 보낸 input value를 name에 추가
	member := a.db.AddMember(id, pswd, name, birth, gender, email, mobile) //model -> a.db로 바꾼다
	rd.JSON(w, http.StatusCreated, member)                                 // JSON으로 member 값을 반환
}

func (a *AppHandler) removeMemberHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]            // id값을 가져온다
	ok := a.db.RemoveMember(id) //model -> a.db로 바꾼다
	if ok {
		rd.JSON(w, http.StatusOK, Success{true}) //ok 성공시
	} else { //없는 경우
		rd.JSON(w, http.StatusOK, Success{false}) // 실패를 알려준다
	}
} // 넣고쓰고업뎃하는거 밖에 없다.

//-------------------------------------------------------Done

func (a *AppHandler) getDoneListHandler(w http.ResponseWriter, r *http.Request) {
	list := a.db1.GetDones() //model -> a.db로 바꾼다
	rd.JSON(w, http.StatusOK, list)
}

//핸들러들을 (a *AppHandler)메소드로 바꾼다
func (a *AppHandler) addDoneHandler(w http.ResponseWriter, r *http.Request) { //member list add 해주는 핸들러

	id := r.FormValue("id")
	pswd := r.FormValue("pswd")
	name := r.FormValue("name")
	birth := r.FormValue("birth")
	gender := r.FormValue("gender")
	email := r.FormValue("email")
	mobile := r.FormValue("mobile")                                     // js에서 보낸 input value를 name에 추가
	done := a.db1.AddDone(id, pswd, name, birth, gender, email, mobile) //model -> a.db로 바꾼다
	rd.JSON(w, http.StatusCreated, done)                                // JSON으로 member 값을 반환
}

type Success struct { //(클라이언트) 응답 결과를 알려주기 위한 구조체
	Success bool `json:"success"`
}

func (a *AppHandler) removeDoneHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]           // id값을 가져온다
	ok := a.db1.RemoveDone(id) //model -> a.db로 바꾼다
	if ok {
		rd.JSON(w, http.StatusOK, Success{true}) //ok 성공시
	} else { //없는 경우
		rd.JSON(w, http.StatusOK, Success{false}) // 실패를 알려준다
	}
} // 넣고쓰고업뎃하는거 밖에 없다.

//----------------------------------------------------Account

func (a *AppHandler) getAccountListHandler(w http.ResponseWriter, r *http.Request) {
	list := a.db2.GetAccounts() //model -> a.db로 바꾼다
	rd.JSON(w, http.StatusOK, list)
}

//핸들러들을 (a *AppHandler)메소드로 바꾼다
func (a *AppHandler) addAccountHandler(w http.ResponseWriter, r *http.Request) { //member list add 해주는 핸들러

	id := r.FormValue("id")
	pswd := r.FormValue("pswd")
	name := r.FormValue("name")
	birth := r.FormValue("birth")
	gender := r.FormValue("gender")
	email := r.FormValue("email")
	mobile := r.FormValue("mobile")                                           // js에서 보낸 input value를 name에 추가
	Account := a.db2.AddAccount(id, pswd, name, birth, gender, email, mobile) //model -> a.db로 바꾼다
	rd.JSON(w, http.StatusCreated, Account)                                   // JSON으로 member 값을 반환
}

func (a *AppHandler) removeAccountHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]              // id값을 가져온다
	ok := a.db2.RemoveAccount(id) //model -> a.db로 바꾼다
	if ok {
		rd.JSON(w, http.StatusOK, Success{true}) //ok 성공시
	} else { //없는 경우
		rd.JSON(w, http.StatusOK, Success{false}) // 실패를 알려준다
	}
} // 넣고쓰고업뎃하는거 밖에 없다.

//----------------------------------------------Stock

func (a *AppHandler) getStockListHandler(w http.ResponseWriter, r *http.Request) {
	list := a.db3.GetStocks() //model -> a.db로 바꾼다
	rd.JSON(w, http.StatusOK, list)
}

//핸들러들을 (a *AppHandler)메소드로 바꾼다
func (a *AppHandler) addStockHandler(w http.ResponseWriter, r *http.Request) { //member list add 해주는 핸들러

	id := r.FormValue("id")
	pswd := r.FormValue("pswd")
	name := r.FormValue("name")
	birth := r.FormValue("birth")
	gender := r.FormValue("gender")
	email := r.FormValue("email")
	mobile := r.FormValue("mobile")                                       // js에서 보낸 input value를 name에 추가
	stock := a.db3.AddStock(id, pswd, name, birth, gender, email, mobile) //model -> a.db로 바꾼다
	rd.JSON(w, http.StatusCreated, stock)                                 // JSON으로 member 값을 반환
}

func (a *AppHandler) removeStockHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]            // id값을 가져온다
	ok := a.db3.RemoveStock(id) //model -> a.db로 바꾼다
	if ok {
		rd.JSON(w, http.StatusOK, Success{true}) //ok 성공시
	} else { //없는 경우
		rd.JSON(w, http.StatusOK, Success{false}) // 실패를 알려준다
	}
} // 넣고쓰고업뎃하는거 밖에 없다.

//핸들러들을 (a *AppHandler)메소드로 바꾼다
func (a *AppHandler) Close() { //새롭게 Close()를 외부에서 만들어 준 것.
	a.db.Close() //model -> a.db로 바꾼다
	a.db1.Close()
	a.db2.Close()
	a.db3.Close()
}

func CheckSignin(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// if request URL is /signin.html, then next()
	if strings.Contains(r.URL.Path, "/signin.html") ||
		strings.Contains(r.URL.Path, "/auth") {
		next(w, r)
		return
	}

	// if user already signed in
	sessionID := getSesssionID(r)
	if sessionID != "" {
		next(w, r)
		return
	}

	// if not user sign in
	// redirect singin.html
	http.Redirect(w, r, "/signin.html", http.StatusTemporaryRedirect)
}

func MakeHandler(filepath string) *AppHandler {
	r := mux.NewRouter()
	n := negroni.New(
		negroni.NewRecovery(),
		negroni.NewLogger(),
		negroni.HandlerFunc(CheckSignin),
		negroni.NewStatic(http.Dir("public")))
	n.UseHandler(r)

	a := &AppHandler{
		Handler: n,
		db:      model.NewDBHandler(filepath),
		db1:     model.NewDBHandler1(filepath),
		db2:     model.NewDBHandler2(filepath),
		db3:     model.NewDBHandler3(filepath),
	}

	r.HandleFunc("/members", a.getMemberListHandler).Methods("GET")
	r.HandleFunc("/members", a.addMemberHandler).Methods("POST")
	r.HandleFunc("/members/{id:[0-9]+}", a.removeMemberHandler).Methods("DELETE")
	r.HandleFunc("/Dones", a.getDoneListHandler).Methods("GET")
	r.HandleFunc("/Dones", a.addDoneHandler).Methods("POST")
	r.HandleFunc("/Dones/{id:[0-9]+}", a.removeDoneHandler).Methods("DELETE")
	r.HandleFunc("/Accounts", a.getAccountListHandler).Methods("GET")
	r.HandleFunc("/Accounts", a.addAccountHandler).Methods("POST")
	r.HandleFunc("/Accounts/{id:[0-9]+}", a.removeAccountHandler).Methods("DELETE")
	r.HandleFunc("/Stocks", a.getStockListHandler).Methods("GET")
	r.HandleFunc("/Stocks", a.addStockHandler).Methods("POST")
	r.HandleFunc("/Stocks/{id:[0-9]+}", a.removeStockHandler).Methods("DELETE")
	r.HandleFunc("/auth/google/login", googleLoginHandler)
	r.HandleFunc("/auth/google/callback", googleAuthCallback)
	r.HandleFunc("/", a.indexHandler)

	return a
}
