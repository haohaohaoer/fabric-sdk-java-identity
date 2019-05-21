package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	//"strconv"
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// IdentityData example simple Chaincode implementation
type IdentityData struct {
}

type identity struct {
	ObjectType string `json:"docType"` //docType is used to distinguish the various types of objects in state database
	Name       string `json:"name"`
	Address    string `json:"address"`
	Itype      string `json:"itype"`
}

// ===================================================================================
// Main
// ===================================================================================
func main() {
	err := shim.Start(new(IdentityData))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// Init initializes chaincode
// ===========================
func (t *IdentityData) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

// Invoke - Our entry point for Invocations
// ========================================
func (t *IdentityData) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "initIdentity" { //初始化标识数据
		return t.initIdentity(stub, args)
	} else if function == "delete" { //删除数据
		return t.delete(stub, args)
	} else if function == "quaryIdentityByName" { //read a Identity
		return t.quaryIdentityByName(stub, args)
	} else if function == "queryIdentityByType" { //find marbles for owner X using rich query
		return t.queryIdentityByType(stub, args)
	} 

	fmt.Println("invoke did not find func: " + function) //error
	return shim.Error("Received unknown function invocation")
}

// ============================================================
// initMarble - create a new marble, store into chaincode state
// ============================================================
func (t *IdentityData) initIdentity(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	//  2        0                  	1
	// "handle", "10/ditu",    			"www.baidu.com"
	// "handle", "20/weichat", 			"www.qq.com"
	// "handle", "50/weibo",   			"www.sina.com"
	// "OID",    "1.2.156.1",  			"www.bupt.edu.cn"
	// "OID",    "1.2.156.1.2", 		"www.fnl.cn"
	// "Ecode",  "100036901234567892",  "www.nongye.cn"

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	// ==== Input sanitation ====
	fmt.Println("- start init identity")
	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}
	if len(args[1]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}
	if len(args[2]) <= 0 {
		return shim.Error("3rd argument must be a non-empty string")
	}
	identityName := args[0]
	address := args[1]
	itype := strings.ToLower(args[2])

	// ==== Check if marble already exists ====
	identityAsBytes, err := stub.GetState(identityName)
	if err != nil {
		return shim.Error("Failed to get identity: " + err.Error())
	} else if identityAsBytes != nil {
		fmt.Println("This identity already exists: " + identityName)
		return shim.Error("This identity already exists: " + identityName)
	}

	// ==== Create marble object and marshal to JSON ====
	objectType := "identity"
	identity := &identity{objectType, identityName, address, itype}
	identityJSONasBytes, err := json.Marshal(identity)
	if err != nil {
		return shim.Error(err.Error())
	}
	//Alternatively, build the marble json string manually if you don't want to use struct marshalling
	//marbleJSONasString := `{"docType":"Marble",  "name": "` + marbleName + `", "color": "` + color + `", "size": ` + strconv.Itoa(size) + `, "owner": "` + owner + `"}`
	//marbleJSONasBytes := []byte(str)

	// === Save marble to state ===
	err = stub.PutState(identityName, identityJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	//  ==== Index the marble to enable color-based range queries, e.g. return all blue marbles ====
	//  An 'index' is a normal key/value entry in state.
	//  The key is a composite key, with the elements that you want to range query on listed first.
	//  In our case, the composite key is based on indexName~color~name.
	//  This will enable very efficient state range queries based on composite keys matching indexName~color~*
	indexName := "type~name"
	typeNameIndexKey, err := stub.CreateCompositeKey(indexName, []string{identity.Itype, identity.Name})
	if err != nil {
		return shim.Error(err.Error())
	}
	//  Save index entry to state. Only the key name is needed, no need to store a duplicate copy of the marble.
	//  Note - passing a 'nil' value will effectively delete the key from state, therefore we pass null character as value
	value := []byte{0x00}
	stub.PutState(typeNameIndexKey, value)

	// ==== Marble saved and indexed. Return success ====
	fmt.Println("- end init identity")
	return shim.Success(nil)
}

// ===============================================
// readMarble - read a marble from chaincode state
// ===============================================
func (t *IdentityData) quaryIdentityByName(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var name, jsonResp string
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting name of the Identity to query")
	}

	name = args[0]
	valAsbytes, err := stub.GetState(name) //get the marble from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + name + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"Identity does not exist: " + name + "\"}"
		return shim.Error(jsonResp)
	}

	return shim.Success(valAsbytes)
}

// ==================================================
// delete - remove a marble key/value pair from state
// ==================================================
func (t *IdentityData) delete(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var jsonResp string
	var identityJSON identity
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	identityName := args[0]

	// to maintain the color~name index, we need to read the marble first and get its color
	valAsbytes, err := stub.GetState(identityName) //get the marble from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + identityName + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"Identity does not exist: " + identityName + "\"}"
		return shim.Error(jsonResp)
	}

	err = json.Unmarshal([]byte(valAsbytes), &identityJSON)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to decode JSON of: " + identityName + "\"}"
		return shim.Error(jsonResp)
	}

	err = stub.DelState(identityName) //remove the marble from chaincode state
	if err != nil {
		return shim.Error("Failed to delete state:" + err.Error())
	}

	// maintain the index
	indexName := "type~name"
	typeNameIndexKey, err := stub.CreateCompositeKey(indexName, []string{identityJSON.Itype, identityJSON.Name})
	if err != nil {
		return shim.Error(err.Error())
	}

	//  Delete index entry to state.
	err = stub.DelState(typeNameIndexKey)
	if err != nil {
		return shim.Error("Failed to delete state:" + err.Error())
	}
	return shim.Success(nil)
}

// ===========================================================================================
// constructQueryResponseFromIterator constructs a JSON array containing query results from
// a given result iterator
// ===========================================================================================
func constructQueryResponseFromIterator(resultsIterator shim.StateQueryIteratorInterface) (*bytes.Buffer, error) {
	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	return &buffer, nil
}

func getQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	fmt.Printf("- getQueryResultForQueryString queryString:\n%s\n", queryString)

	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	buffer, err := constructQueryResponseFromIterator(resultsIterator)
	if err != nil {
		return nil, err
	}

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())

	return buffer.Bytes(), nil
}
// =======Rich queries =========================================================================
// Two examples of rich queries are provided below (parameterized query and ad hoc query).
// Rich queries pass a query string to the state database.
// Rich queries are only supported by state database implementations
//  that support rich query (e.g. CouchDB).
// The query string is in the syntax of the underlying state database.
// With rich queries there is no guarantee that the result set hasn't changed between
//  endorsement time and commit time, aka 'phantom reads'.
// Therefore, rich queries should not be used in update transactions, unless the
// application handles the possibility of result set changes between endorsement and commit time.
// Rich queries can be used for point-in-time queries against a peer.
// ============================================================================================

// ===== Example: Parameterized rich query =================================================
// queryMarblesByOwner queries for marbles based on a passed in owner.
// This is an example of a parameterized query where the query logic is baked into the chaincode,
// and accepting a single query parameter (owner).
// Only available on state databases that support rich query (e.g. CouchDB)
// =========================================================================================
func (t *IdentityData) queryIdentityByType(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	//   0
	// "handle"
	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	itype := strings.ToLower(args[0])

	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"identity\",\"itype\":\"%s\"}}", itype)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}
