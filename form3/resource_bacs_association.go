package form3

import (
	"encoding/json"
	"errors"
	"fmt"
	form3 "github.com/form3tech-oss/terraform-provider-form3/api"
	"github.com/form3tech-oss/terraform-provider-form3/client/associations"
	"github.com/form3tech-oss/terraform-provider-form3/models"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"gitlab.com/c0b/go-ordered-json"
	"log"
	"strings"
)

func resourceForm3BacsAssociation() *schema.Resource {
	return &schema.Resource{
		Create: resourceBacsAssociationCreate,
		Read:   resourceBacsAssociationRead,
		Delete: resourceBacsAssociationDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"association_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"organisation_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"service_user_number": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"account_number": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"sorting_code": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"account_type": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"bank_code": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"centre_number": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"input_key_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"input_certificate_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"input_tsu_number": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"output_key_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"output_certificate_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"output_tsu_number": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"messaging_key_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"messaging_certificate_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"messaging_tsu_number": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"test_file_submission": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
				Default:  false,
			},
			"service_user_numbers_config": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func resourceBacsAssociationCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	serviceUserNumber := d.Get("service_user_number").(string)
	log.Printf("[INFO] Creating Bacs association with service user number: %s", serviceUserNumber)

	association, err := createBacsNewAssociationFromResourceData(d)
	if err != nil {
		return fmt.Errorf("failed to create Bacs association: %s", form3.JsonErrorPrettyPrint(err))
	}

	createdAssociation, err := client.AssociationClient.Associations.PostBacs(associations.NewPostBacsParams().
		WithCreationRequest(&models.BacsAssociationCreation{
			Data: association,
		}))

	if err != nil {
		return fmt.Errorf("failed to create Bacs association: %s", form3.JsonErrorPrettyPrint(err))
	}

	d.SetId(createdAssociation.Payload.Data.ID.String())
	log.Printf("[INFO] Bacs association key: %s", d.Id())

	return resourceBacsAssociationRead(d, meta)
}

func resourceBacsAssociationRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	associationId, _ := GetUUIDOK(d, "association_id")

	if associationId == "" {
		associationId = strfmt.UUID(d.Id())
		log.Printf("[INFO] Importing bacs association with resource id: %s.", associationId)
	} else {
		log.Printf("[INFO] Reading bacs association with resource id: %s.", associationId)
	}

	bacsAssociation, err := client.AssociationClient.Associations.GetBacsID(associations.NewGetBacsIDParams().
		WithID(associationId))

	if err != nil {
		if !form3.IsJsonErrorStatusCode(err, 404) {
			return fmt.Errorf("couldn't find bacs association: %s", form3.JsonErrorPrettyPrint(err))
		}
		d.SetId("")
		return nil
	}

	_ = d.Set("association_id", bacsAssociation.Payload.Data.ID.String())
	_ = d.Set("organisation_id", bacsAssociation.Payload.Data.OrganisationID.String())
	_ = d.Set("sorting_code", bacsAssociation.Payload.Data.Attributes.SortingCode)
	_ = d.Set("account_type", bacsAssociation.Payload.Data.Attributes.AccountType)
	_ = d.Set("bank_code", bacsAssociation.Payload.Data.Attributes.BankCode)
	_ = d.Set("centre_number", bacsAssociation.Payload.Data.Attributes.CentreNumber)
	_ = d.Set("account_number", bacsAssociation.Payload.Data.Attributes.AccountNumber)
	_ = d.Set("test_file_submission", bacsAssociation.Payload.Data.Attributes.TestFileSubmission)

	if bacsAssociation.Payload.Data.Relationships != nil {
		if bacsAssociation.Payload.Data.Relationships.InputCertificate != nil && bacsAssociation.Payload.Data.Relationships.InputCertificate.Data != nil {
			_ = d.Set("service_user_number", bacsAssociation.Payload.Data.Attributes.ServiceUserNumber)
			_ = d.Set("input_key_id", bacsAssociation.Payload.Data.Relationships.InputCertificate.Data.KeyID)
			_ = d.Set("input_certificate_id", bacsAssociation.Payload.Data.Relationships.InputCertificate.Data.CertificateID)
			_ = d.Set("input_tsu_number", bacsAssociation.Payload.Data.Relationships.InputCertificate.Data.TsuNumber)
		}
		if bacsAssociation.Payload.Data.Relationships.OutputCertificate != nil && bacsAssociation.Payload.Data.Relationships.OutputCertificate.Data != nil {
			_ = d.Set("output_key_id", bacsAssociation.Payload.Data.Relationships.OutputCertificate.Data.KeyID)
			_ = d.Set("output_certificate_id", bacsAssociation.Payload.Data.Relationships.OutputCertificate.Data.CertificateID)
			_ = d.Set("output_tsu_number", bacsAssociation.Payload.Data.Relationships.OutputCertificate.Data.TsuNumber)
		}
		if bacsAssociation.Payload.Data.Relationships.MessagingCertificate != nil && bacsAssociation.Payload.Data.Relationships.MessagingCertificate.Data != nil {
			_ = d.Set("messaging_key_id", bacsAssociation.Payload.Data.Relationships.MessagingCertificate.Data.KeyID)
			_ = d.Set("messaging_certificate_id", bacsAssociation.Payload.Data.Relationships.MessagingCertificate.Data.CertificateID)
			_ = d.Set("messaging_tsu_number", bacsAssociation.Payload.Data.Relationships.MessagingCertificate.Data.TsuNumber)
		}
	}

	multiSunConfig := make([]string, 0)

	for idx, sunConfig := range bacsAssociation.Payload.Data.Attributes.ServiceUserNumbersConfig {
		unorderJson, jsonErr := json.Marshal(*sunConfig)
		if jsonErr != nil {
			return errors.New("Couldn't serialize SUN config element - idx: " + fmt.Sprint(idx) + ", error: " + jsonErr.Error())
		}
		orderedMap := ordered.NewOrderedMap()
		orderErr := orderedMap.UnmarshalJSON(unorderJson)
		if orderErr != nil {
			return errors.New("Couldn't sort SUN config - error: " + orderErr.Error())
		}
		orderedJson, jsonErr := orderedMap.MarshalJSON()
		if jsonErr != nil {
			return errors.New("Couldn't serialize sorted SUN config element - idx: " + fmt.Sprint(idx) + ", error: " + jsonErr.Error())
		}
		multiSunConfig = append(multiSunConfig, string(orderedJson))
	}

	if err := d.Set("service_user_numbers_config", multiSunConfig); err != nil {
		return err
	}

	return nil
}

func resourceBacsAssociationDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	bacsAssociation, err := client.AssociationClient.Associations.GetBacsID(associations.NewGetBacsIDParams().
		WithID(strfmt.UUID(d.Id())))

	if err != nil {
		return fmt.Errorf("error deleting Bacs association: %s", form3.JsonErrorPrettyPrint(err))
	}

	log.Printf("[INFO] Deleting Bacs association for id: %s service user number: %s",
		bacsAssociation.Payload.Data.ID,
		bacsAssociation.Payload.Data.Attributes.ServiceUserNumber)

	_, err = client.AssociationClient.Associations.DeleteBacsID(associations.NewDeleteBacsIDParams().
		WithID(bacsAssociation.Payload.Data.ID).
		WithVersion(*bacsAssociation.Payload.Data.Version))

	if err != nil {
		return fmt.Errorf("error deleting Bacs association: %s", form3.JsonErrorPrettyPrint(err))
	}

	return nil
}

func createBacsNewAssociationFromResourceData(d *schema.ResourceData) (*models.BacsNewAssociation, error) {

	association := models.BacsNewAssociation{
		Attributes:    &models.BacsAssociationAttributes{},
		Relationships: &models.BacsAssociationRelationships{},
	}

	association.Type = "associations"
	if attr, ok := GetUUIDOK(d, "association_id"); ok {
		association.ID = attr
	}

	if attr, ok := GetUUIDOK(d, "organisation_id"); ok {
		association.OrganisationID = attr
	}

	if attr, ok := d.GetOk("service_user_number"); ok {
		association.Attributes.ServiceUserNumber = attr.(string)
	}

	if attr, ok := d.GetOk("account_number"); ok {
		association.Attributes.AccountNumber = attr.(string)
	}

	if attr, ok := d.GetOk("sorting_code"); ok {
		association.Attributes.SortingCode = attr.(string)
	}

	if attr, ok := d.GetOkExists("account_type"); ok {
		accountType := int64(attr.(int))
		association.Attributes.AccountType = &accountType
	}

	if attr, ok := d.GetOk("bank_code"); ok {
		association.Attributes.BankCode = attr.(string)
	}

	if attr, ok := d.GetOk("centre_number"); ok {
		association.Attributes.CentreNumber = attr.(string)
	}

	if attr, ok := d.GetOk("test_file_submission"); ok {
		b := attr.(bool)

		association.Attributes.TestFileSubmission = &b
	}

	association.Relationships.InputCertificate = buildRelationship(d, "input")
	association.Relationships.OutputCertificate = buildRelationship(d, "output")
	association.Relationships.MessagingCertificate = buildRelationship(d, "messaging")

	if attr, ok := d.GetOk("service_user_numbers_config"); ok {
		sunArray, sunArrayTypeOk := attr.([]interface{})
		if !sunArrayTypeOk {
			return nil, errors.New("service_user_numbers_config must be an array of string")
		}

		var multiSunConfig []*models.BacsServiceUserNumber

		for idx, sunElement := range sunArray {
			stringValue, valueOk := sunElement.(string)
			if !valueOk {
				return nil, errors.New("service_user_numbers_config must be an array of string, invalid element at: " + fmt.Sprint(idx))
			}
			sunConfig := models.BacsServiceUserNumber{}
			jsonErr := json.Unmarshal([]byte(strings.ReplaceAll(stringValue, "\\\"", "\"")), &sunConfig)
			if jsonErr != nil {
				return nil, errors.New("invalid SUN config element at: " + fmt.Sprint(idx) + ", error: " + jsonErr.Error())
			}
			multiSunConfig = append(multiSunConfig, &sunConfig)
		}

		association.Attributes.ServiceUserNumbersConfig = multiSunConfig
	}

	return &association, nil
}

func buildRelationship(d *schema.ResourceData, relation string) *models.BacsAssociationCertificateRelationship {
	tsuNumber := ""
	if attr, ok := d.GetOk(relation + "_tsu_number"); ok {
		tsuNumber = attr.(string)
	}
	if keyId, ok := GetUUIDOK(d, relation+"_key_id"); ok {
		if certId, certOk := GetUUIDOK(d, relation+"_certificate_id"); certOk {
			return &models.BacsAssociationCertificateRelationship{
				Data: &models.BacsAssociationCertificateRelationshipData{KeyID: keyId, CertificateID: certId, TsuNumber: tsuNumber},
			}
		}
		return &models.BacsAssociationCertificateRelationship{
			Data: &models.BacsAssociationCertificateRelationshipData{KeyID: keyId, TsuNumber: tsuNumber},
		}
	}
	return nil
}
