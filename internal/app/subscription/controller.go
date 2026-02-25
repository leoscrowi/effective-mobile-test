package subscription

import "net/http"

type Controller interface {
	CreateSubscription(w http.ResponseWriter, r *http.Request)
	ReadSubscription(w http.ResponseWriter, r *http.Request)
	EditSubscription(w http.ResponseWriter, r *http.Request)
	DeleteSubscription(w http.ResponseWriter, r *http.Request)
	ReadSubscriptionsList(w http.ResponseWriter, r *http.Request)
}
