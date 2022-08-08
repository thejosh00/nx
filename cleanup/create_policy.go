package cleanup

import (
	"fmt"
	"org/sonatype/nx/api"
	"org/sonatype/nx/util"
)

type CleanupCreatePolicyCommand struct {
	Verbose    bool   `short:"v" long:"verbose" description:"log verbose debug information"`
	Format     string `short:"f" long:"format" required:"true" description:"format that cleanup policy applies to"`
	Age        int    `short:"a" long:"componentAge" description:"remove components that were publish over this many days ago"`
	Usage      int    `short:"u" long:"componentUsage" description:"remove components that haven't been downloaded in this many days"`
	Positional struct {
		Name string `positional-arg-name:"name"`
	} `positional-args:"yes"`
}

func (cmd *CleanupCreatePolicyCommand) Execute(args []string) error {
	if !cmd.Verbose {
		util.StopLogging()
	}

	name := "cleanup"
	if cmd.Positional.Name != "" {
		name = cmd.Positional.Name
	}

	fmt.Println("NAME", name)

	err := createPolicy(name, cmd.Format, cmd.Age, cmd.Usage)
	if err != nil {
		return err
	}
	fmt.Println("Created cleanup policy", name)
	return nil
}

type payload struct {
	Name                    string `json:"name"`
	Notes                   string `json:"notes"`
	Format                  string `json:"format"`
	CriteriaLastBlobUpdated int    `json:"criteriaLastBlobUpdated,omitempty"`
	CriteriaLastDownloaded  int    `json:"criteriaLastDownloaded,omitempty"`
	CriteriaReleaseType     string `json:"criteriaReleaseType,omitempty"`
	CriteriaAssetRegex      string `json:"criteriaAssetRegex,omitempty"`
}

func createPolicy(name string, format string, age int, usage int) error {
	payload := payload{
		Name:                    name,
		Notes:                   "",
		Format:                  format,
		CriteriaLastBlobUpdated: age,
		CriteriaLastDownloaded:  usage,
	}

	return api.Post("internal/cleanup-policies", payload, 200)
}
