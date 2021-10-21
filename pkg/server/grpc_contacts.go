package server

import (
	"context"
	"database/sql"
	"grpc_task/pkg/proto"

	_ "github.com/lib/pq"
	"google.golang.org/protobuf/types/known/emptypb"
)

const connStr = "user=postgres password=7900 dbname=grpc_tasks sslmode=disable"

var cDb, _ = sql.Open("postgres", connStr)

type GRPCContactSrever struct {
	proto.UnimplementedContacListServer
}

// Create and add are THE SAME
func (grpc GRPCContactSrever) Create(c context.Context, cr *proto.ContactRequest) (*proto.ContactResponse, error) {
	_, err := cDb.Exec(`insert into contacts (first_name, last_name, phone, email)
	values ( $1, $2, $3, $4);`, cr.C.GetFirstName(), cr.C.GetLastName(), cr.C.GetPhone(), cr.C.GetEmail())

	return &proto.ContactResponse{C: cr.C}, err
}

func (grpc GRPCContactSrever) Add(c context.Context, cr *proto.ContactRequest) (*proto.ContactResponse, error) {
	_, err := cDb.Exec(`insert into contacts (first_name, last_name, phone, email)
	values ( $1, $2, $3, $4);`, cr.C.GetFirstName(), cr.C.GetLastName(), cr.C.GetPhone(), cr.C.GetEmail())

	return &proto.ContactResponse{C: cr.C}, err
}

func (grpc GRPCContactSrever) Update(c context.Context, cr *proto.ContactRequest) (*proto.ContactResponse, error) {
	_, err := cDb.Exec(`update contacts 
	set first_name = $1, last_name = $2, phone = $3, email = $4
	where id = $5;`, cr.C.GetFirstName(), cr.C.GetLastName(), cr.C.GetPhone(), cr.C.GetEmail(), cr.C.GetId())

	return &proto.ContactResponse{C: cr.C}, err
}

func (grpc GRPCContactSrever) Delete(c context.Context, cr *proto.ContactRequest) (*proto.ContactResponse, error) {
	_, err := cDb.Exec(`delete from contacts where id = $1;`, cr.C.GetId())

	return &proto.ContactResponse{C: cr.C}, err
}

func (grpc GRPCContactSrever) Get(c context.Context, crid *proto.ContactRequestID) (*proto.ContactResponse, error) {
	row := cDb.QueryRow(`select * from contacts where id = $1;`, crid.GetId())

	var tempCont proto.Contact
	row.Scan(&tempCont.Id, &tempCont.FirstName, &tempCont.LastName, &tempCont.Phone, &tempCont.Email)

	return &proto.ContactResponse{C: &tempCont}, nil
}

func (grpc GRPCContactSrever) GetAll(c context.Context, emp *emptypb.Empty) (*proto.ContactResponseSlice, error) {
	rows, err := cDb.Query(`select * from contacts;`)

	var conts = make([]*proto.Contact, 0)

	for rows.Next() {
		var tempCont proto.Contact
		rows.Scan(&tempCont.Id, &tempCont.FirstName, &tempCont.LastName, &tempCont.Phone, &tempCont.Email)
		conts = append(conts, &tempCont)
	}
	return &proto.ContactResponseSlice{C: conts}, err
}
