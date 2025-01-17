// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package snippets

// [START compute_get_instance_serial_port]
import (
	"context"
	"fmt"
	"io"

	compute "cloud.google.com/go/compute/apiv1"
	computepb "cloud.google.com/go/compute/apiv1/computepb"
)

// getInstanceSerialPort prints the serial port output for an instance.
func getInstanceSerialPort(w io.Writer, projectID, zone, instanceName string) error {
	// projectID := "your_project_id"
	// zone := "europe-central2-b"
	// instanceName := "your_instance_name"

	ctx := context.Background()
	instancesClient, err := compute.NewInstancesRESTClient(ctx)
	if err != nil {
		return fmt.Errorf("NewInstancesRESTClient: %w", err)
	}
	defer instancesClient.Close()

	req := &computepb.GetSerialPortOutputInstanceRequest{
		Project:  projectID,
		Zone:     zone,
		Instance: instanceName,
	}

	output, err := instancesClient.GetSerialPortOutput(ctx, req)
	if err != nil {
		return fmt.Errorf("unable to get serial port output: %w", err)
	}

	fmt.Fprintf(w, "Output from instance serial port: %s", output.GetContents())

	return nil
}

// [END compute_get_instance_serial_port]
