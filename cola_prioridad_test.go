package cola_prioridad

import (
	"cmp"
	"fmt"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func compararCadenas(a, b string) int {
	return strings.Compare(a, b)
}

func compararNumeros(a, b int) int {
	return cmp.Compare(a, b)
}

func TestColaConPrioridadVacio(t *testing.T) {
	t.Log("Creamos un Heap vacio y se comporta como un Heap recien creado")
	datos := CrearHeap[int](compararNumeros)
	require.True(t, datos.EstaVacia())
	require.PanicsWithValue(t, PANIC_COLA_VACIA, func() {
		datos.Desencolar()
	})
	require.PanicsWithValue(t, PANIC_COLA_VACIA, func() {
		datos.VerMax()
	})
	require.Equal(t, 0, datos.Cantidad())
}

func TestColaConPrioridadEncolar(t *testing.T) {
	t.Log("Creamos un heap vacio y vamos encolando, se tiene que cumplir la propiedad de UpHeap")
	heap := CrearHeap[string](compararCadenas)
	heap.Encolar("B")
	heap.Encolar("H")
	heap.Encolar("A")
	heap.Encolar("Z")
	heap.Encolar("Y")
	require.Equal(t, 5, heap.Cantidad())
	require.False(t, heap.EstaVacia())
	require.Equal(t, "Z", heap.VerMax(), "El maximo de los datos encolados")
}

func TestColaConPrioridadDesencolar(t *testing.T) {
	t.Log("Desencolamos los elementos y que cumpla con la propiedad de heap, downHeap")
	var numeros = []int{5, 6, 2, 4, 9, 10, 21}
	heap := CrearHeap(compararNumeros)
	for _, valor := range numeros {
		heap.Encolar(valor)
	}
	for i := 0; i < len(numeros); i++ {
		maximo := heap.VerMax()
		datoMaximo := heap.Desencolar()
		require.Equal(t, maximo, datoMaximo, "ver maximo con la primitiva desencolar tienen que coincidir")
	}
	require.True(t, heap.EstaVacia())
	require.Panics(t, func() {
		heap.Desencolar()
	}, "El heap tiene que cumplir su comportamiento de estar vacio una ves desencolado todos los elementos")
	require.Panics(t, func() {
		heap.VerMax()
	}, "El heap tiene que cumplir su comportamiento de estar vacio una ves desencolado todos los elementos")
	require.Equal(t, 0, heap.Cantidad())

}

func TestColaConPrioridadArregloVacio(t *testing.T) {
	t.Log("Creamos un heap con un arreglo nulo y tiene que comportarse como un heap vacio")
	heapArr := CrearHeapArr(nil, compararNumeros)
	require.True(t, heapArr.EstaVacia())
	require.Equal(t, 0, heapArr.Cantidad())
	require.PanicsWithValue(t, PANIC_COLA_VACIA, func() {
		heapArr.Desencolar()
	})
	require.PanicsWithValue(t, PANIC_COLA_VACIA, func() {
		heapArr.VerMax()
	})

}

func TestColaConPrioridadArreglConElementos(t *testing.T) {
	t.Log("Recibimos un arreglo con elementos por parametro y le damos comportamiento de heap, heapify")
	numeros := []int{4, 2, 7, 10, 4, 3, 8, 5}
	heapArr := CrearHeapArr(numeros, compararNumeros)
	require.Equal(t, 10, heapArr.VerMax())
}

func TestColaConPrioridadHeapSort(t *testing.T) {
	t.Log("Ordenamos un arreglo usando los metodos de heap")
	numeros := []int{3, 4, 5, 1, 2, 6, 23, 90}
	HeapSort(numeros, compararNumeros)
	fmt.Println(numeros)
}
