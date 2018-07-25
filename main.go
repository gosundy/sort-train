package main

import (
	"fmt"
)
type BinaryTreeNode struct{
	data int
	left *BinaryTreeNode
	right *BinaryTreeNode
	leftHight int
	rightHight int
}
func main(){
	//data:=[]int{0,1,2,3,4,5,6,7}
	//fmt.Println(zheBanChanZhao(data,8))
	//该data可以测试ll,lr,rl,rr
	//data:=[]int{16,3,7,11,9,26,18,14,15}
	data:=[]int{1}
	//data:=[]int{4,2,5,1,3}
	//quickSort(data,0,len(data)-1)
	//bubbleSort(data)
	//insertSort(data)
	//selectSort(data)
	//xiErSort(data,len(data)/2)
	//bucketSort(data,7)
	//merge(data,2)
	//duiSort(data)
	//zuiXiaoNgeShu(data,5)
	//tmp:=make([]int,len(data))
	//erluguibingSort(data,tmp,0,len(data))
	//waiPaiSort(data,3)
	//initDui(data)
	//fmt.Println(data)
	//var root *BinaryTreeNode
	//bianLiBinaryTree(root)
	//root,_:=createBinaryTree(data,0,-1)
	//qianXubianLiBinaryTree(root)
	//fmt.Println("zhongxu")
	//zhongXuBianLiBinaryTree(root)
	//fmt.Println("houxu")
	root:=createSortedBinaryTree(data)
	//InsertSortedBinaryTree(root,9)
	//InsertSortedBinaryTree(root,0)
	//sortedBinaryLayer(root)
	fmt.Println("中序遍历")
	zhongXuBianLiBinaryTree(root)
	fmt.Println("中序遍历结束")
	DeleteSortedBinaryTree(&root,1)

	sortedBinaryLayer(root)

	fmt.Println("删除节点")
	zhongXuBianLiBinaryTree(root)
	fmt.Println("删除节点结束")
}
func zheBanChanZhao(data []int,findData int)int{
	low,high:=0,len(data)-1
	mid:=0
	for low<=high{
		mid=(low+high)/2
		fmt.Println(mid)
		if data[mid]==findData{
			return mid
		}
		if data[mid]<findData{
		    low=mid+1
		}else{
			high=mid-1
		}
	}
	return -1
}
func quickSort(data []int,low,high int){
	if low>=high{
		return
	}
	i,j:=low,high
	zhongzhou:=data[low]
	for i!=j{
		for data[j]>=zhongzhou&&j!=i{
			j--
		}
		if j==i{
			data[j]=zhongzhou
			break
		}
		data[i]=data[j]
		i++
		for data[i]<=zhongzhou&&j!=i{
			i++
		}
		if j==i{
			data[j]=zhongzhou
			break
		}
		data[j]=data[i]
		j--
	}
	quickSort(data,low,i-1)
	quickSort(data,i+1,high)

}
func bubbleSort(data []int){
	lenData:=len(data)
	for i:=lenData-1;i>=0;i--{
		for j:=1;j<=i;j++{
			if data[j-1]>data[j]{
				tmp:=data[j-1]
				data[j-1]=data[j]
				data[j]=tmp
			}
		}
	}
}
func insertSort(data []int){
	lenData:=len(data)
	j:=0
	for i:=1;i<lenData;i++{
		tmp:=data[i]
		for j=i;j>0&&data[j-1]>tmp;j--{
				data[j]=data[j-1]
		}
		data[j]=tmp
	}
}
func selectSort(data []int){
	lenData:=len(data)
	maxIndex:=0
	for i:=lenData-1;i>=1;i--{
		maxIndex=0
		for j:=1;j<=i;j++{
			if data[maxIndex]<data[j]{
				maxIndex=j
			}
		}
		tmp:=data[maxIndex]
		data[maxIndex]=data[i]
		data[i]=tmp
	}
}
func xiErSort(data[]int,d int){
	if d==0{
		return
	}
	lenData:=len(data)
	j:=0
	for i:=0;i<d;i++{
		for k:=i+d;k<lenData;k=k+d{
			tmp:=data[k]
			for j=k;(j-d)>=0&&data[j-d]>tmp;j=j-d{
				data[j]=data[j-d]
			}
			data[j]=tmp
		}
	}
	xiErSort(data,d/2)
}
func bucketSort(data[]int,d int){

	k:=1
	lenData:=len(data)
	for i:=0;i<d;i++{
		bucket:=make(map[int][]int)
		for j:=0;j<lenData;j++{

			m:=data[j]/k
			mm:=m%10
			bucket[mm]=append(bucket[mm],data[j])
		}
		m1:=0
		for i1:=0;i1<10;i1++{
			for _,j1:=range bucket[i1]{
				data[m1]=j1
				m1=m1+1
			}
		}
		k=k*10
	}

}
func merge(data[]int,d int){
	lenData:=len(data)
	if d>lenData{
		return
	}
	for i:=d;i<lenData;i=i+d{
		insertSort(data[i-d:i])
	}
	m:=lenData%d
	insertSort(data[lenData-m:lenData])
	merge(data,d*2)
}

func duiSort(data[]int){
	duiTiaoZheng(data)
	maxIndex:=len(data)-1
	for maxIndex>1{
		tmp:=data[1]
		data[1]=data[maxIndex]
		data[maxIndex]=tmp
		maxIndex=maxIndex-1
		shiftMax(data,1,maxIndex)

	}
}
func shiftMax(data[]int,parent int,maxIndex int){
	for parent<=maxIndex{
		left:=parent*2
		right:=left+1
		tmpMax:=0
		if right<=maxIndex{
			if data[right]<data[left]{
				tmpMax=left
			}else{
				tmpMax=right
			}
		}else if left<=maxIndex{
			tmpMax=left
		}else{
			break
		}
		if data[parent]<data[tmpMax]{
			tmp:=data[parent]
			data[parent]=data[tmpMax]
			data[tmpMax]=tmp
			parent=tmpMax
		}else{
			break
		}

	}
}
func shiftMin(data[]int,parent int,maxIndex int){
	for parent<=maxIndex{
		left:=parent*2
		right:=left+1
		tmpMin:=0
		if right<=maxIndex{
			if data[right]>data[left]{
				tmpMin=left
			}else{
				tmpMin=right
			}
		}else if left<=maxIndex{
			tmpMin=left
		}else{
			break
		}
		if data[parent]>data[tmpMin]{
			tmp:=data[parent]
			data[parent]=data[tmpMin]
			data[tmpMin]=tmp
			parent=tmpMin
		}else{
			break
		}

	}
}
func duiTiaoZheng(data []int){
	maxIndex:=len(data)-1
	firstParent:=maxIndex/2
	//fmt.Println(firstParent)
	for firstParent!=0{

		parent:=firstParent
		shiftMax(data,parent,maxIndex)
		firstParent=firstParent-1

	}
}
func initDui(data[]int){
	maxIndex:=len(data)-1
	firstParent:=maxIndex/2
	for firstParent!=0{

		parent:=firstParent
		shiftMin(data,parent,maxIndex)
		firstParent=firstParent-1
	}
}
func zuiXiaoNgeShu(data[]int,n int){
	dui:=make([]int,n+1)
	copy(dui[1:],data)
	duiTiaoZheng(dui)
	for i:=n;i<len(data);i++{
		if data[i]<dui[1]{
			dui[1]=data[i]
			shiftMax(dui,1,n)
		}
	}
	fmt.Println(dui)
}
func erluguibingSort(data ,tmp[]int,start,end int)  {
	if (end-start)<=1{
		tmp[start]=data[start]
		return
	}
	mid:=(start+end)/2
	erluguibingSort(data,tmp,start,mid)
	erluguibingSort(data,tmp,mid,end)

	start1:=start
	start2:=mid
	index:=start
	for start1<mid&&start2<end{
		if data[start1]<data[start2]{
			tmp[index]=data[start1]
			index++
			start1++
		}else{
			tmp[index]=data[start2]
			index++
			start2++
		}
	}
	for start1<mid{
		tmp[index]=data[start1]
		start1++
		index++
	}
	for start2<end{
		tmp[index]=data[start2]
		start2++
		index++
	}
	//将拍好序的tmp复制一份给data，使这一阶段的序列有序
	copy(data[start:end],tmp[start:end])
}
//外排序
func waiPaiSort(data[]int,n int){
	//小顶堆
	dui:=make([]int,n+1)
	copy(dui[1:],data)
	index:=n
	for index<len(data){
		maxIndex:=n
		initDui(dui)
		tmpData:=make([]int,0)
		for _,v:=range data[index:]{
			if v>dui[1]{
				tmpData=append(tmpData,dui[1])
				dui[1]=v
				shiftMin(dui,1,maxIndex)
			}else if v==dui[1]{
				tmpData=append(tmpData,dui[1])
			}else{
				tmpData=append(tmpData,dui[1])
				dui[1]=dui[maxIndex]
				dui[maxIndex]=v
				maxIndex=maxIndex-1
				if maxIndex==0{
					fmt.Println(tmpData)
					break
				}
				shiftMin(dui,1,maxIndex)
			}
			index++
			if index==len(data){
				fmt.Println(tmpData)
				tmpData=make([]int,0)
				initDui(dui)
				maxIndex=n
				for maxIndex!=0{
					tmpData=append(tmpData,dui[1])
					dui[1]=dui[maxIndex]
					maxIndex=maxIndex-1
					shiftMin(dui,1,maxIndex)
				}
				fmt.Println(tmpData)

			}

		}
	}

}
//创建二叉树
func createBinaryTree(data[]int,index int,stopFlag int)(*BinaryTreeNode,int){
	if index>=len(data)||data[index]==stopFlag{
		return nil,index
	}
	node:=new(BinaryTreeNode)
	node.data=data[index]
	node.left,index=createBinaryTree(data,index+1,stopFlag)
	node.right,index=createBinaryTree(data,index+1,stopFlag)
	return node,index
}
//二叉排序树创建
func createSortedBinaryTree(data[]int) *BinaryTreeNode{
	if data==nil{
		return nil
	}
	root:=new(BinaryTreeNode)
	root.data=data[0]
	root.leftHight=0
	root.rightHight=0
	for _,v:=range data[1:]{
		InsertSortedBinaryTree(root,v)
		//建立平衡因子
		sortedBinaryLayer(root)
		//zhongXuBianLiBinaryTree(root)
		//检查是否平衡
		//leftOrRight用来判断node是nodeParent的左孩子还是右孩子
		node,nodeParent,flag,leftOrRight:=sortedBinaryTreeIsBalance(root,nil,0)
		fmt.Println("添加节点",v,"是否平衡",flag)

		//出现不平衡进行调整
		if !flag{
			fmt.Println("不平衡节点",node.data)
			//分四种情况
			//LL型
			if (node.leftHight-node.rightHight)==2&&(node.left.leftHight-node.left.rightHight)==1{
				//分离出，数值第二大节点的右指针，让其移动到调整后的右节点的最左边，因为它比
				//第二大节点大，比第一大节点小
				tmpRight:=node.left.right
				//调整三个节点的位置
				//如果父节点为nil,则表示是root节点不平衡
				if nodeParent==nil{
					//data:=[]int{6,7,4,3,5,1}
					node.left.right=root
					root=node.left
					node.left=nil
					if tmpRight!=nil{
						tmpNode:=root.right
						for tmpNode.left!=nil{
							tmpNode=tmpNode.left
						}
						tmpNode.left=tmpRight
					}

				}else{
					//如果node节点是nodeParent节点的左子节点
					//data:=[]int{10,8,11,6,9,12,5,7,4}
					if leftOrRight==1{

						nodeParent.left=node.left
						nodeParent.left.right=node
						node.left=nil
						if tmpRight!=nil{
							tmpNode:=nodeParent.left.right
							for tmpNode.left!=nil{
								tmpNode=tmpNode.left
							}
							tmpNode.left=tmpRight
						}
					}else if leftOrRight==2{

						//data:=[]int{4,3,8,6,5}
						nodeParent.right=node.left
						nodeParent.right.right=node
						node.left=nil
						if tmpRight!=nil{
							tmpNode:=nodeParent.right.right
							for tmpNode.left!=nil{
								tmpNode=tmpNode.left
							}
							tmpNode.left=tmpRight
						}
					}
				}
				fmt.Println("平衡后中序遍历")
				zhongXuBianLiBinaryTree(root)
				fmt.Println("平衡后中序遍历结束")
				continue

			}
			//RR型
			if (node.leftHight-node.rightHight)==-2&&(node.right.leftHight-node.right.rightHight)==-1{
				tmpLeft:=node.right.left
				//如果不平衡节点为root节点，则其nodeParent为nil
				if nodeParent==nil{
					//data:=[]int{1,2,3}
					root=node.right
				}else{
					//如果nodeParent的左节点不平衡
					if leftOrRight==1{
						//data:=[]int{8,4,9,5,6}
						nodeParent.left=node.right
						//如果nodeParent的右节点不平衡
						//data:=[]int{6,5,7,8,9}
					}else if leftOrRight==2{
						nodeParent.right=node.right
					}
				}
				node.right.left=node
				node.right=tmpLeft
				fmt.Println("平衡后中序遍历")
				zhongXuBianLiBinaryTree(root)
				fmt.Println("平衡后中序遍历结束")
				continue
			}

			//LR型

			if (node.leftHight-node.rightHight)==2&&(node.left.leftHight-node.left.rightHight)==-1{
				tmpLeft:=node.left.right.left
				tmpRight:=node.left.right.right
				if nodeParent!=nil{
					fmt.Println("LR",node.data,nodeParent.data)
					zhongXuBianLiBinaryTree(root)
					qianXubianLiBinaryTree(root)
					fmt.Println("LEFTORRIGHT",leftOrRight)
				}

				//如果nodeParent为nil，则表示root节点不平衡
				if nodeParent==nil{
					root=node.left.right
					root.left=node.left
					root.right=node
					node.left.right=nil
					node.left=nil
					tmpLeftNode:=root.left
					for tmpLeftNode.right!=nil{
						tmpLeftNode=tmpLeftNode.right
					}
					tmpLeftNode.right=tmpLeft
					tmpRightNode:=root.right
					for tmpRightNode.left!=nil{
						tmpRightNode=tmpRightNode.left
					}
					tmpRightNode.left=tmpRight
				}else{
					//如果nodeParent的左节点为不平衡节点
					if leftOrRight==1{
						//data:=[]int{8,6,10,2,7,9,11,1,4,12,3}
						fmt.Println("12312312312312")
						fmt.Println(nodeParent.data)
						nodeParent.left=node.left.right
						nodeParent.left.left=node.left
						nodeParent.left.right=node
						node.left.right=nil
						node.left=nil
						tmpLeftNode:=nodeParent.left.left
						for tmpLeftNode.right!=nil{
							tmpLeftNode=tmpLeftNode.right
						}
						tmpLeftNode.right=tmpLeft
						tmpRightNode:=nodeParent.left.right
						for tmpRightNode.left!=nil{
							tmpRightNode=tmpRightNode.left
						}
						tmpRightNode.left=tmpRight
					}else if leftOrRight==2{
						fmt.Println("22222222222222222")
						nodeParent.right=node.left.right
						nodeParent.right.left=node.left
						nodeParent.right.right=node
						node.left.right=nil
						node.left=nil
						tmpLeftNode:=nodeParent.right.left
						for tmpLeftNode.right!=nil{
							tmpLeftNode=tmpLeftNode.right
						}
						tmpLeftNode.right=tmpLeft
						tmpRightNode:=nodeParent.right.right
						for tmpRightNode.left!=nil{
							tmpRightNode=tmpRightNode.left
						}
						tmpRightNode.left=tmpRight
					}

				}
				fmt.Println("平衡后中序遍历")
				zhongXuBianLiBinaryTree(root)
				fmt.Println("平衡后中序遍历结束")
				continue

			}

			//RL型
			if (node.leftHight-node.rightHight)==-2&&(node.right.leftHight-node.right.rightHight)==1{
				tmpLeft:=node.right.left.left
				tmpRight:=node.right.left.right

				if nodeParent==nil{
					//data:=[]int{1,3,2}
					root=node.right.left
					root.left=node
					root.right=node.right
					node.right.left=nil
					node.right=nil
					tmpNode:=root.left
					for tmpNode.right!=nil{
						tmpNode=tmpNode.right
					}
					tmpNode.right=tmpLeft
					tmpNode=root.right
					for tmpNode.left!=nil{
						tmpNode=tmpNode.left
					}
					tmpNode.left=tmpLeft
				}else{
					if leftOrRight==1{
						//data:=[]int{4,1,5,3,2}
						nodeParent.left=node.right.left
						nodeParent.left.left=node
						nodeParent.left.right=node.right
						node.right.left=nil
						node.right=nil
						nodeParent.left.left.right=tmpLeft
						nodeParent.left.right.left=tmpRight
					}else if leftOrRight==2{
						//data:=[]int{2,1,5,7,6}
						nodeParent.right=node.right.left
						nodeParent.right.left=node
						nodeParent.right.right=node.right
						node.right.left=nil
						node.right=nil
						nodeParent.right.left.right=tmpLeft
						nodeParent.right.right.left=tmpRight
					}
				}
			}
			fmt.Println("平衡后中序遍历")
			zhongXuBianLiBinaryTree(root)
			fmt.Println("平衡后中序遍历结束")
			continue
		}
		fmt.Println("平衡后中序遍历")
		zhongXuBianLiBinaryTree(root)
		fmt.Println("平衡后中序遍历结束")
	}

	//重新建立平衡
	sortedBinaryLayer(root)

	return root
}
//二叉排序树的节点插入
func InsertSortedBinaryTree(root *BinaryTreeNode,val int){
	node:=root
	for node!=nil{
		if val>=node.data{
			if node.right!=nil{
				node=node.right
			}else{
				node.right=&BinaryTreeNode{data:val}
				break
			}
		}else{
			if node.left!=nil{
				node=node.left

			}else{
				node.left=&BinaryTreeNode{data:val}
				break
			}
		}

	}
}
//平衡二叉树的节点删除
func DeleteSortedBinaryTree(root **BinaryTreeNode,val int)(deletedVal int,status bool){
	node:=*root
	var parent *BinaryTreeNode
	leftOrRight:=0
	for node!=nil{
		if val==node.data{
			if node.left!=nil{
				if node.left.right==nil{
					node.data=node.left.data
					node.left=nil
				}else{
					tmpNode:=node.left
					for tmpNode.right.right!=nil{
						tmpNode=tmpNode.right
					}
					node.data=tmpNode.right.data
					tmpNode.right=nil
				}
			}else if node.right!=nil{
				if node.right.left==nil{
					node.data=node.right.data
					node.right=nil
				}else{
					tmpNode:=node.right
					for tmpNode.left.left!=nil{
						tmpNode=tmpNode.left
					}
					node.data=tmpNode.left.data
					tmpNode.left=nil
				}
				return val,true
			}else{
				if leftOrRight==0{
					*root=nil
				}
				if leftOrRight==1{
					parent.left=nil
				}
				if leftOrRight==2{
					parent.right=nil
				}
			}
			return  val,true
		}else{
			if val<node.data{
				parent=node
				leftOrRight=1
				node=node.left
			}
			if val>node.data{
				leftOrRight=2
				parent=node
				node=node.right
			}
		}

	}
	return val,false
}
//每个节点建立左子树层数和右子树层数
func sortedBinaryLayer(root *BinaryTreeNode){
	sortedBinaryLayer2(root)
}
//后序遍历建立层数关系
func sortedBinaryLayer2(node *BinaryTreeNode)(leftLayer,rightLayer int){
	if node==nil{
		return -1,-1
	}
	l_leftLayer,l_rightLayer:=sortedBinaryLayer2(node.left)
	r_leftLayer,r_rightLayer:=sortedBinaryLayer2(node.right)
	if l_leftLayer>l_rightLayer{
		leftLayer=l_leftLayer+1
	}else{
		leftLayer=l_rightLayer+1
	}
	if r_leftLayer>r_rightLayer{
		rightLayer=r_leftLayer+1
	}else{
		rightLayer=r_rightLayer+1
	}
	node.leftHight=leftLayer
	node.rightHight=rightLayer
	return leftLayer,rightLayer
}
//判断二叉排序树是否平衡，使用后序遍历，可以找到最近的不平衡点，返回不平衡节点和该节点的父节点
func sortedBinaryTreeIsBalance(node *BinaryTreeNode,nodeParent *BinaryTreeNode,leftOrRight int)(noBalanceNode *BinaryTreeNode,noBalanceNodepParent *BinaryTreeNode,flag bool,leftOrRightTmp int){
	if node==nil{
		return nil,nil,true,0
	}
	leftNode,leftNodeParent,leftFlag,left:=sortedBinaryTreeIsBalance(node.left,node,1)
	rightNode,rightNodeParent,rightFlag,right:=sortedBinaryTreeIsBalance(node.right,node,2)
	if !leftFlag{

		return leftNode,leftNodeParent,leftFlag,left
	}
	if !rightFlag{
		return rightNode,rightNodeParent,rightFlag,right
	}
	if ((node.rightHight-node.leftHight)>1) || ((node.rightHight-node.leftHight)< -1){

		return node,nodeParent,false,leftOrRight
	}

	return nil,nil,true,0
}
//前序遍历
func qianXubianLiBinaryTree(treeNode *BinaryTreeNode)  {
	if treeNode==nil{
		return
	}
	fmt.Println(treeNode.data)
	qianXubianLiBinaryTree(treeNode.left)
	qianXubianLiBinaryTree(treeNode.right)
}
//中序遍历
func zhongXuBianLiBinaryTree(treeNode *BinaryTreeNode){
	if treeNode==nil{
		return
	}
	zhongXuBianLiBinaryTree(treeNode.left)
	fmt.Println(treeNode.data,treeNode.leftHight,treeNode.rightHight)
	zhongXuBianLiBinaryTree(treeNode.right)
}
//后序遍历
func houXuBianLiBinaryTree(treeNode *BinaryTreeNode){
	if treeNode==nil{
		return
	}
	houXuBianLiBinaryTree(treeNode.left)
	houXuBianLiBinaryTree(treeNode.right)
	fmt.Println(treeNode.data)
}



