package main

import (
	"fmt"
	"os"
   "./file"
	"net/http"
)
type appHandler func (writer http.ResponseWriter,
					 request *http.Request) error
func errWrapper (handler appHandler) func (
	    http.ResponseWriter, *http.Request)  {
	return func (writer http.ResponseWriter,
		      request *http.Request)  {
		defer func ()  {
			r := recover()
			if r != nil {
         fmt.Println(r)
			   http.Error(writer,
						http.StatusText(http.StatusInternalServerError),
					  http.StatusInternalServerError)
			}
			// log.Printf("Panic: %v", r)
			
		}()

		err := handler(writer, request)
		if err != nil {
			if userError, ok := err.(userError); ok {
				http.Error(writer,
						 userError.Message(),
						 http.StatusBadRequest)
				return
			}
			code := http.StatusOK
       switch  {
			 case os.IsNotExist(err):
				 code = http.StatusNotFound
			 case os.IsPermission(err):
				code = http.StatusForbidden
			 default:
				  code = http.StatusInternalServerError
			 }
			 http.Error(
				 writer,
				 http.StatusText(code),code)
		}
	}
}
type 	userError interface{
	error
	Message() string
}
func main()  {
	http.HandleFunc("/", errWrapper(file.Halder))
    
			err := http.ListenAndServe(":8080", nil)
			if err != nil {
				panic(err)
			}
}