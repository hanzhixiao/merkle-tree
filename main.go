package main

import (
    "crypto/md5"
    "fmt"
)

var LeafNodeNumber int=8
var LeafNode=[] string {"1","2","3","4","5","6","7","8"}
//用数组代表Merkle树（没用结构体），数组最后一个值为root
var MerkleTree[15]string

func BuildMerkleTree(LeafNode []string,LeafNodeNumber int) [15]string {
	var MerkleTree[15]string
	for n:=0;n<LeafNodeNumber;n++{
		MerkleTree[n]=LeafNode[n]
	}
	var j int=0;
	var p int=LeafNodeNumber;
	for nsize := LeafNodeNumber;nsize>1
	{
		for i:=0;i<nsize;i+=2{
			var data[16] byte
			var i2 int
			if i+1>nsize-1{
				i2=nsize-1
			}else{
				i2=i+1
			}
			data=md5.Sum([]byte(MerkleTree[j+i]+MerkleTree[j+i2]))
			MerkleTree[p]=string(data[:])
			p=p+1
		}
		j+=nsize
		nsize=(nsize+1)/2
	}
    return MerkleTree
}


func APIgetmerklebranch(HashInput string) []string {
    var MerkleBrach[] string
    var index int
    //可以把叶子节点排序，用二分查找
    //这里只用了普通方式
    for index=0;index<len(MerkleTree);index++{
        if MerkleTree[index]==HashInput{
            break;
        }
    }
    MerkleBrach=append(MerkleBrach,MerkleTree[index])
    var j int=0
    for nsize:=LeafNodeNumber;nsize>1;{
        var i int
        if index^1>nsize-1{
			i=nsize-1
		}else{
			i=index^1
		}
		MerkleBrach=append(MerkleBrach,MerkleTree[j+i])
		index>>=1
		j+=nsize
        nsize=(nsize+1)/2
    }
    return MerkleBrach
}


func main() {
    MerkleTree=BuildMerkleTree(LeafNode,LeafNodeNumber)
    fmt.Print("叶子节点值：")
    for _, value := range LeafNode{
      fmt.Printf("%s ", value)
    }
    
    fmt.Print("\n")
    fmt.Print("树根：")
	fmt.Println(fmt.Sprintf("%x",MerkleTree[len(MerkleTree)-1]))
	
	
    fmt.Print("值为1的叶子节点的默克尔证明路径：")
	var MerkleBrach[] string=APIgetmerklebranch("1")
	for key, value := range MerkleBrach{
	  if key<2{
	      fmt.Println(value)
	      continue
	  }
      fmt.Println(fmt.Sprintf("%x",value))
    }
}
