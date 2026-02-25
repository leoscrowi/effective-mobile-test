package v1

import "github.com/go-chi/chi/v5"

func (c *SubscriptionController) SetupRoutes(r chi.Router) {
	r.Route("/subscriptions", func(r chi.Router) {
		r.Post("/", c.CreateSubscription)
		r.Get("/{id}", c.ReadSubscription)
		r.Get("/", c.ReadSubscriptionsList)
		r.Patch("/{id}", c.EditSubscription)
		r.Delete("/{id}", c.DeleteSubscription)
		r.Get("/amount", c.GetSubscriptionsAmount)
	})
}
