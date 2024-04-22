package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// User defines the structure for a user
type User struct {
	Name         string `json:"name"`
	Role         string `json:"role"`
	Organization string `json:"organization"`
}

// RegistryContract implements the smart contract
type RegistryContract struct {
	contractapi.Contract
}

// RegisterUser registers a new user
func (rc *RegistryContract) RegisterUser(ctx contractapi.TransactionContextInterface, name string, role string, organization string) error {
	// Validate role and organization
	err := validateRoleAndOrg(role, organization)
	if err != nil {
		return err
	}

	// Create a new user
	user := User{
		Name:         name,
		Role:         role,
		Organization: organization,
	}

	// Marshal user to JSON
	userJSON, err := json.Marshal(user)
	if err != nil {
		return fmt.Errorf("failed to marshal user: %v", err)
	}

	// Put the user state in the ledger
	err = ctx.GetStub().PutState(name, userJSON)
	if err != nil {
		return fmt.Errorf("failed to put user state: %v", err)
	}

	return nil
}

// validateRoleAndOrg validates the role and organization
func validateRoleAndOrg(role string, organization string) error {
	validRoles := map[string]bool{
		"doctor":        true,
		"patient":       true,
		"pathologist":   true,
		"insuranceAgent": true,
	}

	validOrgs := map[string]bool{
		"HospitalOrg":  true,
		"PatientOrg":   true,
		"LabsOrg":      true,
		"InsuranceOrg": true,
	}

	if !validRoles[role] {
		return fmt.Errorf("invalid role: %s", role)
	}

	if organization == "" || !validOrgs[organization] {
		return fmt.Errorf("invalid organization: %s", organization)
	}

	return nil
}

func main() {
	registryContract := new(RegistryContract)

	cc, err := contractapi.NewChaincode(registryContract)
	if err != nil {
		fmt.Printf("Error creating registry chaincode: %v\n", err)
		return
	}

	if err := cc.Start(); err != nil {
		fmt.Printf("Error starting registry chaincode: %v\n", err)
	}
}
