package cronjob

import (
	"context"
	"log"
	"time"

	"github.com/robfig/cron/v3"
	"grpc-report-service/proto"
)

var predefinedUsers = []string{"user1", "user2", "user3"}

func StartCron(reportClient proto.ReportServiceClient) {
	c := cron.New()

	// Run every 10 seconds
	c.AddFunc("@every 10s", func() {
		log.Println("⏱️ Cron triggered: Generating reports...")

		for _, userID := range predefinedUsers {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
			defer cancel()

			res, err := reportClient.GenerateReport(ctx, &proto.ReportRequest{UserId: userID})
			if err != nil {
				log.Printf("❌ Failed to generate report for %s: %v\n", userID, err)
				continue
			}
			log.Printf("✅ Report generated for %s -> ReportID: %s\n", userID, res.ReportId)
		}
	})

	c.Start()
}
