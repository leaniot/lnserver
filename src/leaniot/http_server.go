package leaniot
import (
    "net/http"
    "log"
 
    "github.com/gorilla/mux"
)

type HttpServer struct {
    ListenPort string
}

const (
    SUCCESS_STR = "{\"isSuccess\":true}"
    ERR_STR = "{\"isSuccess\":false, \"errMsg\":\"TidbOpError\"}"
)
func (hs *HttpServer) PutSensor(w http.ResponseWriter, r *http.Request) {

    err := writeToTidb()
    if err != nil {
        w.Write([]byte(ERR_STR))
    } else {
        w.Write([]byte(SUCCESS_STR))
    }
}

func (hs *HttpServer) GetAllSensor(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("getsensor"))
}


func (hs *HttpServer) Start() {
    r := mux.NewRouter()
    // Routes consist of a path and a handler function.

    r.HandleFunc("/api/v1/sensor", hs.GetAllSensor).Methods("GET")
    r.HandleFunc("/api/v1/sensor", hs.PutSensor).Methods("PUT")

    err := http.ListenAndServe(hs.ListenPort, r)
    log.Fatal(err)
}
