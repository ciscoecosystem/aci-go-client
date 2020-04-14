package tests

// import (
// 	"fmt"
// 	"testing"

// 	"github.com/ciscoecosystem/aci-go-client/client"
// 	"github.com/ciscoecosystem/aci-go-client/models"
// )

// func createSubject(c *client.Client, dn string, desc string, tenantName, contractName string, subjectAttrs models.SubjectAttributes) (*models.Subject, error) {
// 	subject := models.NewSubject(dn, desc, tenantName, contractName, subjectAttrs)
// 	err := c.Save(subject)
// 	return subject, err
// }

// func deleteSubject(c *client.Client, tenant *models.Tenant, contract *models.Contract, subject *models.Subject) error {
// 	err := c.Delete(subject)
// 	if err != nil {
// 		return err
// 	}
// 	err = deleteContarct(c, tenant, contract)
// 	if err != nil {
// 		return err
// 	}
// 	return deleteTenant(c, tenant)
// }

// func TestSubjectCreation(t *testing.T) {
// 	client := GetTestClient()
// 	tenant, err := createTenant(client, "terraform-test-tenant", "Test tenant created with golang aci-client-sdk.")
// 	contract, err := createContract(client, "terraform-test-contract", "Test Contarct created with golang aci-client-sdk", models.GetMOName(tenant.DistinguishedName), models.ContractAttributes{})

// 	subject, err := createSubject(client, "terraform-test-subject", "Test subject created with golang aci-client-sdk", models.GetMOName(tenant.DistinguishedName), models.GetMOName(contract.DistinguishedName), models.SubjectAttributes{})

// 	if err != nil {
// 		t.Error(err)
// 	}
// 	err = deleteSubject(client, tenant, contract, subject)
// 	if err != nil {
// 		t.Error(err)
// 	}

// }

// func TestDupliateSubject(t *testing.T) {
// 	client := GetTestClient()
// 	tenant, err := createTenant(client, "terraform-test-tenant", "Test tenant created with golang aci-client-sdk.")
// 	contract, err := createContract(client, "terraform-test-contract", "Test Contarct created with golang aci-client-sdk", models.GetMOName(tenant.DistinguishedName), models.ContractAttributes{})

// 	subject, err := createSubject(client, "terraform-test-subject", "Test subject created with golang aci-client-sdk", models.GetMOName(tenant.DistinguishedName), models.GetMOName(contract.DistinguishedName), models.SubjectAttributes{})

// 	if err != nil {
// 		t.Error(err)
// 	}
// 	_, err = createSubject(client, "terraform-test-subject", "Test subject created with golang aci-client-sdk", models.GetMOName(tenant.DistinguishedName), models.GetMOName(contract.DistinguishedName), models.SubjectAttributes{})
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	err = deleteSubject(client, tenant, contract, subject)
// 	if err != nil {
// 		t.Error(err)
// 	}

// }
// func TestGetSubjectContainer(t *testing.T) {
// 	client := GetTestClient()
// 	tenant, err := createTenant(client, "terraform-test-tenant", "Test tenant created with golang aci-client-sdk.")
// 	contract, err := createContract(client, "terraform-test-contract", "Test Contarct created with golang aci-client-sdk", models.GetMOName(tenant.DistinguishedName), models.ContractAttributes{})

// 	subject, err := createSubject(client, "terraform-test-subject", "Test subject created with golang aci-client-sdk", models.GetMOName(tenant.DistinguishedName), models.GetMOName(contract.DistinguishedName), models.SubjectAttributes{})

// 	cont, err := client.Get(subject.DistinguishedName)

// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	t.Logf("%+v", cont)
// 	err = deleteSubject(client, tenant, contract, subject)
// 	if err != nil {
// 		t.Error(err)
// 	}
// }

// func TestSubjectFromContainer(t *testing.T) {
// 	client := GetTestClient()
// 	tenant, err := createTenant(client, "terraform-test-tenant", "Test tenant created with golang aci-client-sdk.")
// 	contract, err := createContract(client, "terraform-test-contract", "Test Contarct created with golang aci-client-sdk", models.GetMOName(tenant.DistinguishedName), models.ContractAttributes{})

// 	subject, err := createSubject(client, "terraform-test-subject", "Test subject created with golang aci-client-sdk", models.GetMOName(tenant.DistinguishedName), models.GetMOName(contract.DistinguishedName), models.SubjectAttributes{})

// 	cont, err := client.Get(subject.DistinguishedName)

// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	subjectCon := models.SubjectFromContainer(cont)
//
// 	if subjectCon.DistinguishedName == "" {
// 		t.Error("the subject dn was empty")
// 	}
// 	err = deleteSubject(client, tenant, contract, subject)
// 	if err != nil {
// 		t.Error(err)
// 	}
// }

// func TestSubjectUpdate(t *testing.T) {
// 	client := GetTestClient()
// 	tenant, err := createTenant(client, "terraform-test-tenant", "Test tenant created with golang aci-client-sdk.")
// 	contract, err := createContract(client, "terraform-test-contract", "Test Contarct created with golang aci-client-sdk", models.GetMOName(tenant.DistinguishedName), models.ContractAttributes{})

// 	subject, err := createSubject(client, "terraform-test-subject", "Test subject created with golang aci-client-sdk", models.GetMOName(tenant.DistinguishedName), models.GetMOName(contract.DistinguishedName), models.SubjectAttributes{})

// 	cont, err := client.Get(subject.DistinguishedName)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	subjectCon := models.SubjectFromContainer(cont)
//
// 	if subjectCon.DistinguishedName == "" {
// 		t.Error("the subject dn was empty")
// 	}

// 	subjectCon.ProviderMatch = "AtleastOne"
// 	err = client.Save(subjectCon)
// 	if err != nil {
// 		t.Error(err)
// 	}
//
// 	err = deleteSubject(client, tenant, contract, subject)
// 	if err != nil {
// 		t.Error(err)
// 	}

// }

// func TestSubjectDelete(t *testing.T) {
// 	client := GetTestClient()
// 	tenant, err := createTenant(client, "terraform-test-tenant", "Test tenant created with golang aci-client-sdk.")
// 	contract, err := createContract(client, "terraform-test-contract", "Test Contarct created with golang aci-client-sdk", models.GetMOName(tenant.DistinguishedName), models.ContractAttributes{})

// 	subject, err := createSubject(client, "terraform-test-subject", "Test subject created with golang aci-client-sdk", models.GetMOName(tenant.DistinguishedName), models.GetMOName(contract.DistinguishedName), models.SubjectAttributes{})

// 	cont, err := client.Get(subject.DistinguishedName)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	subjectCon := models.SubjectFromContainer(cont)
//
// 	if subjectCon.DistinguishedName == "" {
// 		t.Error("the subject dn was empty")
// 	}

// 	err = client.Delete(subjectCon)
// 	if err != nil {
// 		t.Error("the subject was not removed")
// 	}
// 	err = deleteSubject(client, tenant, contract, subject)
// 	if err != nil {
// 		t.Error(err)
// 	}
// }
