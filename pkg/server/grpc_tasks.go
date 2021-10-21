package server

import (
	"context"
	"database/sql"
	"grpc_task/pkg/proto"

	_ "github.com/lib/pq"
	"google.golang.org/protobuf/types/known/emptypb"
)

//const connStr = "user=postgres password=7900 dbname=grpc_tasks sslmode=disable"

// opening taskDB
var tDb, _ = sql.Open("postgres", connStr)

type GRPCTasksSrever struct {
	proto.UnimplementedTaskListServer
}

// Create and Add are THE SAME
func (grpc GRPCTasksSrever) Create(c context.Context, tr *proto.TaskRequest) (*proto.TaskResponse, error) {
	_, err := tDb.Exec(`insert into tasks (name, status, priority, created_at, created_by, due_date)
	values ($1, $2, $3, $4, $5, $6);`, tr.T.GetName(), tr.T.GetStatus(), tr.T.GetPriority(),
		tr.T.GetCreatedAt(), tr.T.GetCreatedBy(), tr.T.GetDueDate())

	return &proto.TaskResponse{T: tr.T}, err
}

func (grpc GRPCTasksSrever) Add(c context.Context, tr *proto.TaskRequest) (*proto.TaskResponse, error) {
	_, err := tDb.Exec(`insert into tasks (name, status, priority, created_at, created_by, due_date)
	values ($1, $2, $3, $4, $5, $6);`, tr.T.GetName(), tr.T.GetStatus(), tr.T.GetPriority(),
		tr.T.GetCreatedAt(), tr.T.GetCreatedBy(), tr.T.GetDueDate())

	return &proto.TaskResponse{T: tr.T}, err
}

func (grpc GRPCTasksSrever) Update(c context.Context, tr *proto.TaskRequest) (*proto.TaskResponse, error) {
	_, err := tDb.Exec(`update tasks
	set name = $1, status = $2, priority = $3, due_date = $4
	where id = $5;`, tr.T.GetName(), tr.T.GetStatus(), tr.T.GetPriority(), tr.T.GetDueDate(), tr.T.GetId())

	return &proto.TaskResponse{T: tr.T}, err
}

func (grpc GRPCTasksSrever) Delete(c context.Context, tr *proto.TaskRequest) (*proto.TaskResponse, error) {
	_, err := tDb.Exec(`delete from tasks where id = $1`, tr.T.GetId())

	return &proto.TaskResponse{T: tr.T}, err
}

func (grpc GRPCTasksSrever) Get(c context.Context, trid *proto.TaskRequestID) (*proto.TaskResponse, error) {
	row := tDb.QueryRow(`select * from tasks where id = $1;`, trid.GetId())

	var result proto.Task

	row.Scan(&result.Id, &result.Name, &result.Status, &result.Priority, &result.CreatedAt,
		&result.CreatedBy, &result.DueDate)

	return &proto.TaskResponse{T: &result}, nil
}

func (grpc GRPCTasksSrever) GetAll(c context.Context, emp *emptypb.Empty) (*proto.TaskResponseSlice, error) {
	rows, err := tDb.Query(`select * from tasks;`)

	var result = make([]*proto.Task, 0)
	for rows.Next() {
		var tempTask proto.Task
		rows.Scan(&tempTask.Id, &tempTask.Name, &tempTask.Status, &tempTask.Priority,
			&tempTask.CreatedAt, &tempTask.CreatedBy, &tempTask.CreatedBy)
		result = append(result, &tempTask)
	}

	return &proto.TaskResponseSlice{T: result}, err
}
