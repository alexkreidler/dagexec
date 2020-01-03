package api

//go:generate go-enum -f=$GOFILE
//
// --marshal --lower

// State is an enumeration of the possible states of a node
// ENUM(Waiting=1,Running,Failed,Succeeded)
type State int32
