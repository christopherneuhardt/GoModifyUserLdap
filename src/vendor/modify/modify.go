package modify

import (
	"fmt"
	"log"
	"gopkg.in/ldap.v2"
	"os"
	"net/http"
)

func Modify(res http.ResponseWriter, req *http.Request){
	
	uid := req.FormValue("uid")
	firstname := req.FormValue("firstname")
	lastname := req.FormValue("lastname")
	email := req.FormValue("email")
	gidnumber := req.FormValue("gidnumber")
	uidnumber := req.FormValue("uidnumber")
	homedirectory := req.FormValue("firstname")
	displayname := req.FormValue("displayname")
	loginshell := req.FormValue("loginshell")
	mobile := req.FormValue("mobile")
	
	
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
	
	var fields []string
	var values []string
	
	if (firstname != ""){
		fields = append(fields, "givenName")
		values = append(values, firstname)
	}
	if (lastname != ""){
		fields = append(fields, "sn")
		values = append(values, lastname)
	}
	if (email != ""){
		fields = append(fields, "mail")
		values = append(values, email)
	}
	if (displayname != ""){
		fields = append(fields, "displayName")
		values = append(values, displayname)
	}
	if (mobile != ""){
		fields = append(fields, "mobile")
		values = append(values, mobile)
	}
	if (gidnumber != ""){
		fields = append(fields, "gidNumber")
		values = append(values, gidnumber)
	}
	if (uidnumber != ""){
		fields = append(fields, "uidNumber")
		values = append(values, uidnumber)
	}
	if (homedirectory != ""){
		fields = append(fields, "homeDirectory")
		values = append(values, homedirectory)
	}
	if (loginshell != ""){
		fields = append(fields, "loginShell")
		values = append(values, loginshell)
	}
	
	
	// Add a description, and replace the mail attributes
	modify := ldap.NewModifyRequest("uid=" + uid + ",ou=People,dc=spg,dc=cgi,dc=com")
	//modify.Add("description", []string{"An example user"})
	//modify.Replace("givenName", []string{firstname})
	//modify.Replace("sn", []string{lastname})
	//modify.Replace("mail", []string{email})
	//modify.Replace("displayName", []string{displayname})
	//modify.Replace("mobile", []string{mobile})
	
	for i := 0; i < len(fields); i++ {
        modify.Replace(fields[i], []string{values[i]})
    }

	err = l.Modify(modify)
	if err != nil {
		log.Fatal(err)
	}
}