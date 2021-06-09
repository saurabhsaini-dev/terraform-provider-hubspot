package hubspot

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"strings"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-provider-hubspot/client"
)

func validateEmail(v interface{}, k string) (ws []string, es []error) {
	var errs []error
	var warns []string
	value := v.(string)

	var emailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

	if !(emailRegex.MatchString(value)) {
		errs = append(errs, fmt.Errorf("Expected EmailId is not valid  %s", k))
		return warns, errs
	}
	return
}

func resourceUser() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceUserCreate,
		ReadContext:   resourceUserRead,
		UpdateContext: resourceUserUpdate,
		DeleteContext: resourceUserDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"email": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateEmail,
			},
			"roleid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceUserCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	apiClient := m.(*client.Client)
	user := client.User{
		Email:  d.Get("email").(string),
		RoleId: d.Get("roleid").(string),
	}

	err := apiClient.CreateUser(&user)
	if err != nil {
		log.Println("[ERROR]: ", err)
		return diag.FromErr(err)
	}
	d.SetId(user.Email)
	resourceUserRead(ctx, d, m)
	return diags
}

func resourceUserRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	apiClient := m.(*client.Client)
	userId := d.Id()
	user, err := apiClient.GetUser(userId)
	if err != nil {
		log.Println("[ERROR]: ", err)
		if strings.Contains(err.Error(), "not found") {
			d.SetId("")
		} else {
			return diag.FromErr(err)
		}
	}

	d.Set("email", user.Email)
	d.Set("roleid", user.RoleId)

	return diags
}

func resourceUserUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	apiClient := m.(*client.Client)
	var diags diag.Diagnostics
	if d.HasChange("email") {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "User not allowed to change email",
			Detail:   "User not allowed to change email",
		})

		return diags
	}

	if d.HasChange("roleid") {
		user := client.User{
			Email:  d.Get("email").(string),
			RoleId: d.Get("roleid").(string),
		}

		err := apiClient.UpdateUser(&user)
		if err != nil {
			log.Printf("[Error] Error updating user :%s", err)
			return diag.FromErr(err)
		}
	}
	return resourceUserRead(ctx, d, m)
}

func resourceUserDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	apiClient := m.(*client.Client)
	userId := d.Id()
	err := apiClient.DeleteUser(userId)
	if err != nil {
		log.Printf("[Error] Error deleting user :%s", err)
		return diag.FromErr(err)
	}
	d.SetId("")
	return diags
}
