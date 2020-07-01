package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
)

//GetUserProfile handler
func GetUserProfile(c echo.Context) error {
	response, err := callAuthenticatedMicroService(c, "/auth")
	if response.StatusCode == 401 {
		return echo.NewHTTPError(401)
	} else if response.StatusCode != 200 {
		return echo.NewHTTPError(response.StatusCode)
	}
	response, err = callUserMicroService(c, "/user/profile")
	result := make(map[string]interface{})
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(contents, &result)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer response.Body.Close()
	return c.JSON(200, result)
}

//GetMicroServiceName handler
func GetMicroServiceName(c echo.Context) error {
	response, err := callUserMicroService(c, "/microservice/name")
	result := make(map[string]interface{})
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(contents, &result)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer response.Body.Close()
	return c.JSON(200, result)
}

func callAuthenticatedMicroService(c echo.Context, url string) (*http.Response, error) {
	req, err := http.NewRequest("POST", "http://127.0.0.1:8081"+url, nil)
	if err != nil {
		return &http.Response{}, err
	}
	req.Header = c.Request().Header
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return &http.Response{}, err
	}
	return resp, nil
}

func callUserMicroService(c echo.Context, url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", "http://127.0.0.1:8082"+url, nil)
	if err != nil {
		return &http.Response{}, err
	}
	req.Header = c.Request().Header
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return &http.Response{}, err
	}
	return resp, nil
}
