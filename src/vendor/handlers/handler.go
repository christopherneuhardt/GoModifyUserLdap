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
	Disabled string
	SeTeam string
	Jira string
	Jrebel string
	Nagios string
	Owncloud string
	RocketChat string
	SassyDev string
	SassyProd string
	SavvyServiceDesk string
	Solaris_Linux string
	Subversion string
	VNC string
	Wiki string
}

func LoadPage() *Page {
	return &Page{}
}

func LoadEditPage(uid, first, last, email, displayname, gnum, mobile, uidnum, homedir, logshell, disabled, seteam, jira, jrebel, nagios, owncloud, rocketchat, sassydev, sassyprod, savvyservicedesk, solaris_linux, subversion, vnc, wiki string) *Page {
	return &Page{Uid: uid, First: first, Last: last, Email: email, GNum: gnum, UidNum: uidnum, HomeDir: homedir, DisplayName: displayname, LogShell: logshell, Mobile: mobile, Disabled: disabled, SeTeam: seteam, Jira: jira, Jrebel: jrebel, Nagios: nagios, Owncloud: owncloud, RocketChat: rocketchat, SassyDev: sassydev, SassyProd: sassyprod, SavvyServiceDesk: savvyservicedesk, Solaris_Linux: solaris_linux, Subversion: subversion, VNC: vnc, Wiki: wiki}
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
    []string{"givenName", "sn", "mail", "displayName", "gidNumber", "mobile", "uidNumber", "homeDirectory", "loginShell", "employeeType" },                    // A list attributes to retrieve
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
		disabled := ""
		seteam := ""
		jira := ""
		jrebel := ""
		nagios := ""
		owncloud := ""
		rocketchat := ""
		sassydev := ""
		sassyprod := ""
		savvyservicedesk := ""
		solaris_linux := ""
		subversion := ""
		vnc := ""
		wiki := ""
		for _, empType := range entry.GetAttributeValues("employeeType") {
			if(empType == "disabled") {
				disabled = "checked"
			}
			if(empType == "se-team") {
				seteam = "checked"
			}
			if(empType == "jira-user") {
				jira = "checked"
			}
			if(empType == "jrebel-user") {
				jrebel = "checked"
			}
			if(empType == "nagios-user") {
				nagios = "checked"
			}
			if(empType == "owncloud-user") {
				owncloud = "checked"
			}
			if(empType == "rocketchat-user") {
				rocketchat = "checked"
			}
			if(empType == "sassy-dev-user") {
				sassydev = "checked"
			}
			if(empType == "sassy-prod-user") {
				sassyprod = "checked"
			}
			if(empType == "savvyservicedesk-user") {
				savvyservicedesk = "checked"
			}
			if(empType == "linux-user") {
				solaris_linux = "checked"
			}
			if(empType == "svn-user") {
				subversion = "checked"
			}
			if(empType == "vnc") {
				vnc = "checked"
			}
			if(empType == "wiki-user") {
				wiki = "checked"
			}
		}
		return LoadEditPage(uid, given, last, mail, display, gnum, mobile, uidNumber, homeDirectory, loginShell, disabled, seteam, jira, jrebel, nagios, owncloud, rocketchat, sassydev, sassyprod, savvyservicedesk, solaris_linux, subversion, vnc, wiki)
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
