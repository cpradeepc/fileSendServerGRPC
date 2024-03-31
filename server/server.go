package main

import (
	pb "apps/proto"
	"io"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedStreamUploadServer
}

func main() {
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Println("tcp error :", err)

	}
	log.Println("listerner :", listener)
	srv := grpc.NewServer() //engine
	log.Println("server struct befor: ", server{})
	pb.RegisterStreamUploadServer(srv, &server{})
	reflection.Register(srv)
	log.Println("server struct after: ", server{})
	log.Println("srv: ", srv)
	err = srv.Serve(listener)
	if err != nil {
		log.Println("server error :", err)
	}
	log.Println("server is ready stage")

}

// implemet the Upload method for server type
func (s server) Upload(stream pb.StreamUpload_UploadServer) error {
	var fileBytes []byte
	var fileSize int64 = 0

	//var fileName string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			//fileName = req.GetFilePath()
			break
		}
		chunks := req.GetChunk()
		fileBytes = append(fileBytes, chunks...)
		fileSize += int64(len(chunks))
	}

	// f, err := os.Create("./1GB.mp4")
	f, err := os.Create("./3GB.zip")
	if err != nil {
		log.Println("error in file create :", err)
		return err
	}
	defer f.Close()
	_, err = f.Write(fileBytes)
	if err != nil {
		log.Println("error in file write :", err)
		return err
	}
	err = stream.SendAndClose(&pb.UploadResp{FileSize: fileSize,
		Message: "File written successfully"})
	return err

}
