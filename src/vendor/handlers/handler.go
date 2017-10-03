package handlers

//Created By: Ricky Clevinger
//Updated On: 8/17/2017
//Last Updated By: Ricky Clevinger

import (
	"html/template"
	"net/http"
	"regexp"
	"modify"
	"fmt"
	"os"
	"gopkg.in/ldap.v2"
	"log"
)

var validPath = regexp.MustCompile("^/(index.html|edit.html)$")
var templates = template.Must(template.ParseFiles("views/index.html", "views/edit.html"))

type Page struct {
	Uid string
	First string
	Last string
	Email string
	GNum string
	UidNum string
	HomeDir string
	DisplayName string
	LogShell string
	Mobile string
	Disabled bool
	SeTeam bool
	Jira bool
	Jrebel bool
	Nagios bool
	Owncloud bool
	RocketChat bool
	SassyDev bool
	SassyProd bool
	SavvyServiceDesk bool
	Solaris_Linux bool
	Subversion bool
	VNC bool
	Wiki bool
}

func LoadPage() *Page {
	return &Page{}
}

func LoadEditPage(uid, first, last, email, displayname, gnum, mobile, uidnum, homedir, logshell string) *Page {
	return &Page{Uid: uid, First: first, Last: last, Email: email, DisplayName: displayname, GNum: gnum, Mobile: mobile, UidNum: uidnum, HomeDir: homedir, LogShell: logshell}
}


//Renders HTML page
func RenderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//Validates path and calls handler
func MakeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r)
	}
}

//Handles the index page
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	//uid := "rclevinger"
	p := LoadPage()
	RenderTemplate(w, "index", p)
}

//Handles the edit page
func EditHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	uid := r.Form["uid"]
	
	p := Search(uid[0])
	RenderTemplate(w, "edit", p)
}

func Search(uid string) *Page{
	
	username := os.Args[1]
	password := os.Args[2]

	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", os.Args[3], 389))
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	err = l.Bind(username, password)
	if err != nil {
		log.Fatal(err)
	}
	
	searchRequest := ldap.NewSearchRequest(
    "uid=rclevinger,ou=People,dc=spg,dc=cgi,dc=com", // The base dn to search
    ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
	"(&(uid="+ uid +"))",// The filter to apply
    []string{"givenName", "sn", "mail", "displayName", "gidNumber", "mobile", "uidNumber", "homeDirectory", "loginShell" },                    // A list attributes to retrieve
    nil,
	)
	
	sr, err := l.Search(searchRequest)
	if err != nil {
    log.Fatal(err)
	}
	
	
	for _, entry := range sr.Entries {
	given := entry.GetAttributeValue("givenName")
	last := entry.GetAttributeValue("sn")
	mail := entry.GetAttributeValue("mail")
	display := entry.GetAttributeValue("displayName")
	gnum := entry.GetAttributeValue("gidNumber")
	mobile := entry.GetAttributeValue("mobile")
	uidNumber := entry.GetAttributeValue("uidNumber")
	homeDirectory := entry.GetAttributeValue("homeDirectory")
	loginShell := entry.GetAttributeValue("loginShell")
	return LoadEditPage(uid, given, last, mail, display, gnum, mobile, uidNumber, homeDirectory, loginShell)
	}
	
	return LoadPage()

}


//Redirect to login.html
func Redirect(w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r, "/index.html", 301)
}


//Handles
func Handles() {

	http.HandleFunc("/index.html", MakeHandler(IndexHandler))
	http.HandleFunc("/edit.html", MakeHandler(EditHandler))
	http.HandleFunc("/modify", modify.Modify)
	http.HandleFunc("/", Redirect)
}
