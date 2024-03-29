package Database


import (
	"log"
)

func DatabaseByEnvironment(Environment string) (string){
	db := DatabaseInitAll("/etc/dm/","GenService", "GenService.username", "GenService.password", "GenService.dbhost")
	var db_ip string;
	//var databaseip string;
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	defer tx.Rollback()
	stmt, err := tx.Query("SELECT db_ip FROM devops_tools.tbl_database_info  WHERE db_env = ?", Environment)
	if err != nil {
		log.Fatal("There was a problem up the db.")
	}
	defer stmt.Close()



	for stmt.Next() {
		err := stmt.Scan(&db_ip)
		if err != nil {
			stmt.Close()
			log.Fatal(err)
		}
	}

	return db_ip

}

// Returns ip address of all databases
func DatabaseAllIP() []string {
	var databaseip  []string
	db := DatabaseInitAll("/etc/dm/","GenService", "GenService.username", "GenService.password", "GenService.dbhost")
	var db_ip string;
	//var databaseip string;
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Query("SELECT db_ip FROM devops_tools.tbl_database_info")
	if err != nil {
		log.Fatal("There was a problem up the db.")
	}
	defer stmt.Close()



	for stmt.Next() {
		err := stmt.Scan(&db_ip)
		if err != nil {
			stmt.Close
			log.Fatal(err)
		}
		databaseip = append(databaseip, db_ip)
	}

	return databaseip

}

func DatabaseAllHostEnabled() []string {
	var databaseip  []string
	db := DatabaseInitAll("/etc/dm/","GenService", "GenService.username", "GenService.password", "GenService.dbhost")
	var db_ip string;
	//var databaseip string;
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	defer tx.Rollback()
	stmt, err := tx.Query("SELECT db_name FROM devops_tools.tbl_database_info where Enabled  = 1")
	if err != nil {
		stmt.Close()
		log.Fatal("There was a problem up the db.")
	}
	defer stmt.Close()



	for stmt.Next() {
		err := stmt.Scan(&db_ip)
		if err != nil {
			log.Fatal(err)
		}
		databaseip = append(databaseip, db_ip)
	}

	return databaseip

}


func DatabaseDemo() []string {
	var databaseip  []string
	db := DatabaseInitAll("/etc/dm/","GenService", "GenService.username", "GenService.password", "GenService.dbhost")
	var db_ip string;
	//var databaseip string;
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	defer tx.Rollback()
	stmt, err := tx.Query("SELECT db_name FROM devops_tools.tbl_database_info_test where Enabled = 1")
	if err != nil {
		log.Fatal("There was a problem up the db.")
	}
	defer stmt.Close()



	for stmt.Next() {
		err := stmt.Scan(&db_ip)
		if err != nil {
			stmt.Close()
			log.Fatal(err)
		}
		databaseip = append(databaseip, db_ip)
	}

	return databaseip

}


// Garbage Function that does a really simple thing that is way over complicated
func GetEnvironmentByHostname(database_name string) string {
	var db_env string;
	var EnvironmentName string;
	db := DatabaseInitAll("/etc/dm/","GenService", "GenService.username", "GenService.password", "GenService.dbhost")
	//var databaseip string;
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	defer tx.Rollback()
	stmt, err := tx.Query("SELECT db_env FROM devops_tools.tbl_database_info where db_name = ?", database_name)
	if err != nil {
		log.Fatal("There was a problem up the db.")
	}
	defer stmt.Close()

	for stmt.Next() {
		err := stmt.Scan(&db_env)
		if err != nil {
			stmt.Close()
			log.Fatal(err)
		}
		EnvironmentName = db_env

	}
	 return EnvironmentName

}




