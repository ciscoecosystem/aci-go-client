package tests

// import (
// 	"fmt"
// 	"testing"

// 	"github.com/ciscoecosystem/aci-go-client/client"
// 	"github.com/ciscoecosystem/aci-go-client/models"
// )

// func createSubnet(c *client.Client, dn string, desc string, tenantName, bdName string, subnetAttrs models.SubnetAttributes) (*models.Subnet, error) {

// 	subnet := models.NewSubnet(dn, desc, tenantName, bdName, subnetAttrs)
// 	err := c.Save(subnet)
// 	return subnet, err
// }

// func deleteSubnet(c *client.Client, tenant *models.Tenant, bd *models.BridgeDomain, subnet *models.Subnet) error {
// 	err := c.Delete(subnet)
// 	if err != nil {
// 		return err
// 	}
// 	err = deleteBd(c, tenant, bd)
// 	if err != nil {
// 		return err
// 	}
// 	return deleteTenant(c, tenant)
// }

// func TestSubnetCreation(t *testing.T) {
// 	client := GetTestClient()
// 	tenant, err := createTenant(client, "terraform-test-tenant", "Test tenant created with golang aci-client-sdk.")
// 	bdattr := models.BridgeDomainAttributes{}
// 	bd, err := createBd(client, "terraform-test-bd", "Test Bridge Domain created with golang aci-client-sdk", models.GetMOName(tenant.DistinguishedName), bdattr)
// 	subnet, err := createSubnet(client, "[10.0.0.29/27]", "Test subnet created with golang aci-client-sdk", models.GetMOName(tenant.DistinguishedName), models.GetMOName(bd.DistinguishedName), models.SubnetAttributes{})

// 	if err != nil {
// 		t.Error(err)
// 	}
// 	err = deleteSubnet(client, tenant, bd, subnet)
// 	if err != nil {
// 		t.Error(err)
// 	}
// }

// func TestDuplicateSubnet(t *testing.T){
// 	client := GetTestClient()
// 	tenant, err := createTenant(client, "terraform-test-tenant", "Test tenant created with golang aci-client-sdk.")
// 	bdattr := models.BridgeDomainAttributes{}
// 	bd, err := createBd(client, "terraform-test-bd", "Test Bridge Domain created with golang aci-client-sdk", models.GetMOName(tenant.DistinguishedName), bdattr)
// 	subnet, err := createSubnet(client, "[10.0.0.29/27]", "Test subnet created with golang aci-client-sdk", models.GetMOName(tenant.DistinguishedName), models.GetMOName(bd.DistinguishedName), models.SubnetAttributes{})

// 	if err != nil {
// 		t.Error(err)
// 	}
// 	_, err = createSubnet(client, "[10.0.0.29/27]", "Test subnet created with golang aci-client-sdk", models.GetMOName(tenant.DistinguishedName), models.GetMOName(bd.DistinguishedName), models.SubnetAttributes{})
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	err = deleteSubnet(client, tenant, bd, subnet)
// 	if err != nil {
// 		t.Error(err)
// 	}

// }

// func TestGetSubnetContainer(t *testing.T) {
// 	client := GetTestClient()
// 	tenant, err := createTenant(client, "terraform-test-tenant", "Test tenant created with golang aci-client-sdk.")
// 	bdattr := models.BridgeDomainAttributes{}
// 	bd, err := createBd(client, "terraform-test-bd", "Test Bridge Domain created with golang aci-client-sdk", models.GetMOName(tenant.DistinguishedName), bdattr)
// 	subnet, err := createSubnet(client, "[10.0.0.29/27]", "Test subnet created with golang aci-client-sdk", models.GetMOName(tenant.DistinguishedName), models.GetMOName(bd.DistinguishedName), models.SubnetAttributes{})
// 	cont, err := client.Get(subnet.DistinguishedName)

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	t.Logf("%+v", cont)
// 	err = deleteSubnet(client, tenant, bd, subnet)
// 	if err != nil {
// 		t.Error(err)
// 	}
// }

// func TestSubnetFromContainer(t *testing.T) {
// 	client := GetTestClient()
// 	tenant, err := createTenant(client, "terraform-test-tenant", "Test tenant created with golang aci-client-sdk.")
// 	bdattr := models.BridgeDomainAttributes{}
// 	bd, err := createBd(client, "terraform-test-bd", "Test Bridge Domain created with golang aci-client-sdk", models.GetMOName(tenant.DistinguishedName), bdattr)
// 	subnet, err := createSubnet(client, "[10.0.0.29/27]", "Test subnet created with golang aci-client-sdk", models.GetMOName(tenant.DistinguishedName), models.GetMOName(bd.DistinguishedName), models.SubnetAttributes{})
// 	cont, err := client.Get(subnet.DistinguishedName)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	subnetCon := models.SubnetFromContainer(cont)
//
// 	if subnetCon.DistinguishedName == "" {
// 		t.Error("The subnet dn was empty")
// 	}
// 	err = deleteSubnet(client, tenant, bd, subnet)
// 	if err != nil {
// 		t.Error(err)
// 	}
// }

// func TestSubnetUpdate(t *testing.T) {
// 	client := GetTestClient()
// 	tenant, err := createTenant(client, "terraform-test-tenant", "Test tenant created with golang aci-client-sdk.")
// 	bdattr := models.BridgeDomainAttributes{}
// 	bd, err := createBd(client, "terraform-test-bd", "Test Bridge Domain created with golang aci-client-sdk", models.GetMOName(tenant.DistinguishedName), bdattr)
// 	subnet, err := createSubnet(client, "[10.0.0.29/27]", "Test subnet created with golang aci-client-sdk", models.GetMOName(tenant.DistinguishedName), models.GetMOName(bd.DistinguishedName), models.SubnetAttributes{})
// 	cont, err := client.Get(subnet.DistinguishedName)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	subnetCon := models.SubnetFromContainer(cont)
//
// 	if subnetCon.DistinguishedName == "" {
// 		t.Error("The subnet dn was empty")
// 	}

// 	var scopes []string
// 	scopes = append(scopes, "public")
// 	subnetCon.Scope = scopes
// 	err = client.Save(subnetCon)
// 	if err != nil {
// 		t.Error(err)
// 	}
//
// 	err = deleteSubnet(client, tenant, bd, subnet)
// 	if err != nil {
// 		t.Error(err)
// 	}
// }

// func TestSubnetDelete(t *testing.T) {
// 	client := GetTestClient()
// 	tenant, err := createTenant(client, "terraform-test-tenant", "Test tenant created with golang aci-client-sdk.")
// 	bdattr := models.BridgeDomainAttributes{}
// 	bd, err := createBd(client, "terraform-test-bd", "Test Bridge Domain created with golang aci-client-sdk", models.GetMOName(tenant.DistinguishedName), bdattr)
// 	subnet, err := createSubnet(client, "[10.0.0.29/27]", "Test subnet created with golang aci-client-sdk", models.GetMOName(tenant.DistinguishedName), models.GetMOName(bd.DistinguishedName), models.SubnetAttributes{})
// 	cont, err := client.Get(subnet.DistinguishedName)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	subnetCon := models.SubnetFromContainer(cont)
//
// 	if subnetCon.DistinguishedName == "" {
// 		t.Error("The subnet dn was empty")
// 	}

// 	err = client.Delete(subnetCon)
// 	if err != nil {
// 		t.Error("The subnet was not removed")
// 	}
// 	err = deleteSubnet(client, tenant, bd, subnet)
// 	if err != nil {
// 		t.Error(err)
// 	}
// }
