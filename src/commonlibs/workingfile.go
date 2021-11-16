package commonlibs

import (
	"fmt"
	"log"
)

func DataDemo() {
	var err error
	var stmt string
	dbconnxn := new(DBConnection)
	if err = dbconnxn.InitDBConnection("127.0.0.1", "", ""); err != nil {
		fmt.Printf("%s", err)
	}
	defer dbconnxn.CloseDBConnection()

	stmt = `CREATE KEYSPACE IF NOT EXISTS camldemo`
	if err = dbconnxn.ExecStatement(stmt); err != nil {
		fmt.Printf("%s", err)
		return
	}

	stmt = `DROP TABLE IF EXISTS camldemo.employee`
	if err = dbconnxn.ExecStatement(stmt); err != nil {
		fmt.Printf("%s", err)
		return
	}

	stmt = `CREATE TABLE camldemo.employee (id int PRIMARY KEY,
		name varchar,
		language varchar)`
	if err = dbconnxn.ExecStatement(stmt); err != nil {
		fmt.Printf("%s", err)
		return
	}

	{
		stmt = `INSERT INTO camldemo.employee (id, name, language) Values (?,?,?)`
		insertValues := map[string]interface{}{
			"id":       100,
			"name":     "Raj",
			"language": "python",
		}
		if err = dbconnxn.InsertStatement(stmt, insertValues); err != nil {
			fmt.Printf("%s", err)
			return
		}
	}

	var returnRecords []map[string]interface{}
	sqlStmt := `SELECT id,name,language FROM camldemo.employee`
	if returnRecords, err = dbconnxn.SelectStatement(sqlStmt); err != nil {
		fmt.Printf("%s", err)
		return
	}
	fmt.Printf("%v", returnRecords)
}

func LoadFileData() {
	controlsFile := "../etc/ccm_controls.json"
	metricsFile := "../etc/ccm_metrics.json"

	controlsVO := new(ControlsVO)
	if err := controlsVO.LoadControlsData(controlsFile, "json"); err != nil {
		log.Printf("%s", err)
	} else {
		log.Printf("%s\n", controlsVO)
	}

	metricsVO := new(MetricsVO)
	if err := metricsVO.LoadMetricsData(metricsFile, "json"); err != nil {
		log.Printf("%s", err)
	} else {
		fmt.Printf("%s\n", metricsVO)
	}

}

func GenerateSampleData() {

	// if _, err := generateControlsData(); err != nil {
	// 	log.Fatalf("Error generating ccm controls data\n")
	// }

	// if _, err := generateMetricsData(); err != nil {
	// 	log.Fatalf("Error generating ccm controls data\n")
	// }

	// if _, err := generateMetricsConfiguration(); err != nil {
	// 	log.Fatalf("Error generating ccm controls data\n")
	// }

	if _, err := GenerateMetricsRuntime(); err != nil {
		log.Fatalf("Error generating ccm controls data\n")
	}

}
