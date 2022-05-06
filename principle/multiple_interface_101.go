package principle

import (
	"fmt"
	"sync"
)

type TestDetails interface {
	Details()
}

// Interface 2
type TestArticles interface {
	Articles()
	Details()
}

type FinalMultiPleInterface interface {
	TestDetails
	TestArticles
	Stops()
}

type PricipleStore struct {
	Name  string
	Money int
	Thip  int
}

type PricipleStoreAgruement struct {
	PricipleStore PricipleStore
}

func NewMultiplesInterface(pricipleStore PricipleStore) FinalMultiPleInterface {
	return &PricipleStoreAgruement{
		PricipleStore: pricipleStore,
	}
}

func (psa *PricipleStoreAgruement) Details() {
	fmt.Printf("PricipleStore detail [OBJECT] -->> %+v\n", psa.PricipleStore)
}

func (psa *PricipleStoreAgruement) Articles() {
	fmt.Println("In Articles")
	var mu sync.Mutex
	var wg sync.WaitGroup
	n := 4
	stackArticles := 0
	psa.Details()

	wg.Add(1)
	go func() {
		defer wg.Done()
		mu.Lock()
		stackArticles += 1
		mu.Unlock()
	}()
	wg.Wait()

	for i := 0; i < n; i++ {
		fmt.Println("stackArticles :", stackArticles)
		if stackArticles == n {
			psa.Stops()
		}
		stackArticles++
	}
}

func (psa *PricipleStoreAgruement) Stops() {
	fmt.Println("<---- stop ---->")
}
