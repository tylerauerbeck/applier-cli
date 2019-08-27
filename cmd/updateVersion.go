package cmd

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v2"
)

var updateVersionCmd = &cobra.Command{
	Use: "update-version",
	Short: "Update version of openshift-applier".
	Long: "Update the current version of the openshift-applier to a specific release",
	Run: func(cmd *cobra.Command, args []string) {
		version, _ := cmd.Flags().GetString("version")
		backup, _ := cmd.Flags().GetBool("backup")
		updateVer(version, backup)
	},
}

func updateVer(version string, backup bool) {
	yamlFile, err := ioutil.ReadFile("./requirements.yml")
	if err != nil {
		return
	}

	var reqs []map[string]interface{}
	err = yaml.Unmarshal(yamlFile, &reqs)
	if err != nil {
		log.Fatal(err)
	}

	var dirname string

	for _, e := range reqs {
		if strings.ToLower(e["name"].(string)) == "openshift-applier" {
			e["version"] = version
			dirname = e["name"].(string)
			break
		} else {
			log.Fatal("Cannot find openshift-applier in your requirements file.")
		}
	}

	out, err := yaml.Marshal(&reqs)

	err = ioutil.WriteFile("./requirements.yml", out, 0744)
	if err != nil {
		log.Fatal(err)
	}

	if !backup {
		os.Remove("./roles/" + dirname)
	} else {
		os.Rename("./roles/"dirname, "./roles/"+dirname+"-backup")
	}

	installGalaxyRequirements()
}

func init() {
	rootCmd.AddCommand(updateVersionCmd)

	updateVersionCmd.Flags().StringP("version", "v", "", "provide version of applier that you would like to use")
	updateVersionCmd.Flags().BoolP("backup", "b", "false", "backup existing applier role (defaults to false)")
}