package file

import (
	"io/ioutil"
	"os"
	"strings"
	"net/http"
)

type userError string

func (e userError) Error() string {
	 return e.Message()
}

func (e userError) Message() string {
	 return string(e)
}
func Halder(writer http.ResponseWriter,
	request *http.Request)  error{
		if strings.Index(request.URL.Path, "/list/") != 0 {
       return userError("path must true")
		}
		
		path := request.URL.Path[len("/list/"):]
		file, err := os.Open(path)
		if err != nil{
					return err
		}
			defer file.Close()
			all, err := ioutil.ReadAll(file)
			if err != nil {
				return err
			}
			writer.Write(all)
			return nil
		}