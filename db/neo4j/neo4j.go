package neo4j

import "github.com/neo4j/neo4j-go-driver/v4/neo4j"

func HelloWorld(uri, username, password string) (string, error) {
	driver, err := neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		return "", err
	}
	defer driver.Close()

	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	greeting, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		result, err := tx.Run(
			"CREATE (a:Greeting) SET a.message = $message RETURN a.message + ', from node ' + id(a)",
			map[string]interface{}{"message": "hello, world"})
		if err != nil {
			return nil, err
		}
		if result.Next() {
			return result.Record().Values[0], nil
		}
		return nil, result.Err()
	})
	if err != nil {
		return "", err
	}
	return greeting.(string), nil
}
