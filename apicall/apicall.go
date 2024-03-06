package apicall

// import (
// 	"encoding/json"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// )

// type WeatherResponse struct {

// 	Main Main `json:"current"`
// }

// type Main struct {
// 	Temp float64 `json:"temp_c"`
// 	Humidity float64 `json:"humidity"`
// 	TempFeelsLike float64 `json:"feelslike_c"`
// 	Cloud float64 `json:"cloud"`
// }


// func main(){
// 	response, err := http.Get(`http://api.weatherapi.com/v1/current.json?key=d4c47ebde11148cabc232448233112&q=Jakarta&aqi=no`)
// 	if err != nil{
// 		log.Fatal(err)
// 	}

// 	log.Print(response.StatusCode)
// 	bytes, errRead := ioutil.ReadAll(response.Body)

// 	defer func(){
// 		e := response.Body.Close()
// 		if e != nil {
// 			log.Fatal(e)
// 		}
// 	}()
// 	if errRead != nil {
// 		log.Fatal(errRead)
// 	}

// 	// log.Print(string(bytes))
// 	var weatherResponse WeatherResponse
// 	errUnmarshal := json.Unmarshal(bytes, &weatherResponse)
// 	if errUnmarshal != nil {
// 		log.Fatal(errUnmarshal)
// 	}

// 	log.Printf("%+v",weatherResponse )
// }

