package main

import (
	pb "apps/proto"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var client pb.StreamUploadClient

func main() {
	//connection to internal grpc server
	con, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("error in dial :", err)
	}
	client = pb.NewStreamUploadClient(con)

	mb := 1024 * 1024 * 2
	// uploadStreamFile("./1GB.mp4", mb)
	uploadStreamFile("./3GB.zip", mb)

}

func uploadStreamFile(path string, batchSize int) {
	t := time.Now()
	file, err := os.Open(path)
	if err != nil {
		log.Println("error in open file: ", err)
	}

	//settiing up buffer size
	buf := make([]byte, batchSize)
	batchNumber := 1
	stream, err := client.Upload(context.TODO())
	if err != nil {
		log.Println("error in open file: ", err)
	}
	for {
		num, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("error in read file: ", err)
			return
		}
		chunk := buf[:num]

		err = stream.Send(&pb.UploadReq{FilePath: path, Chunk: chunk})
		if err != nil {
			log.Println("error in send stream: ", err)
			return
		}
		log.Printf("Send - batch #%v - size - %v\n", batchNumber, len(chunk))

		batchNumber += 1
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Println("error in close stream: ", err)
		return
	}
	fmt.Println("total time took :", time.Since(t))
	log.Printf("send - %v bytes - %s\n ", res.GetFileSize(), res.GetMessage())
}
