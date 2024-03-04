package controller

import (
	"log"
	"math/rand"
	"time"

	"github.com/ffalconesmera/bank-simulator/utils"
	"github.com/gin-gonic/gin"
)

type TransactionInput struct {
	CardNumber       string `json:"card_number"`
	PaymentReference string `json:"payment_reference"`
}

type RefundInput struct {
	RefundCase       string `json:"refund_case"`
	PaymentReference string `json:"payment_reference"`
}

type TransactionResponse struct {
	Status    string `json:"status"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Reference string `json:"reference"`
}

type bankController struct {
}

func NewBankController() *bankController {
	return &bankController{}
}

func randSleep() time.Duration {
	return time.Duration(int64(rand.Intn(1200)+200) * time.Hour.Milliseconds())
}

func (b bankController) ProcessPayment(c *gin.Context) {
	time.Sleep(randSleep())

	cardNumbersPaymentsCases := map[string]string{
		"card_success":               "payment processed successullfy",
		"card_insufficient_founds":   "insufficient funds",
		"card_incorrect":             "credit card information is incorrect",
		"card_bad_request":           "bad request",
		"card_timeout":               "payment request is timeout",
		"card_service_not_available": "service not available",
		"card_server_error":          "internal server error",
	}

	cardNumbersPaymentsCasesCode := map[string]int{
		"card_success":               1000,
		"card_insufficient_founds":   2000,
		"card_incorrect":             3000,
		"card_bad_request":           4000,
		"card_timeout":               5000,
		"card_service_not_available": 6000,
	}

	log.Println("mapping transaction data..")
	var transaction *TransactionInput
	if err := c.ShouldBindJSON(&transaction); err != nil {
		res := TransactionResponse{
			Status:    "failure",
			Code:      cardNumbersPaymentsCasesCode["card_bad_request"],
			Message:   cardNumbersPaymentsCases["card_bad_request"],
			Reference: "",
		}

		log.Println(res)
		c.JSON(200, res)
		return
	}

	caseFound, ok := cardNumbersPaymentsCases[transaction.CardNumber]

	log.Println(gin.H{
		"case":        caseFound,
		"card_number": transaction.CardNumber,
	})

	if !ok {
		res := TransactionResponse{
			Status:    "failure",
			Code:      cardNumbersPaymentsCasesCode["card_incorrect"],
			Message:   cardNumbersPaymentsCases["card_incorrect"],
			Reference: "",
		}

		log.Println(res)
		c.JSON(200, res)
		return
	}

	// card_number -> "refund_server_error" send internal server error
	if transaction.CardNumber == "card_server_error" {
		c.JSON(500, "internal server error by user")
		return
	}

	reference := ""
	status := "failure"
	if transaction.CardNumber == "card_success" {
		reference = utils.NewUUIDString()
		status = "succeeded"
	}

	res := TransactionResponse{
		Status:    status,
		Code:      cardNumbersPaymentsCasesCode[transaction.CardNumber],
		Message:   cardNumbersPaymentsCases[transaction.CardNumber],
		Reference: reference,
	}

	log.Println(res)
	c.JSON(200, res)
}

func (b bankController) RefundPayment(c *gin.Context) {
	time.Sleep(randSleep())

	refundCases := map[string]string{
		"refund_success":               "refund processed successullfy",
		"refund_already_refunded":      "payment already refunded",
		"refund_incorrect":             "payment information is incorrect",
		"refund_bad_request":           "bad request",
		"refund_timeout":               "refund request is timeout",
		"refund_service_not_available": "service not available",
		"refund_server_error":          "internal server error",
	}

	refundCasesCode := map[string]int{
		"refund_success":               1000,
		"refund_already_refunded":      2000,
		"refund_incorrect":             3000,
		"refund_bad_request":           4000,
		"refund_timeout":               5000,
		"refund_service_not_available": 6000,
	}

	log.Println("mapping refund data..")
	var refund *RefundInput
	if err := c.ShouldBindJSON(&refund); err != nil {
		res := TransactionResponse{
			Status:    "failure",
			Code:      refundCasesCode["refund_bad_request"],
			Message:   refundCases["refund_bad_request"],
			Reference: "",
		}

		log.Println(res)
		c.JSON(200, res)
		return
	}

	if refund.PaymentReference == "" {
		res := TransactionResponse{
			Status:    "failure",
			Code:      refundCasesCode["refund_incorrect"],
			Message:   "payment reference is missing",
			Reference: "",
		}

		log.Println(res)
		c.JSON(200, res)
		return
	}

	_, ok := refundCases[refund.RefundCase]

	if !ok {
		res := TransactionResponse{
			Status:    "failure",
			Code:      refundCasesCode["refund_bad_request"],
			Message:   refundCases["refund_bad_request"],
			Reference: "",
		}

		log.Println(res)
		c.JSON(200, res)
		return
	}

	// refund_case -> "refund_server_error" send internal server error
	if refund.RefundCase == "refund_server_error" {
		c.JSON(500, "internal server error by user")
		return
	}

	reference := ""
	status := "failure"
	if refund.RefundCase == "refund_success" {
		reference = utils.NewUUIDString()
		status = "refunded"
	}

	res := TransactionResponse{
		Status:    status,
		Code:      refundCasesCode[refund.RefundCase],
		Message:   refundCases[refund.RefundCase],
		Reference: reference,
	}

	log.Println(res)
	c.JSON(200, res)
}
