package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	compute "google.golang.org/api/compute/v1"
)

func getComputeService() (*compute.Service, error) {

	raw, err := ioutil.ReadFile("tigerworks-kube-9492d6708b81.json")
	if err != nil {
		return nil, err
	}

	conf, err := google.JWTConfigFromJSON(raw, compute.ComputeScope)
	if err != nil {
		return nil, err
	}

	client := conf.Client(oauth2.NoContext)

	computeService, err := compute.New(client)
	if err != nil {
		return nil, err
	}

	return computeService, nil
}

func instanceList(ctx context.Context, computeService *compute.Service, project string, zone string) {

	req := computeService.Instances.List(project, zone)

	if err := req.Pages(ctx, func(page *compute.InstanceList) error {
		for _, instance := range page.Items {
			fmt.Printf("%v\n", instance.Name)
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
}

func instanceCreate(ctx context.Context, computeService *compute.Service, project, zone, name, imageName string) {

	machineType := fmt.Sprintf("projects/%v/zones/%v/machineTypes/n1-standard-1", project, zone)
	srcImage := fmt.Sprintf("projects/%v/global/images/%v", project, imageName)

	requestBody := &compute.Instance{
		Name:        name,
		MachineType: machineType,
		NetworkInterfaces: []*compute.NetworkInterface{
			{
				Network: fmt.Sprintf("projects/%v/global/networks/%v", project, "default"),
				AccessConfigs: []*compute.AccessConfig{
					{
						Type: "ONE_TO_ONE_NAT",
						Name: "External NAT",
					},
				},
			},
		},
		Disks: []*compute.AttachedDisk{
			{
				InitializeParams: &compute.AttachedDiskInitializeParams{
					SourceImage: srcImage,
				},
				Mode:       "READ_WRITE",
				Boot:       true,
				AutoDelete: true,
			},
		},
	}

	resp, err := computeService.Instances.Insert(project, zone, requestBody).Context(ctx).Do()

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(resp)

}

func main() {

	ctx := context.Background()
	computeService, err := getComputeService()

	if err != nil {
		log.Fatal(err)
		return
	}

	project := "tigerworks-kube"
	zone := "us-central1-f"
	name := "onboarding-gce"
	imageName := "appscode-containeros"

	fmt.Println("creating instance...")
	instanceCreate(ctx, computeService, project, zone, name, imageName)
	fmt.Println("instance created")

	fmt.Println("instance list:")
	instanceList(ctx, computeService, project, zone)
}
