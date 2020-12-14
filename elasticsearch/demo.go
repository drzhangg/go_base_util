package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic"
	"reflect"
)

var client *elastic.Client

var host = "http://47.103.9.218:9200"

type Employee struct {
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Age       int      `json:"age"`
	Abort     string   `json:"abort"`
	Interests []string `json:"interests"`
}

func init() {
	var err error

	client, err = elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(host))
	if err != nil {
		panic(err)
	}

	_, _, err = client.Ping(host).Do(context.Background())
	if err != nil {
		panic(err)
	}

	_, err = client.ElasticsearchVersion(host)
	if err != nil {
		panic(err)
	}
}

func main() {
	//create()
	i, t := "test", "employee"
	gets(i, t, "1")
	update(i, t, "1")
	gets(i, t, "1")
}

func create() {
	e1 := Employee{
		FirstName: "zhang",
		LastName:  "san",
		Age:       24,
		Abort:     "一个屌丝",
		Interests: []string{"打游戏"},
	}

	put1, err := client.Index().Index("test").Type("employee").Id("1").BodyJson(e1).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed tweet %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)

	e2 := `{"first_name":"tome","last_name":"jerry","age":22,"abort":"he is a very nice man","interests":["music","coding"]}`
	put2, err := client.Index().Index("test").Type("employee").Id("2").BodyString(e2).Do(context.TODO())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed tweet %s to index %s, type %s\n", put2.Id, put2.Index, put2.Type)

	e3 := `{"first_name":"tommy","last_name":"bob","age":32,"abort":"he is a very nice man","interests":["movie"]}`
	put3, err := client.Index().Index("test").Type("employee").Id("3").BodyString(e3).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed tweet %s to index %s, type %s, result %v\n", put3.Id, put3.Index, put3.Type, put3.Result)
}

func gets(index, types, id string) {
	get1, err := client.Get().Index(index).Type(types).Id(id).Do(context.Background())
	if err != nil {
		panic(err)
	}

	if get1.Found {
		fmt.Printf("document %s in version %d from index %s, type %s\n", get1.Id, get1.Version, get1.Index, get1.Type)

		data, _ := get1.Source.MarshalJSON()
		fmt.Println(string(data))

		var e Employee
		json.Unmarshal(data, &e)

		fmt.Println(e.FirstName, e.Abort, e.Interests, e.Age, e.LastName)
	}
}

func delete(index, types, id string) {
	del, err := client.Delete().Index(index).Type(types).Id(id).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println("delete result:", del.Result)
}

func update(index, types, id string) {
	up, err := client.Update().Index(index).Type(types).Id(id).Doc(map[string]interface{}{"age": 100}).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println("update result:", up.Result)
}

func query(index, types, id string, queryString string) {
	var res *elastic.SearchResult
	var err error

	//取所有
	res, err = client.Search(index).Type(types).Do(context.Background())
	printEmployee(res, err)

	//字段相等
	q := elastic.NewQueryStringQuery(queryString)
	res, err = client.Search(index).Type(types).Query(q).Do(context.Background())
	if err != nil {
		panic(err)
	}
	printEmployee(res, err)

	//条件查询
	//年龄大于30岁的
	boolQ := elastic.NewBoolQuery()
	boolQ.Must(elastic.NewMatchQuery("last_name", "tom"))
	boolQ.Filter(elastic.NewRangeQuery("age").Gt(30))
	res, err = client.Search(index).Type(types).Query(q).Do(context.Background())
	printEmployee(res, err)


	//短语搜索 搜索abort字段中有
	matchPhraseQuery := elastic.NewMatchPhraseQuery("about","一个")
	res, err = client.Search(index).Type(types).Query(matchPhraseQuery).Do(context.Background())
	printEmployee(res, err)

}

// 打印查询到的Employee
func printEmployee(res *elastic.SearchResult, err error) {
	if err != nil {
		print(err.Error())
		return
	}

	var typ Employee
	for _, item := range res.Each(reflect.TypeOf(typ)) {
		t := item.(Employee)
		fmt.Printf("%#v\n", t)
	}
}
