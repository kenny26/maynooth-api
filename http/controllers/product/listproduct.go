package product

import (
	"github.com/valyala/fasthttp"
	"github.com/kenny26/maynooth-api/utils"
	"log"
	"strconv"
	"github.com/kenny26/maynooth-api/models"
)

type ProductVariantResponse struct {
	ID	uint `json:"id"`
	Name	string `json:"name"`
	Sku	string `json:"sku"`
	Stock	int `json:"stock"`
}

type ProductListResponse struct {
	ID	uint `json:"id"`
	Name	string `json:"name"`
	Size	string `json:"size"`
	Description	string `json:"description"`
	Price	float64 `json:"price"`
	ImageUrl	string `json:"imageUrl"`
	ProductImages []string `json:"images"`
	ProductVariants []ProductVariantResponse `json:"variants"`
}

func ListProduct(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")

	queryParams := ctx.QueryArgs()

	name := string(queryParams.Peek("name"))
	order := string(queryParams.Peek("order"))
	categoryId, _ := strconv.Atoi(string(queryParams.Peek("categoryId")))

	log.Println("params", categoryId, name, order)
	products := models.ListActiveProducts(categoryId, name, order)

	var productList []ProductListResponse

	for _, product := range *products {
		if product.ID != 0 {
			parsedProduct := ProductListResponse{}
			parsedProduct.ID = product.ID
			parsedProduct.Name = product.Name
			parsedProduct.Size = product.Size
			parsedProduct.Description = product.Description
			parsedProduct.Price = product.Price

			var productVariants []ProductVariantResponse

			for _, variant := range product.ProductVariants {
				if (variant.ID != 0) {
					productVariant := ProductVariantResponse{}
					productVariant.ID = variant.ID
					productVariant.Name = variant.Name
					productVariant.Sku = variant.Sku
					productVariant.Stock = variant.Stock

					productVariants = append(productVariants, productVariant)
				}
			}

			var productImages []string

			for _, image := range product.ProductImages {
				if (image.ID != 0) {
					productImages = append(productImages, image.Url)

					if (image.IsThumbnail) {
						parsedProduct.ImageUrl = image.Url
					}
				}
			}

			parsedProduct.ProductImages = productImages
			parsedProduct.ProductVariants = productVariants

			productList = append(productList, parsedProduct)
		}
	}

	ctx.SetStatusCode(fasthttp.StatusOK)

	resp := utils.Message(true, "success")
	resp["data"] = productList
	utils.Respond(ctx, resp)
}
