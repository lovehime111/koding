package main

import (
	"encoding/json"
	"fmt"
	"socialapi/workers/payment/paymentemail"
	"socialapi/workers/payment/paymentwebhook/webhookmodels"
	"socialapi/workers/payment/stripe"
)

func stripeInvoiceCreated(raw []byte, c *Controller) error {
	var invoice *webhookmodels.StripeInvoice

	err := json.Unmarshal(raw, &invoice)
	if err != nil {
		return err
	}

	err = stripe.InvoiceCreatedWebhook(invoice)
	if err != nil {
		return err
	}

	return sendInvoiceCreatedEmail(invoice, c)
}

func sendInvoiceCreatedEmail(req *webhookmodels.StripeInvoice, c *Controller) error {
	emailAddress, err := getEmailForCustomer(req.CustomerId)
	if err != nil {
		return err
	}

	if len(req.Lines.Data) < 0 {
		return fmt.Errorf(
			"Invoice: %s for %s has 0 line items", req.ID, req.CustomerId,
		)
	}

	opts := map[string]string{
		"planName": req.Lines.Data[0].Plan.Name,
		"price":    formatStripeAmount(req.CustomerId, req.AmountDue),
	}

	return paymentemail.Send(
		c.Email, paymentemail.InvoiceCreated, emailAddress, opts,
	)
}
