package leaniot

func Main() {

    hs := HttpServer{ListenPort: ":8080"}
    hs.Start()
}
