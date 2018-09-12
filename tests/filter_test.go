package tests

import (
	"fmt"
	"testing"

	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/ciscoecosystem/aci-go-client/models"
)

func createFilter(c *client.Client, dn, desc, parentDn string) (*models.Filter, error) {
	filter := models.NewFilter(dn, desc, parentDn)
	err := c.Save(filter)
	return filter, err
}

func deleteFilter(c *client.Client, tenant *models.Tenant, filter *models.Filter) error {
	err := c.Delete(filter)

	if err != nil {
		return err
	}
	return deleteTenant(c, tenant)

}

func TestFilterCreation(t *testing.T) {
	client := GetTestClient()
	tenant, err := createTenant(client, "terraform-test-tenant", "Test tenant created with golang aci-client-sdk.")
	if err != nil {
		t.Error(err)
	}

	filter, err := createFilter(client, "terrafrom-test-filter", "Test Filter created with golang aci-clienr-sdk", models.GetMOName(tenant.DistinguishedName))
	if err != nil {
		t.Error(err)
	}

	err = deleteFilter(client, tenant, filter)
	if err != nil {
		t.Error(err)
	}
}

func TestDuplicateFilter(t *testing.T) {
	client := GetTestClient()
	tenant, err := createTenant(client, "terraform-test-tenant", "Test tenant created with golang aci-client-sdk.")
	if err != nil {
		t.Error(err)
	}

	filter, err := createFilter(client, "terrafrom-test-filter", "Test Filter created with golang aci-clienr-sdk", models.GetMOName(tenant.DistinguishedName))
	if err != nil {
		t.Error(err)
	}
	_, err = createFilter(client, "terrafrom-test-filter", "Test Filter created with golang aci-clienr-sdk", models.GetMOName(tenant.DistinguishedName))
	if err != nil {
		t.Error(err)
	}

	err = deleteFilter(client, tenant, filter)
	if err != nil {
		t.Error(err)
	}

}

func TestGetFilterContainer(t *testing.T) {
	client := GetTestClient()
	tenant, err := createTenant(client, "terraform-test-tenant", "Test tenant created with golang aci-client-sdk.")
	if err != nil {
		t.Error(err)
	}

	filter, err := createFilter(client, "terrafrom-test-filter", "Test Filter created with golang aci-clienr-sdk", models.GetMOName(tenant.DistinguishedName))

	if err != nil {
		t.Error(err)
	}
	cont, err := client.Get(filter.DistinguishedName)

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", cont)
	err = deleteFilter(client, tenant, filter)
	if err != nil {
		t.Error(err)
	}
}

func TestFilterFromContainer(t *testing.T) {
	client := GetTestClient()
	tenant, err := createTenant(client, "terraform-test-tenant", "Test tenant created with golang aci-client-sdk.")
	if err != nil {
		t.Error(err)
	}

	filter, err := createFilter(client, "terrafrom-test-filter", "Test Filter created with golang aci-clienr-sdk", models.GetMOName(tenant.DistinguishedName))

	if err != nil {
		t.Error(err)
	}
	cont, err := client.Get(filter.DistinguishedName)
	if err != nil {
		t.Fatal(err)
	}

	filterCon := models.FilterFromContainer(cont)
	fmt.Println(filterCon.DistinguishedName)
	if filterCon.DistinguishedName == "" {
		t.Error("the filter dn was empty")

	}
	err = deleteFilter(client, tenant, filter)
	if err != nil {
		t.Error(err)
	}

}

func TestUpdateFilter(t *testing.T) {
	client := GetTestClient()
	tenant, err := createTenant(client, "terraform-test-tenant", "Test tenant created with golang aci-client-sdk.")
	if err != nil {
		t.Error(err)
	}

	filter, err := createFilter(client, "terrafrom-test-filter", "Test Filter created with golang aci-clienr-sdk", models.GetMOName(tenant.DistinguishedName))

	if err != nil {
		t.Error(err)
	}
	cont, err := client.Get(filter.DistinguishedName)
	if err != nil {
		t.Fatal(err)
	}

	filterCon := models.FilterFromContainer(cont)
	fmt.Println(filterCon.DistinguishedName)
	if filterCon.DistinguishedName == "" {
		t.Error("the filter dn was empty")

	}

	filterCon.Description = "Updated the description"
	err = client.Save(filterCon)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("Updated Filter")
	err = deleteFilter(client, tenant, filter)
	if err != nil {
		t.Error(err)
	}
}

func TestFilterDelete(t *testing.T) {
	client := GetTestClient()
	tenant, err := createTenant(client, "terraform-test-tenant", "Test tenant created with golang aci-client-sdk.")
	if err != nil {
		t.Error(err)
	}

	filter, err := createFilter(client, "terrafrom-test-filter", "Test Filter created with golang aci-clienr-sdk", models.GetMOName(tenant.DistinguishedName))

	if err != nil {
		t.Error(err)
	}
	cont, err := client.Get(filter.DistinguishedName)
	if err != nil {
		t.Fatal(err)
	}

	filterCon := models.FilterFromContainer(cont)
	fmt.Println(filterCon.DistinguishedName)
	if filterCon.DistinguishedName == "" {
		t.Error("the filter dn was empty")

	}
	err = client.Delete(filterCon)
	if err != nil {
		t.Error("the filter was not removed")
	}
	err = deleteFilter(client, tenant, filter)
	if err != nil {
		t.Error(err)
	}

}
