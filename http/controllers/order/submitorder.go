package order

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"github.com/kenny26/maynooth-api/models"
	"github.com/kenny26/maynooth-api/utils"
	"github.com/satori/go.uuid"
)

type OrderItemBody struct {
	VariantID uint `json:"variantId"`
	Quantity int `json:"quantity"`
}

type SubmitOrderBody struct {
	CityID uint `json:"cityId"`
	Items []OrderItemBody `json:"items"`
}

func SubmitOrder(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")

	var param []SubmitOrderBody

	err := json.Unmarshal(ctx.PostBody(), &param)
	if err != nil {
		// If the structure of the body is wrong, return an HTTP error
		ctx.Error("BadRequest", fasthttp.StatusBadRequest)
		return
	}

	userID := ctx.UserValue("ID").(uint)

	var cityIds []uint
	var variantIds []uint

	for _, cart := range param {
		cityIds = append(cityIds, cart.CityID)

		for _, item := range cart.Items {
			variantIds = append(variantIds, item.VariantID)
		}
	}

	cities := models.GetActiveCities(cityIds)
	variants := models.GetActiveProductVariants(variantIds)

	var totalAmount float64
	order := &models.Order{
		Number: uuid.NewV4().String(),
		Status: "Pending",
		UserID: userID,
	}
	models.GetDB().Debug().Create(&order)

	for _, cart := range param {
		orderDetail := &models.OrderDetail{
			OrderID: order.ID,
			CityID: cart.CityID,
			Status: "Pending",
		}

		for _, city := range *cities {
			if city.ID == cart.CityID {
				orderDetail.CityID = cart.CityID
				orderDetail.ShippingPrice = city.ShippingPrice

				totalAmount += city.ShippingPrice
				break
			}
		}
		models.GetDB().Debug().Create(&orderDetail)

		for _, item := range cart.Items {
			orderItem := &models.OrderItem{
				Quantity: item.Quantity,
				OrderDetailID: orderDetail.ID,
			}

			for _, variant := range *variants {
				if variant.ID == item.VariantID {
					orderItem.ProductName = variant.Product.Name
					orderItem.ProductPrice = variant.Product.Price
					orderItem.ProductVariantID = variant.ID
					orderItem.ProductVariantName = variant.Name

					totalAmount += variant.Product.Price * float64(item.Quantity)
					break
				}
			}

			models.GetDB().Debug().Create(&orderItem)
		}
	}

	order.Amount = totalAmount
	models.GetDB().Save(&order)

	resp := utils.Message(true, "success")
	resp["data"] = order
	utils.Respond(ctx, resp)
	return
}
