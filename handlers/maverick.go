//package controllers
//
//import (
//	"github.com/melisource/fury_people-manager/internal/services"
//	"net/http"
//	"testing"
//)
//
//func TestPersonController_CreatePerson(t *testing.T) {
//	type fields struct {
//		service *services.PersonService
//	}
//	type args struct {
//		w http.ResponseWriter
//		r *http.Request
//	}
//	var tests []struct {
//		name    string
//		fields  fields
//		args    args
//		wantErr bool
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			p := &PersonController{
//				service: tt.fields.service,
//			}
//			if err := p.CreatePerson(tt.args.w, tt.args.r); (err != nil) != tt.wantErr {
//				t.Errorf("CreatePerson() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}
//
//func TestPersonController_GetPeopleByTerm(t *testing.T) {
//	type fields struct {
//		service *services.PersonService
//	}
//	type args struct {
//		w http.ResponseWriter
//		r *http.Request
//	}
//	var tests []struct {
//		name    string
//		fields  fields
//		args    args
//		wantErr bool
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			p := &PersonController{
//				service: tt.fields.service,
//			}
//			if err := p.GetPeopleByTerm(tt.args.w, tt.args.r); (err != nil) != tt.wantErr {
//				t.Errorf("GetPeopleByTerm() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}
//
//func TestPersonController_GetPersonByID(t *testing.T) {
//	type fields struct {
//		service *services.PersonService
//	}
//	type args struct {
//		w http.ResponseWriter
//		r *http.Request
//	}
//	var tests []struct {
//		name    string
//		fields  fields
//		args    args
//		wantErr bool
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			p := &PersonController{
//				service: tt.fields.service,
//			}
//			if err := p.GetPersonByID(tt.args.w, tt.args.r); (err != nil) != tt.wantErr {
//				t.Errorf("GetPersonByID() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}
//
