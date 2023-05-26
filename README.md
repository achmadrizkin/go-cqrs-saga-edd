# go-cqrs-saga-edd

## This is using 2 Database (MySQL and MongoDB)
Product = localhost (MySQL)
Order = Cloud Server (MongoDB)

## Generate Proto
protoc --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. product.proto

## Running Evans
### evans -r repl


        func startGRPCServer(db *gorm.DB) {
            s := grpc.NewServer()

            pb.RegisterProductServiceServer(s, &server.Server{
                UnimplementedProductServiceServer: pb.UnimplementedProductServiceServer{},
                Db:                                db,
            })
            reflection.Register(s)

            // gRPC server
            lis, err := net.Listen("tcp", ":50051")
            if err != nil {
                log.Fatalf("Failed to listen: %v\n", err)
            }

            go func() {
                fmt.Println("Starting server...")
                if err := s.Serve(lis); err != nil {
                    log.Fatalf("Failed to serve: %v", err)
                }
            }()

            // Menunggu hingga dihentikan dengan Ctrl + C
            ch := make(chan os.Signal, 1)
            signal.Notify(ch, os.Interrupt)

            // Lakukan block hingga sinyal sudah didapatkan
            <-ch
            fmt.Println("Stopping the server..")
            s.Stop()
            fmt.Println("Stopping listener...")
            lis.Close()
            fmt.Println("End of Program")
        }

