package dcos

import (
	"errors"
	"os"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceGsilPokemon() *schema.Resource {
	return &schema.Resource{
		Create: gsilPokemonCreate,
		Read:   gsilPokemonRead,
		Update: gsilPokemonUpdate,
		Delete: gsilPokemonDelete,

		Schema: map[string]*schema.Schema{
			// New GSIL custom values used for installing gsil specific services
			"gsil_filename": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gsil_directory": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gsil_which_poke": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"gsil_pokemon_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gsil_register_key_response": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func gsilPokemonCreate(d *schema.ResourceData, meta interface{}) error {
	d.Partial(true)
	// gsil add
	updatePokemonForGsil(d, meta)
	d.Partial(false)
	return gsilInjectorRead(d, meta)
}

func gsilPokemonRead(d *schema.ResourceData, meta interface{}) error {
	err := d.Set("gsil_filename", "done.zip")
	if err != nil {
		return errors.New("Failed to set gsil_filename: " + err.Error())
	}
	d.SetPartial("gsil_filename")

	err = d.Set("gsil_directory", "foo")
	if err != nil {
		return errors.New("Failed to set gsil_directory: " + err.Error())
	}
	d.SetPartial("gsil_directory")

	pokeName := Read("foo"+"/"+"configuration.properties", "pokemon")
	var pokeNumber int

	if pokeName != "" {
		pokeNumber = FindPokeNumber(pokeName)
	} else {
		pokeNumber = 0
	}

	err = d.Set("gsil_which_poke", pokeNumber)
	if err != nil {
		return errors.New("Failed to set gsil_which_poke: " + err.Error())
	}
	d.SetPartial("gsil_which_poke")

	err = d.Set("gsil_injector_location", "")
	if err != nil {
		return errors.New("Faild to set gsil_injector_location: " + err.Error())
	}
	d.SetPartial("gsil_injector_location")

	return nil
}

func gsilPokemonUpdate(d *schema.ResourceData, meta interface{}) error {
	d.Partial(true)
	// gsil add
	updatePokemonForGsil(d, meta)
	d.Partial(false)
	return gsilInjectorRead(d, meta)
}

func gsilPokemonDelete(d *schema.ResourceData, meta interface{}) error {

	directory := d.Get("gsil_directory").(string)
	err := os.RemoveAll(directory)

	if err != nil {
		return err
	}

	return nil
}

func updatePokemonForGsil(d *schema.ResourceData, m interface{}) {
	directory := d.Get("gsil_directory").(string)
	whichPoke := d.Get("gsil_which_poke").(int)
	filename := d.Get("gsil_filename").(string)

	Unzip(filename, directory)

	var pokemon string
	if whichPoke != 0 {
		pokemon = FindPoke(whichPoke)
	} else {
		pokemon = FindPoke(16) // find any random poke just so it doesn't fail.
	}

	Write(directory+"/"+"configuration.properties", "pokemon", pokemon)

	var configValue string
	configValue = Read(directory+"/"+"configuration.properties", "pokemon")
	d.Set("gsil_pokemon_name", configValue)

	/*
		if v, injectorOk := d.GetOk("gsil_injector.0.location"); injectorOk {
			processInjector(d, v, false)
		}
	*/
	//return nil
}
