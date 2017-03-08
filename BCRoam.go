/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License .
*/

package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

// This is our structure for the broadcaster creating bulk inventory

type rsDetailInit struct {
	MSISDN     string  `json:"msisdn"`
	Name       string  `json:"name"`
	Address    string  `json:"address"`
	HO         string  `json:"ho"`
	RP         string  `json:"rp"`
	Roaming    string  `json:"roaming"`
	Location   string  `json:"location"`
} 


type rsDetail struct {
	MSISDN     string  `json:"msisdn"`
	Name       string  `json:"name"`
	Address    string  `json:"address"`
	HO         string  `json:"ho"`
	RP         string  `json:"rp"`
	Roaming    string  `json:"roaming"`
	Location   string  `json:"location"`
	Plan       string  `json:"plan"`
	VoinceOutL string `json:"voinceOutL"`
	VoinceInL  string `json:"voinceInL"`
	DataL      string  `json:"float64"`
	VoiceOutR  string `json:"voiceOutR"`
	VoiceInR   string `json:"voiceInR"`
	DataR      string `json:"dataR"`
} 

//This is a helper structure to point to allPeers
type AllPeers struct {
	PeerName []string `json:"peerName"`
}

//For Debugging
func showArgs(args []string) {

	for i := 0; i < len(args); i++ {
		fmt.Printf("\n %d) : [%s]", i, args[i])
	}
	fmt.Printf("\n")
}

// Init function
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	//var err error

	fmt.Println("Launching Init Function")
	
	//Inventory hard coded here
	rs1:=rsDetailInit{"14691234567","A","DALLAS","AT&T","","FALSE","DALLAS"}
	rs2:=rsDetailInit{"14691234568","A","DALLAS","AT&T","","FALSE","DALLAS"}
	rs3:=rsDetailInit{"14691234569","A","DALLAS","AT&T","","FALSE","DALLAS"}
	rs4:=rsDetailInit{"14691234570","A","DALLAS","AT&T","","FALSE","DALLAS"}
	rs5:=rsDetailInit{"349091234567","A","BARCELONA","VODAFONE","","FALSE","DALLAS"}
	rs6:=rsDetailInit{"349091234568","A","BARCELONA","VODAFONE","","FALSE","DALLAS"}
	rs7:=rsDetailInit{"349091234569","A","BARCELONA","VODAFONE","","FALSE","DALLAS"}

	//Create array for all adspots in ledger
	//var AllPeersArray AllPeers

	t.putMSIDN(stub,rs1, rs1.MSISDN)
	t.putMSIDN(stub,rs2, rs2.MSISDN)
	t.putMSIDN(stub,rs3, rs3.MSISDN)
	t.putMSIDN(stub,rs4, rs4.MSISDN)
	t.putMSIDN(stub,rs5, rs5.MSISDN)
	t.putMSIDN(stub,rs6, rs6.MSISDN)
	t.putMSIDN(stub,rs7, rs7.MSISDN)

	fmt.Println("Init Function Complete")
	return nil, nil
}

//Invoke function

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Printf("Invoke called, determining function :%v",function)

	showArgs(args)
	var msisdn,val string

	// Handle different functions
	if function == "discoverRP" {
		fmt.Printf("Function is discoverRP")
		val=args[0]
		msisdn=args[1]
		return t.discoverRP(stub,msisdn,val)
	} else if function == "enterData" {
		fmt.Printf("Function is enterData")
		return t.enterData(stub,args)
	} else if function == "roamOnOff" {
		fmt.Printf("Function is roamOnOff")
		msisdn=args[0]
		return t.roamOnOff(stub, msisdn)
	}else if function == "updateRates" {
		fmt.Printf("Function is updateRates")
		val=args[0]
		msisdn=args[1]
		return t.updateRates(stub,msisdn,val)
	} 
        return nil, errors.New("Received unknown function invocation")
}

//QUERY FUNCTION
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Printf("======== Query called, determining function")

	showArgs(args)

	if function == "queryPeers" {
		fmt.Printf("Function is queryPeers")
		return t.queryPeers(stub, args)
	} else if function == "queryMSISDN" {
		fmt.Printf("Function is queryMSISDN")
		return t.queryMSISDN(stub, args)
	} else {
		fmt.Printf("Invalid Function!")
	}

	return nil, nil
}
////////////////////////////////////////////////////

//Redirect FUNCTIONS

//Query peers in our network
func (t *SimpleChaincode) queryPeers(stub shim.ChaincodeStubInterface,args []string) ([]byte, error) {
	fmt.Println("queryPeers called")
	var user string
	user= args[0]
	fmt.Println("User name: %v",user)
	bytes,_:= stub.GetState(user)
	var peer string
	err := json.Unmarshal(bytes, &peer)
	if err != nil{
		fmt.Printf("Error in Unmarshalling")
	} else {
		fmt.Printf("Peer name: %v",peer)
	}
	return nil,nil
}


//Query MSISDN in our network
func (t *SimpleChaincode) queryMSISDN(stub shim.ChaincodeStubInterface,args []string) ([]byte, error) {
	fmt.Println("queryMSISDN called")
	var msisdn string
	msisdn= args[0]
	fmt.Println("MSISDN: %v",msisdn)
	bytes,_:= stub.GetState(msisdn)
	fmt.Println(string(bytes))
	fmt.Printf("%x",bytes)
	/*var rsDetailObj rsDetail
	err := json.Unmarshal(bytes, &rsDetailobj)
	if err != nil{
		fmt.Printf("Error in Unmarshalling")
	} else {
		fmt.Printf("Name:%v",rsDetailObj.name)
		fmt.Printf("Addr:%v",rsDetailObj.address)
		fmt.Printf("HO:%v",rsDetailObj.ho)
		fmt.Printf("RP:%v",rsDetailObj.rp)
		fmt.Printf("Roaming:%v",rsDetailObj.roaming)
		fmt.Printf("Location:%v",rsDetailObj.location)
		fmt.Printf("Plan:%v",rsDetailObj.plan)
		fmt.Printf("VoiceOutLocal:%v",rsDetailObj.voinceOutL)
		fmt.Printf("VoiceInLocal:%v",rsDetailObj.voinceInL)
		fmt.Printf("DataLocal:%v",rsDetailObj.dataL)  
		fmt.Printf("VoiceOutRoam:%v",rsDetailObj.voiceOutR)
		fmt.Printf("VoiceInRoam:%v",rsDetailObj.voiceInR)
		fmt.Printf("DataRoam:%v",rsDetailObj.dataR)
	}*/
	return bytes,nil
}

func (t *SimpleChaincode) enterData(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	var MSISDN=args[0]
	var rsDetailObj rsDetail
	        rsDetailObj.MSISDN = args[0]
		rsDetailObj.Name = args[1]
		rsDetailObj.Address = args[2]
		rsDetailObj.HO = args[3]
		rsDetailObj.RP = args[4]   
		rsDetailObj.Roaming = args[5]
		rsDetailObj.Location = args[6]
		rsDetailObj.Plan = args[7]
		rsDetailObj.VoinceOutL = args[8]
		rsDetailObj.VoinceInL = args[9]
		rsDetailObj.DataL = args[10]     
		rsDetailObj.VoiceOutR = args[11]
		rsDetailObj.VoiceInR = args[12]
		rsDetailObj.DataR = args[13]
	
	        Currbytes, err := stub.GetState(MSISDN)
		if err != nil {
			fmt.Println("Error - Could not get User details : %s", MSISDN)
			//return nil, err
		} else {
			fmt.Println("Success - User details found %s %v", MSISDN ,string(Currbytes) )
		}
	
	        fmt.Println(rsDetailObj)
		bytes, _ := json.Marshal(rsDetailObj)
	        fmt.Println(string(bytes))
	        fmt.Printf("%x",bytes)
		err2 := stub.PutState(MSISDN, bytes)
		if err2 != nil {
			fmt.Println("Error - could not Marshall in rsDetailObj")
		} else {
			fmt.Println("Success -  works")
		}
	
	//showArgs(args)

	return nil, errors.New("Received unknown function invocation")
}


//putNetworkPeers: To put an array containing pointers to all blocks for a particular user(or peer) on the ledger
func (t *SimpleChaincode) putMSIDN(stub shim.ChaincodeStubInterface,rs rsDetailInit,msisdn string) ([]byte, error) {
	//marshalling
	fmt.Println(" Initializing msisdn: ", msisdn)
	fmt.Printf("put details: %+v ", rs)
	fmt.Printf("\n")
	bytes, _ := json.Marshal(rs)
	fmt.Println(string(bytes))
	err2 := stub.PutState(msisdn, bytes)
	if err2 != nil {
		fmt.Println("Error - could not Marshall in msisdn")
		//return nil, err
	} else {
		fmt.Println("Success - Marshall in msisdn details")
	}
	return nil, nil
}

//Remote Partner Discovery
func (t *SimpleChaincode) discoverRP(stub shim.ChaincodeStubInterface, msisdn string,val string) ([]byte, error) {

	bytes, err := stub.GetState(msisdn)
	if err != nil {
		fmt.Println("Error - Could not get User details : %s", msisdn)
		//return nil, err
	} else {
		fmt.Println("Success - User details found %s", msisdn)
	}

	var rsDetailobj rsDetail
	err = json.Unmarshal(bytes, &rsDetailobj)
	rsDetailobj.RP=val
	bytes2, _ := json.Marshal(rsDetailobj)
	err2 := stub.PutState(msisdn,bytes2)
	if err2 != nil {
		fmt.Println("Error - could not Marshall in msisdn")
		//return nil, err
	} else {
		fmt.Println("Success, updated record")
	}
	
	return nil,nil
}

//Roaming on and off
func (t *SimpleChaincode) roamOnOff(stub shim.ChaincodeStubInterface, msisdn string) ([]byte, error) {
	return nil,nil
}

//Update voice and data rates
func (t *SimpleChaincode) updateRates(stub shim.ChaincodeStubInterface, msisdn string, val string) ([]byte, error) {
	
	
	
	return nil,nil
}


	
//MAIN FUNCTION
func main() {
	err := shim.Start(new(SimpleChaincode))

	fmt.Printf("IN MAIN of TelcoChaincode")
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
