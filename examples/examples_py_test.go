//go:build python || all
// +build python all

package examples

import (
        "os"
	"path/filepath"
	"testing"
	"crypto/tls"
    	"net/http"

        "github.com/stretchr/testify/assert"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
	"github.com/AviatrixSystems/terraform-provider-aviatrix/v2/goaviatrix"
)

//example: https://github.com/pulumi/pulumi-google-native/blob/master/examples/examples_py_test.go

/*
The following environment variable need to be set in order to execute the tests:
AVIATRIX_CONTROLLER_IP: The Aviatrix controller's IP address
AVIATRIX_USERNAME: The Aviatrix controller's username
AVIATRIX_PASSWORD: The Aviatrix controller's password
AZURE_ACCT_NAME: Name of the Azure account on the controller
AZURE_REGION: Region of Azure workloads
*/

func getPythonBaseOptions(t *testing.T) integration.ProgramTestOptions {
	base := getBaseOptions()
	basePython := base.With(integration.ProgramTestOptions{
		Dependencies: []string{
			filepath.Join("..", "sdk", "python", "bin"),
		},
	})

	return basePython
}

//Test function name must start with Test
func TestAviatrixSpoke(t *testing.T) {
	cwd, err := os.Getwd()
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	options := getPythonBaseOptions(t).With(integration.ProgramTestOptions{
		Dir:         filepath.Join(cwd, "azure", "spoke-py", "step1"),
		SkipRefresh: true,
		ExtraRuntimeValidation: validateAviatrixSpoke,
	})
	integration.ProgramTest(t, &options)
}

func validateAviatrixSpoke(t *testing.T, info integration.RuntimeValidationStackInfo) {
        //Invoke the goaviatrix SDK to check wether resources have actually been created
        //Log in to the controller
        tr := &http.Transport{
            TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
        }

        client, err := goaviatrix.NewClient(os.Getenv("AVIATRIX_USERNAME"), os.Getenv("AVIATRIX_PASSWORD"), os.Getenv("AVIATRIX_CONTROLLER_IP"), &http.Client{Transport: tr}, &goaviatrix.IgnoreTagsConfig{})
        if err != nil {
                t.Fatal(err)
        }

	foundGateway := &goaviatrix.Gateway{
			GwName:      "pulumi-spk1",
			AccountName: os.Getenv("AZURE_ACCT_NAME"),
		}

	foundGateway2, err := client.GetGateway(foundGateway)
	if err != nil {
		t.Fatal(err)
	}
	if foundGateway2.GwName != "pulumi-spk1" {
		t.Fatal("Spoke gateway not found")
	}
}

func TestAviatrixTransit(t *testing.T) {
        cwd, err := os.Getwd()
        if !assert.NoError(t, err) {
                t.FailNow()
        }
        options := getPythonBaseOptions(t).With(integration.ProgramTestOptions{
                Dir:         filepath.Join(cwd, "azure", "transit-py", "step1"),
                SkipRefresh: true,
                ExtraRuntimeValidation: validateAviatrixTransit,
        })
        integration.ProgramTest(t, &options)
}

func validateAviatrixTransit(t *testing.T, info integration.RuntimeValidationStackInfo) {
        //Invoke the goaviatrix SDK to check wether resources have actually been created
	tr := &http.Transport{
            TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
        }

        client, err := goaviatrix.NewClient(os.Getenv("AVIATRIX_USERNAME"), os.Getenv("AVIATRIX_PASSWORD"), os.Getenv("AVIATRIX_CONTROLLER_IP"), &http.Client{Transport: tr}, &goaviatrix.IgnoreTagsConfig{})
        if err != nil {
                t.Fatal(err)
        }

        foundGateway := &goaviatrix.Gateway{
                        GwName:      "pulumi-tr2",
                        AccountName: os.Getenv("AZURE_ACCT_NAME"),
                }

        foundGateway2, err := client.GetGateway(foundGateway)
        if err != nil {
                t.Fatal(err)
        }
        if foundGateway2.GwName != "pulumi-tr2" {
                t.Fatal("Transit gateway not found")
        }
}

func TestAviatrixSpokeTransitAttach(t *testing.T) {
        cwd, err := os.Getwd()
        if !assert.NoError(t, err) {
                t.FailNow()
        }
        options := getPythonBaseOptions(t).With(integration.ProgramTestOptions{
                Dir:         filepath.Join(cwd, "azure", "spoke-transit-attach-py", "step1"),
                SkipRefresh: true,
                ExtraRuntimeValidation: validateAviatrixSpokeTransitAttach,
        })
        integration.ProgramTest(t, &options)
}

func validateAviatrixSpokeTransitAttach(t *testing.T, info integration.RuntimeValidationStackInfo) {
        //Invoke the goaviatrix SDK to check wether resources have actually been created
	tr := &http.Transport{
            TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
        }

        client, err := goaviatrix.NewClient(os.Getenv("AVIATRIX_USERNAME"), os.Getenv("AVIATRIX_PASSWORD"), os.Getenv("AVIATRIX_CONTROLLER_IP"), &http.Client{Transport: tr}, &goaviatrix.IgnoreTagsConfig{})
        if err != nil {
                t.Fatal(err)
        }
	
	foundSpokeTransitAttachment := &goaviatrix.SpokeTransitAttachment{
			SpokeGwName:   "pulumi-spk3",
			TransitGwName: "pulumi-tr3",
	}
	
	foundSpokeTransitAttachment2, err := client.GetSpokeTransitAttachment(foundSpokeTransitAttachment)
	if err != nil {
		t.Fatal(err)
	}
	if foundSpokeTransitAttachment2.SpokeGwName+"~"+foundSpokeTransitAttachment2.TransitGwName != "pulumi-spk3~pulumi-tr3" {
		t.Fatal("Spoke transit attachment not found")
	}
}


func TestAviatrixVnet(t *testing.T) {
        cwd, err := os.Getwd()
        if !assert.NoError(t, err) {
                t.FailNow()
        }
        options := getPythonBaseOptions(t).With(integration.ProgramTestOptions{
                Dir:         filepath.Join(cwd, "azure", "vnet-py", "step1"),
                SkipRefresh: true,
                ExtraRuntimeValidation: validateAviatrixVnet,
        })
        integration.ProgramTest(t, &options)
}

func validateAviatrixVnet(t *testing.T, info integration.RuntimeValidationStackInfo) {
        //Invoke the goaviatrix SDK to check wether resources have actually been created
        //Log in to the controller
        tr := &http.Transport{
            TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
        }

        client, err := goaviatrix.NewClient(os.Getenv("AVIATRIX_USERNAME"), os.Getenv("AVIATRIX_PASSWORD"), os.Getenv("AVIATRIX_CONTROLLER_IP"), &http.Client{Transport: tr}, &goaviatrix.IgnoreTagsConfig{})
        if err != nil {
                t.Fatal(err)
        }

        //Look for VPC, check name
        foundVpc := &goaviatrix.Vpc{
                Name: "pulumi-tr1",
        }

        foundVpc2, err := client.GetVpc(foundVpc)
        if err != nil {
                t.Fatal(err)
        }
        if foundVpc2.Name != "pulumi-tr1" {
                t.Fatal("VPC not found")
        }
}
