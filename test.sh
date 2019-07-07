 grpc_cli call 127.0.0.1:9090 TaskService.GetAll ""
 echo "===================================================================================================="

 grpc_cli call localhost:9090 TaskService.Upsert "task:{name: 'Task1'}" -metadata "request-uuid: 1"
 grpc_cli call localhost:9090 TaskService.Upsert "task:{name: 'Task2'}" -metadata "request-uuid: 2"
 grpc_cli call 127.0.0.1:9090 TaskService.GetAll ""
 echo "===================================================================================================="

 grpc_cli call localhost:9090 TaskService.Delete "uuid: '823deeb6-8198-5e5d-8b7a-7326735d6846'"
 grpc_cli call 127.0.0.1:9090 TaskService.GetAll ""
 echo "===================================================================================================="

 grpc_cli call localhost:9090 TaskService.Upsert "task:{uuid: 'bce1ce33-c05b-56de-9f13-af91eed8fbe1', name: 'MODIFIED'}" -metadata "request-uuid: 3"
 grpc_cli call localhost:9090 TaskService.Upsert "task:{uuid: 'bce1ce33-c05b-56de-9f13-af91eed8fbe1', name: 'MODIFIED'}" -metadata "request-uuid: 3"
 grpc_cli call 127.0.0.1:9090 TaskService.GetAll ""
 echo "===================================================================================================="

 grpc_cli call localhost:9090 TaskService.Delete "uuid: 'bce1ce33-c05b-56de-9f13-af91eed8fbe1'"
 grpc_cli call 127.0.0.1:9090 TaskService.GetAll ""
