package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

type businessLogicService struct {
	client http.Client
}

type BusinessLogicService interface {
	TestsService(ctx context.Context) (interface{}, error)
}

func NewBusinessLogicService(client http.Client) BusinessLogicService {
	return &businessLogicService{
		client: client,
	}
}

func (bls *businessLogicService) TestsService(ctx context.Context) (interface{}, error) {
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	stackError := []error{}
	stackResponse := [][]byte{}
	errorResponses := make(chan error)
	response := make(chan []byte)

	url := "https://api.nasa.gov/planetary/apod?api_key=DEMO_KEY"
	headers := map[string]string{
		"Content-Type": "application/json",
	}

	n := 5
	i := 0
	for i < n {
		wg.Add(1)
		go func() {
			nctx, cancel := context.WithTimeout(ctx, 5*time.Minute)
			defer func() {
				cancel()
			}()
			result, err := bls.internalHttpRequests(nctx, url, http.MethodGet, headers)
			if err != nil {
				// log.Println("[ERROR] internalHttpRequests ->", err.Error())
				mu.Lock()
				errorResponses <- err
				mu.Unlock()
			}
			response <- result
		}()
		i++
	}

	go func() {
		for e := range errorResponses {
			fmt.Println("e ->", e.Error())
			stackError = append(stackError, e)
			wg.Done()
		}
	}()

	go func() {
		for r := range response {
			fmt.Println("response ->", string(r))
			stackResponse = append(stackResponse, r)
			wg.Done()
		}
	}()
	wg.Wait()
	close(errorResponses)
	close(response)

	if len(stackError) > 0 {
		return "", stackError[0]
	}

	result := []ResponseApiNasa{}
	for _, v := range stackResponse {

		responseApinasa := ResponseApiNasa{}
		errorRateLimit := ResponseErrorRateLimit{}

		if err := json.NewDecoder(bytes.NewReader(v)).Decode(&errorRateLimit); err != nil {
			return "", err
		}

		if !EmptyString(errorRateLimit.Error.Code) {
			errMeassage := errors.New(errorRateLimit.Error.Message)
			return "", errMeassage
		}

		if err := json.Unmarshal(v, &responseApinasa); err != nil {
			log.Println("err Unmarshal ->", err.Error())
			return "", err
		}
		result = append(result, responseApinasa)
	}
	return result, nil
}

func (bls *businessLogicService) internalHttpRequests(ctx context.Context, url string, methods string, headers map[string]string) ([]byte, error) {
	req, err := http.NewRequest(methods, url, nil)
	if err != nil {
		return nil, err
	}
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	req = req.WithContext(ctx)
	resp, err := bls.client.Do(req)
	if err != nil {
		return nil, err
	}
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func EmptyString(s string) bool {
	return s == ""
}
