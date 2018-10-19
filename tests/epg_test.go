package tests

// import (
// 	"fmt"
// 	"testing"

// 	"github.com/ciscoecosystem/aci-go-client/client"
// 	"github.com/ciscoecosystem/aci-go-client/models"
// )

// func createEpg(c *client.Client, dn string, desc string, tenantName, apName string, epgAttr models.EPGAttributes) (*models.EPG, error) {
// 	epg := models.NewEPG(dn, desc, tenantName, apName, epgAttr)
// 	err := c.Save(epg)
// 	return epg, err
// }

// func deleteEpg(c *client.Client, tenant *models.Tenant, ap *models.ApplicationProfile, epg *models.EPG) error {
// 	err := c.Delete(epg)
// 	if err != nil {
// 		return err
// 	}
// 	err = deleteAp(c, tenant, ap)
// 	if err != nil {
// 		return err
// 	}
// 	return deleteTenant(c, tenant)
// }

// func TestEPGCreation(t *testing.T) {
// 	client := GetTestClient()
// 	tenant, err := createTenant(client, "terraform-test-tenant", "Test tenant created with golang aci-client-sdk.")
// 	ap, err := createAp(client, "terrafrom-test-ap", "Test AP created with golang aci-clienr-sdk", models.GetMOName(tenant.DistinguishedName))

// 	epgAttrs := models.DefaultEPGAttributes
// 	epg, err := createEpg(client, "terraform-test-epg", "Test EPG created with golang aci-client-sdk", models.GetMOName(tenant.DistinguishedName), models.GetMOName(ap.DistinguishedName), epgAttrs)

// 	if err != nil {
// 		t.Error(err)
// 	}
// 	err = deleteEpg(client, tenant, ap, epg)
// 	if err != nil {
// 		t.Error(err)
// 	}

// }

// func TestDuplicateEPG(t *testing.T) {
// 	client := GetTestClient()
// 	tenant, err := createTenant(client, "terraform-test-tenant", "Test tenant created with golang aci-client-sdk.")
// 	ap, err := createAp(client, "terrafrom-test-ap", "Test AP created with golang aci-clienr-sdk", models.GetMOName(tenant.DistinguishedName))

// 	epgAttrs := models.DefaultEPGAttributes
// 	epg, err := createEpg(client, "terraform-test-epg", "Test EPG created with golang aci-client-sdk", models.GetMOName(tenant.DistinguishedName), models.GetMOName(ap.DistinguishedName), epgAttrs)

// 	if err != nil {
// 		t.Error(err)
// 	}
// 	_, err = createEpg(client, "terraform-test-epg", "Test EPG created with golang aci-client-sdk", models.GetMOName(tenant.DistinguishedName), models.GetMOName(ap.DistinguishedName), epgAttrs)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	err = deleteEpg(client, tenant, ap, epg)
// 	if err != nil {
// 		t.Error(err)
// 	}

// }

// func TestGetEPGContainer(t *testing.T) {
// 	client := GetTestClient()
// 	tenant, err := createTenant(client, "terraform-test-tenant", "Test tenant created with golang aci-client-sdk.")
// 	ap, err := createAp(client, "terrafrom-test-ap", "Test AP created with golang aci-clienr-sdk", models.GetMOName(tenant.DistinguishedName))

// 	epgAttrs := models.DefaultEPGAttributes
// 	epg, err := createEpg(client, "terraform-test-epg", "Test EPG created with golang aci-client-sdk", models.GetMOName(tenant.DistinguishedName), models.GetMOName(ap.DistinguishedName), epgAttrs)
// 	cont, err := client.Get(epg.DistinguishedName)

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	t.Logf("%+v", cont)
// 	err = deleteEpg(client, tenant, ap, epg)
// 	if err != nil {
// 		t.Error(err)
// 	}
// }

// func TestEPGFromContainer(t *testing.T) {
// 	client := GetTestClient()
// 	tenant, err := createTenant(client, "terraform-test-tenant", "Test tenant created with golang aci-client-sdk.")
// 	ap, err := createAp(client, "terrafrom-test-ap", "Test AP created with golang aci-clienr-sdk", models.GetMOName(tenant.DistinguishedName))

// 	epgAttrs := models.DefaultEPGAttributes
// 	epg, err := createEpg(client, "terraform-test-epg", "Test EPG created with golang aci-client-sdk", models.GetMOName(tenant.DistinguishedName), models.GetMOName(ap.DistinguishedName), epgAttrs)
// 	cont, err := client.Get(epg.DistinguishedName)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	epgCon := models.EPGFromContainer(cont)
// 	fmt.Println(epgCon.DistinguishedName)
// 	if epgCon.DistinguishedName == "" {
// 		t.Error("the epg dn was empty")

// 	}
// 	err = deleteEpg(client, tenant, ap, epg)
// 	if err != nil {
// 		t.Error(err)
// 	}
// }

// func TestEPGUpdate(t *testing.T) {
// 	client := GetTestClient()
// 	tenant, err := createTenant(client, "terraform-test-tenant", "Test tenant created with golang aci-client-sdk.")
// 	ap, err := createAp(client, "terrafrom-test-ap", "Test AP created with golang aci-clienr-sdk", models.GetMOName(tenant.DistinguishedName))

// 	epgAttrs := models.DefaultEPGAttributes
// 	epg, err := createEpg(client, "terraform-test-epg", "Test EPG created with golang aci-client-sdk", models.GetMOName(tenant.DistinguishedName), models.GetMOName(ap.DistinguishedName), epgAttrs)
// 	cont, err := client.Get(epg.DistinguishedName)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	epgCon := models.EPGFromContainer(cont)
// 	fmt.Println(epgCon.DistinguishedName)
// 	if epgCon.DistinguishedName == "" {
// 		t.Error("the epg dn was empty")

// 	}

// 	epgCon.LabelMatchCriteria = "AtmostOne"
// 	err = client.Save(epgCon)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	fmt.Println("EPG was updated")
// 	err = deleteEpg(client, tenant, ap, epg)
// 	if err != nil {
// 		t.Error(err)
// 	}
// }

// func TestEPGDelete(t *testing.T) {
// 	client := GetTestClient()
// 	tenant, err := createTenant(client, "terraform-test-tenant", "Test tenant created with golang aci-client-sdk.")
// 	ap, err := createAp(client, "terrafrom-test-ap", "Test AP created with golang aci-clienr-sdk", models.GetMOName(tenant.DistinguishedName))

// 	epgAttrs := models.DefaultEPGAttributes
// 	epg, err := createEpg(client, "terraform-test-epg", "Test EPG created with golang aci-client-sdk", models.GetMOName(tenant.DistinguishedName), models.GetMOName(ap.DistinguishedName), epgAttrs)
// 	cont, err := client.Get(epg.DistinguishedName)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	epgCon := models.EPGFromContainer(cont)
// 	fmt.Println(epgCon.DistinguishedName)
// 	if epgCon.DistinguishedName == "" {
// 		t.Error("the epg dn was empty")

// 	}

// 	err = client.Delete(epgCon)
// 	if err != nil {
// 		t.Error("the epg was not removed")
// 	}
// 	err = deleteEpg(client, tenant, ap, epg)
// 	if err != nil {
// 		t.Error(err)
// 	}
// }
