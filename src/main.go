package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
	"encoding/json"
	"math/rand"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Bienvenido!!!!!!!")
	fmt.Println("Ingrese las opciones")
	fmt.Println("---------------------")
	for {
		fmt.Println("--------------------------------------------------Ingrese url del balanceador de carga-----------------------------------------")

		url, _ := reader.ReadString('\n')
		// convert CRLF to LF
		url = strings.Replace(url, "\n", "", -1)
		fmt.Println("Ingrese la cantidad de hilos---------------------------------")
		hilos, _ := reader.ReadString('\n')
		hilos = strings.Replace(hilos, "\n", "", -1)
		fmt.Println("Ingrese la cantidad de solicitudes---------------------------")
		solicitudes, _ := reader.ReadString('\n')
		solicitudes = strings.Replace(solicitudes, "\n", "", -1)
		fmt.Println("Ruta del archivo---------------------------------------------")
		ruta, _ := reader.ReadString('\n')
		ruta = strings.Replace(ruta, "\n", "", -1)
		fmt.Println()
		fmt.Println("Iniciando....................................................")
		fmt.Println()
		enviar(url, hilos, solicitudes, ruta)
	}
}
func enviar(url string, hilos string, solicitudes string, archivo string) {
	inthilos, err := strconv.Atoi(hilos)
	intsolicitudes, err2 := strconv.Atoi(solicitudes)
	if err != nil || err2 != nil {
		fmt.Println("Error en los valores de solicitudes")
		return
	}
	contenido, err3 := ioutil.ReadFile(archivo)
	//fmt.Println("Contenido ", string(contenido))
	if err3 != nil {
		fmt.Println("Error, no se pudo abrir el archivo")
		return
	}
	txt := string(contenido)
	for i := 0; i < inthilos; i++ {
		fmt.Println("Hilo ", i)
		go thread(url, intsolicitudes, txt)
	}

}
type persona struct {
    Nombre   string
	Departamento string
	Edad int
	Forma string
	Estado string
}
func thread(url string, cantidad int, archivo string) {
	bytes := []byte(archivo)
    // Unmarshal string into structs.
    var people []persona
	json.Unmarshal(bytes, &people)
	for i := 0; i < cantidad; i++ {
		indi:=rand.Intn(len(people))
		fmt.Println(people[indi])
	}
	time.Sleep(120)
}
