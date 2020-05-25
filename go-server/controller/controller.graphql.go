package controller

import (
	"bytes"
	"log"
	"net/http"
	"runtime"
	"time"

	"go-server/gql"
	"go-server/model"

	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func ExecQuery(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var mem_s runtime.MemStats
	var mem_e runtime.MemStats
	time_s := time.Now()
	runtime.ReadMemStats(&mem_s)

	bufBody := new(bytes.Buffer)
	bufBody.ReadFrom(r.Body)

	status := http.StatusOK
	payload := model.GraphQLResponse{}
	payload.Success = true
	payload.Message = "Succeeded! your query was executed :)"

	q := bufBody.String()
	log.Println(q)

	sc := gql.GetSchemaConfig(db)
	s := gql.GetSchema(sc)

	result := graphql.Do(graphql.Params{
		Schema:        s,
		RequestString: q,
	})
	if len(result.Errors) > 0 {
		payload.Success = false
		payload.Message = "Failed! some error occured :("
		for _, e := range result.Errors {
			log.Println(e.Message)
		}
		status = http.StatusBadRequest
	}
	payload.Result = result

	time_e := time.Now()
	time_diff := time_e.Sub(time_s).Seconds()
	runtime.ReadMemStats(&mem_e)
	mem_diff := mem_e.TotalAlloc - mem_s.TotalAlloc
	measurement := model.Measurement{
		MemoryUsageChange: mem_diff,
		ProcessTime:       time_diff,
	}
	payload.Measurement = &measurement

	respondJSON(w, status, payload)
}
