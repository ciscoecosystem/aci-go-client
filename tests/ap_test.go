package tests

// import "testing"
// import "fmt"
// import "github.com/ciscoecosystem/aci-go-client/client"
// import "github.com/ciscoecosystem/aci-go-client/models"

// func createAp(c *client.Client, dn string, desc string, tenantName string) (*models.ApplicationProfile, error) {

// 	ap := models.NewApplicationProfile(dn, desc, tenantName)
// 	err := c.Save(ap)
// 	return ap, err
// }

// func deleteAp(c *client.Client, tenant *models.Tenant, ap *models.ApplicationProfile) error {
// 	err := c.Delete(ap)

// 	if err != nil {
// 		return err
// 	}
// 	return deleteTenant(c, tenant)

// }

// func TestAPCreation(t *testing.T) {
// 	client := GetTestClient()
// 	tenant, err := createTenant(client, "terraform-test-tenant", "Test tenant created with golang aci-client-sdk.")
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	ap, err := createAp(client, "terrafrom-test-ap", "Test AP created with golang aci-client-sdk", models.GetMOName(tenant.DistinguishedName))
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	err = deleteAp(client, tenant, ap)
// 	if err != nil {
// 		t.Error(err)
// 	}
// }

// func TestDuplicateAP(t *testing.T) {
// 	client := GetTestClient()
// 	tenant, err := createTenant(client, "terraform-test-tenant", "Test tenant created with golang aci-client-sdk.")
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	ap, err := createAp(client, "terrafrom-test-ap", "Test AP created with golang aci-client-sdk", models.GetMOName(tenant.DistinguishedName))
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	_, err = createAp(client, "terrafrom-test-ap", "Test AP created with golang aci-client-sdk", models.GetMOName(tenant.DistinguishedName))
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	err = deleteAp(client, tenant, ap)
// 	if err != nil {
// 		t.Error(err)
// 	}

// }

// func TestGetAPContainer(t *testing.T) {
// 	client := GetTestClient()
// 	tenant, err := createTenant(client, "terraform-test-tenant", "Test tenant created with golang aci-client-sdk.")
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	ap, err := createAp(client, "terrafrom-test-ap", "Test AP created with golang aci-clienr-sdk", models.GetMOName(tenant.DistinguishedName))
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	cont, err := client.Get(ap.DistinguishedName)

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	t.Logf("%+v", cont)
// 	err = deleteAp(client, tenant, ap)
// 	if err != nil {
// 		t.Error(err)
// 	}
// }

// func TestAPFromContainer(t *testing.T) {
// 	client := GetTestClient()
// 	tenant, err := createTenant(client, "terraform-test-tenant", "Test tenant created with golang aci-client-sdk.")
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	ap, err := createAp(client, "terrafrom-test-ap", "Test AP created with golang aci-clienr-sdk", models.GetMOName(tenant.DistinguishedName))
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	cont, err := client.Get(ap.DistinguishedName)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	apCon := models.ApplicationProfileFromContainer(cont)
//
// 	if apCon.DistinguishedName == "" {
// 		t.Error("the application profile dn was empty")

// 	}
// 	err = deleteAp(client, tenant, ap)
// 	if err != nil {
// 		t.Error(err)
// 	}

// }

// func TestUpdateAP(t *testing.T) {
// 	client := GetTestClient()
// 	tenant, err := createTenant(client, "terraform-test-tenant", "Test tenant created with golang aci-client-sdk.")
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	ap, err := createAp(client, "terrafrom-test-ap", "Test AP created with golang aci-clienr-sdk", models.GetMOName(tenant.DistinguishedName))
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	cont, err := client.Get(ap.DistinguishedName)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	apCon := models.ApplicationProfileFromContainer(cont)
//
// 	if apCon.DistinguishedName == "" {
// 		t.Error("the application profile dn was empty")

// 	}

// 	apCon.Description = "Updated the description"
// 	err = client.Save(apCon)
// 	if err != nil {
// 		t.Error(err)
// 	}

//
// 	err = deleteAp(client, tenant, ap)
// 	if err != nil {
// 		t.Error(err)
// 	}

// }

// func TestAPDelete(t *testing.T) {
// 	client := GetTestClient()
// 	tenant, err := createTenant(client, "terraform-test-tenant", "Test tenant created with golang aci-client-sdk.")
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	ap, err := createAp(client, "terrafrom-test-ap", "Test AP created with golang aci-clienr-sdk", models.GetMOName(tenant.DistinguishedName))
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	cont, err := client.Get(ap.DistinguishedName)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	apCon := models.ApplicationProfileFromContainer(cont)
//
// 	if apCon.DistinguishedName == "" {
// 		t.Error("the application profile dn was empty")

// 	}

// 	err = client.Delete(apCon)
// 	if err != nil {
// 		t.Error("the application profile was not removed")
// 	}
// 	err = deleteAp(client, tenant, ap)
// 	if err != nil {
// 		t.Error(err)
// 	}
// }
