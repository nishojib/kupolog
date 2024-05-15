package server

// Envelope is a generic map type.
// Used to represent json objects in the API in the form
//
//	{
//	  	"users": [
//		 	{
//				"id": 1,
//				"name": "Alice"
//	  	 	},
//	  	]
//	}
type Envelope[T any] map[string]T
