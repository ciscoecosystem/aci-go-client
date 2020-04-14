package tests

// import (
// 	"fmt"
// 	"testing"

// 	"github.com/ciscoecosystem/aci-go-client/client"
// 	"github.com/ciscoecosystem/aci-go-client/models"
// )

// func createBd(c *client.Client, dn string, desc string, parentDn string, bdattr models.BridgeDomainAttributes) (*models.BridgeDomain, error) {

// 	return c.CreateBridgeDomain(dn, desc, parentDn, bdattr)
// }

// func deleteBd(c *client.Client, tenant *models.Tenant, bd *models.BridgeDomain) error {
// 	err := c.DeleteBridgeDomain(models.GetMOName(bd.DistinguishedName), models.GetMOName(tenant.DistinguishedName))
// 	if err != nil {
// 		return err
// 	}
// 	return deleteTenant(c, tenant)

// }

// func TestBDCreation(t *testing.T) {
// 	client := GetTestClient()
// 	tenant, err := createTenant(client, "terraform-test-tenant", "Test tenant created with golang aci-client-sdk.")
// 	bdattr := models.BridgeDomainAttributes{}
// 	bd, err := createBd(client, "terraform-test-bd", "Test Bridge Domain created with golang aci-client-sdk", models.GetMOName(tenant.DistinguishedName), bdattr)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	err = deleteBd(client, tenant, bd)
// 	if err != nil {
// 		t.Error(err)
// 	}
// }

// func TestDuplicateBD(t *testing.T) {
// 	client := GetTestClient()
// 	tenant, err := createTenant(client, "terraform-test-tenant", "Test tenant created with golang aci-client-sdk.")
// 	bdattr := models.BridgeDomainAttributes{}
// 	bd, err := createBd(client, "terraform-test-bd", "Test Bridge Domain created with golang aci-client-sdk", models.GetMOName(tenant.DistinguishedName), bdattr)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	_, err = createBd(client, "terraform-test-bd", "Test Bridge Domain created with golang aci-client-sdk", models.GetMOName(tenant.DistinguishedName), bdattr)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	err = deleteBd(client, tenant, bd)
// 	if err != nil {
// 		t.Error(err)
// 	}

// }

// func TestBDUpdate(t *testing.T) {
// 	client := GetTestClient()
// 	tenant, err := createTenant(client, "terraform-test-tenant", "Test tenant created with golang aci-client-sdk.")
// 	bdattr := models.BridgeDomainAttributes{}
// 	bd, err := createBd(client, "terraform-test-bd", "Test Bridge Domain created with golang aci-client-sdk", models.GetMOName(tenant.DistinguishedName), bdattr)
// 	bdCon, err := client.ReadBridgeDomain(models.GetMOName(bd.DistinguishedName), models.GetMOName(tenant.DistinguishedName))
// 	if err != nil {
// 		t.Fatal(err)
// 	}

//
// 	if bdCon.DistinguishedName == "" {
// 		t.Error("the Bridge Domain dn was empty")

// 	}

// 	bdCon.MAC = "00:22:BD:88:AA:BB"
// 	_, err = client.UpdateBridgeDomain(models.GetMOName(bdCon.DistinguishedName), bdCon.Description, models.GetMOName(tenant.DistinguishedName), bdCon.BridgeDomainAttributes)
// 	if err != nil {
// 		t.Error(err)
// 	}
//
// 	err = deleteBd(client, tenant, bd)
// 	if err != nil {
// 		t.Error(err)
// 	}
// }

// func TestBDDelete(t *testing.T) {
// 	client := GetTestClient()
// 	tenant, err := createTenant(client, "terraform-test-tenant", "Test tenant created with golang aci-client-sdk.")
// 	bdattr := models.BridgeDomainAttributes{}
// 	bd, err := createBd(client, "terraform-test-bd", "Test Bridge Domain created with golang aci-client-sdk", models.GetMOName(tenant.DistinguishedName), bdattr)
// 	bdCon, err := client.ReadBridgeDomain(models.GetMOName(bd.DistinguishedName), models.GetMOName(tenant.DistinguishedName))
// 	if err != nil {
// 		t.Fatal(err)
// 	}

//
// 	if bdCon.DistinguishedName == "" {
// 		t.Error("the Bridge Domain dn was empty")

// 	}

// 	err = client.DeleteBridgeDomain(models.GetMOName(bdCon.DistinguishedName), models.GetMOName(tenant.DistinguishedName))
// 	if err != nil {
// 		t.Error("the Bridge Domain was not removed")
// 	}
// 	err = deleteBd(client, tenant, bd)
// 	if err != nil {
// 		t.Error(err)
// 	}
// }

// func TestListBD(t *testing.T) {
// 	client := GetTestClient()
// 	tenant, err := createTenant(client, "terraform-test-tenant", "Test tenant created with golang aci-client-sdk.")
// 	bdattr := models.BridgeDomainAttributes{}
// 	bd, err := createBd(client, "terraform-test-bd", "Test Bridge Domain created with golang aci-client-sdk", models.GetMOName(tenant.DistinguishedName), bdattr)
// 	bd1, err := createBd(client, "terraform-test-bd1", "Test Bridge Domain created with golang aci-client-sdk", models.GetMOName(tenant.DistinguishedName), bdattr)
// 	bd2, err := createBd(client, "terraform-test-bd2", "Test Bridge Domain created with golang aci-client-sdk", models.GetMOName(tenant.DistinguishedName), bdattr)

// 	if err != nil {
// 		t.Error(err)
// 	}

// 	cont, err := client.ListBridgeDomain(models.GetMOName(tenant.DistinguishedName))

//
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	err = deleteBd(client, tenant, bd)
// 	err = deleteBd(client, tenant, bd1)
// 	err = deleteBd(client, tenant, bd2)

// 	t.Error(err)

// }
