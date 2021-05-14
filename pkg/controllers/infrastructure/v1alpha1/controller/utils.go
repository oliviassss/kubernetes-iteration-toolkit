/*
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

const (
	TagKeyNameForAWSResources = "kit.k8s.amazonaws.com/cluster-name"
	vpcCIDR                   = "10.0.0.0/16" // TODO hardcoded for now, make defaults
)

var (
	// TODO hardcoded for now, make defaults
	privateSubnetCIDRs = []string{"10.0.1.0/24", "10.0.2.0/24", "10.0.3.0/24"}
	publicSubnetCIDRs  = []string{"10.0.11.0/24", "10.0.12.0/24", "10.0.13.0/24"}

	resourceReconcileFailed    = reconcile.Result{RequeueAfter: 5 * time.Second}
	resourceReconcileSucceeded = reconcile.Result{RequeueAfter: 30 * time.Second}
)

func generateEC2Tags(svcName, clusterName string) []*ec2.TagSpecification {
	return []*ec2.TagSpecification{
		&ec2.TagSpecification{
			ResourceType: aws.String(svcName),
			Tags: []*ec2.Tag{
				&ec2.Tag{
					Key:   aws.String(TagKeyNameForAWSResources),
					Value: aws.String(clusterName),
				},
				&ec2.Tag{
					Key:   aws.String("Name"),
					Value: aws.String(fmt.Sprintf("%s-%s", clusterName, svcName)),
				},
			},
		},
	}
}
