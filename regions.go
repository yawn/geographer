package geographer

import (
	"context"
	"sort"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/pkg/errors"
)

type Regions []string

// Intersection returns only regions which are present in the available list (or all if none)
func (r Regions) Intersection(available ...string) []string {

	if len(available) == 0 {
		return r
	}

	var (
		out []string
		set = make(map[string]interface{}, len(available))
	)

	for _, region := range available {
		set[region] = struct{}{}
	}

	for _, region := range r {

		if _, ok := set[region]; ok {
			out = append(out, region)
		}

	}

	return out

}

// GetRegions returns configured regions in alphabetical order.
func GetRegions(ctx context.Context) (Regions, error) {

	cfg, err := config.LoadDefaultConfig(ctx)

	if err != nil {
		return nil, errors.Wrapf(err, "failed to load default config")
	}

	client := ec2.NewFromConfig(cfg)

	req := &ec2.DescribeRegionsInput{
		AllRegions: false,
	}

	res, err := client.DescribeRegions(ctx, req)

	if err != nil {
		return nil, errors.Wrapf(err, "failed to describe regions")
	}

	var regions []string

	for _, region := range res.Regions {
		regions = append(regions, *region.RegionName)
	}

	sort.Strings(regions)

	return regions, nil

}
