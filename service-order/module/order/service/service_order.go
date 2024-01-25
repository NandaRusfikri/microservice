package order

import (
	"context"
	"fmt"
	"service-order/dto"
	repoorder "service-order/module/order"
	"service-order/module/order/entity"
	repoproduct "service-order/module/product/repository"
	"service-order/pkg"
)

type orderService struct {
	OrderRepository repoorder.RepositoryInterface
	OrderRepoRPC    repoproduct.RepositoryGRPCInterface
	Kafka           *pkg.KafkaProducer
}

func NewOrderService(
	repository repoorder.RepositoryInterface,
	repoProduct repoproduct.RepositoryGRPCInterface,
	kafka *pkg.KafkaProducer,
) repoorder.OrderServiceInterface {
	return &orderService{
		OrderRepository: repository,
		OrderRepoRPC:    repoProduct,
		Kafka:           kafka,
	}

}

func (s *orderService) Create(input *dto.SchemaOrder) (*entities.Order, dto.ResponseError) {

	var product dto.SchemaOrder

	Product, err := s.OrderRepoRPC.FindProductByIdRepository(context.Background(), input.ProductId)
	if Product.ID < 1 {
		return nil, err
	}
	fmt.Println("oke")
	ayam := s.Kafka.KirimPesan("sarama", "hola", 0)
	if ayam != nil {
		fmt.Println(ayam.Error())
	}

	fmt.Println("mantap")
	product.ProductId = Product.ID
	product.UserId = input.UserId
	product.Amount = Product.Price

	res, err := s.OrderRepository.Create(&product)
	return res, err
}

func (s *orderService) GetById(input *dto.SchemaOrder) (*entities.Order, dto.ResponseError) {

	var student dto.SchemaOrder
	student.ID = input.ID

	res, err := s.OrderRepository.GetById(&student)
	return res, err
}

func (s *orderService) GetList() ([]*entities.Order, dto.ResponseError) {

	res, err := s.OrderRepository.GetList()
	return res, err
}
