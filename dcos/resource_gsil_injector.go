package dcos

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
)

// GsilInjector is the definition for a gsil_injector type in marathon
type GsilInjector struct {
	Username         string `json:"username,omitempty"`
	Password         string `json:"password,omitempty"`
	Location         string `json:"location,omitempty"`
	Token            []byte
	IntentManagerURL string
}

// GsilInjectorEntry is the definition for the gsil_injector type in marathon
type GsilInjectorEntry struct {
	KeyName  string `json:"key_name,omitempty"`
	Value    string `json:"value,omitempty"`
	Path     string `json:"path,omitempty"`
	Secret   string `json:"secret,omitempty"`
	Function string `json:"function,omitempty"`
}

func resourceGsilInjector() *schema.Resource {
	return &schema.Resource{
		Create: gsilInjectorCreate,
		Read:   gsilInjectorRead,
		Update: gsilInjectorUpdate,
		Delete: gsilInjectorDelete,

		Schema: map[string]*schema.Schema{
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"location": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"entries": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: false,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"entry": {
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: false,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key_name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"value": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"path": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"secret": &schema.Schema{
										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
			"gsil_injector_token": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gsil_intent_manager_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func updateForGsil(d *schema.ResourceData, m interface{}) {
	if v, injectorOk := d.GetOk("location"); injectorOk {
		processInjector(d, v, false)
	} else {
		log.Printf("[DEBUG] - location.# - %d", v)
	}
}

func processInjector(d *schema.ResourceData, v interface{}, deleting bool) *GsilInjector {
	injector := new(GsilInjector)
	t := v.(string)
	injector.Location = t

	if v, ok := d.GetOk("username"); ok {
		injector.Username = v.(string)
	}

	if v, ok := d.GetOk("password"); ok {
		injector.Password = v.(string)
	}

	token := Login(injector.Location, injector.Username, injector.Password)
	d.Set("gsil_injector_token", string(token))
	injector.Token = token

	respURL := IntentManagerURL(injector.Location, string(token))
	d.Set("gsil_intent_manager_url", string(respURL))
	injector.IntentManagerURL = string(respURL)

	if v, ok := d.GetOk("entries.#"); ok {
		log.Printf("[DEBUG] - entries.# - %d", v)

		for i := 0; i < v.(int); i++ {
			if v, ok := d.GetOk("entries." + strconv.Itoa(i) + ".entry.#"); ok {
				log.Printf("[DEBUG] - entries."+strconv.Itoa(i)+".entry.#- %d", v)
				processEntries(d, v, i, injector, deleting)
			}
		}
	}

	return injector
}

func processEntries(d *schema.ResourceData, v interface{}, i int, injector *GsilInjector, deleting bool) {
	for j := 0; j < v.(int); j++ {
		log.Printf("[DEBUG] - creating entry : %d", j)
		var entry InjectorEntry //entry := new(InjectorEntry)

		if v, ok := d.GetOk("entries." + strconv.Itoa(i) + ".entry." + strconv.Itoa(j) + ".key_name"); ok {
			entry.keyName = v.(string)
			log.Println("[DEBUG] - entry.0.key_name : " + entry.keyName)
		} else {
			log.Printf("[DEBUG] - NO entries." + strconv.Itoa(i) + ".entry." + strconv.Itoa(j) + ".key_name")
		}

		if v, ok := d.GetOk("entries." + strconv.Itoa(i) + ".entry." + strconv.Itoa(j) + ".value"); ok {
			entry.value = v.(string)
			log.Println("[DEBUG] - entry.0.value : " + entry.value)
		}

		if v, ok := d.GetOk("entries." + strconv.Itoa(i) + ".entry." + strconv.Itoa(j) + ".path"); ok {
			entry.path = v.(string)
			log.Println("[DEBUG] - entry.0.path : " + entry.path)
		}

		if v, ok := d.GetOk("entries." + strconv.Itoa(i) + ".entry." + strconv.Itoa(j) + ".secret"); ok {
			entry.secret = v.(bool)
			//log.Println("[DEBUG] - entry.0.secret : " + entry.secret)
		}

		//if v, ok := d.GetOk("gsil_injector.0.entries." + strconv.Itoa(i) + ".entry." + strconv.Itoa(j) + ".function"); ok {
		if deleting {
			entry.function = "DELETE"
		} else {
			entry.function = "PUT"
		}
		log.Println("[DEBUG] - entry.0.function : " + entry.function)

		token2 := RegisterKey(injector.Location, string(injector.Token), entry.keyName)
		d.Set("gsil_register_key_response", string(token2))
		log.Printf("\n[DEBUG]%v\n\n", string(token2))

		token3 := SetInjectorEntry(injector.Location, string(injector.Token), entry)
		fmt.Printf("\n%v\n\n", string(token3))
		//paramMap := d.Get(fmt.Sprintf("gsil_inject.0.entries.%d", i)).(map[string]interface{})
		//docker.AddParameter(paramMap["key"].(string), paramMap["value"].(string))
	}
}

func gsilInjectorCreate(d *schema.ResourceData, meta interface{}) error {
	/*
		config, err := genMarathonConf(d, meta)
		if err != nil {
			return err
		}

		client := config.Client

		c := make(chan deploymentEvent, 100)
		ready := make(chan bool)
		go readDeploymentEvents(&client, c, ready)
		select {
		case <-ready:
		case <-time.After(60 * time.Second):
			return errors.New("Timeout getting an EventListener")
		}
	*/

	//application := mapResourceToApplication(d)

	//application, err = client.CreateApplication(application)
	//if err != nil {
	//	log.Println("[ERROR] creating application", err)
	//	return err
	//}

	d.Partial(true)

	// gsil add
	updateForGsil(d, meta)

	/*
		d.SetId(application.ID)
		err = setSchemaFieldsForApp(application, d)
		if err != nil {
			log.Println("[ERROR] setSchemaFieldsForApp", err)
			return err
		}

		for _, deploymentID := range application.DeploymentIDs() {
			err = waitOnSuccessfulDeployment(c, deploymentID.DeploymentID, config.DefaultDeploymentTimeout)
			if err != nil {
				log.Println("[ERROR] waiting for application for deployment", deploymentID, err)
				return err
			}
		}
	*/
	d.Partial(false)

	return gsilInjectorRead(d, meta)
}

func gsilInjectorRead(d *schema.ResourceData, meta interface{}) error {

	return nil
}

func gsilInjectorUpdate(d *schema.ResourceData, meta interface{}) error {
	d.Partial(true)

	// gsil add
	updateForGsil(d, meta)

	d.Partial(false)

	return nil
}

func gsilInjectorDelete(d *schema.ResourceData, meta interface{}) error {

	if v, injectorOk := d.GetOk("location"); injectorOk {
		processInjector(d, v, true)
	}

	return nil
}
