package trie

import (
	"errors"
	"log"

	"github.com/cbergoon/merkletree"
	poseidon "github.com/iden3/go-iden3-crypto/poseidon"
)

//TrieContent implements the Content interface provided by merkletree and represents the content stored in the tree.
type TrieContent struct {
	x string
}

//CalculateHash hashes the values of a TrieContent
func (t TrieContent) CalculateHash() ([]byte, error) {
	res, err := poseidon.HashBytes([]byte(t.x))
	if err != nil {
		return nil, err
	} else {
		return res.Bytes(), nil
	}
}

//Equals tests for equality of two Contents
func (t TrieContent) Equals(other merkletree.Content) (bool, error) {
	otherTC, ok := other.(TrieContent)
	if !ok {
		return false, errors.New("value is not of type TrieContent")
	}		
	return t.x == otherTC.x, nil
}

type CommitteeTrie struct {
	list []merkletree.Content
	mt *merkletree.MerkleTree
	mr []byte
}

func (t *CommitteeTrie) MerkleRoot() []byte {
	return t.mr
}

func (t *CommitteeTrie) Hash() []byte {
	return t.mr
}

func (t *CommitteeTrie) Add(s []byte) {
	t.list = append(t.list, TrieContent{x: string(s)})
	//Create a new Merkle Tree from the list of Content
	mt, err := merkletree.NewTree(t.list)
	if err != nil {
		log.Fatal(err)
	}
	t.mt = mt
	mr := mt.MerkleRoot()
	t.mr = mr
}

func NewTrie() CommitteeTrie {
	return CommitteeTrie{}
}

func (t *CommitteeTrie) VerifyTree() {
	//Verify the entire tree (hashes for each node) is valid
	vt, err := t.mt.VerifyTree()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Verify Tree: ", vt)
}

func (t *CommitteeTrie) VerifyContent(i int) {
	//Verify a specific content in in the tree
	vc, err := t.mt.VerifyContent(t.list[i])
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Verify Content: ", vc)
}
