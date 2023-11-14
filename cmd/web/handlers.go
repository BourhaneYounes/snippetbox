package main

import(
  "fmt"
  "net/http"
  "strconv"
)

func home(w http.ResponseWriter, r *http.Request){

  if r.URL.Path != "/" {
    http.NotFound(w, r)
    return
  }
  w.Write([]byte("home page"))
}

func snippetView(w http.ResponseWriter, r *http.Request){
  
  id, err := strconv.Atoi(r.URL.Query().Get("id"))
  if err != nil || id < 1{
    http.NotFound(w, r)
    return
  }
  fmt.Fprintf(w, "Display snippet with ID %d..", id)
}

func snippetCreate(w http.ResponseWriter, r *http.Request){
  if r.Method != http.MethodPost {
    
    w.Header().Set("Allow", http.MethodPost)
    http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    return
  }
  w.Write([]byte("create a new snippet"))
}