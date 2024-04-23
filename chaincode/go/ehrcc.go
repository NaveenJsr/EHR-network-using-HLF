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
	EHR          string `json:"ehr,omitempty"`       // Electronic Health Record
	LabReport    string `json:"labReport,omitempty"` // Lab Report
}

// RegistryContract implements the smart contract
type RegistryContract struct {
	contractapi.Contract
}

// Initialize initializes the chaincode
func (rc *RegistryContract) Init(ctx contractapi.TransactionContextInterface) error {
    // You can add any initial state setup here
    // For example, let's add a dummy user during initialization
    user := User{
        Name:         "John Doe",
        Role:         "patient",
        Organization: "PatientOrg",
        EHR:          "Initial EHR",
        LabReport:    "Initial Lab Report",
    }

    // Marshal user to JSON
    userJSON, err := json.Marshal(user)
    if err != nil {
        return fmt.Errorf("failed to marshal user: %v", err)
    }

    // Put the user state in the ledger
    err = ctx.GetStub().PutState(user.Name, userJSON)
    if err != nil {
        return fmt.Errorf("failed to put user state: %v", err)
    }

    return nil
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

// AddEHR adds Electronic Health Record for a patient
func (rc *RegistryContract) AddEHR(ctx contractapi.TransactionContextInterface, patientName string, ehr string) error {
	// Get the patient's current state
	patientBytes, err := ctx.GetStub().GetState(patientName)
	if err != nil {
		return fmt.Errorf("failed to read patient state: %v", err)
	}
	if patientBytes == nil {
		return fmt.Errorf("patient does not exist")
	}

	var patient User
	err = json.Unmarshal(patientBytes, &patient)
	if err != nil {
		return fmt.Errorf("failed to unmarshal patient: %v", err)
	}

	// Add the EHR to the patient's state
	patient.EHR = ehr
	patientJSON, err := json.Marshal(patient)
	if err != nil {
		return fmt.Errorf("failed to marshal patient: %v", err)
	}

	// Put the updated patient state in the ledger
	err = ctx.GetStub().PutState(patientName, patientJSON)
	if err != nil {
		return fmt.Errorf("failed to put patient state: %v", err)
	}

	return nil
}

// AddLabReport allows a pathologist to upload lab reports for a patient
func (rc *RegistryContract) AddLabReport(ctx contractapi.TransactionContextInterface, patientName string, labReport string) error {
	// Get the patient's current state
	patientBytes, err := ctx.GetStub().GetState(patientName)
	if err != nil {
		return fmt.Errorf("failed to read patient state: %v", err)
	}
	if patientBytes == nil {
		return fmt.Errorf("patient does not exist")
	}

	var patient User
	err = json.Unmarshal(patientBytes, &patient)
	if err != nil {
		return fmt.Errorf("failed to unmarshal patient: %v", err)
	}

	// Add the lab report to the patient's state
	patient.LabReport = labReport
	patientJSON, err := json.Marshal(patient)
	if err != nil {
		return fmt.Errorf("failed to marshal patient: %v", err)
	}

	// Put the updated patient state in the ledger
	err = ctx.GetStub().PutState(patientName, patientJSON)
	if err != nil {
		return fmt.Errorf("failed to put patient state: %v", err)
	}

	return nil
}

// ViewEHR allows authorized users to view Electronic Health Record of a patient
func (rc *RegistryContract) ViewEHR(ctx contractapi.TransactionContextInterface, patientName string, invokingUserRole string) (string, error) {
	// Get the patient's current state
	patientBytes, err := ctx.GetStub().GetState(patientName)
	if err != nil {
		return "", fmt.Errorf("failed to read patient state: %v", err)
	}
	if patientBytes == nil {
		return "", fmt.Errorf("patient does not exist")
	}

	var patient User
	err = json.Unmarshal(patientBytes, &patient)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal patient: %v", err)
	}

	// Check if the invoking user is authorized to view EHR
	switch invokingUserRole {
	case "doctor", "patient", "pathologist", "insuranceAgent":
		return patient.EHR, nil
	default:
		return "", fmt.Errorf("unauthorized to view EHR")
	}
}



// validateRoleAndOrg validates the role and organization
func validateRoleAndOrg(role string, organization string) error {
	validRoles := map[string]bool{
		"doctor":         true,
		"patient":        true,
		"pathologist":    true,
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
