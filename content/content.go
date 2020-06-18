package content

type BranchNode struct {
	Node
	Build Build
}

type Build struct {
	Render           bool
	List             string // never always local
	PublishResources bool
}

type Item struct {
	Params  Params
	Content Content
}

type LeafNode struct {
	Node
}

type Node struct {
	Item
	Resources Resources
}

type Cascade struct {
}

type Content struct {
	Path string
	URI  string
	Raw  string
}

type Params map[string]interface{}

type Resource struct {
	Item
}
type Resources []Resource
