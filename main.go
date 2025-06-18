package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"grpc-report-service/proto"
	"grpc-report-service/server"
	"grpc-report-service/cronjob"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("❌ Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	reportService := &server.ReportServer{}
	proto.RegisterReportServiceServer(grpcServer, reportService)

	// Enable reflection for debugging (optional)
	reflection.Register(grpcServer)

	go func() {
		log.Println("🚀 Starting gRPC server on :50051")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("❌ Failed to serve: %v", err)
		}
	}()

	// Start cron client to call server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("❌ Failed to connect to server: %v", err)
	}
	defer conn.Close()

	client := proto.NewReportServiceClient(conn)
	cronjob.StartCron(client)

	select {} // keep the main goroutine alive
}
