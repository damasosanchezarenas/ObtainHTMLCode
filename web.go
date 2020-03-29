//In this class we are gonna get the html code of a URL.

package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
)

//This struc implements de interface of http
type desktopWeb struct{}

//We need implements this metod.
func (desktopWeb) Write(p []byte) (int, error) {
	fmt.Println(string(p))
	return len(p), nil
}

func leerEntrada() string {

	scanner := bufio.NewScanner(os.Stdin) //Scanner declaration
	scanner.Scan()
	operacion := scanner.Text() //Stores what we type on the keyboard.

	return operacion
}

func main() {

	request := leerEntrada()

	response, err := http.Get(request)
	if err != nil {
		fmt.Println("Error getting response: ", err)
	} else {
		d := desktopWeb{}
		io.Copy(d, response.Body)
	}
}
