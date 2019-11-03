//Api rest sencilla golang

package main

import(
	"encoding/json"
	"./db"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)
//Get products in database



//Datos de productos
type Product struct {
	ID string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Price float64 `json:"price,omitempty"`
	Category *Category `json:"category,omitempty"` //endpoint to Category products in database
}

//Gato crying
type Category struct {
	Services string `json:"services,omitempty"`
	Electronics string `json:"electronics,omitempty"`
	Fitness string `json:"fitness,omitempty"`
	Other string `json:"other,omitempty"`
}
//Obtiene listado de productos
func GetProducts(w http.ResponseWriter, req http.Request){

}


//Obtiene un producto especifico por su id
func GetOneproduct(w http.ResponseWriter, req http.Request){

}

//Crea un producto y lo asigna a la database
func CreateOneProducts(w http.ResponseWriter, req http.Request) {


}

func DeleteOneProduct(){

}
func main(){
	router:=mux.NewRouter()

	//Aca se hacen consultas a la base de datos
	router.HandleFunc("/products",GetProducts).Methods("GET")//Muestra listado de productos
	router.HandleFunc("/products/{id}",GetOneproduct).Methods("GET")//Sirve para buscar un producto especifico poor su id
	//Aca se añade los nuevos datos
	router.HandleFunc("/products/{id}",CreateOneProducts).Methods("POST")
	router.HandleFunc("/products/{id}",DeleteOneProduct).Methods("POST")

	log.Fatal(http.ListenAndServe("3000",router))

}