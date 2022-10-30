package main

import (
	"fmt"
	"reflect"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/sirupsen/logrus"
)

const (
	username = "neo4j"
	password = "123456"
	port     = "11005"
	dbName   = "neo4j"
)

type sorder struct {
	Name   string
	Age    int
	Gender string
	Career string
}

var session neo4j.Session

func initital() {
	//URL:neo4j+s是加证书,具体参考文档
	uri := "neo4j://localhost:" + port + "/neo4j"
	token := neo4j.BasicAuth(username, password, "")
	driver, err := neo4j.NewDriver(uri, token)
	if err != nil {
		panic(err)
	}
	fmt.Println("driver success!")
	session = driver.NewSession(neo4j.SessionConfig{DatabaseName: dbName})
	fmt.Println("session success!")
}
func main() {
	initital()
	//写入,用闭包传送, 主要是为了能够失败时重传
	reina := sorder{
		Name:   "深见玲奈",
		Age:    18,
		Career: "枪手",
		Gender: "女",
	}
	haruto := sorder{
		Name:   "苍井春人",
		Age:    19,
		Career: "剑士",
		Gender: "男",
	}

	AddTwoNodes(reina, haruto, "master", "dog")
}

/*
	func writeTrans(tx neo4j.Transaction) (interface{}, error) {
		createRelationshipBetweenPeopleQuery := `
					MERGE (p1:Person { name: $person1_name })
					MERGE (p2:Person { name: $person2_name })
					MERGE (p1)-[:KNOWS]->(p2)
					RETURN p1, p2`
		//使用类似于正则替换的方式
		result, err := tx.Run(createRelationshipBetweenPeopleQuery, map[string]interface{}{
			"person1_name": "Alice",
			"person2_name": "David",
		})
		if err != nil {
			return nil, err
		}
		return result.Collect()
	}

	func readTrans(tx neo4j.Transaction) (interface{}, error) {
		readPersonByName := `
				MATCH (p:Person)
				WHERE p.name = $person_name
				RETURN p.name AS name`
		result, err := tx.Run(readPersonByName, map[string]interface{}{
			"person_name": "Alice",
		})
		if err != nil {
			return nil, err
		}
		// Iterate over the result within the transaction instead of using
		// Collect (just to show how it looks...). Result.Next returns true
		// while a record could be retrieved, in case of error result.Err()
		// will return the error.
		for result.Next() {
			fmt.Printf("Person name: '%s' \n", result.Record().Values[0].(string))
		}
		// Again, return any error back to driver to indicate rollback and
		// retry in case of transient error.
		return result, result.Err()
	}
*/
func getString(s sorder) string {
	str := fmt.Sprintf("name:%s, age:%d, gender:%s, career:%s", s.Name, s.Age, s.Gender, s.Career)
	return str
}

// 添加两个节点
func AddTwoNodes(A interface{}, B interface{}, AtoB string, BtoA string) error {
	_, err := session.WriteTransaction(
		func(tx neo4j.Transaction) (interface{}, error) {
			CQL := `MERGE (p1:Person { ` + EnumStruct(A) + ` })
				MERGE (p2:Person { ` + EnumStruct(B) + ` })`
			if AtoB != "" {
				CQL += `MERGE (p1)-[:` + AtoB + `]->(p2)`
			}
			if AtoB != "" {
				CQL += `MERGE (p1)-[:` + BtoA + `]->(p2)`
			}

			result, err := tx.Run(CQL, nil)
			if err != nil {
				logrus.Errorf("[AddTwoNodes] merge fail, %+v", err)
				return nil, err
			}
			return result.Collect()
		})
	if err != nil {
		logrus.Errorf("[AddTwoNodes] merge fail, %+v", err)
	}
	return nil
}

/*
func read() {
	_, err := session.ReadTransaction(readTrans)
	if err != nil {
		panic(err)
	}
}*/

func EnumStruct(src interface{}) string {
	_type := reflect.TypeOf(src)
	_value := reflect.ValueOf(src)
	// 不是结构体
	if _type.Kind() != reflect.Struct {
		return ""
	}
	var str string
	for i := 0; i < _type.NumField(); i++ {
		typeField := _type.Field(i)
		valueField := _value.Field(i)

		switch typeField.Type.Kind() {
		case reflect.String:
			str += fmt.Sprintf("%s:\"%s\"", typeField.Name, valueField.String())
		case reflect.Int:
			str += fmt.Sprintf("%s:%d", typeField.Name, valueField.Int())
		case reflect.Bool:
			if valueField.Bool() {
				str += fmt.Sprintf("%s:true", typeField.Name)
			} else {
				str += fmt.Sprintf("%s:false", typeField.Name)
			}
		}
		if i != _type.NumField()-1 {
			str += ", "
		}
	}
	return str
}
