package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
)

var gen2Csbms = []string{
	"f8981d09-bb9a-4dad-8947-75ea5bd3a4dc", "cf24ab33-32e2-47cb-9468-4b000f95ec97", "f82b8ac5-70f2-4b1c-9225-79e0f1d0431b",
	"7a9c63ad-90ff-44f2-bb70-36b79834e9bc", "65fb5970-7fbd-4537-ac67-2d7575509f6e", "3670b2c0-8de3-44d5-ab87-9f26d4ed9708",
	"f51bd7b2-f59d-4add-8113-d9fef334091d", "1b8c1419-e899-4061-9bd3-9168586a24c3", "d3a7fd9f-00af-4a19-a4cb-643871d7cca9",
	"e28e8e11-9bd0-4b52-a840-d4190a7a2694", "14730793-756a-467e-bb1d-d317d0930093", "be96c92e-3716-433c-90e7-520857473e18",
	"14b58fcc-89b3-48d5-ab03-4308a6284cf7", "55b68496-8f1a-44d0-beca-d1bc30f20147", "d5d08793-f5e8-4433-ba9b-bfa7f1ae6f78",
	"f1af3987-1573-47a3-bb90-6bd3dbadab55", "9cf104d6-3eb8-43e7-99ce-aee7de598c2c", "f75ec466-4219-412c-9983-1897e0c72b0a",
	"c55087c6-d74d-4cb3-bcad-6840cf4fc0d3", "cd49e1a4-c234-493c-a1ee-59e9335035fe", "27309d3e-0b7a-4244-9c82-796bf2f0378d",
	"05ab3885-ac7b-410d-a5e2-cf5eb40a19b7", "1df498d2-dca9-4653-9411-b07a59414d60", "b6a75033-f95b-49ed-a25f-22f6ea7a9fc8",
	"7a9c63ad-90ff-44f2-bb70-36b79834e9bc",
}

var gen3Csbms = []string{
	"a554f3b3-e0b6-4b06-88cb-51ad686aef61", "ff324bf5-0ca9-4be9-9753-fcdf06476b10", "308df16d-46c1-437b-8ede-28412900c61a",
	"0372c590-22d1-4aeb-ac5a-ea6dfe385e39", "ca06d1dc-361d-4bca-b68f-f557670ddb27", "83a4ce52-6b2d-47c3-a6ff-35864e3d30ce",
	"499ee4ee-3e9f-4375-bbd8-dde3d8ea0e3f", "697f07b1-dd98-457d-a8d9-ae9f83430c1a", "5498c81a-e792-4328-8339-4635ad274e1b",
	"10a5fb4e-e1eb-4e2f-bd90-105e7c64a644", "152e5be8-315d-4ffe-9355-8fdef0a5c8f3", "4cfb3a60-7318-4dd2-8278-2d435f709351",
	"95c9efb5-7b8c-4f99-956e-e712067a4736",
}

var stageCsbms = []string{
	"92743f4a-a194-4fc9-ab65-eece09128bfe", "678986ba-02c6-4eb7-9928-7d657714f84c", "cf955142-42c7-4c75-936e-c41af8085e3c",
	"b1af613f-35a7-49b4-88bc-b5c0f22c120b", "8b7ba07c-4f73-45dd-847e-553839096446",
}

var alphaCsbms = []string{
	"10b11363-5d68-40a9-931a-974c1afc8623", "3bf55a34-c137-4a4d-b7c1-ebf13a3f8d4f", "92a7d776-0e75-4847-8302-021e1ddd2f9b",
	"8e289f04-c5a8-47a6-a10c-ab12eb1a8fa1",
}

func getProdLiveConfig(csbm string) string {
	res, err := http.Get("http://<someip>:<someport>/r1rmGA/csbm/" + csbm + "/liveConfig")
	if err != nil {
		fmt.Printf("Unable to access proxy host: %s", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Unable to read response body: %s", err)
	}

	res.Body.Close()
	return fmt.Sprintf("%s", body)
}

func getStagLiveConfig(csbm string) string {
	res, err := http.Get("http://<someip>:<someport>/r1rmHouston/csbm/" + csbm + "/liveConfig")
	if err != nil {
		fmt.Printf("Unable to access proxy host: %s", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Unable to read response body: %s", err)
	}

	res.Body.Close()
	return fmt.Sprintf("%s", body)
}

func getAlphLiveConfig(csbm string) string {
	res, err := http.Get("http://<someip>:<someport>/r1rmQA2/csbm/" + csbm + "/liveConfig")
	if err != nil {
		fmt.Printf("Unable to access proxy host: %s", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Unable to read response body: %s", err)
	}

	res.Body.Close()
	return fmt.Sprintf("%s", body)
}

func equalJSON(s1, s2 string) (bool, error) {
	var o1 interface{}
	var o2 interface{}

	var err error

	err = json.Unmarshal([]byte(s1), &o1)
	if err != nil {
		return false, fmt.Errorf("Error mashalling string 1 :: %s", err.Error())
	}

	err = json.Unmarshal([]byte(s2), &o2)
	if err != nil {
		return false, fmt.Errorf("Error mashalling string 2 :: %s", err.Error())
	}

	return reflect.DeepEqual(o1, o2), nil
}

func main() {
	//
	// GOLD liveconfigs for production, staging, and alpha.
	//
	// 	Production
	//	----------
	// 	MaxSpools 	= 30
	//	SOAP timeout 	= 3600
	//	DeltaCalcFreq	= 432000
	//	MaxVMs		= 20
	//
	// 	GEN2 Storage    = /storage01/replication ... /storage09/replication, 9 total volumes
	// 	GEN3 Storage    = /storage01/replication ... /storage06/replication, 6 total volumes
	//
	//	Staging
	//	-------
	//	MaxSpools	= 30
	//	SOAP timeout	= 3600
	//	MaxVMs		= 20
	//
	//	STAGING Storage	= /storage01/replication, and or /storage02/replication
	//
	//	Alpha
	//	-----
	//	MaxSpools	= 10
	//	SOAP timeout	= 300
	//	MaxVMs		= 10
	//
	//	ALPHA Storage	= /storage01/replication, and or /storage02/replication
	//
	GEN2GOLD := getProdLiveConfig("d3a7fd9f-00af-4a19-a4cb-643871d7cca9") // wdcsbm10 GOLD LIVECONFIG
	GEN3GOLD := getProdLiveConfig("ca06d1dc-361d-4bca-b68f-f557670ddb27") // wdcsbm18 GOLD LIVECONFIG
	STAGGOLD := getStagLiveConfig("92743f4a-a194-4fc9-ab65-eece09128bfe") // housbm03 GOLD LIVECONFIG
	ALPHGOLD := getAlphLiveConfig("10b11363-5d68-40a9-931a-974c1afc8623") // noname   GOLD LIVECONFIG

	numsbms := 0

	fmt.Printf("\nGen2 csbms:\n\n")

	for i := range gen2Csbms {

		s1 := getProdLiveConfig(gen2Csbms[i])
		s2 := GEN2GOLD

		areEqual, err := equalJSON(s1, s2)

		if err != nil {
			fmt.Println("Error marshalling strings", err.Error())
		}
		fmt.Println(gen2Csbms[i]+" Valid JSON?  :: ", areEqual)

		numsbms++
	}

	fmt.Printf("\nGen3 csbms:\n\n")

	for i := range gen3Csbms {

		s1 := getProdLiveConfig(gen3Csbms[i])
		s2 := GEN3GOLD

		areEqual, err := equalJSON(s1, s2)

		if err != nil {
			fmt.Println("Error marshalling strings", err.Error())
		}
		fmt.Println(gen3Csbms[i]+" Valid JSON?  :: ", areEqual)

		numsbms++
	}

	fmt.Printf("\nStaging csbms:\n\n")

	for i := range stageCsbms {

		s1 := getStagLiveConfig(stageCsbms[i])
		s2 := STAGGOLD

		areEqual, err := equalJSON(s1, s2)

		if err != nil {
			fmt.Println("Error marshalling strings", err.Error())
		}
		fmt.Println(stageCsbms[i]+" Valid JSON?  :: ", areEqual)

		numsbms++
	}

	fmt.Printf("\nAlpha csbms:\n\n")

	for i := range alphaCsbms {

		s1 := getAlphLiveConfig(alphaCsbms[i])
		s2 := ALPHGOLD

		areEqual, err := equalJSON(s1, s2)

		if err != nil {
			fmt.Println("Error marshalling strings", err.Error())
		}
		fmt.Println(alphaCsbms[i]+" Valid JSON?  :: ", areEqual)

		numsbms++
	}

	fmt.Printf("\nTotal number of sbms in prod, staging, and alpha: %d\n\n", numsbms)

}
