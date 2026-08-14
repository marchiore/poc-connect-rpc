package motherboard

import (
	"context"
	motherboardv1 "poc/gen/proto/Motherboard"
	"poc/gen/proto/Motherboard/motherboardv1connect"
)

type MotherboardServer struct {
	service *MotherboardService
}

var _ motherboardv1connect.MotherboardServiceHandler = (*MotherboardServer)(nil)

func NewMotherboardServer(service *MotherboardService) *MotherboardServer {
	return &MotherboardServer{
		service: service,
	}
}

func (m *MotherboardServer) PostMotherboard(ctx context.Context, req *motherboardv1.PostMotherboardRequest) (*motherboardv1.PostMotherboardResponse, error) {
	id, err := m.service.CreateMotherboard(req.Description, req.ColorId, req.ManufacturerId)
	if err != nil {
		return nil, err
	}

	return &motherboardv1.PostMotherboardResponse{
		Id: id,
	}, nil
}

func (m *MotherboardServer) GetMotherboard(ctx context.Context, req *motherboardv1.MotherboardRequest) (*motherboardv1.MotherboardResponse, error) {
	record, err := m.service.GetMotherboardByName(req.Name)

	if err != nil {
		return nil, err
	}

	// Mapeia só os campos que devem mostrar pro front
	return &motherboardv1.MotherboardResponse{
		Name:  record.Name,
		Color: record.Color,
		Price: float32(record.Price),
	}, nil
}
