package p2p

// Peer is a representation of a remote node
type Peer interface{

}

// Transport is anything that handles the
// communication in a network between the nodes
// This can be of the form (TCP,UDP,Websockets, ...)
type Transport interface{
	ListenAndAccept() error
}

