package cola_prioridad_test

import (
	"cmp"
	"fmt"
	"math/rand"
	"strings"
	TDAHeap "tdas/cola_prioridad"
	"testing"

	"github.com/stretchr/testify/require"
)

var TAMS_VOLUMEN = []int{12500, 25000, 50000, 100000, 200000, 400000}

func compararCadenas(a, b string) int {
	return strings.Compare(a, b)
}

func compararNumeros(a, b int) int {
	return cmp.Compare(a, b)
}

func numeroAleatorio(n int) int {
	return rand.Intn(n)
}

func TestHeapVacio(t *testing.T) {
	t.Log("Creamos un Heap vacio y se comporta como un Heap recien creado")
	datos := TDAHeap.CrearHeap[int](compararNumeros)
	require.True(t, datos.EstaVacia())
	require.PanicsWithValue(t, TDAHeap.PANIC_COLA_VACIA, func() {
		datos.Desencolar()
	})
	require.PanicsWithValue(t, TDAHeap.PANIC_COLA_VACIA, func() {
		datos.VerMax()
	})
	require.Equal(t, 0, datos.Cantidad())
}

func TestHeapEncolar(t *testing.T) {
	t.Log("Creamos un heap vacio y vamos encolando, se tiene que cumplir la propiedad de UpHeap")
	heap := TDAHeap.CrearHeap(compararCadenas)
	heap.Encolar("B")
	heap.Encolar("H")
	heap.Encolar("A")
	heap.Encolar("Z")
	heap.Encolar("Y")
	require.Equal(t, 5, heap.Cantidad())
	require.False(t, heap.EstaVacia())
	require.Equal(t, "Z", heap.VerMax(), "El maximo elemento de la tabla del abecedario es la Z")
}

func TestHeapDesencolar(t *testing.T) {
	t.Log("Desencolamos los elementos y que cumpla con la propiedad de heap, downHeap")
	var numeros = []int{5, 6, 2, 4, 9, 10, 21}
	heap := TDAHeap.CrearHeap(compararNumeros)
	for _, valor := range numeros {
		heap.Encolar(valor)
	}
	for i := 0; i < len(numeros); i++ {
		heap.Desencolar()
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

func TestHeapEncolarDesencolar(t *testing.T) {
	t.Log("Probamos que Encolar, Desencolar varias veces no se rompa y coincida el VER MAX")
	heap := TDAHeap.CrearHeap(compararNumeros)
	arrPrueba := make([]int, TAMS_VOLUMEN[0])
	for i := 0; i < len(arrPrueba); i++ {
		require.True(t, heap.EstaVacia())
		heap.Encolar(i)
		require.EqualValues(t, i, heap.VerMax(), "VER MAX con un primer elemento encolado tiene que ser el mismo que se encolo")
		require.False(t, heap.EstaVacia())
		heap.Desencolar()
		require.True(t, heap.EstaVacia())
	}
	require.Equal(t, 0, heap.Cantidad())
	require.Panics(t, func() {
		heap.Desencolar()
	}, "El heap tiene que cumplir su comportamiento de estar vacio una ves desencolado todos los elementos")
	require.Panics(t, func() {
		heap.VerMax()
	}, "El heap tiene que cumplir su comportamiento de estar vacio una ves desencolado todos los elementos")
}

func ejecutarPruebaVolumenHeap(b *testing.B, n int) {
	heap := TDAHeap.CrearHeap[int](compararNumeros)
	arrVolumen := make([]int, n)

	//Encolamos varios elementos en el heap
	for i := 0; i < n; i++ {
		datoAleatorio := numeroAleatorio(n)
		arrVolumen[i] = datoAleatorio
		heap.Encolar(datoAleatorio)
	}
	require.EqualValues(b, n, heap.Cantidad(), "La cantidad de elementos es incorrecta")
	require.False(b, heap.EstaVacia(), "El heap no puede quedar vacio")

	//Desencolamos todos los elementos
	for i := 0; i < n; i++ {
		heap.Desencolar()
	}
	require.Equal(b, 0, heap.Cantidad())
	require.True(b, heap.EstaVacia())
	require.PanicsWithValue(b, "La cola esta vacia", func() {
		heap.Desencolar()
	})
	require.PanicsWithValue(b, "La cola esta vacia", func() {
		heap.VerMax()
	})
}

func BenchmarkHeap(b *testing.B) {
	b.Log("Prueba de stress del heap. Prueba guardando distinta cantidad de elementos(muy grandes)," +
		"ejecutando muchas veces las pruebas para generar un benchmark. Valida que la CANTIDAD sea la adecuadad," +
		"Las redimensiones cumplan con su funcionamiento," + "y luego podemos Desencolar sin ningun problema dejando el heap VACIO")
	for _, n := range TAMS_VOLUMEN {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarPruebaVolumenHeap(b, n)
			}
		})
	}

}

func TestHeapifyArregloVacio(t *testing.T) {
	t.Log("Creamos un heap con un arreglo nulo y tiene que comportarse como un heap vacio")
	heapArr := TDAHeap.CrearHeapArr([]int{}, compararNumeros)
	require.True(t, heapArr.EstaVacia())
	require.Equal(t, 0, heapArr.Cantidad())
	require.PanicsWithValue(t, TDAHeap.PANIC_COLA_VACIA, func() {
		heapArr.Desencolar()
	})
	require.PanicsWithValue(t, TDAHeap.PANIC_COLA_VACIA, func() {
		heapArr.VerMax()
	})

}

func TestHeapifyArregloVacioEncolar(t *testing.T) {
	t.Log("Recibimos un arreglo vacio y encolamos elementos")
	heapArr := TDAHeap.CrearHeapArr([]string{}, compararCadenas)
	heapArr.Encolar("vaca")
	heapArr.Encolar("burro")
	heapArr.Encolar("caballo")
	heapArr.Encolar("lobo")
	heapArr.Encolar("gallina")
	require.False(t, heapArr.EstaVacia(), "El heap con un arreglo vacio y encolado los elementos no puede ser vacio")
	require.Equal(t, "vaca", heapArr.VerMax())
	require.Equal(t, 5, heapArr.Cantidad())
}

func TestHeapifyArregloVacioDesencolar(t *testing.T) {
	t.Log("Recibimos un arreglo vacio y desencolamos todos los elementos")
	heapArr := TDAHeap.CrearHeapArr([]int{}, compararNumeros)
	for i := 0; i < 200; i++ {
		heapArr.Encolar(i)
	}
	require.False(t, heapArr.EstaVacia())
	require.Equal(t, 200, heapArr.Cantidad())
	for i := 0; i < 200; i++ {
		heapArr.Desencolar()
	}
	require.True(t, heapArr.EstaVacia())
	require.Equal(t, 0, heapArr.Cantidad())
	require.PanicsWithValue(t, TDAHeap.PANIC_COLA_VACIA, func() {
		heapArr.Desencolar()
	})
	require.PanicsWithValue(t, TDAHeap.PANIC_COLA_VACIA, func() {
		heapArr.VerMax()
	})

}

func TestHeapifyArregloConUnElemento(t *testing.T) {
	t.Log("Recibe un arreglo con un unico elemento, encolamos datos, no debe romper la propiedad de heap")
	arrNum := []int{5}
	heapArr := TDAHeap.CrearHeapArr(arrNum, compararNumeros)
	for i := 0; i < 21; i++ {
		heapArr.Encolar(i)
	}
	require.False(t, heapArr.EstaVacia())
	require.Equal(t, 22, heapArr.Cantidad())
	require.Equal(t, 20, heapArr.VerMax())
}

func TestHeapifylArregloConElementos(t *testing.T) {
	t.Log("Recibimos un arreglo con elementos por parametro y le damos comportamiento de heap, el arreglo dado no tiene que ser modificado")
	numeros := []int{4, 2, 7, 10, 4, 3, 8, 5}
	arregloOrdenado := []int{10, 8, 7, 5, 4, 4, 3, 2}
	heapArr := TDAHeap.CrearHeapArr(numeros, compararNumeros)
	require.EqualValues(t, len(numeros), heapArr.Cantidad(), "la cantidad de elementos del slices tiene que ser igual al heap")
	ok := true
	for i := 0; i < len(arregloOrdenado); i++ {
		dato := heapArr.Desencolar()
		if arregloOrdenado[i] != dato {
			ok = false
			break
		}
	}
	require.True(t, ok, "Los elementos desencolados tiene que ir de mayor prioridad a menor prioridad")
	require.True(t, heapArr.EstaVacia())
	require.Equal(t, 0, heapArr.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() {
		heapArr.VerMax()
	})
	require.PanicsWithValue(t, "La cola esta vacia", func() {
		heapArr.Desencolar()
	})
}
func ejecutarPruebaVolumenHeapify(b *testing.B, n int) {
	arrVolumen := make([]int, n)
	//Creamos el arrVolumen para el heapify
	for i := 0; i < n; i++ {
		arrVolumen[i] = i + 1
	}
	heapArr := TDAHeap.CrearHeapArr(arrVolumen, compararNumeros)
	require.False(b, heapArr.EstaVacia())
	require.Equal(b, n, heapArr.Cantidad())
	require.Equal(b, n, heapArr.VerMax())
	//Desencolamos todos los elementos
	for i := 0; i < n; i++ {
		heapArr.Desencolar()
	}
	require.Equal(b, 0, heapArr.Cantidad())
	require.True(b, heapArr.EstaVacia())
	require.PanicsWithValue(b, "La cola esta vacia", func() {
		heapArr.Desencolar()
	})
	require.PanicsWithValue(b, "La cola esta vacia", func() {
		heapArr.VerMax()
	})
}

func BenchmarkHeapifyConElementos(b *testing.B) {
	b.Log("Prueba stress del heapify dado un array con elementos," +
		" Probamos dado un array (grande) con elementos y le damos comportamiento de heap," + "Dencolamos todos los elementos y el heap se comporta vacio")
	for _, n := range TAMS_VOLUMEN {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarPruebaVolumenHeapify(b, n)
			}
		})
	}
}

func TestHeapSortVacio(t *testing.T) {
	t.Log("Damos un arreglo vacio y no deberia generar problemas")
	TDAHeap.HeapSort(nil, compararCadenas)

}
func TestHeapSortUnElemento(t *testing.T) {
	t.Log("Ordenamos un arreglo con un elemento")
	arrAnimal := []string{"perro"}
	TDAHeap.HeapSort(arrAnimal, compararCadenas)
	require.Equal(t, "perro", arrAnimal[0])
}
func TestHeapSort(t *testing.T) {
	t.Log("Ordenamos un arreglo usando la funcion de HeapSort")
	numeros := []int{4, 2, 7, 10, 4, 3, 8, 5}
	arrOrdenado := []int{2, 3, 4, 4, 5, 7, 8, 10}
	TDAHeap.HeapSort(numeros, compararNumeros)
	ok := true
	for i := 0; i < len(arrOrdenado); i++ {
		if arrOrdenado[i] != numeros[i] {
			ok = false
			break
		}
	}
	require.True(t, ok, "El array tiene que estar ordenado de menor a mayor")
}

func ejecutarPruebaVolumenHeapSort(b *testing.B, n int) {
	arrVolumen := make([]int, n+1)
	//Creamos el arrVolumen para el heapify
	for i := n; i > 0; i-- {
		arrVolumen[i] = i
	}
	TDAHeap.HeapSort(arrVolumen, compararNumeros)
	ok := true
	for i := 0; i < n; i++ {
		if i != arrVolumen[i] {
			ok = false
			break
		}
	}
	require.True(b, ok, "El array tiene que estar ordenado de menor a mayor")
}

func BenchmarkHeapSort(b *testing.B) {
	b.Log("Prueba stress HeapSort ordenando varios arrays de tamanios grandes")
	for _, n := range TAMS_VOLUMEN {
		b.Run(fmt.Sprintf("Prueba de %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarPruebaVolumenHeapSort(b, n)
			}
		})
	}
}
