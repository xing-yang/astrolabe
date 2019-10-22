// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// CreateSnapshotHandlerFunc turns a function with the right signature into a create snapshot handler
type CreateSnapshotHandlerFunc func(CreateSnapshotParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateSnapshotHandlerFunc) Handle(params CreateSnapshotParams) middleware.Responder {
	return fn(params)
}

// CreateSnapshotHandler interface for that can handle valid create snapshot params
type CreateSnapshotHandler interface {
	Handle(CreateSnapshotParams) middleware.Responder
}

// NewCreateSnapshot creates a new http.Handler for the create snapshot operation
func NewCreateSnapshot(ctx *middleware.Context, handler CreateSnapshotHandler) *CreateSnapshot {
	return &CreateSnapshot{Context: ctx, Handler: handler}
}

/*CreateSnapshot swagger:route POST /arachne/{service}/{protectedEntityID}/snapshots createSnapshot

Creates a new snapshot for this protected entity

*/
type CreateSnapshot struct {
	Context *middleware.Context
	Handler CreateSnapshotHandler
}

func (o *CreateSnapshot) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateSnapshotParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
