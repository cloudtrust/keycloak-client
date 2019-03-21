package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/cloudtrust/keycloak-client"
	"github.com/spf13/pflag"
)

const (
	tstRealm = "__internal"
	user     = "version"
)

// This should be oncverted into
// GetClient(accessToken string, realmName, idClient string) (kc.ClientRepresentation, error)
// GetClientRoleMappings(accessToken string, realmName, userID, clientID string) ([]kc.RoleRepresentation, error)
// AddClientRolesToUserRoleMapping(accessToken string, realmName, userID, clientID string, roles []kc.RoleRepresentation) error
// GetRealmLevelRoleMappings(accessToken string, realmName, userID string) ([]kc.RoleRepresentation, error)
// ResetPassword(accessToken string, realmName string, userID string) error
// SendVerifyEmail(accessToken string, realmName string, userID string) error

// GetRoles(accessToken string, realmName string) ([]kc.RoleRepresentation, error)
// GetRole(accessToken string, realmName string, roleID string) (kc.RoleRepresentation, error)
// GetClientRoles(accessToken string, realmName, idClient string) ([]kc.RoleRepresentation, error)
// CreateClientRole(accessToken string, realmName, clientID string, role kc.RoleRepresentation) (string, error)

func main() {
	var conf = getKeycloakConfig()
	var client, err = keycloak.New(*conf)
	if err != nil {
		log.Fatalf("could not create keycloak client: %v", err)
	}

	// Get access token
	accessToken, err := client.GetToken("master", "admin", "admin")
	if err != nil {
		log.Fatalf("could not get access token: %v", err)
	}

	err = client.VerifyToken("master", accessToken)
	if err != nil {
		log.Fatalf("could not validate access token: %v", err)
	}

	// Delete test realm
	client.DeleteRealm(accessToken, tstRealm)

	// Check existing realms
	var initialRealms []keycloak.RealmRepresentation
	{
		var err error
		initialRealms, err = client.GetRealms(accessToken)
		if err != nil {
			log.Fatalf("could not get realms: %v", err)
		}
		for _, r := range initialRealms {
			if *r.Realm == tstRealm {
				log.Fatalf("test realm should not exists yet")
			}
		}
	}

	// Create test realm.
	{
		var realm = tstRealm
		var err error
		_, err = client.CreateRealm(accessToken, keycloak.RealmRepresentation{
			Realm: &realm,
		})
		if err != nil {
			log.Fatalf("could not create keycloak client: %v", err)
		}
		fmt.Println("Test realm created.")
	}

	// Check getRealm.
	{
		var realmR, err = client.GetRealm(accessToken, tstRealm)
		if err != nil {
			log.Fatalf("could not get test realm: %v", err)
		}
		if *realmR.Realm != tstRealm {
			log.Fatalf("test realm has wrong name")
		}
		if realmR.DisplayName != nil {
			log.Fatalf("test realm should not have a field displayName")
		}
		fmt.Println("Test realm exists.")
	}

	// Update Realm
	{
		var displayName = "updated realm"
		var err = client.UpdateRealm(accessToken, tstRealm, keycloak.RealmRepresentation{
			DisplayName: &displayName,
		})
		if err != nil {
			log.Fatalf("could not update test realm: %v", err)
		}
		// Check update
		{
			var realmR, err = client.GetRealm(accessToken, tstRealm)
			if err != nil {
				log.Fatalf("could not get test realm: %v", err)
			}
			if *realmR.DisplayName != displayName {
				log.Fatalf("test realm update failed")
			}
		}
		fmt.Println("Test realm updated.")
	}

	// Count users.
	{
		var nbrUser, err = client.CountUsers(accessToken, tstRealm)
		if err != nil {
			log.Fatalf("could not count users: %v", err)
		}
		if nbrUser != 0 {
			log.Fatalf("there should be 0 users")
		}
	}

	// Create test users.
	{
		for _, u := range tstUsers {
			var username = strings.ToLower(u.firstname + "." + u.lastname)
			var email = username + "@cloudtrust.ch"
			var err error
			_, err = client.CreateUser(accessToken, tstRealm, keycloak.UserRepresentation{
				Username:  &username,
				FirstName: &u.firstname,
				LastName:  &u.lastname,
				Email:     &email,
			})
			if err != nil {
				log.Fatalf("could not create test users: %v", err)
			}

		}
		// Check that all users where created.
		{
			var nbrUser, err = client.CountUsers(accessToken, tstRealm)
			if err != nil {
				log.Fatalf("could not count users: %v", err)
			}
			if nbrUser != 50 {
				log.Fatalf("there should be 50 users")
			}
		}
		fmt.Println("Test users created.")
	}

	// Get users
	{
		{
			// No parameters.
			var users, err = client.GetUsers(accessToken, tstRealm)
			if err != nil {
				log.Fatalf("could not get users: %v", err)
			}
			if len(users) != 50 {
				log.Fatalf("there should be 50 users")
			}

			user, err := client.GetUser(accessToken, tstRealm, *(users[0].Id))
			if err != nil {
				log.Fatalf("could not get user")
			}

			if !(*(user.Username) != "") {
				log.Fatalf("Username should not be empty")
			}

			fmt.Println("Test user retrieved.")
		}
		{
			// email.
			var users, err = client.GetUsers(accessToken, tstRealm, "email", "john.doe@cloudtrust.ch")
			if err != nil {
				log.Fatalf("could not get users: %v", err)
			}
			if len(users) != 1 {
				log.Fatalf("there should be 1 user matched by email")
			}
		}
		{
			// firstname.
			var users, err = client.GetUsers(accessToken, tstRealm, "firstName", "John")
			if err != nil {
				log.Fatalf("could not get users: %v", err)
			}
			// Match John and Johnny
			if len(users) != 2 {
				log.Fatalf("there should be 2 user matched by firstname")
			}
		}
		{
			// lastname.
			var users, err = client.GetUsers(accessToken, tstRealm, "lastName", "Wells")
			if err != nil {
				log.Fatalf("could not get users: %v", err)
			}
			if len(users) != 3 {
				log.Fatalf("there should be 3 users matched by lastname")
			}
		}
		{
			// username.
			var users, err = client.GetUsers(accessToken, tstRealm, "username", "lucia.nelson")
			if err != nil {
				log.Fatalf("could not get users: %v", err)
			}
			if len(users) != 1 {
				log.Fatalf("there should be 1 user matched by username")
			}
		}
		{
			// first.
			var users, err = client.GetUsers(accessToken, tstRealm, "max", "7")
			if err != nil {
				log.Fatalf("could not get users: %v", err)
			}
			if len(users) != 7 {
				log.Fatalf("there should be 7 users matched by max")
			}
		}
		{
			// search.
			var users, err = client.GetUsers(accessToken, tstRealm, "search", "le")
			if err != nil {
				log.Fatalf("could not get users: %v", err)
			}
			if len(users) != 7 {
				log.Fatalf("there should be 7 users matched by search")
			}
		}

		fmt.Println("Test users retrieved.")
	}

	// Update user.
	{
		// Get user ID.
		var userID string
		{
			var users, err = client.GetUsers(accessToken, tstRealm, "search", "Maria")
			if err != nil {
				log.Fatalf("could not get Maria: %v", err)
			}
			if len(users) != 1 {
				log.Fatalf("there should be 1 users matched by search Maria")
			}
			if users[0].Id == nil {
				log.Fatalf("user ID should not be nil")
			}
			userID = *users[0].Id
		}
		// Update user.
		var username = "Maria"
		var updatedLastname = "updated"
		{

			var err = client.UpdateUser(accessToken, tstRealm, userID, keycloak.UserRepresentation{
				FirstName: &username,
				LastName:  &updatedLastname,
			})
			if err != nil {
				log.Fatalf("could not update user: %v", err)
			}
		}
		// Check that user was updated.
		{
			var users, err = client.GetUsers(accessToken, tstRealm, "search", "Maria")
			if err != nil {
				log.Fatalf("could not get Maria: %v", err)
			}
			if len(users) != 1 {
				log.Fatalf("there should be 1 users matched by search Maria")
			}
			if users[0].LastName == nil || *users[0].LastName != updatedLastname {
				log.Fatalf("user was not updated")
			}
		}
		fmt.Println("User updated.")
	}

	// Delete user.
	{
		// Get user ID.
		var userID string
		{
			var users, err = client.GetUsers(accessToken, tstRealm, "search", "Toni")
			if err != nil {
				log.Fatalf("could not get Toni: %v", err)
			}
			if len(users) != 1 {
				log.Fatalf("there should be 1 users matched by search Toni")
			}
			if users[0].Id == nil {
				log.Fatalf("user ID should not be nil")
			}
			userID = *users[0].Id
		}
		// Delete user.
		{
			var err = client.DeleteUser(accessToken, tstRealm, userID)
			if err != nil {
				log.Fatalf("could not delete user: %v", err)
			}
		}
		// Check that user was deleted.
		{
			var nbrUser, err = client.CountUsers(accessToken, tstRealm)
			if err != nil {
				log.Fatalf("could not count users: %v", err)
			}
			if nbrUser != 49 {
				log.Fatalf("there should be 49 users")
			}
		}
		fmt.Println("User deleted.")
	}

	// Delete test realm.
	{
		var err = client.DeleteRealm(accessToken, tstRealm)
		if err != nil {
			log.Fatalf("could not delete test realm: %v", err)
		}
		// Check that the realm was deleted.
		{
			var realms, err = client.GetRealms(accessToken)
			if err != nil {
				log.Fatalf("could not get realms: %v", err)
			}
			for _, r := range realms {
				if *r.Realm == tstRealm {
					log.Fatalf("test realm should be deleted")
				}
			}
		}
		fmt.Println("Test realm deleted.")
	}
}

func getKeycloakConfig() *keycloak.Config {
	var apiAddr = pflag.String("urlKc", "http://localhost:8080", "keycloak address")
	var tokenAddr = pflag.String("url", "http://localhost:8080", "token address")
	pflag.Parse()

	return &keycloak.Config{
		AddrTokenProvider: *tokenAddr,
		AddrAPI:           *apiAddr,
		Timeout:           10 * time.Second,
	}
}

var tstUsers = []struct {
	firstname string
	lastname  string
}{
	{"John", "Doe"},
	{"Johnny", "Briggs"},
	{"Karen", "Sutton"},
	{"Cesar", "Mathis"},
	{"Ryan", "Kennedy"},
	{"Kent", "Phillips"},
	{"Loretta", "Curtis"},
	{"Derrick", "Cox"},
	{"Greg", "Wilkins"},
	{"Andy", "Reynolds"},
	{"Toni", "Meyer"},
	{"Joyce", "Sullivan"},
	{"Johanna", "Wells"},
	{"Judith", "Barnett"},
	{"Joanne", "Ward"},
	{"Bethany", "Johnson"},
	{"Maria", "Murphy"},
	{"Mattie", "Quinn"},
	{"Erick", "Robbins"},
	{"Beulah", "Greer"},
	{"Patty", "Wong"},
	{"Gayle", "Garrett"},
	{"Stewart", "Floyd"},
	{"Wilbur", "Schneider"},
	{"Diana", "Logan"},
	{"Eduardo", "Mitchell"},
	{"Lela", "Wells"},
	{"Homer", "Miles"},
	{"Audrey", "Park"},
	{"Rebecca", "Fuller"},
	{"Jeremiah", "Andrews"},
	{"Cedric", "Reyes"},
	{"Lee", "Griffin"},
	{"Ebony", "Knight"},
	{"Gilbert", "Franklin"},
	{"Jessie", "Norman"},
	{"Cary", "Wells"},
	{"Arlene", "James"},
	{"Jerry", "Chavez"},
	{"Marco", "Weber"},
	{"Celia", "Guerrero"},
	{"Faye", "Massey"},
	{"Jorge", "Mccarthy"},
	{"Jennifer", "Colon"},
	{"Angel", "Jordan"},
	{"Bennie", "Hubbard"},
	{"Terrance", "Norris"},
	{"May", "Sharp"},
	{"Glenda", "Hogan"},
	{"Lucia", "Nelson"},
}
