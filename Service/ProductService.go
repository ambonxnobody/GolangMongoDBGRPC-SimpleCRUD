package Service

import (
	productPb "GolangMongoDBGRPCSimpleCRUD/Controller"
	helperPb "GolangMongoDBGRPCSimpleCRUD/Controller/Helper"
	"GolangMongoDBGRPCSimpleCRUD/Entity"
	"GolangMongoDBGRPCSimpleCRUD/Helper"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/anypb"
	"log"
)

type ProductService struct {
	productPb.UnimplementedProductControllerServer
	DB *mongo.Client
}

func (p *ProductService) environment() *mongo.Collection {
	return p.DB.Database("GogRPC").Collection("Product")
}

func (p *ProductService) GetProducts(ctx context.Context, filter *helperPb.RequestOnAllData) (*helperPb.Response, error) {
	var allData Helper.ResponseOnAllData
	var products []*anypb.Any
	var results *mongo.Cursor
	var dataCount int64
	var err error
	var response helperPb.Response

	if !filter.GetIsDeleted() {
		dataCount, err = p.environment().CountDocuments(context.TODO(), bson.M{"DeletedAt": bson.M{"$exists": false}})

		results, err = p.environment().Find(context.TODO(), bson.M{"DeletedAt": bson.M{"$exists": false}}, options.Find().SetLimit(filter.GetLimit()).SetSkip((filter.GetPage()-1)*filter.GetLimit()))
	} else {
		dataCount, err = p.environment().CountDocuments(context.TODO(), bson.M{})

		results, err = p.environment().Find(context.TODO(), bson.M{}, options.Find().SetLimit(filter.GetLimit()).SetSkip((filter.GetPage()-1)*filter.GetLimit()))
	}

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	defer func(results *mongo.Cursor, ctx context.Context) {
		err := results.Close(ctx)
		if err != nil {
			log.Fatalf("Failed to close cursor: %v", err)
		}
	}(results, context.Background())

	for results.Next(context.Background()) {
		var productEntity Entity.ProductEntity
		err := results.Decode(&productEntity)

		product := &productPb.Product{
			ID:    productEntity.ID.Hex(),
			Name:  productEntity.Name,
			Price: float64(productEntity.Price),
			Stock: uint32(productEntity.Stock),
			Category: &productPb.Category{
				ID: productEntity.CategoryID,
			},
		}

		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}

		details, err := anypb.New(product)

		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}

		products = append(products, details)
	}

	returnData := allData.ProcessResponseOnAllData(dataCount, filter.GetLimit(), filter.GetPage())

	returnData.Data = products

	response.StatusCode = 200
	response.Message = "Berhasil mengambil data semua produk"
	response.Data, err = anypb.New(&returnData)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &response, nil
}

func (p *ProductService) GetProduct(ctx context.Context, id *helperPb.ID) (*helperPb.Response, error) {
	var product Entity.ProductEntity
	var response helperPb.Response

	productID, err := primitive.ObjectIDFromHex(id.GetID())

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	err = p.environment().FindOne(context.TODO(), bson.D{{"_id", productID}}).Decode(&product)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	productResponse := &productPb.Product{
		ID:    product.ID.Hex(),
		Name:  product.Name,
		Price: float64(product.Price),
		Stock: uint32(product.Stock),
	}

	response.StatusCode = 200
	response.Message = "Berhasil mengambil data produk"
	response.Data, err = anypb.New(productResponse)

	return &response, nil
}

func (p *ProductService) CreateProduct(ctx context.Context, product *productPb.Product) (*helperPb.Response, error) {
	return nil, nil
}

func (p *ProductService) UpdateProduct(ctx context.Context, product *productPb.Product) (*helperPb.Response, error) {
	return nil, nil
}

func (p *ProductService) DeleteProduct(ctx context.Context, id *helperPb.ID) (*helperPb.Response, error) {
	return nil, nil
}

func (p *ProductService) DeleteProductPermanently(ctx context.Context, id *helperPb.ID) (*helperPb.Response, error) {
	return nil, nil
}
