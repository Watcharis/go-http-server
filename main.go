package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	p "watcharis/so/go-tutorial/gopointer"
	"watcharis/so/go-tutorial/handler"
	"watcharis/so/go-tutorial/principle"
)

func Run() {
	wg := sync.WaitGroup{}
	srv := &http.Server{
		Addr:    ":8080",
		Handler: nil,
	}

	go func() {
		fmt.Printf("\nsrv -> %+v\n", srv)
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("[START SERVER ERROR] -> %+v\n", err)
		}
	}()

	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)
	wg.Add(1)
	go func() {
		defer wg.Done()
		sigInt := <-sigs
		log.Println(sigInt)
	}()
	wg.Wait()
}

func main() {

	start := time.Now().Unix()
	a := time.Unix(start, 0)
	fmt.Println("a ->", a)
	end := time.Now()

	b := a.Sub(end)
	fmt.Println("b ->", b)

	data := 10
	result := p.PlusPointer(&data)
	fmt.Println("result :", result)

	// gennerate excel
	// e.GennerateExcel()

	testkeys := p.TestGetKeysFormStruct()
	fmt.Println("testkeys :", testkeys)

	prime := p.FindPrimeNumber(100)
	fmt.Println("prime :", prime)

	sliceNumber := []int{5, 78, 23, 45, 9, 6, 12, 0}

	sortNumberAsc := p.SortNumberAsc(sliceNumber)
	fmt.Println("sortNumberAsc ->", sortNumberAsc)

	sortNumberDesc := p.SortNumberDesc(sliceNumber)
	fmt.Println("sortNumberDesc ->", sortNumberDesc)

	// tr := p.Test()
	// fmt.Println("tr :", tr)
	slice := []int{5, 78, 23, 45, 9, 6, 12, 0}
	ra := p.ReversArray(slice)
	fmt.Println("ra :", ra)

	arrayNumber := []int{1, 2, 3, 4, 5, 6, 7}
	sa := p.ShiftLeftLoop(4, arrayNumber)
	fmt.Println("sa :", sa)

	testsOverridInGo := principle.Police{}
	caseOne := testsOverridInGo.TestOne()
	caseTwo := testsOverridInGo.Student.TestOne()
	fmt.Println("caseOne ->", caseOne)
	fmt.Println("caseTwo ->", caseTwo)

	fmt.Printf("/n--------------[ Embedded Interface ]---------------/n")
	values := principle.Author{
		A_name:    "Mickey",
		Branch:    "Computer science",
		College:   "XYZ",
		Year:      2012,
		Salary:    50000,
		Particles: 209,
		Tarticles: 309,
	}

	ttuEmbeddedInterface := principle.NewPrincipleTwo(values)
	fmt.Printf("\nttuEmbeddedInterface -> %+v\n", ttuEmbeddedInterface)

	ttuEmbeddedInterface.Details()
	ttuEmbeddedInterface.Articles()
	salary, err := ttuEmbeddedInterface.Sums()
	if err != nil {
		log.Println("[ERROR] -->", err.Error())
	}
	fmt.Printf("\nsalary -> %d\n", salary)

	defer func() {
		ttuEmbeddedInterface.Cleans()
	}()

	fmt.Printf("/n--------------[ Mulitiples Interface ]---------------/n")
	principleStore := principle.PricipleStore{
		Name:  "Test",
		Money: 20,
		Thip:  1000,
	}
	ttuMultiplesInterface := principle.NewMultiplesInterface(principleStore)
	fmt.Println("\nttuMultiplesInterface ->", ttuMultiplesInterface)
	ttuMultiplesInterface.Details()
	ttuMultiplesInterface.Articles()

	fmt.Printf("/n--------------[ 	HTTP SERVER ]---------------/n")

	ctx := context.Background()
	client := http.Client{}
	initServicesHttp := handler.NewHttp()

	initService := handler.NewBusinessLogicService(client)
	initEndpoint := handler.NewEndpointService(initService)
	initTransport := handler.NewTransportService(initEndpoint)

	http.Handle("/signin", initServicesHttp.Signin(ctx))
	http.Handle("/extract", initServicesHttp.ExtractToken(ctx))
	http.Handle("/refresh", initServicesHttp.Refresh(ctx))
	http.Handle("/images", initServicesHttp.SendImageFormDirectory(ctx))
	http.Handle("/url-images", initServicesHttp.GetUrlImages(ctx))
	http.Handle("/tests", initTransport.Tests(ctx))
	Run()
}
