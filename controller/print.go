package controller

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"

	"github.com/kataras/golog"

	"github.com/gorilla/mux"
)

type PrintController struct {
	print printServiceIface
}

func NewPrintController(printService printServiceIface) *PrintController {
	return &PrintController{
		print: printService,
	}
}

func (c *PrintController) OrderByInvoiceToPdfHandler(w http.ResponseWriter, r *http.Request) {
	buffer := new(bytes.Buffer)
	vars := mux.Vars(r)
	invoice := vars["invoice"]
	golog.Infof("GET - Printer: OrderByInvoiceToPdfHandler (/print/order/%s)", invoice)
	textAttachment := fmt.Sprintf("attachment; filename=\"%s.pdf\";", invoice)

	pdf, err := c.print.OrderByInvoiceToPdf(invoice)
	if err != nil {
		renderJSONError(w, http.StatusInternalServerError, err, err.Error())
	}

	err = pdf.Output(buffer)
	if err != nil {
		message := fmt.Sprintf("GET - Printer: OrderByInvoiceToPdfHandler (/print/order/%s)", invoice)
		renderJSONError(w, http.StatusInternalServerError, err, message)
		return
	}
	pdf.Close()

	_, err = buffer.WriteTo(w)
	if err != nil {
		message := fmt.Sprintf("GET - Printer: OrderByInvoiceToPdfHandler (/print/order/%s)", invoice)
		renderJSONError(w, http.StatusInternalServerError, err, message)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Pragma", "public")
	w.Header().Set("Expires", "0")
	w.Header().Set("Cache-Control", "must-revalidate, post-check=0, pre-check=0")
	w.Header().Set("Cache-Control", "private") // required for certain browsers
	w.Header().Set("Content-Type", "application/pdf; charset=utf-8")
	w.Header().Set("Content-Disposition", textAttachment)
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	w.Header().Set("Content-Transfer-Encoding", "binary")
}
