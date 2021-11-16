package commonlibs

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/yugabyte/gocql"
)

// InitDBConnection : Creates a cluster and a session and maintains pointers on the dbConnection receiver
func (dbConnection *DBConnection) InitDBConnection(hostName, userName, password string) error {

	cluster := gocql.NewCluster(hostName)
	cluster.Timeout = 12 * time.Second
	if strings.TrimSpace(userName) != "" && strings.TrimSpace(password) != "" {
		cluster.Authenticator = gocql.PasswordAuthenticator{
			Username: userName,
			Password: password,
		}
	}
	session, err := cluster.CreateSession()
	if err != nil {
		return fmt.Errorf("cant create db session: %s", err)
	}
	dbConnection.cluster = cluster
	dbConnection.session = session

	return nil
}

// CloseDBConnection: Allows the client to defer close the connection after initDBConnection
func (dbConnection *DBConnection) CloseDBConnection() error {
	if dbConnection.session.Closed() {
		return fmt.Errorf("connection already closed")
	}
	dbConnection.session.Close()
	return nil
}

// ExecStatement: Takes the wholly formed SQL statement as parameter.
// Use this method to invoke any DDL calls such as CREATE, DROP etc
// Refer to insertStatement() and selectStatement() methods for inserting records into table and for querying tables respectively
func (dbConnection *DBConnection) ExecStatement(sqlStatement string) (err error) {
	if dbConnection.session.Closed() {
		return fmt.Errorf("connection already closed")
	}
	err = dbConnection.session.Query(sqlStatement).Exec()
	if err != nil {
		return fmt.Errorf("cant exec statement %s. Error: %s", sqlStatement, err)
	}
	return nil
}

// InsertStatement: contains 2 parts: INSERT statement containing columns and map[string]interface{} containing values
// Example: sqlStatement = INSERT INTO example.table (pk, value) VALUES (?, ?)
// values should be a []interface{} with exactly the same number of values and matching types as INSERT INTO
func (dbConnection *DBConnection) InsertStatement(sqlStatement string, inputMap map[string]interface{}) (err error) {
	if dbConnection.session.Closed() {
		return fmt.Errorf("connection already closed")
	}
	fieldNames, err := dbConnection.fetchColumns(sqlStatement)
	if err != nil {
		return fmt.Errorf("cannot fetch column names %s. Error: %s", sqlStatement, err)
	}
	if len(fieldNames) <= 0 {
		return fmt.Errorf("Where are the columns? %s. Error: %s", sqlStatement, err)
	}

	values := make([]interface{}, 0)
	for _, value := range fieldNames {
		values = append(values, inputMap[strings.TrimSpace(value)])
	}

	if len(fieldNames) != len(values) {
		return fmt.Errorf("Values not matching with the columns:: %s. Error: %s", sqlStatement, err)
	}
	err = dbConnection.session.Query(sqlStatement, values...).Exec()
	if err != nil {
		return fmt.Errorf("cant exec statement %s. Error: %s", sqlStatement, err)
	}
	return nil
}

// SelectStatement: For querying records. Passes the "fully formed" sql statement as a parameter
// Example: sqlStatement = SELECT id,name FROM example.table WHERE language = 'go'
// https://github.com/yugabyte/gocql/blob/master/example_dynamic_columns_test.go
func (dbConnection *DBConnection) SelectStatement(sqlStatement string) (records []map[string]interface{}, err error) {
	if dbConnection.session.Closed() {
		return nil, fmt.Errorf("connection already closed")
	}
	iter := dbConnection.session.Query(sqlStatement).Iter()
	for {
		rowData, err := iter.RowData()
		if err != nil {
			return nil, fmt.Errorf("error fetching rowdata")
		}
		if !iter.Scan(rowData.Values...) {
			break
		}

		holdingMap := make(map[string]interface{}, 0)
		for index, rowDataVal := range rowData.Values {
			colName := rowData.Columns[index]
			rowVal := reflect.Indirect(reflect.ValueOf(rowDataVal)).Interface()
			holdingMap[colName] = rowVal
		}
		records = append(records, holdingMap)
	}

	return records, nil
}

func (dbConnection *DBConnection) fetchColumns(sqlStatement string) (fieldNames []string, err error) {
	// check if the length of values matches the INSERT INTO columns
	const filepattern = `\([a-z,\s]+\)`
	var ruleregex = regexp.MustCompile(filepattern)
	isMatch := ruleregex.MatchString(sqlStatement)
	if !isMatch {
		return []string{}, fmt.Errorf("cannot match the insert pattern. It should follow INSERT INTO schema.table (field1, field2, field3) Values (?,?,?)")
	}
	match := strings.TrimSpace(ruleregex.FindStringSubmatch(sqlStatement)[0])
	match = strings.Replace(strings.Replace(match, "(", "", -1), ")", "", -1)
	fieldNames = strings.Split(match, ",")
	if len(fieldNames) <= 0 {
		fieldNames = []string{match}
	}
	return fieldNames, nil
}
