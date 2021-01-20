# geographer

Minimal API to determine the effective regions in which a given AWS service is available. The effective regions are determined by general service availability and the regions that are enabled for a given account.

## Example

Determine the effectively available regions for `EC2`:

```
regions, err := geographer.GetRegions(context.TODO())

if err != nil {
  panic(err)
}

ec2 := geographer.Services["ec2"].Intersection(regions...)
```
