package main

import (
	"encoding/json"
	"fmt"
)

// fuera de funciones solo puede declararse mas no asignar por separado
//
var global_inicializada int
var edad int = 23

func main() {

	//fmt.Println("Gianlucca edad ->", edad)
	//declaraciones()
	//fmt.Println("Primer global -> ", global_inicializada)
	//global_inicializada = 8 // segundo valor asiganado a la global
	//fmt.Println("Segunda global -> ", global_inicializada)
	//format()
	//verbos()
	//tipos_propios()
	//funcion_iota()
	//f_for()
	//f_switch()
	//f_arrays()
	//f_slice()
	//f_make()
	//f_matrices()
	//f_map()
	//f_struct() // basico y se encuentra todo en su interior
	//f_func()
	//f_interfaz_metodos_estructuras()
	//f_anonima()

	//numeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//f_callback(parametros_variados, numeros...)

	//f_closure_main()
	//f_punteros()
	f_json_marsharl()
}

func declaraciones() {

	var inicializada int // si o si la tengo que usar,pero dentro de funciones
	fmt.Println("inicializada ->", inicializada)

	x := 5                  // declarar e inizializar, := infiere el tipo
	x = 3                   //reasignar valor
	fmt.Println("x -> ", x) //

	global_inicializada = 5 // primer valor asigando a la global

}

func format() {
	// string literal
	a := "este texto no admite lineas nuevas. \n"
	//row lteral
	b := `		estas comillas
		respetan los saltos 
		de linea.`

	fmt.Println(a)
	fmt.Println(b)

}

func verbos() {
	// v para valor
	// T para tipo
	// d para int
	// s para Strings

	a := 5
	b := "letra b"

	fmt.Printf("el valor de a es : %v \n", a)
	fmt.Printf("el tipo de a es : %T \n", a)

	// asignarle a un variable todo un conjunto de Strings

	s := fmt.Sprint("primero ", b, " y termina.")
	fmt.Println(s)

}

type mi_tipo2 int

var b mi_tipo2

func tipos_propios() {

	type mi_tipo int
	var a bool

	fmt.Printf(" mi propio tipo %T \n", a)
	fmt.Printf(" mi propio tipo %T \n", b)

}

// _____________- IOTA -______________

func funcion_iota() {

	const a = 0
	//incrementa de a una unidad cada constante relacionada,
	//si se hace otro const se reinicia el iota
	const (
		b = iota
		c
		d
	)

	var e int8 = 4
	var f int

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	fmt.Println("c = ", c)
	fmt.Println("d = ", d)
	fmt.Println("e = ", e)

	f = 666
	fmt.Println("f = ", f)

}

func f_for() {
	var x int
	//primera definicion de for: inicializador;predicado;aumento
	for i := 0; i <= 5; i++ {
		fmt.Print(i, ",")
		x = i
	}
	//for como while

	for x > 0 {
		fmt.Print(x, ",")
		x--
	}
	mi_slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // no pueden ser de tipo diferente

	//var mi_slice_2 []int
	//mi_slice_2 = make(int, 0) {1,2,3,4,5}

	fmt.Println(mi_slice)
	fmt.Println(cap(mi_slice))
	fmt.Printf("%T", mi_slice)

	//fmt.Printf("%T",mi_slice_2)
	//fmt.Println(mi_slice_2)

	for i, value := range mi_slice {
		fmt.Println(i, value)
	}

	for value := range mi_slice {
		fmt.Println(value)
	}

}

func f_switch() {
	var animal = "perro"

	switch animal {
	case "gato":
		fmt.Println("Tiene 7 vidas")
		break // no tiene utilidad ya que lo incorpora
	case "perro":
		fmt.Println("Mejor amigo del hombre")
		fallthrough //comando que  ejecuta el case de abajo
	case "mico":
		fmt.Println("Primate")
	default:
		fmt.Println("El animal no se encuetra")

	}

	// SIN PREDICADO

	switch {
	//case (true && false):
	//	fmt.Println("Case uno")
	case (true):
		fmt.Println("Case dos")
	case (false):
		fmt.Println("Case uno")
	}
}

func f_arrays() {

	var arreglo [5]int // forma de declarar un array de manera estatica, se llena de valores Cero
	arreglo[3] = 666
	fmt.Println(arreglo)
	fmt.Println(len(arreglo))
}

func f_slice() {

	mi_slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // no pueden ser de tipo diferente

	//var mi_slice_2 []int
	//mi_slice_2 = make(int, 0) {1,2,3,4,5}

	fmt.Println(mi_slice)
	fmt.Println(cap(mi_slice))
	fmt.Printf("%T", mi_slice)

	//fmt.Printf("%T",mi_slice_2)
	//fmt.Println(mi_slice_2)

	for i, value := range mi_slice {
		fmt.Println(i, value)
	}

	for value := range mi_slice {
		fmt.Println(value)
	}

	// ---------------------  DIVIDIR SLICE ----------------------------

	fmt.Println(mi_slice[3:])  // sin incluir
	fmt.Println(mi_slice[:3])  // incluido
	fmt.Println(mi_slice[3:8]) // sin incluir, incluido

	// ---------------------  AÑADIR SLICE ----------------------------

	mi_slice = append(mi_slice, 11, 12, 13, 14, 15)
	fmt.Println(" append uno ", mi_slice)

	mi_slice_2 := []int{16, 17, 18, 19, 20}

	mi_slice = append(mi_slice, mi_slice_2...)
	fmt.Println(" append concatenado 2 con 1  ", mi_slice)

	mi_slice = append(mi_slice[:5], mi_slice[9:]...)
	fmt.Println(" append con remove  ", mi_slice)

	// ----------------------- Relacion de slice y array Subyacente--------

	x := []int{1, 2, 3, 4, 5}
	fmt.Println("x => ", x)

	y := append(x[:2], x[4:]...)
	fmt.Println("x => ", x)
	fmt.Println("y => ", y)

}
func f_make() {

	slice := make([]int, 10, 12) // tipo, tamaño, capacidad

	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	slice = append(slice, 300)
	slice = append(slice, 400)
	slice = append(slice, 500) // se recomienda no permitir que
	// el lenguaje duplique por su cuenta la capacidad de l slice.

	fmt.Println(len(slice))
	fmt.Println(cap(slice))

}

func f_matrices() {
	pares := []int{2, 4, 6, 8}
	impares := []int{1, 3, 5, 7}

	enteros := [][]int{pares, impares}

	fmt.Println(enteros)

	slice1 := [][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 0}}

	fmt.Println(slice1)

}

// diccionarios en Python...
func f_map() {
	mi_map := map[string]int{

		"ave": 1,
		"key": 2,
	}

	fmt.Println(mi_map)

	// SABER SI UN VALOR ESTA EN EL MAPA

	if valor, ok := mi_map["No existo"]; ok {
		fmt.Println("el valor de la llave es : ", valor)
	} else {
		fmt.Println("el valor no existe en el mapa")
	}

	// Agragar un dato al MAPA

	mi_map["agregado"] = 123
	mi_map["ave"] = 444

	for key, value := range mi_map {

		fmt.Println(value, key)
	}

	// BORRAR DATO DE MAPA

	delete(mi_map, "agregado") // no retorna nada
	fmt.Println(" se borro agregado -> ", mi_map)
	// entonces
	if v, ok := mi_map["agregado"]; ok {
		fmt.Println("se borro correctamente", v)
	} else {
		fmt.Println("la llave no existe, llorelo")
	}

}

func f_struct() {

	//se puede definir afuera para que su Scope sea Global y no solo en la funcion
	type persona struct {
		nombre   string
		apellido string
		edad     int
	}

	persona_uno := persona{
		nombre:   "Gianlucca",
		apellido: "Aguado",
		edad:     23,
	}
	persona_dos := persona{
		nombre:   "Aura",
		apellido: "Cadena",
		edad:     21,
	}

	fmt.Println(persona_uno)
	fmt.Println(persona_dos)
	fmt.Println("----------")
	fmt.Println("El nombre de la primera persona es ", persona_uno.nombre)

	// STRUCT DE STRUCT --------------------------------------------------

	type embebida struct {
		persona
		frase   string
		boleano bool
	}

	embebida_uno := embebida{
		persona: persona{
			nombre:   "pepito",
			apellido: "Melechea",
		},
		frase:   "Saludo desde Go",
		boleano: true,
	}

	fmt.Println("Embebida = ", embebida_uno)
	fmt.Println("El nombre de la persona embebida es ", embebida_uno.nombre, embebida_uno.apellido)

	// STRUCTS ANONIMO ----------------------------------------------------------------------------

	anonimo := struct {
		campo_uno string
		campo_dos int
	}{
		campo_uno: "Saludame Esta",
		campo_dos: 24,
	}

	fmt.Println("Anonimo -> ", anonimo)

	// Struct CON SLICES Y MAP

	type estructura_lista struct {
		array []int
		lista []string
	}

	e_l := estructura_lista{
		array: []int{1, 2, 3, 4, 5},
		lista: []string{"H", "O", "l", "a"},
	}

	fmt.Println("Estructura con lista -> ", e_l)

	// Mapa de estructuras donde key es string y value es estructura

	m := map[string]estructura_lista{
		persona_uno.nombre: e_l,
	}

	fmt.Println(" Mapa m de String:Struct ->", m) // hacemos una mejor impresion
	for key, value := range m {
		fmt.Println(key, value.lista[0])
	}
}

// Sintaxis
// fun (r receptor) identificador (parametros) retorno(0...n) { code }
func f_func() {

	defer f_defer() // por lo tanto se ejecuta al final de f_func()

	fmt.Println(func_uno("Hola"))
	n, e := func_dos("gian", 23)
	fmt.Println(n, e)

	suma_numeros := parametros_variados(2, 4, 6, 8) //Convierte los argumentos de la funcion en un slice
	fmt.Println(suma_numeros)

	numeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	suma_numeros = parametros_variados(numeros...) //los tres puntos ayudan a hacer el despliegue individual
	fmt.Println(suma_numeros)
}
func func_uno(saludo string) string {
	return saludo
}
func func_dos(nombre string, edad int) (string, int) {
	return nombre, edad
}
func parametros_variados(numeros_pares ...int) int {

	fmt.Printf("%T \n", numeros_pares)

	var suma int

	for _, val := range numeros_pares {
		suma += val
	}

	return suma

}

func f_defer() {
	fmt.Println(" Esta funcion Difiera al final de la que la invoca...")
}

//
//
//
//
//
//
//
//
// ______________________________________________________________________
//
//
//

//
//
//
//funcion a la que se le adjunta un tipo de dato
// la structura persona recibe al metodo f_metodo e implementa a la clase humano
type persona struct {
	nombre   string
	apellido string
	edad     int
}

// la structura estudiante se extiendo de persona y recibe al metodo f_metodo e implementa a la clase humano
type estudiante struct {
	persona // estructura persona embebidad
	codigo  int
	carrera string
	activo  bool
}

func (e estudiante) f_metodos() (string, string) {
	informacion_personal := fmt.Sprint("", e.nombre, e.apellido, e.edad)
	informacion_academica := fmt.Sprint("", e.codigo, e.carrera, e.activo)

	return informacion_personal, informacion_academica
}

func (e estudiante) f_presentar() {
	fmt.Println("HOLA! soy la estudiante ", e.nombre)
}
func (p persona) f_presentar() {
	fmt.Println("HOLA! soy la persona", p.nombre)
}

//las interfaces permiten definir comportamientos

type humano interface {
	f_presentar()
	//f_metodos() (string, string)
	// las interfaces solo resiven metodos
	// ? QUE PASA SI LOS METODOS RETORNAN ALGO?
}

// Las ionterfaces definen comportamientos
func f_interfaz(h humano) {

	// AFIRMACION DE TIPO
	switch h.(type) {
	case persona:
		fmt.Println(" TIPO FORZADO persona", h.(persona).apellido)
	case estudiante:
		fmt.Println(" TIPO FORZADO estudiante", h.(estudiante).apellido)
	}

	fmt.Println("Me llamaron Interfaz ", h)
}

func f_interfaz_metodos_estructuras() {
	/* ----------------------------------------------------------- */
	// ejemplo para metodos y structuras
	// le ingresamos los datos a la estructura
	matriculado_uno := estudiante{

		persona: persona{
			nombre:   "Aura",
			apellido: "Maria",
			edad:     21,
		},

		codigo:  223344,
		carrera: "Administracion",
		activo:  true,
	}
	// accion sobre la estructura
	datos_uno, datos_dos := matriculado_uno.f_metodos()

	fmt.Println("datos uno ", datos_uno)
	fmt.Println("datos dos ", datos_dos)

	// ------------------------------------------------ FIN ejemplo

	matriculado_dos := persona{
		nombre: "Jeus",
		edad:   2021,
	}

	matriculado_dos.f_presentar()
	fmt.Println("Matriculado dos ", matriculado_dos)

	f_interfaz(matriculado_dos)

}

//____________________________________________________________________Fin

func f_anonima() {
	var saludo string
	saludo = "Y te estoy saludando"

	func(saludo string) {
		fmt.Println("Soy una funcion anonima ", saludo)
	}(saludo)

	// Exprecion Funcion, es ciudadano de primer orden

	f := func(saludo string) string {
		fmt.Println("Soy una funcion anonima ", saludo)
		return saludo
	}

	fmt.Println(f(saludo))

}

func f_callback(f func(enteros ...int) int, numeros ...int) int {
	// funcion para sumar numeros pares
	var pares []int

	for _, valor := range numeros {
		if valor%2 == 0 {
			pares = append(pares, valor)
		}
	}
	return f(pares...)
}

func f_closure_main() {

	// Modifica la seccion en memoria donde vive la funcion
	//por lo tanto se alteran los mismo valores
	// en este caso x
	a := f_closure()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

}

func f_closure() func() int {
	var x int
	return func() int { // con estas funcion se retorna una copia del entorno

		x++
		return x
	}
}

// un puntero es una direccion de memoria
//se usa cuando un valor es muy grande para estarlo pasando, es mas eficiente pasar la dirrecion de memoria
// tambien cuando se desea cambiar el valor directo en la memoria de alguna variable
// & -> para la direccion
// * -> para el valor en una direccion *int es el tipo
func f_punteros() {
	a := 42
	fmt.Println(" Valor inicial de A ", a)
	fmt.Println(" Direccion  de A ", &a)

	var b *int = &a
	fmt.Println(" Valor inicial de B ", b)
	fmt.Println(" Puntero de *b que trae el valor en esa direccion de memoria ", *b)

	*b = 999
	fmt.Println(" Valor alterado de la memoria de A ", a)

}

func f_json_marsharl() {

	// debe ir si o si la primera en mayuscula
	type person struct {
		Name    string
		Profile string
		Gender  int
	}

	persona_uno := person{

		Name:    "jesus",
		Profile: "apostol",
		Gender:  1,
	}

	persona_dos := person{

		Name:    "laura",
		Profile: "ingeniera",
		Gender:  0,
	}

	personas := []person{persona_uno, persona_dos}

	//fmt.Print(personas)

	js, err := json.Marshal(personas) // Marshal transforma la estructura en JSON

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(js))

}
