package cola_prioridad

const (
	CAPACIDAD_INICIAL   = 20
	FACTOR_REDIMENSION  = 2
	LIMITE_CANT_OCUPADA = FACTOR_REDIMENSION * FACTOR_REDIMENSION
	PANIC_COLA_VACIA    = "La cola esta vacia"
	VALOR_INICIAL       = 0
)

/*
****************************************************************
-----------------ESTRUCTURA DEL TDA-----------------------------
****************************************************************
*/

type colaConPrioridad[T any] struct {
	datos []T
	cant  int
	cmp   func(T, T) int
}

/*
****************************************************************
-----------------CREACION DEL TDA-------------------------------
****************************************************************
*/

func CrearHeap[T any](func_cmp func(T, T) int) ColaPrioridad[T] {
	return &colaConPrioridad[T]{cmp: func_cmp, datos: make([]T, CAPACIDAD_INICIAL)}
}

func CrearHeapArr[T any](arr []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	cola := new(colaConPrioridad[T])
	cola.cmp = funcion_cmp
	cola.cant = len(arr)
	if len(arr) == VALOR_INICIAL {
		cola.datos = make([]T, CAPACIDAD_INICIAL)
		return cola
	}
	arrCopia := make([]T, len(arr))
	copy(arrCopia, arr)
	cola.datos = arrCopia
	posicion := cola.cant - 1
	cola.crearHeapify(posicion)
	return cola
}

func HeapSort[T any](elementos []T, funcion_cmp func(T, T) int) {
	cola := new(colaConPrioridad[T])
	cola.cmp = funcion_cmp
	cola.datos = elementos
	cola.cant = len(elementos)
	posicion := cola.cant - 1
	reducirPosicion := posicion - 1
	cola.crearHeapify(posicion)
	for i := 0; i < posicion; i++ {
		cola.swap(VALOR_INICIAL, posicion-i)
		cola.downHeap(VALOR_INICIAL, reducirPosicion-i)
	}
	elementos = cola.datos
}

/*
****************************************************************
-----------------IMPLEMENTACION PRIMITIVAS----------------------
****************************************************************
*/

func (cola colaConPrioridad[T]) EstaVacia() bool {
	return cola.cant == VALOR_INICIAL
}

func (cola *colaConPrioridad[T]) Encolar(dato T) {
	cola.verAumentarCapacidad()
	cola.datos[cola.cant] = dato
	cola.upHeap(cola.cant)
	cola.cant++
}

func (cola colaConPrioridad[T]) VerMax() T {
	if cola.EstaVacia() {
		panic(PANIC_COLA_VACIA)
	}
	return cola.datos[VALOR_INICIAL]
}

func (cola *colaConPrioridad[T]) Desencolar() T {
	if cola.EstaVacia() {
		panic(PANIC_COLA_VACIA)
	}

	cola.verDisminuirCapacidad()
	elemento := cola.datos[VALOR_INICIAL]
	cola.cant--
	cola.swap(VALOR_INICIAL, cola.cant)
	reducirPosicion := cola.cant - 1
	cola.downHeap(VALOR_INICIAL, reducirPosicion)
	return elemento
}

func (cola colaConPrioridad[T]) Cantidad() int {
	return cola.cant
}

/*
****************************************************************
-----------------METODOS AUXILIARES INTERNOS--------------------
****************************************************************
*/

func (cola *colaConPrioridad[T]) crearHeapify(posicion int) {
	//realizar un downHeap desde el ultimo elemento hasta la raiz
	if posicion < VALOR_INICIAL {
		return
	}
	var (
		posicionMaxima  = cola.cant - 1
		reducirPosicion = posicion - 1
	)
	cola.downHeap(posicion, posicionMaxima)
	cola.crearHeapify(reducirPosicion)
}
func (cola *colaConPrioridad[T]) upHeap(posicionHijo int) {
	//upHeap : si el padre es menor que el hijo , intercambiamos
	if posicionHijo == VALOR_INICIAL {
		return
	}
	posicionPadre := (posicionHijo - 1) / 2
	if cola.cmp(cola.datos[posicionPadre], cola.datos[posicionHijo]) > 0 {
		return
	}
	cola.swap(posicionPadre, posicionHijo)
	cola.upHeap(posicionPadre)
}

func (cola *colaConPrioridad[T]) downHeap(inicio, fin int) {
	//downHeap : Buscamos los hijos del padre e intercambiamos entre ellos el mayor,
	posHijo := (2 * inicio) + 1
	//si la posicion del hijo es mayor al rango de elementos,cortamos
	if posHijo > fin {
		return
	}
	if posHijo < fin {
		posHijo = cola.obtenerHijoMayor(posHijo, posHijo+1)
	}
	if cola.cmp(cola.datos[posHijo], cola.datos[inicio]) < 0 {
		return
	}
	cola.swap(inicio, posHijo)
	cola.downHeap(posHijo, fin)
}

func (cola *colaConPrioridad[T]) obtenerHijoMayor(hijoIzq, hijoDer int) int {
	// obtiene el hijo mayor, recibe por parametro las posiciones
	if cola.cmp(cola.datos[hijoIzq], cola.datos[hijoDer]) > 0 {
		return hijoIzq
	}
	return hijoDer
}

func (cola *colaConPrioridad[T]) verAumentarCapacidad() {
	//Si la Cantidad de elementos es igual a la capacidad, redimensiono hacia arriba
	if cola.cant == cap(cola.datos) {
		cola.redimensionar(cap(cola.datos) * FACTOR_REDIMENSION)
	}
}
func (cola *colaConPrioridad[T]) verDisminuirCapacidad() {
	//Si la cantidad ocupada es 1/4 menos que la capacidad del array y la capacidad es mayor que la que iniciamos , redimensiono hacia abajo
	if cola.cant*LIMITE_CANT_OCUPADA <= cap(cola.datos) && cap(cola.datos) > CAPACIDAD_INICIAL {
		cola.redimensionar(cap(cola.datos) / FACTOR_REDIMENSION)
	}
}

func (cola *colaConPrioridad[T]) redimensionar(tamanio int) {
	//Creo un nuevo array con diferente tamanio copiando los elementos del array anterior
	arr := make([]T, tamanio)
	copy(arr, cola.datos)
	cola.datos = arr
}

func (cola *colaConPrioridad[T]) swap(posicionPadre, posicionHijo int) {
	//intercambio posiciones del array
	cola.datos[posicionPadre], cola.datos[posicionHijo] = cola.datos[posicionHijo], cola.datos[posicionPadre]
}
