package tests

// import (
// 	"fmt"
// 	"testing"

// 	"github.com/ciscoecosystem/aci-go-client/client"
// 	"github.com/ciscoecosystem/aci-go-client/models"
// )

// func createEntry(c *client.Client, dn, desc, tenantName, filterName string, entryAttr models.EntryAttributes) (*models.Entry, error) {
// 	entry := models.NewEntry(dn, desc, tenantName, filterName, entryAttr)
// 	err := c.Save(entry)
// 	return entry, err
// }

// func deleteEntry(c *client.Client, tenant *models.Tenant, filter *models.Filter, entry *models.Entry) error {
// 	err := c.Delete(entry)
// 	if err != nil {
// 		return err
// 	}
// 	err = deleteFilter(c, tenant, filter)
// 	if err != nil {
// 		return err
// 	}
// 	return deleteTenant(c, tenant)
// }

// func TestEntryCreation(t *testing.T) {
// 	client := GetTestClient()
// 	tenant, err := createTenant(client, "terraform-test-tenant", "Test tenant created with golang aci-client-sdk.")
// 	filter, err := createFilter(client, "terrafrom-test-filter", "Test filter created with golang aci-clienr-sdk", models.GetMOName(tenant.DistinguishedName))

// 	entry, err := createEntry(client, "terraform-test-entry", "Test Entry created with golang aci-client-sdk", models.GetMOName(tenant.DistinguishedName), models.GetMOName(filter.DistinguishedName), models.EntryAttributes{})

// 	if err != nil {
// 		t.Error(err)
// 	}
// 	err = deleteEntry(client, tenant, filter, entry)
// 	if err != nil {
// 		t.Error(err)
// 	}

// }

// func TestDuplicateEntry(t *testing.T) {
// 	client := GetTestClient()
// 	tenant, err := createTenant(client, "terraform-test-tenant", "Test tenant created with golang aci-client-sdk.")
// 	filter, err := createFilter(client, "terrafrom-test-filter", "Test filter created with golang aci-clienr-sdk", models.GetMOName(tenant.DistinguishedName))

// 	entry, err := createEntry(client, "terraform-test-entry", "Test Entry created with golang aci-client-sdk", models.GetMOName(tenant.DistinguishedName), models.GetMOName(filter.DistinguishedName), models.EntryAttributes{})

// 	if err != nil {
// 		t.Error(err)
// 	}
// 	_, err = createEntry(client, "terraform-test-entry", "Test Entry created with golang aci-client-sdk", models.GetMOName(tenant.DistinguishedName), models.GetMOName(filter.DistinguishedName), models.EntryAttributes{})
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	err = deleteEntry(client, tenant, filter, entry)
// 	if err != nil {
// 		t.Error(err)
// 	}

// }

// func TestGetEntryContainer(t *testing.T) {
// 	client := GetTestClient()
// 	tenant, err := createTenant(client, "terraform-test-tenant", "Test tenant created with golang aci-client-sdk.")
// 	filter, err := createFilter(client, "terrafrom-test-filter", "Test filter created with golang aci-clienr-sdk", models.GetMOName(tenant.DistinguishedName))

// 	entry, err := createEntry(client, "terraform-test-entry", "Test Entry created with golang aci-client-sdk", models.GetMOName(tenant.DistinguishedName), models.GetMOName(filter.DistinguishedName), models.EntryAttributes{})

// 	if err != nil {
// 		t.Error(err)
// 	}

// 	cont, err := client.Get(entry.DistinguishedName)

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	t.Logf("%+v", cont)

// 	err = deleteEntry(client, tenant, filter, entry)
// 	if err != nil {
// 		t.Error(err)
// 	}

// }

// func TestEntryFromContainer(t *testing.T) {

// 	client := GetTestClient()
// 	tenant, err := createTenant(client, "terraform-test-tenant", "Test tenant created with golang aci-client-sdk.")
// 	filter, err := createFilter(client, "terrafrom-test-filter", "Test filter created with golang aci-clienr-sdk", models.GetMOName(tenant.DistinguishedName))

// 	entry, err := createEntry(client, "terraform-test-entry", "Test Entry created with golang aci-client-sdk", models.GetMOName(tenant.DistinguishedName), models.GetMOName(filter.DistinguishedName), models.EntryAttributes{})

// 	if err != nil {
// 		t.Error(err)
// 	}

// 	cont, err := client.Get(entry.DistinguishedName)

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	entryCon := models.EntryFromContainer(cont)
// 	fmt.Println(entryCon.DistinguishedName)
// 	if entryCon.DistinguishedName == "" {
// 		t.Error("the entry dn was empty")

// 	}

// 	err = deleteEntry(client, tenant, filter, entry)
// 	if err != nil {
// 		t.Error(err)
// 	}
// }

// func TestEntryUpdate(t *testing.T) {
// 	client := GetTestClient()
// 	tenant, err := createTenant(client, "terraform-test-tenant", "Test tenant created with golang aci-client-sdk.")
// 	filter, err := createFilter(client, "terrafrom-test-filter", "Test filter created with golang aci-clienr-sdk", models.GetMOName(tenant.DistinguishedName))

// 	entry, err := createEntry(client, "terraform-test-entry", "Test Entry created with golang aci-client-sdk", models.GetMOName(tenant.DistinguishedName), models.GetMOName(filter.DistinguishedName), models.EntryAttributes{})

// 	if err != nil {
// 		t.Error(err)
// 	}

// 	cont, err := client.Get(entry.DistinguishedName)

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	entryCon := models.EntryFromContainer(cont)
// 	fmt.Println(entryCon.DistinguishedName)
// 	if entryCon.DistinguishedName == "" {
// 		t.Error("the entry dn was empty")

// 	}
// 	entryCon.Description = "Updated the Entry"
// 	err = client.Save(entryCon)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	fmt.Println("Entry was updated")

// 	err = deleteEntry(client, tenant, filter, entry)
// 	if err != nil {
// 		t.Error(err)
// 	}

// }

// func TestEntryDelete(t *testing.T) {
// 	client := GetTestClient()
// 	tenant, err := createTenant(client, "terraform-test-tenant", "Test tenant created with golang aci-client-sdk.")
// 	filter, err := createFilter(client, "terrafrom-test-filter", "Test filter created with golang aci-clienr-sdk", models.GetMOName(tenant.DistinguishedName))

// 	entry, err := createEntry(client, "terraform-test-entry", "Test Entry created with golang aci-client-sdk", models.GetMOName(tenant.DistinguishedName), models.GetMOName(filter.DistinguishedName), models.EntryAttributes{})

// 	if err != nil {
// 		t.Error(err)
// 	}

// 	cont, err := client.Get(entry.DistinguishedName)

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	entryCon := models.EntryFromContainer(cont)
// 	fmt.Println(entryCon.DistinguishedName)
// 	if entryCon.DistinguishedName == "" {
// 		t.Error("the entry dn was empty")

// 	}
// 	err = client.Delete(entryCon)
// 	if err != nil {
// 		t.Error("the entry was not removed")
// 	}

// 	err = deleteEntry(client, tenant, filter, entry)
// 	if err != nil {
// 		t.Error(err)
// 	}

// }
