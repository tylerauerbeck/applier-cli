package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	dash "github.com/tylerauerbeck/dash/dash"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Runs the current OpenShift-Applier inventory.",
	Long: `Runs the current OpenShift-Applier inventory. Default is 
to run using local Ansible, but can also run OpenShift-Applier 
in a Docker container.`,
	Run: func(cmd *cobra.Command, args []string) {

		envFile := dash.GetInventoryConfig(envDir)
		inventory := dash.GetInventoryContent(envDir, envFile)
		version := inventory.Version

		if version == "" {
			log.Println("Unable to identify version. Defaulting to 2.X")
			version = "2.0"
			//TODO: Should we default to 2.X? Do we leave this configurable? What's the safer/saner default?
		}

		switch string(version[0]) {
		case "3":
			log.Println("Running Dash..")
			//TODO: RunV3()
			dash.ExecV3(inventory)
		case "2":
			log.Println("Running Openshift-Applier..")
			//TODO: Do something here to find Inventory that contains openshift_cluster_content
		default:
			log.Println("Unable to determine version. Exiting.")
			os.Exit(1)
		}

		//var data, _ = ioutil.ReadFile("/Users/tylerauerbeck/workspace/test-cli/dash.yml")
		//t := dash.Inventory{}

		//err := yaml.Unmarshal([]byte(data), &t)
		//if err != nil {
		//	log.Fatalf("error: %v", err)
		//}
		//fmt.Printf("--- t:\n%v\n\n", t)

		//fmt.Println(t.Resources)
		//fmt.Println("hello")

		//d, err := yaml.Marshal(&t)
		//if err != nil {
		//	log.Fatalf("error: %v", err)
		//}
		//fmt.Printf("--- t dump:\n%s\n\n", string(d))

	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
