// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
)

func ListCloudformationStackSet(client *Client) ([]Resource, error) {
	req := client.Cloudformationconn.ListStackSetsRequest(&cloudformation.ListStackSetsInput{})

	var result []Resource

	p := cloudformation.NewListStackSetsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Summaries {

			result = append(result, Resource{
				Type:      "aws_cloudformation_stack_set",
				ID:        *r.StackSetName,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
