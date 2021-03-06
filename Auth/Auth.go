package Auth

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/DMEvanCT/GoBase/Database"
	"github.com/spf13/viper"
	"log"
)

type Authenticated struct {
	Authenticated string
}


// This was taken from STACKOVERFLOW because I'm lazy
func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func GetSalt() string {
	viper.AddConfigPath("/etc/dm/")
	viper.SetConfigName("GenService")
	viper.ReadInConfig()
	salt := viper.GetString("GenService.salt")

	return salt
}


func AuthenticatedUser(akey, username string) (bool) {

	var apikey string;
	var authenticated bool;

	db := Database.DatabaseInitAuth()

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	defer tx.Rollback()
	stmt, err := tx.Query("SELECT apikey FROM accontrol.tbl_users WHERE username = ?", username )
	if err != nil {
		log.Fatal("There was a problem looking you up.")
	}
	defer stmt.Close()



	for stmt.Next() {
		err := stmt.Scan(&apikey)
		if err != nil {
			log.Fatal(err)
		}
		if apikey != akey {
			log.Println("You are not authenticated")
			authenticated := false

			return authenticated
		}

		if apikey == akey {
			log.Println("You are authenticated!")
			authenticated := true
			return authenticated
		}

	}

	return authenticated

}

func AuthorizedUser(username, service  string) bool {

	var Authorized int;
	var Authed bool;
	db := Database.DatabaseInitAll("/etc/dm/", "GenService", "GenService.username", "GenService.password", "GenService.dbhost")
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	defer tx.Rollback()
	stmt, err := tx.Query("SELECT Authorized FROM accontrol.vw_authcheck WHERE username = ? and MicroserviceName = ?", username, service)
	if err != nil {
		log.Fatal("There was a problem looking you up.")
	}
	defer stmt.Close()



	for stmt.Next() {
		err := stmt.Scan(&Authorized)
		if err != nil {
			log.Fatal(err)
		}
		if Authorized == 0 {
			log.Println("You are not Authorized")
			Authed := false

			return Authed
		}

		if Authorized == 1 {
			log.Println("You are Authorized!")
			authenticated := true
			return authenticated
		}

	}

	return Authed

}

// Used to authorize users to specific environments per service.
func AuthorizeByEnv(username, service, envfromdb string) bool {
	var Authorized int;
	var Authed bool;
	db := Database.DatabaseInitAll("/etc/dm/", "GenService", "GenService.username", "GenService.password", "GenService.dbhost")
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	defer tx.Rollback()
	stmt, err := tx.Query("SELECT Authorized FROM accontrol.vw_authcheck WHERE Username = ? and MicroServiceName = ? and Environment = ?", username, service, envfromdb)
	if err != nil {
		log.Fatal("There was a problem looking you up.", err)
	}
	defer stmt.Close()



	for stmt.Next() {
		err := stmt.Scan(&Authorized)
		if err != nil {
			log.Fatal(err)
		}
		if  Authorized == 0 {
			log.Println("You are not Authorized")
			Authed := false

			return Authed
		}
		if Authorized == 1  {
			log.Println("You are Authorized by environment!")
			authenticated := true
			return authenticated
		} else {
			Authed := false
			return Authed
		}

	}

	return Authed

}




func AuthorizeAuthenticate(apikey, username, service  string) bool {
	var AllowUser bool = true
	Authenticated := AuthenticatedUser(apikey, username)
	Authorized := AuthorizedUser(username, service)

	if !Authenticated  {
		AllowUser = false
		log.Println(username + " " + "was not able to authenticate")
		return  AllowUser
	}

	if !Authorized {
		log.Println(username + " " + "was not authorized for " + service)
		AllowUser = false
		return AllowUser
	}
	return AllowUser
}

func AuthorizeAuthenticateWithEnv(apikey, username, service, envbydb  string) bool {
	var AllowUser bool = true
	salt := GetSalt()
	apikeymd5 := GetMD5Hash(salt + apikey + salt)
	Authenticated := AuthenticatedUser(apikeymd5, username)
	Authorized := AuthorizeByEnv(username, service, envbydb)

	if !Authenticated  {
		AllowUser = false
		log.Println(username + " " + "was not able to authenticate")
		return  AllowUser
	}

	if !Authorized {
		log.Println(username + " " + "was not authorized for " + service)
		AllowUser = false
		return AllowUser
	}
	return AllowUser
}
