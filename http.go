package myORMTools

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func RegisterAll(r *mux.Router, rutas []Route) {
	for _, ruta := range rutas {
		rr := r.HandleFunc(ruta.Url, JsonResponse(ruta.Handler))
		if ruta.Methods != nil {
			rr.Methods(*ruta.Methods...)
		}
	}
}

type Route = struct {
	Url     string
	Handler http.HandlerFunc
	Methods *[]string
}

//func ToJSONError(error JsonableError) ([]byte, error) {
//	fmt.Printf("%v", error+"!!!")
//	return json.Marshal(&struct{ message string }{error})
//}

type JsonError struct {
	Err string `json:"error"`
}

//	func (e JsonError) MarshalJSON() ([]byte, error) {
//		println("ACA!")
//		return json.Marshal(&struct {
//			error string
//		}{
//			e.Err,
//		})
//	}
func JSONError(w http.ResponseWriter, err string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	//j := JsonError{error(err)}
	er := json.NewEncoder(w).Encode(JsonError{err})
	if er != nil {
		println(er.Error())
		return
	}
}
