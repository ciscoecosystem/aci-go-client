package tests

import (
	"fmt"
	"testing"
	"time"

	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/ciscoecosystem/aci-go-client/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

const privKey = `
-----BEGIN PRIVATE KEY-----
MIICeAIBADANBgkqhkiG9w0BAQEFAASCAmIwggJeAgEAAoGBANCgcS0CtMkwTDFG
E1bYJ3Gt3XzHJhV75i8TjzGswzwBasXqFTLsxaqAf7NvrHu/O+k2KN+2kOLhLBAp
f1FdhJaTeualowF9UAUGY6HX28jaXjUkKEUNVWasl7hYiKSQaTKKh5eCJ0kwn6Ta
zNeZs1NoW3JjnA7WolC85Z5tdJ3RAgMBAAECgYEAkuOUK+vO3CScUMkVDr1MMFcJ
LjSNlOzSaezskj4gjBw3UDf7Swq4Nc8Zmn4TRGGlzhKq/rBtHMELpCmDkgc1NQT0
aHsgR+RFtJnYkafuEISuneRZEIz7b2AwUCiDw2z5uJ7xxgCM9MTibXIRRtjTgq1k
KmllK06vgg2Y9S/28pECQQD7cV1auzUFVojL5/E+ZofUAMtqLduVnwsQiQa9A683
5jOgRXJCwOyXJrEyfHUBAt+FW3uIt8iIo0IOLjb0y/NdAkEA1GhruuR/YGRx/rPt
RO+tt3BPznpo5sb411bVBWn2WIh34ocigr3I7KyeQqePlZjZTHIJhVa5sTHvBnJI
gMOBBQJBAJOCAwFKWCWEiYYN0OIJpl+CA9OLiYlyHzyZFoHmWnGRs/GvLAPrSCC4
SzXR+5YXSGfkrkkbgCJgnEzTYdwkleUCQQCs0zm0m264s4G9SBDqYknqU8vbqOXp
wEOAkvpIqWrzpjZFbsa5sknlqJ4shcHiareD59WvVF1Ku+JMUHiFrI0xAkA/r1il
G8oH/CjSLRukLRkEPfNgVPVk4ZjJiOXR7H6kc/KAJ084Ba5PuvjpC/Z99AHYAjYS
+ZXowhw0oCckCjLf
-----END PRIVATE KEY-----
`

func GetTestClient() *client.Client {
	return client.GetClient("https://192.168.10.102", "admin", client.Insecure(true), client.PrivateKey("/Users/nirav.katarmal/Documents/github/aci_test/admin.key"), client.AdminCert("admin"))

}

// func GetTestClient() *client.Client {
// 	return client.GetClient("https://192.168.10.102", "admin", client.Insecure(true), client.Password("cisco123"))

// }

// func TestTenantPrepareModel(t *testing.T) {
// 	c := GetTestClient()

// 	cont, _, err := c.PrepareModel(models.NewTenant("terraform-test-tenant", "A test tenant created with aci-client-sdk."))

// 	if err != nil {
// 		t.Error(err)
// 	}
// 	if !cont.ExistsP("FvTenant.attributes.dn") {
// 		t.Error("malformed model")
// 	}
// }

func createTenant(c *client.Client, dn string, desc string) (*models.Tenant, error) {
	tenant := models.NewTenant(fmt.Sprintf("tn-%s", dn), "uni", desc, models.TenantAttributes{})
	err := c.Save(tenant)
	return tenant, err
}

func deleteTenant(c *client.Client, tenant *models.Tenant) error {
	return c.Delete(tenant)
}

func TestTenantCreation(t *testing.T) {
	c := GetTestClient()
	tenant, err := createTenant(c, "terraform-test-tenant", "A test tenant created with aci-client-sdk.")
	tenant2, err := createTenant(c, "terraform-test-tenantwe", "A test tenant created with aci-client-sdk.")
	tenant3, err := createTenant(c, "terraform-test-tenantwert", "A test tenant created with aci-client-sdk.")

	if err != nil {
		t.Error(err)
	}

	err = deleteTenant(c, tenant)
	if err != nil {
		t.Error(err)
	}
	err = deleteTenant(c, tenant2)
	if err != nil {
		t.Error(err)
	}
	err = deleteTenant(c, tenant3)
	if err != nil {
		t.Error(err)
	}
}

func TestDomCreate(t *testing.T) {
	c := GetTestClient()
	dom := models.NewVMMDomain(fmt.Sprintf("dom-%s", "test"), "uni/vmmp-VMware", "", models.VMMDomainAttributes{})
	err := c.Save(dom)
	t.Error(err)
}
func TestGetDN(t *testing.T) {
	c := GetTestClient()
	path := "api/node/mo/uni/tn-tenant_for_bd/out-testext.json"
	className := "l3extOut"
	cont, err := c.GetViaURL(path)
	dn := cont.Search("imdata", className, "attributes", "dn").String()
	fmt.Print("container", cont)
	fmt.Print(dn)
	t.Error(err)

}

func TestDuplicateTenant(t *testing.T) {
	c := GetTestClient()
	tenant1, err := createTenant(c, "terraform-test-tenant", "A test tenant created with aci-client-sdk.")
	if err != nil {
		t.Error(err)
	}
	_, err = createTenant(c, "terraform-test-tenant", "A test tenant created with aci-client-sdk.")
	if err != nil {
		t.Error(err)
	}

	err = deleteTenant(c, tenant1)
	if err != nil {
		t.Error(err)
	}

}

func TestGetTenantContainer(t *testing.T) {

	c := GetTestClient()
	tenant, _ := createTenant(c, "terraform-test-tenant", "A test tenant created with aci-client-sdk.")
	cont, err := c.Get("uni/tn-terraform-test-tenant")

	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Tenant read container %+v", cont)

	err = deleteTenant(c, tenant)
	if err != nil {
		t.Error(err)
	}
	t.Error(err)
}

func TestTenantFromContainer(t *testing.T) {
	c := GetTestClient()
	tenant, _ := createTenant(c, "terraform-test-tenant", "A test tenant created with aci-client-sdk.")
	cont, err := c.Get("uni/tn-terraform-test-tenant")
	if err != nil {
		t.Error(err)
	}
	tenantCon := models.TenantFromContainer(cont)
	fmt.Println(tenantCon.DistinguishedName)
	if tenantCon.DistinguishedName == "" {
		t.Error("the tenant dn was empty")
	}
	err = deleteTenant(c, tenant)
	if err != nil {
		t.Error(err)
	}
}

func TestUpdateTenant(t *testing.T) {
	client := GetTestClient()
	tenant, _ := createTenant(client, "terraform-test-tenant", "A test tenant created with aci-client-sdk.")
	cont, err := client.Get("uni/tn-terraform-test-tenant")
	if err != nil {
		t.Error(err)
	}
	tenantCon := models.TenantFromContainer(cont)
	if tenantCon.DistinguishedName == "" {
		t.Error("the tenant dn was empty")
	}
	tenantCon.Description = "Updated the description "
	err = client.Save(tenantCon)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("Description Updated for tenant")
	err = deleteTenant(client, tenant)
	if err != nil {
		t.Error(err)
	}
}

func TestTenantDelete(t *testing.T) {
	c := GetTestClient()
	tenant, _ := createTenant(c, "terraform-test-tenant", "A test tenant created with aci-client-sdk.")
	cont, err := c.Get("uni/tn-terraform-test-tenant")
	if err != nil {
		t.Error(err)
	}
	tenantCon := models.TenantFromContainer(cont)
	fmt.Println(tenantCon.DistinguishedName)
	if tenantCon.DistinguishedName == "" {
		t.Error("the tenant dn was empty")
	}

	err = c.Delete(tenant)
	if err != nil {
		t.Error("the tenant was not remove")
	}
	err = deleteTenant(c, tenant)
	if err != nil {
		t.Error(err)
	}

}

func TestReadRel(t *testing.T) {
	c := GetTestClient()
	tenant, err := createTenant(c, "terraform-test-tenant", "A test tenant created with aci-client-sdk.")
	time.Sleep(1000 * time.Millisecond)
	c.CreateRelationfvRsTnDenyRuleFromTenant(tenant.DistinguishedName, "uni/tn-terraform-test-tenant/flt-test-rel")
	rel, err := c.ReadRelationfvRsTnDenyRuleFromTenant(tenant.DistinguishedName)
	if err != nil {
		fmt.Println("******Error *********")

		t.Error(err)
	}
	err = deleteTenant(c, tenant)
	if err != nil {

		t.Error(err)
	}
	if rel != nil {
		fmt.Println(rel.(*schema.Set))

	}

}

func TestDeleteAll(t *testing.T) {
	c := GetTestClient()
	tenList, err := c.ListTenant()

	if err != nil {
		t.Error(err)
	}

	for i := 0; i < len(tenList); i++ {
		err = c.DeleteTenant(tenList[i].Name)
		if err != nil {
			t.Error(err)
		}
	}
}
