package tests

// import (
// 	"fmt"
// 	"testing"

// 	"github.com/ciscoecosystem/aci-go-client/client"
// 	"github.com/ciscoecosystem/aci-go-client/models"
// )

// func createContract(c *client.Client, dn string, desc string, parentDn string, contractAttr models.ContractAttributes) (*models.Contract, error) {
// 	contract := models.NewContract(dn, desc, parentDn, contractAttr)
// 	err := c.Save(contract)
// 	return contract, err
// }

// func deleteContarct(c *client.Client, tenant *models.Tenant, contract *models.Contract) error {
// 	err := c.Delete(contract)
// 	if err != nil {
// 		return err
// 	}
// 	return deleteTenant(c, tenant)
// }
// func TestContractCreation(t *testing.T) {
// 	client := GetTestClient()
// 	tenant, err := createTenant(client, "terraform-test-tenant", "Test tenant created with golang aci-client-sdk.")
// 	contract, err := createContract(client, "terraform-test-contract", "Test Contarct created with golang aci-client-sdk", models.GetMOName(tenant.DistinguishedName), models.ContractAttributes{})

// 	if err != nil {
// 		t.Error(err)
// 	}
// 	err = deleteContarct(client, tenant, contract)
// 	if err != nil {
// 		t.Error(err)
// 	}

// }

// func TestDuplicateContract(t *testing.T) {
// 	client := GetTestClient()
// 	tenant, err := createTenant(client, "terraform-test-tenant", "Test tenant created with golang aci-client-sdk.")
// 	contract, err := createContract(client, "terraform-test-contract", "Test Contarct created with golang aci-client-sdk", models.GetMOName(tenant.DistinguishedName), models.ContractAttributes{})

// 	if err != nil {
// 		t.Error(err)
// 	}
// 	_, err = createContract(client, "terraform-test-contract", "Test Contarct created with golang aci-client-sdk", models.GetMOName(tenant.DistinguishedName), models.ContractAttributes{})
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	err = deleteContarct(client, tenant, contract)
// 	if err != nil {
// 		t.Error(err)
// 	}

// }

// func TestGetContractContainer(t *testing.T) {
// 	client := GetTestClient()
// 	tenant, err := createTenant(client, "terraform-test-tenant", "Test tenant created with golang aci-client-sdk.")
// 	contract, err := createContract(client, "terraform-test-contract", "Test Contarct created with golang aci-client-sdk", models.GetMOName(tenant.DistinguishedName), models.ContractAttributes{})
// 	cont, err := client.Get(contract.DistinguishedName)

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	t.Logf("%+v", cont)
// 	err = deleteContarct(client, tenant, contract)
// 	if err != nil {
// 		t.Error(err)
// 	}
// }

// func TestContractFromContainer(t *testing.T) {
// 	client := GetTestClient()
// 	tenant, err := createTenant(client, "terraform-test-tenant", "Test tenant created with golang aci-client-sdk.")
// 	contract, err := createContract(client, "terraform-test-contract", "Test Contarct created with golang aci-client-sdk", models.GetMOName(tenant.DistinguishedName), models.ContractAttributes{})
// 	cont, err := client.Get(contract.DistinguishedName)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	contractCon := models.ContractFromContainer(cont)
// 	fmt.Println(contractCon.DistinguishedName)
// 	if contractCon.DistinguishedName == "" {
// 		t.Error("the contract dn was empty")

// 	}
// 	err = deleteContarct(client, tenant, contract)
// 	if err != nil {
// 		t.Error(err)
// 	}
// }

// func TestContractUpdate(t *testing.T) {
// 	client := GetTestClient()
// 	tenant, err := createTenant(client, "terraform-test-tenant", "Test tenant created with golang aci-client-sdk.")
// 	contract, err := createContract(client, "terraform-test-contract", "Test Contarct created with golang aci-client-sdk", models.GetMOName(tenant.DistinguishedName), models.ContractAttributes{})
// 	cont, err := client.Get(contract.DistinguishedName)

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	contractCon := models.ContractFromContainer(cont)
// 	fmt.Println(contractCon.DistinguishedName)
// 	if contractCon.DistinguishedName == "" {
// 		t.Error("the contract dn was empty")

// 	}

// 	contractCon.Scope = "tenant"
// 	err = client.Save(contractCon)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	fmt.Println("Contract was Updated")
// 	err = deleteContarct(client, tenant, contract)
// 	if err != nil {
// 		t.Error(err)
// 	}

// }
// func TestContractDelete(t *testing.T) {
// 	client := GetTestClient()
// 	tenant, err := createTenant(client, "terraform-test-tenant", "Test tenant created with golang aci-client-sdk.")
// 	contract, err := createContract(client, "terraform-test-contract", "Test Contarct created with golang aci-client-sdk", models.GetMOName(tenant.DistinguishedName), models.ContractAttributes{})
// 	cont, err := client.Get(contract.DistinguishedName)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	contractCon := models.ContractFromContainer(cont)
// 	fmt.Println(contractCon.DistinguishedName)
// 	if contractCon.DistinguishedName == "" {
// 		t.Error("the contract dn was empty")

// 	}

// 	err = client.Delete(contractCon)
// 	if err != nil {
// 		t.Error("the contract was not removed")
// 	}
// 	err = deleteContarct(client, tenant, contract)
// 	if err != nil {
// 		t.Error(err)
// 	}

// }
