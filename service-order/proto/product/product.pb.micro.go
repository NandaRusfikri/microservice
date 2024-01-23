// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/product/product.proto

package product

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for ServiceProductHandler service

func NewServiceProductHandlerEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for ServiceProductHandler service

type ServiceProductHandlerService interface {
	GetProductByIdRPC(ctx context.Context, in *ProductByIdRequest, opts ...client.CallOption) (*EntityProtoProduct, error)
	ListProductRPC(ctx context.Context, in *emptypb.Empty, opts ...client.CallOption) (*ResponseEntityProductList, error)
	CreateProductRPC(ctx context.Context, in *CreateProductRequest, opts ...client.CallOption) (*EntityProtoProduct, error)
	UpdateProductRPC(ctx context.Context, in *EntityProtoProduct, opts ...client.CallOption) (*EntityProtoProduct, error)
	DeleteProductRPC(ctx context.Context, in *ProductByIdRequest, opts ...client.CallOption) (*emptypb.Empty, error)
}

type serviceProductHandlerService struct {
	c    client.Client
	name string
}

func NewServiceProductHandlerService(name string, c client.Client) ServiceProductHandlerService {
	return &serviceProductHandlerService{
		c:    c,
		name: name,
	}
}

func (c *serviceProductHandlerService) GetProductByIdRPC(ctx context.Context, in *ProductByIdRequest, opts ...client.CallOption) (*EntityProtoProduct, error) {
	req := c.c.NewRequest(c.name, "ServiceProductHandler.GetProductByIdRPC", in)
	out := new(EntityProtoProduct)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceProductHandlerService) ListProductRPC(ctx context.Context, in *emptypb.Empty, opts ...client.CallOption) (*ResponseEntityProductList, error) {
	req := c.c.NewRequest(c.name, "ServiceProductHandler.ListProductRPC", in)
	out := new(ResponseEntityProductList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceProductHandlerService) CreateProductRPC(ctx context.Context, in *CreateProductRequest, opts ...client.CallOption) (*EntityProtoProduct, error) {
	req := c.c.NewRequest(c.name, "ServiceProductHandler.CreateProductRPC", in)
	out := new(EntityProtoProduct)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceProductHandlerService) UpdateProductRPC(ctx context.Context, in *EntityProtoProduct, opts ...client.CallOption) (*EntityProtoProduct, error) {
	req := c.c.NewRequest(c.name, "ServiceProductHandler.UpdateProductRPC", in)
	out := new(EntityProtoProduct)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceProductHandlerService) DeleteProductRPC(ctx context.Context, in *ProductByIdRequest, opts ...client.CallOption) (*emptypb.Empty, error) {
	req := c.c.NewRequest(c.name, "ServiceProductHandler.DeleteProductRPC", in)
	out := new(emptypb.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ServiceProductHandler service

type ServiceProductHandlerHandler interface {
	GetProductByIdRPC(context.Context, *ProductByIdRequest, *EntityProtoProduct) error
	ListProductRPC(context.Context, *emptypb.Empty, *ResponseEntityProductList) error
	CreateProductRPC(context.Context, *CreateProductRequest, *EntityProtoProduct) error
	UpdateProductRPC(context.Context, *EntityProtoProduct, *EntityProtoProduct) error
	DeleteProductRPC(context.Context, *ProductByIdRequest, *emptypb.Empty) error
}

func RegisterServiceProductHandlerHandler(s server.Server, hdlr ServiceProductHandlerHandler, opts ...server.HandlerOption) error {
	type serviceProductHandler interface {
		GetProductByIdRPC(ctx context.Context, in *ProductByIdRequest, out *EntityProtoProduct) error
		ListProductRPC(ctx context.Context, in *emptypb.Empty, out *ResponseEntityProductList) error
		CreateProductRPC(ctx context.Context, in *CreateProductRequest, out *EntityProtoProduct) error
		UpdateProductRPC(ctx context.Context, in *EntityProtoProduct, out *EntityProtoProduct) error
		DeleteProductRPC(ctx context.Context, in *ProductByIdRequest, out *emptypb.Empty) error
	}
	type ServiceProductHandler struct {
		serviceProductHandler
	}
	h := &serviceProductHandlerHandler{hdlr}
	return s.Handle(s.NewHandler(&ServiceProductHandler{h}, opts...))
}

type serviceProductHandlerHandler struct {
	ServiceProductHandlerHandler
}

func (h *serviceProductHandlerHandler) GetProductByIdRPC(ctx context.Context, in *ProductByIdRequest, out *EntityProtoProduct) error {
	return h.ServiceProductHandlerHandler.GetProductByIdRPC(ctx, in, out)
}

func (h *serviceProductHandlerHandler) ListProductRPC(ctx context.Context, in *emptypb.Empty, out *ResponseEntityProductList) error {
	return h.ServiceProductHandlerHandler.ListProductRPC(ctx, in, out)
}

func (h *serviceProductHandlerHandler) CreateProductRPC(ctx context.Context, in *CreateProductRequest, out *EntityProtoProduct) error {
	return h.ServiceProductHandlerHandler.CreateProductRPC(ctx, in, out)
}

func (h *serviceProductHandlerHandler) UpdateProductRPC(ctx context.Context, in *EntityProtoProduct, out *EntityProtoProduct) error {
	return h.ServiceProductHandlerHandler.UpdateProductRPC(ctx, in, out)
}

func (h *serviceProductHandlerHandler) DeleteProductRPC(ctx context.Context, in *ProductByIdRequest, out *emptypb.Empty) error {
	return h.ServiceProductHandlerHandler.DeleteProductRPC(ctx, in, out)
}
