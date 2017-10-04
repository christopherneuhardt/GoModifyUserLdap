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
	
	var empValues []string
	var fields []string
	var values []string
	
	var delValues []string
	
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
	
	
	searchRequest := ldap.NewSearchRequest(
    "uid=rclevinger,ou=People,dc=spg,dc=cgi,dc=com", // The base dn to search
    ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
	"(&(uid="+ uid +"))",// The filter to apply
    []string{"employeeType" },                    // A list attributes to retrieve
    nil,
	)
	
	sr, err := l.Search(searchRequest)
	if err != nil {
    log.Fatal(err)
	}
	
	disabled := false
	seteam := false
	jira := false
	jrebel := false
	nagios := false
	owncloud := false
	rocketchat := false
	sassydev := false
	sassyprod := false
	savvyservicedesk := false
	solaris_linux := false
	subversion := false
	vnc := false
	wiki := false
	for _, entry := range sr.Entries {
		for _, empType := range entry.GetAttributeValues("employeeType") {
			if(empType == "disabled") {
				disabled = true
			}
			if(empType == "se-team") {
				seteam = true
			}
			if(empType == "jira-user") {
				jira = true
			}
			if(empType == "jrebel-admin") {
				jrebel = true
			}
			if(empType == "nagios-user") {
				nagios = true
			}
			if(empType == "owncloud-user") {
				owncloud = true
			}
			if(empType == "rocketchat-user") {
				rocketchat = true
			}
			if(empType == "sassy-dev-user") {
				sassydev = true
			}
			if(empType == "sassy-prod-user") {
				sassyprod = true
			}
			if(empType == "savvyservicedesk-user") {
				savvyservicedesk = true
			}
			if(empType == "linux-user") {
				solaris_linux = true
			}
			if(empType == "svn-user") {
				subversion = true
			}
			if(empType == "vnc") {
				vnc = true
			}
			if(empType == "wiki-user") {
				wiki = true
			}
		}
	}
	
	
	
	
	if (req.FormValue("disabled") == "on"){
		empValues = append(empValues, "disabled")
	} else if (disabled == true){
		delValues = append(delValues, "disabled")
	}
	if (req.FormValue("seteamaccess") == "on"){
		empValues = append(empValues, "se-team")
	} else if (seteam == true){
		delValues = append(delValues, "se-team")
	}
	if (req.FormValue("jiraaccess") == "on"){
		empValues = append(empValues, "jira-user")
	} else if (jira == true){
		delValues = append(delValues, "jira-user")
	}
	if (req.FormValue("jrlsaccess") == "on"){
		empValues = append(empValues, "jrebel-admin")
	} else if (jrebel == true){
		delValues = append(delValues, "jrebel-admin")
	}
	if (req.FormValue("nagiosaccess") == "on"){
		empValues = append(empValues, "nagios-user")
	} else if (nagios == true){
		delValues = append(delValues, "nagios-user")
	}
	if (req.FormValue("owncloudaccess") == "on"){
		empValues = append(empValues, "owncloud-user")
	} else if (owncloud == true){
		delValues = append(delValues, "owncloud-user")
	}
	if (req.FormValue("rocketchataccess") == "on"){
		empValues = append(empValues, "rocketchat-user")
	} else if (rocketchat == true){
		delValues = append(delValues, "rocketchat-user")
	}
	if (req.FormValue("sassydevaccess") == "on"){
		empValues = append(empValues, "sassy-dev-user")
	} else if (sassydev == true){
		delValues = append(delValues, "sassy-dev-user")
	}
	if (req.FormValue("sassyprodaccess") == "on"){
		empValues = append(empValues, "sassy-prod-user")
	} else if (sassyprod == true){
		delValues = append(delValues, "sassy-prod-user")
	}
	if (req.FormValue("savvyservicedeskaccess") == "on"){
		empValues = append(empValues, "savvyservicedesk-user")
	} else if (savvyservicedesk == true){
		delValues = append(delValues, "savvyservicedesk-user")
	}
	if (req.FormValue("solaccess") == "on"){
		empValues = append(empValues, "linux-user")
	} else if (solaris_linux == true){
		delValues = append(delValues, "linux-user")
	}
	if (req.FormValue("svnaccess") == "on"){
		empValues = append(empValues, "svn-user")
	} else if (subversion == true){
		delValues = append(delValues, "svn-user")
	}
	if (req.FormValue("vncaccess") == "on"){
		empValues = append(empValues, "vnc")
	} else if (vnc == true){
		delValues = append(delValues, "vnc")
	}
	if (req.FormValue("wikiaccess") == "on"){
		empValues = append(empValues, "wiki-user")
	} else if (wiki == true){
		delValues = append(delValues, "wiki-user")
	}
	
	
	// Add a description, and replace the mail attributes
	modify := ldap.NewModifyRequest("uid=" + uid + ",ou=People,dc=spg,dc=cgi,dc=com")
	//modify.Add("description", []string{"An example user"})
	
	for i := 0; i < len(fields); i++ {
        modify.Replace(fields[i], []string{values[i]})
    }
	
	if (len(empValues) > 0){
        modify.Replace("employeeType", empValues)
	}
		
	if (len(delValues) > 0){
        modify.Delete("employeeType", delValues)
	}

	err = l.Modify(modify)
	if err != nil {
		log.Fatal(err)
	}
}