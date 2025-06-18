package server

import (
	"context"
	"fmt"
	"log"
	"time"

	"grpc-report-service/proto"

	"github.com/google/uuid"
)

// In-memory storage for reports
var reportStore = make(map[string]string) // map[reportID] = userID

type ReportServer struct {
	proto.UnimplementedReportServiceServer
}

func (s *ReportServer) GenerateReport(ctx context.Context, req *proto.ReportRequest) (*proto.ReportResponse, error) {
	userID := req.GetUserId()
	reportID := uuid.New().String()
	reportStore[reportID] = userID

	log.Printf("[%s] Generated report for user %s -> Report ID: %s\n", time.Now().Format(time.RFC3339), userID, reportID)

	return &proto.ReportResponse{
		ReportId: reportID,
		Error:    "",
	}, nil
}

func (s *ReportServer) HealthCheck(ctx context.Context, req *proto.HealthRequest) (*proto.HealthResponse, error) {
	log.Printf("[%s] Health check called\n", time.Now().Format(time.RFC3339))
	return &proto.HealthResponse{Status: "OK"}, nil
}
