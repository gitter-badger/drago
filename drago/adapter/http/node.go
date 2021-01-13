package http

import (
	"net/http"

	structs "github.com/seashell/drago/drago/structs"
	rpc "github.com/seashell/drago/pkg/rpc"
)

// NodeHandler :
type NodeHandler struct {
	rpcClient *rpc.Client
}

// NewNodeHandler :
func NewNodeHandler(rpcClient *rpc.Client) *NodeHandler {
	return &NodeHandler{
		rpcClient: rpcClient,
	}
}

// Handle :
func (h *NodeHandler) Handle(rw http.ResponseWriter, req *http.Request) (interface{}, error) {

	pathParams := parsePathParams(req)
	if len(pathParams) > 1 {
		return nil, NewCodedError(404, ErrNotFound)
	}

	switch req.Method {
	case "GET":
		return h.handleGet(rw, req, pathParams)
	case "PUT", "POST":
		return h.handlePost(rw, req, pathParams)
	default:
		return nil, NewCodedError(405, ErrMethodNotAllowed)
	}
}

func (h *NodeHandler) handleGet(rw http.ResponseWriter, req *http.Request, pathParams []string) (interface{}, error) {

	nodeID := pathParams[0]

	if nodeID == "" {
		return h.handleList(rw, req)
	}

	args := structs.NodeSpecificRequest{
		QueryOptions: parseQueryOptions(req),
		ID:           nodeID,
	}

	var out structs.SingleNodeResponse
	if err := h.rpcClient.Call("Node.GetNode", &args, &out); err != nil {
		return nil, parseError(err)
	}

	return out.Node, nil
}

func (h *NodeHandler) handleList(rw http.ResponseWriter, req *http.Request) (interface{}, error) {

	args := &structs.NodeListRequest{
		QueryOptions: parseQueryOptions(req),
	}

	var out structs.NodeListResponse
	if err := h.rpcClient.Call("Node.ListNodes", &args, &out); err != nil {
		return nil, parseError(err)
	}

	if out.Items == nil {
		out.Items = make([]*structs.NodeListStub, 0)
	}

	return out.Items, nil
}

func (h *NodeHandler) handlePost(rw http.ResponseWriter, req *http.Request, pathParams []string) (interface{}, error) {

	nodeID := pathParams[0]

	var node structs.Node
	err := parseBody(req.Body, &node)
	if err != nil {
		return nil, NewCodedError(500, ErrInternal, err)
	}

	// Make sure the node ID matches
	if node.ID != nodeID {
		return nil, NewCodedError(400, "Node ID does not match request path")
	}

	args := &structs.NodePreregisterRequest{
		Node:         &node,
		WriteRequest: parseWriteRequestOptions(req),
	}

	var out structs.NodePreregisterResponse
	if err := h.rpcClient.Call("Node.PreregisterNode", &args, &out); err != nil {
		return nil, parseError(err)
	}

	return nil, nil
}
