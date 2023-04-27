.
├── api
│   ├── myapi.proto
│   ├── myapi.pb.go
│   ├── myapi.pb.gw.go
│   └── handlers.go
├── cmd
│   └── server
│       └── main.go
├── internal
│   ├── config
│   │   └── config.go
│   ├── service
│   │   ├── service.go
│   │   ├── service_test.go
│   │   ├── repository.go
│   │   └── repository_test.go
│   ├── middleware
│   │   └── middleware.go
│   └── model
│       ├── model.go
│       ├── model_test.go
│       └── db.go
├── Makefile
├── go.mod
├── go.sum
└── README.md

api: contiene el archivo proto para definir la API, así como los archivos generados por el compilador protobuf que definen los mensajes y servicios de la API, y el controlador para manejar las solicitudes entrantes. El archivo myapi.pb.gw.go se genera con protoc y contiene el proxy de gRPC a JSON, que se utiliza para implementar un enrutador RESTful. El archivo handlers.go contiene los controladores que manejan las solicitudes HTTP entrantes y los envían al servicio interno.

cmd: contiene la implementación del programa principal, que inicia el servidor de la API.

internal: contiene la implementación interna del servidor. El paquete config contiene la configuración del servidor, el paquete service implementa la lógica de negocio de la aplicación y se comunica con la base de datos, el paquete middleware contiene funciones de middleware para validar la autenticación y la autorización, y el paquete model contiene las definiciones de la base de datos y las funciones para interactuar con ella.

Makefile: contiene las reglas de construcción para compilar y ejecutar la aplicación.

go.mod y go.sum: archivos que administran las dependencias del proyecto.

README.md: archivo de documentación que describe el proyecto y cómo utilizarlo.