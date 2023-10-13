package main

func getAWSCost(startDate, endDate time.Time) (*costexplorer.GetCostAndUsageOutput, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"), // Replace with your desired AWS region
	})

	if err != nil {
		return nil, err
	}

	svc := costexplorer.New(sess)

	input := &costexplorer.GetCostAndUsageInput{
		TimePeriod: &costexplorer.DateInterval{
			Start: aws.String(startDate.Format("2006-01-02")),
			End:   aws.String(endDate.Format("2006-01-02")),
		},
		Granularity: aws.String("MONTHLY"),
		Metrics: []*string{
			aws.String("UNBLENDED_COST"),
			aws.String("USAGE_QUANTITY"),
		},
		GroupBy: []*costexplorer.GroupDefinition{
			{
				Type: aws.String("DIMENSION"),
				Key:  aws.String("SERVICE"),
			},
			{
				Type: aws.String("DIMENSION"),
				Key:  aws.String("USAGE_TYPE"),
			},
		},
	}

	resp, err := svc.GetCostAndUsage(input)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func main() {
	startDate := time.Now().AddDate(0, 0, -30)
	endDate := time.Now()

	cost, err := getAWSCost(startDate, endDate)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Print(cost)
}
