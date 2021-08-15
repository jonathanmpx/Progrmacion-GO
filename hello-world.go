package  main

import(
  "net/http"
)

func main ()  {
//  println("Hello World")
http.ListenAndServer(":3000, nil")
}
