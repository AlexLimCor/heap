package cola_prioridad

const (
	CAPACACIDAD_INICIAL = 20
	FACTOR_REDIMENSION  = 2
	LIMITE_CANT_OCUPADA = FACTOR_REDIMENSION * FACTOR_REDIMENSION
	PANIC_COLA_VACIA    = "La cola esta vacia"
	VALOR_INICIAL       = 0
)

type colaConPrioridad[T any] struct {
	datos []T
	cant  int
	cmp   func(T, T) int
}

func CrearHeap[T any](func_cmp func(T, T) int) ColaPrioridad[T] {
	return &colaConPrioridad[T]{cmp: func_cmp, datos: make([]T, CAPACACIDAD_INICIAL)}
}

func CrearHeapArr[T any](arr []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	// Hacer el heapify del array
	// hacer downHeap para crear el heapify
	cola := new(colaConPrioridad[T])
	cola.cmp = funcion_cmp
	cola.cant = len(arr)
	cola.datos = arr
	cola.crearHeapify(VALOR_INICIAL, cola.cant-1)
	return cola
}

func HeapSort[T any](elementos []T, funcion_cmp func(T, T) int) {
	heap := new(colaConPrioridad[T])
	heap.cmp = funcion_cmp
	heap.datos = elementos
	heap.cant = len(elementos)
	heap.crearHeapify(VALOR_INICIAL, heap.cant-1)
	for i := 0; i < len(elementos); i++ {
		heap.swap(VALOR_INICIAL, heap.cant-i-1)
		heap.downHeap(VALOR_INICIAL, heap.cant-i-2)
	}
	elementos = heap.datos
}

func (cola colaConPrioridad[T]) EstaVacia() bool {
	return cola.cant == 0
}

func (cola *colaConPrioridad[T]) Encolar(dato T) {
	//Hacer el UpHeap
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
	//Hacer el DownHeap
	if cola.EstaVacia() {
		panic(PANIC_COLA_VACIA)
	}
	cola.verDisminuirCapacidad()
	elemento := cola.datos[VALOR_INICIAL]
	cola.cant--
	cola.swap(VALOR_INICIAL, cola.cant)
	cola.downHeap(VALOR_INICIAL, cola.cant-1)
	return elemento
}

func (cola colaConPrioridad[T]) Cantidad() int {
	return cola.cant
}

// METODOS AUXILIARES INTERNOS
func (cola *colaConPrioridad[T]) crearHeapify(inicio, fin int) {
	//realizar un downHeap desde el ultimo elemento hasta la raiz
	if inicio == fin || fin < 0 {
		return
	}
	posicionPadre := (fin - 1) / 2
	if cola.cmp(cola.datos[posicionPadre], cola.datos[fin]) < 0 {
		cola.swap(posicionPadre, fin)
	}
	cola.crearHeapify(inicio, fin-1)

}
func (cola *colaConPrioridad[T]) upHeap(posicionHijo int) {
	//upHeap : si el padre es menor que el hijo , intercambiamos
	if posicionHijo < 1 {
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
	if inicio > fin {
		return
	}
	posHijo := (2 * inicio) + 1
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
		cola.redimensionar(cap(cola.datos) * 2)
	}
}
func (cola *colaConPrioridad[T]) verDisminuirCapacidad() {
	//Si la cantidad ocupada es 1/4 menos que la capacidad del array y la capacidad es mayor que la que iniciamos , redimensiono hacia abajo
	if cola.cant*LIMITE_CANT_OCUPADA <= cap(cola.datos) && cap(cola.datos) > CAPACACIDAD_INICIAL {
		cola.redimensionar(cap(cola.datos) / FACTOR_REDIMENSION)
	}
}

func (cola *colaConPrioridad[T]) redimensionar(tamanio int) {
	//Creo un nuevo array copiando los elementos a un nuevo tamanio
	arr := make([]T, tamanio)
	copy(arr, cola.datos)
	cola.datos = arr
}

// FUNCIONES AUXILIARES

func (cola *colaConPrioridad[T]) swap(posicionPadre, posicionHijo int) {
	cola.datos[posicionPadre], cola.datos[posicionHijo] = cola.datos[posicionHijo], cola.datos[posicionPadre]
}
