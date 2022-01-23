// Package configuration - читает из флагов или из переменных окружения,
//заполняет структуру - переменную типа Configuration и возвращает эту структуру и ошибку
//go run hw8.go --port=8080 --db_url=postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable --jaeger_url=http://jaeger:16686 --sentry_url=http://sentry:9000 --some_app_id=testid --some_app_key=testkey
package configuration

import (
	"github.com/namsral/flag"
	//"github.com/joho/godotenv"
)

type Configuration struct {
	port         string
	db_url       string
	jaeger_url   string
	sentry_url   string
	kafka_broker string
	some_app_id  string
	some_app_key string

	/* Шаблон:
	port: 8080
	db_url: postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable
	jaeger_url: http://jaeger:16686
	sentry_url: http://sentry:9000
	kafka_broker: kafka:9092
	some_app_id: testid
	some_app_key: testkey
	*/
}

var set Configuration // Переменная set является структурой типа Configuration
//var err error

//Load - читает из флагов или из переменных окружения,
//заполняет структуру - переменную типа Configuration и возвращает эту структуру и ошибку
func Load() (*Configuration, error) {

	//Имя файла передается как аргументы командной строки flag.Arg
	//if flag.NArg() != 1 {
	//	log.Fatal("Неверно задано количество аргументов командной строки. Проверьте, что вы задали имя файла")
	//}
	//Читаем значения из флагов
	set.port = *flag.String("port", "x", "Enter port")
	set.db_url = *flag.String("db_url", "x", "Enter db_url")
	set.jaeger_url = *flag.String("jaeger_url", "x", "Enter jaeger_url")
	set.sentry_url = *flag.String("sentry_url", "x", "Enter sentry_url")
	set.kafka_broker = *flag.String("kafka_broker", "x", "Enter kafka_broker")
	set.some_app_id = *flag.String("some_app_id", "x", "Enter some_app_id")
	set.some_app_key = *flag.String("some_app_key", "x", "Enter some_app_key")
	flag.Parse()     //Сообщаем библиотеке flag, что необходимо считать флаги
	return &set, nil //Вернуть заполненную структуру и ошибку, а где взять ошибку?
}
