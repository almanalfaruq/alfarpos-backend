package service

import (
	"fmt"
	"strings"

	"github.com/almanalfaruq/alfarpos-backend/util"
	"github.com/dustin/go-humanize"
	"github.com/kataras/golog"

	"github.com/jung-kurt/gofpdf"
)

type PrintService struct {
	order orderRepositoryIface
	conf  util.Config
}

func NewPrintService(conf util.Config, orderRepo orderRepositoryIface) *PrintService {
	return &PrintService{
		order: orderRepo,
		conf:  conf,
	}
}

func (service *PrintService) OrderByInvoiceToPdf(invoice string) (*gofpdf.Fpdf, error) {
	invoice = strings.ToLower(invoice)
	order, err := service.order.FindByInvoice(invoice)
	if err != nil {
		golog.Error(err)
		return nil, err
	}

	pdf := gofpdf.NewCustom(&gofpdf.InitType{
		UnitStr: "mm",
		Size:    gofpdf.SizeType{Wd: 76.0, Ht: 83.0},
	})
	pdf.SetMargins(5, 2, 5)
	pdf.SetFont("Helvetica", "B", 12)
	pdf.AddPage()
	pdf.WriteAligned(0, 35, service.conf.ShopProfile.Name, "C")
	pdf.SetFont("Helvetica", "", 10)
	pdf.WriteAligned(0, 35, service.conf.ShopProfile.Address, "C")
	pdf.SetFont("Courier", "", 10)
	pdf.SetDashPattern([]float64{0.8, 0.8}, 0)
	textDate := fmt.Sprintf("Tgl.: %02d-%02d-%d", order.CreatedAt.Day(), order.CreatedAt.Month(), order.CreatedAt.Year())
	pdf.CellFormat(33, 5, textDate, "T", 0, "LM", false, 0, "")
	textCashier := fmt.Sprintf("Kasir: %s", order.User.FullName)
	pdf.CellFormat(33, 5, textCashier, "T", 1, "RM", false, 0, "")
	textInvoice := fmt.Sprintf("No.#: %s", order.Invoice)
	pdf.CellFormat(33, 5, textInvoice, "B", 0, "LM", false, 0, "")
	textTime := fmt.Sprintf("Jam: %02d:%02d:%02d", order.CreatedAt.Hour(), order.CreatedAt.Minute(), order.CreatedAt.Second())
	pdf.CellFormat(33, 5, textTime, "B", 1, "RM", false, 0, "")

	for _, orderDetail := range order.OrderDetails {
		pdf.CellFormat(0, 5, orderDetail.Product.Name, "", 1, "LM", false, 0, "")
		textQty := fmt.Sprintf("%d %s x", orderDetail.Quantity, orderDetail.Product.Unit.Name)
		pdf.CellFormat(22, 5, textQty, "", 0, "RM", false, 0, "")
		textPrice := fmt.Sprintf("%d =", orderDetail.Product.SellPrice)
		pdf.CellFormat(22, 5, textPrice, "", 0, "LM", false, 0, "")
		textSubTotal := fmt.Sprintf("%d", orderDetail.SubTotal)
		pdf.CellFormat(22, 5, textSubTotal, "", 1, "RM", false, 0, "")
	}

	pdf.CellFormat(33, 5, "Total :", "T", 0, "LM", false, 0, "")
	totalText := fmt.Sprintf("Rp %s", humanize.FormatInteger("#.###,", int(order.Total)))
	pdf.CellFormat(33, 5, totalText, "T", 1, "RM", false, 0, "")
	pdf.CellFormat(33, 5, "Total Bayar :", "", 0, "LM", false, 0, "")
	totalPaymentText := fmt.Sprintf("Rp %s", humanize.FormatInteger("#.###,", int(order.AmountPaid)))
	pdf.CellFormat(33, 5, totalPaymentText, "", 1, "RM", false, 0, "")
	pdf.CellFormat(33, 5, "Total Kembali :", "", 0, "LM", false, 0, "")
	totalChangeText := fmt.Sprintf("Rp %s", humanize.FormatInteger("#.###,", int(order.TotalChange)))
	pdf.CellFormat(33, 5, totalChangeText, "", 1, "RM", false, 0, "")

	pdf.WriteAligned(0, 35, "Terima Kasih", "C")
	return pdf, nil
}
