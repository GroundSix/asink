package asink

type Block struct {

}

func NewBlock() Block {
	return Block{}
}

func (b Block) Exec() bool {
	return true
}