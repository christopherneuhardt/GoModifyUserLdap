package main

import (
	"fmt"
	"log"
	"gopkg.in/ldap.v2"
	"os"
)

func main(){
	username := os.Args[1]
	password := os.Args[2]

	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", "162.70.12.25", 389))
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	err = l.Bind(username, password)
	if err != nil {
		log.Fatal(err)
	}
	
	// Add a description, and replace the mail attributes
	modify := ldap.NewModifyRequest("uid=rclevinger,ou=People,dc=spg,dc=cgi,dc=com")
	//modify.Add("description", []string{"An example user"})
	modify.Replace("uidNumber", []string{"56"})

	err = l.Modify(modify)
	if err != nil {
		log.Fatal(err)
	}
}