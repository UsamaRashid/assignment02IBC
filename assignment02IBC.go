package assignment02IBC

// import "fmt"
import (
	"crypto/sha256"
	"fmt"
)

// package assignment02IBC

const miningReward = 100
const rootUser = "Satoshi"

type BlockData struct {
	Title    string
	Sender   string
	Receiver string
	Amount   int
}
type Block struct {
	Data        []BlockData
	PrevPointer *Block
	PrevHash    string
	CurrentHash string
}

func PremineChain(chainHead *Block, numBlocks int) *Block {

	// var premine = make([]BlockData, numBlocks)
	// var newBlock *Block
	// Basetransaction := []BlockData{}
	for i := 0; i < numBlocks; i++ {
		premine := []BlockData{{Title: "Premined", Sender: "nil", Receiver: "nil", Amount: 0}}
		chainHead = InsertBlock(premine, chainHead)
	}
	// fmt.Print("number of blocks: ", numBlocks)
	return chainHead
}

func CalculateBalance(userName string, chainHead *Block) int {
	var currentBalance int = 0
	for i := 0; chainHead != nil; i++ {
		// fmt.Print("\nCalculate balance Func Called, ", chainHead.Data, "\n")
		for i := range chainHead.Data {
			// currentBalance += 50
			if chainHead.Data[i].Sender == userName {
				// fmt.Print("\nSTRING FOUND in Sender!!!")
				currentBalance -= chainHead.Data[i].Amount
			}
			if chainHead.Data[i].Receiver == userName {
				// fmt.Print("\nSTRING FOUND in receiver!!!")
				currentBalance += chainHead.Data[i].Amount
			}
		}
		chainHead = chainHead.PrevPointer
	}
	// fmt.Print("\nCurrent Balance is : ", currentBalance, "\n")
	return currentBalance
}
func ListBlocks(chainHead *Block) {
	fmt.Print("\n############################################################################\n")
	fmt.Print("List Transactions (Most RECENT first) ")
	fmt.Print("\n############################################################################\n")
	fmt.Print("HEAD \n", "↓")
	for i := 0; chainHead != nil; i++ {
		// fmt.Print("############### New BLOCK ################################################", "\n")
		// fmt.Print("BLOCK", chainHead.Data, "\n")
		for j := range chainHead.Data {
			// fmt.Print("\none\n", chainHead.Data[j])
			fmt.Print("\n TRX.No : ", j)
			fmt.Print("  Title : ", chainHead.Data[j].Title)
			fmt.Print("  Sender : ", chainHead.Data[j].Sender)
			fmt.Print("  Receiver : ", chainHead.Data[j].Receiver)
			fmt.Print("  Amount : ", chainHead.Data[j].Amount)
			// fmt.Print("\n##################################################", "\n")
		}
		fmt.Print("\nPrev Hash: ", chainHead.PrevHash)
		fmt.Print("\nCurr Hash: ", chainHead.CurrentHash)
		fmt.Print("\n ↓")
		chainHead = chainHead.PrevPointer
	}
	fmt.Print("\n##################################################################################", "\n\n")
}
func CalculateHash(inputBlock *Block) string {
	// return ""
	var toReturn string
	for i := range inputBlock.Data {
		toReturn = toReturn + (inputBlock.Data[i].Receiver) + (inputBlock.Data[i].Sender) + (inputBlock.Data[i].Title) + fmt.Sprint((inputBlock.Data[i].Amount))
	}
	// fmt.Print("\nHASH TO_RETURN IS : ", toReturn, "\n")
	toReturn = fmt.Sprintf("%x", sha256.Sum256([]byte(toReturn)))
	return toReturn
}
func VerifyTransaction(transaction *BlockData, chainHead *Block) bool {
	// fmt.Print("VERIFY TRANSACTION CALLED!!")

	var balance int = CalculateBalance(transaction.Sender, chainHead)

	if transaction.Amount > balance {
		fmt.Print("ERROR:", transaction.Sender, " has ", balance, " coins -", transaction.Amount, " are needed!\n")
		return false
	}
	return true
}
func InsertBlock(blockData []BlockData, chainHead *Block) *Block {
	// fmt.Print("\nINSERT BLOCK CALLED\n", blockData[0].Amount, "  ", blockData[0].Receiver,
	// 	" ", blockData[0].Sender, " ", blockData[0].Title, " ", "  Root User:", rootUser)
	// fmt.Print("INSERT BLOCK ", rootUser)
	// CHECK IF TRANSACTION IS POSSIBLE OR NOT??
	// VerifyTransaction(blockData)
	for i := range blockData {
		if !VerifyTransaction(&blockData[i], chainHead) {
			fmt.Print("\nTransaction not possible!\n")
			return chainHead
		}

	}
	transaction := BlockData{Title: "Coinbase", Sender: "System", Receiver: rootUser, Amount: miningReward}
	blockData = append(blockData, transaction)
	// newBlock.Data = blockData

	if chainHead == nil {
		// If first Block
		//making a new Block to put in chain
		// fmt.Print("\nFIRST BLOCK ADDED \n")

		var newBlock Block
		newBlock.Data = blockData
		newBlock.PrevPointer = chainHead
		newBlock.PrevHash = ""
		chainHead = &newBlock
		newBlock.CurrentHash = CalculateHash(chainHead)
		// {dataToInsert, chainHead, "", CalculateHash(chainHead)}

		return &newBlock
	} else {
		// fmt.Print("\nInsert block Else\n")
		// newBlock := Block{dataToInsert, chainHead, CalculateHash(chainHead), "chainHead.PrevPointer.CurrentHash"}
		var newBlock Block
		newBlock.Data = blockData
		newBlock.PrevPointer = chainHead
		chainHead = &newBlock
		newBlock.PrevHash = chainHead.PrevPointer.CurrentHash
		newBlock.CurrentHash = CalculateHash(chainHead)
		// fmt.Print("\n123 ::::::: 456")

		return &newBlock

	}

	// return chainHead
}

func VerifyChain(chainHead *Block) {
	for i := 0; chainHead.PrevPointer != nil; i++ {
		var verifyHash string = CalculateHash(chainHead.PrevPointer)
		if verifyHash != chainHead.PrevHash {
			fmt.Print("BlockChain is comprised!")
			return
		}
		chainHead = chainHead.PrevPointer
	}
	println("BlockChain is verified!")
}
