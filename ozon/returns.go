package ozon

import (
	"net/http"
	"time"

	core "github.com/diphantxm/ozon-api-client"
)

type Returns struct {
	client *core.Client
}

type GetFBOReturnsParams struct {
	// Filter
	Filter GetFBOReturnsFilter `json:"filter"`

	// Identifier of the last value on the page. Leave this field blank in the first request.
	//
	// To get the next values, specify the recieved value in the next request in the `last_id` parameter
	LastId int64 `json:"last_id"`

	// Number of values in the response
	Limit int64 `json:"limit"`
}

type GetFBOReturnsFilter struct {
	// Shipment number
	PostingNumber string `json:"posting_number"`

	Status []string `json:"status"`
}

type GetFBOReturnsResponse struct {
	core.CommonResponse

	// Identifier of the last value on the page
	LastId int64 `json:"last_id"`

	// Returns information
	Returns []struct {
		// Time when a return was received from the customer
		AcceptedFromCustomerMoment time.Time `json:"accepted_from_customer_moment"`

		// Seller identifier
		CompanyId int64 `json:"company_id"`

		// Current return location
		CurrentPlaceName string `json:"current_place_name"`

		// Return destination
		DestinationPlaceName string `json:"dst_place_name"`

		// Return identifier
		Id int64 `json:"id"`

		// Indication that the package has been opened. true, if it has been
		IsOpened bool `json:"is_opened"`

		// Shipment number
		PostingNumber string `json:"posting_number"`

		// Return reason
		ReturnReasonName string `json:"return_reason_name"`

		// Return delivery time to the Ozon warehouse
		ReturnedToOzonMoment time.Time `json:"returned_to_ozon_moment"`

		// Product identifier in the Ozon system, SKU
		SKU int64 `json:"sku"`

		// Return status
		Status string `json:"status_name"`
	} `json:"returns"`
}

// Method for getting information on returned products that are sold from the Ozon warehouse
func (c Returns) GetFBOReturns(params *GetFBOReturnsParams) (*GetFBOReturnsResponse, error) {
	url := "/v3/returns/company/fbo"

	resp := &GetFBOReturnsResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type GetFBSReturnsParams struct {
	// Filter
	Filter GetFBSReturnsFilter `json:"filter"`

	// Number of values in the response:
	//   - maximum — 1000,
	//   - minimum — 1
	Limit int64 `json:"limit"`

	// Number of elements that will be skipped in the response.
	// For example, if offset=10, the response will start with the 11th element found
	Offset int64 `json:"offset"`
}

type GetFBSReturnsFilter struct {
	// Time of receiving the return from the customer
	AcceptedFromCustomerMoment GetFBSReturnsFilterTimeRange `json:"accepted_from_customer_moment"`

	// Last day of free storage
	LastFreeWaitingDay []GetFBSReturnsFilterTimeRange `json:"last_free_waiting_dat"`

	// Order ID
	OrderId int64 `json:"order_id"`

	// Shipment ID
	PostingNumber []string `json:"posting_number"`

	// Product name
	ProductName string `json:"product_name"`

	// Product ID
	ProductOfferId string `json:"product_offer_id"`

	// Return status:
	//   - returned_to_seller — returned to seller,
	//   - waiting_for_seller — waiting for seller,
	//   - accepted_from_customer — accepted from customer,
	//   - cancelled_with_compensation — cancelled with compensation,
	//   - ready_for_shipment — ready for shipment
	Status string `json:"status"`
}

type GetFBSReturnsFilterTimeRange struct {
	// The beginning of the period.
	//
	// Format: YYYY-MM-DDTHH:mm:ss.sssZ.
	//
	// Example: 2019-11-25T10:43:06.51
	TimeFrom time.Time `json:"time_from"`

	// The end of the period
	//
	// Format: YYYY-MM-DDTHH:mm:ss.sssZ.
	//
	// Example: 2019-11-25T10:43:06.51
	TimeTo time.Time `json:"time_to"`
}

type GetFBSReturnsResponse struct {
	core.CommonResponse

	Result struct {
		// Elements counter in the response
		Count int64 `json:"count"`

		// Returns information
		Returns []struct {
			// Time of receiving the return from the customer
			AcceptedFromCustomerAmount string `json:"accepted_from_customer_amount"`

			// Bottom barcode on the product label
			ClearingId int64 `json:"clearing_id"`

			// Commission fee
			Commission float64 `json:"commission"`

			// Commission percentage
			CommissionPercent float64 `json:"commission_percent"`

			// Return identifier
			Id int64 `json:"id"`

			// If the product is in transit — true
			IsMoving bool `json:"is_moving"`

			// Indication that the package has been opened. true, if it has been
			IsOpened bool `json:"is_opened"`

			// Last day of free storage
			LastFreeWaitingDay string `json:"last_free_waiting_day"`

			// ID of the warehouse the product is being transported to
			PlaceId int64 `json:"place_id"`

			// Name of the warehouse the product is being transported to
			MovingToPlaceName string `json:"moving_to_place_name"`

			// Delivery cost
			PickingAmount float64 `json:"picking_amount"`

			// Shipment number
			PostingNumber string `json:"posting_number"`

			// Current product price without a discount
			Price float64 `json:"price"`

			// Product price without commission
			PriceWithoutCommission float64 `json:"price_without_commission"`

			// Product identifier
			ProductId int64 `json:"product_id"`

			// Product name
			ProductName string `json:"product_name"`

			// Product quantity
			Quantity int64 `json:"quantity"`

			// Product return date
			ReturnDate string `json:"return_date"`

			// Return reason
			ReturnReasonName string `json:"return_reason_name"`

			// Date when the product is ready to be handed over to the seller
			WaitingForSellerDate string `json:"waiting_for_seller_date_time"`

			// Date of handing over the product to the seller
			ReturnedToSellerDate string `json:"returned_to_seller_date_time"`

			// Return storage period in days
			WaitingForSellerDays int64 `json:"waiting_for_seller_days"`

			// Return storage cost
			ReturnsKeepingCost float64 `json:"returns_keeping_cost"`

			// Product identifier in the Ozon system, SKU
			SKU int64 `json:"sku"`

			// Return status
			Status string `json:"status"`
		} `json:"returns"`
	} `json:"result"`
}

// Method for getting information on returned products that are sold from the seller's warehouse
func (c Returns) GetFBSReturns(params *GetFBSReturnsParams) (*GetFBSReturnsResponse, error) {
	url := "/v2/returns/company/fbs"

	resp := &GetFBSReturnsResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}
