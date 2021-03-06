package autospotting

import (
	"errors"
	"os"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func TestMain(m *testing.M) {
	disableLogging()
	os.Exit(m.Run())
}

func Test_fetch(t *testing.T) {
	tests := []struct {
		name             string
		config           *spotPrices
		product          string
		duration         time.Duration
		availabilityZone *string
		instanceTypes    []*string
		data             []*ec2.SpotPrice
		err              error
	}{
		{
			name: "error",
			config: &spotPrices{
				data: []*ec2.SpotPrice{},
				conn: connections{
					ec2: mockEC2{
						dspho: &ec2.DescribeSpotPriceHistoryOutput{
							SpotPriceHistory: []*ec2.SpotPrice{},
						},
						dspherr: errors.New("error"),
					},
				},
			},
			data: []*ec2.SpotPrice{},
			err:  errors.New("error"),
		},
		{
			name: "ok",
			config: &spotPrices{
				data: []*ec2.SpotPrice{},
				conn: connections{
					ec2: mockEC2{
						dspho: &ec2.DescribeSpotPriceHistoryOutput{
							SpotPriceHistory: []*ec2.SpotPrice{
								{SpotPrice: aws.String("1")},
							},
						},
					},
				},
			},
			data: []*ec2.SpotPrice{
				{SpotPrice: aws.String("1")},
			},
			err: errors.New(""),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.config.fetch(tc.product, tc.duration, tc.availabilityZone, tc.instanceTypes)
			if len(tc.data) != len(tc.config.data) {
				t.Errorf("Price data actual: %v\nexpected: %v", tc.config.data, tc.data)
			}
			if len(tc.data) > 0 {
				str1 := *tc.data[0].SpotPrice
				str2 := *tc.config.data[0].SpotPrice
				if str1 != str2 {
					t.Errorf("Price actual: %s, expected: %s", str2, str1)
				}
			}
			if err != nil && err.Error() != tc.err.Error() {
				t.Errorf("error expected: %s, actual: %s", tc.err.Error(), err.Error())
			}
		})
	}
}
