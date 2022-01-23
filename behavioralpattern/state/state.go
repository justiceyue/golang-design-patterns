package state

import (
	"fmt"
	"sync"
)

/*
允许一个对象在其内部状态发生改变的时候改变他的行为。
该设计模式通过维护和切换状态。可避免过多的if-else判断。
*/

type Vote interface {
	Vote(voter, voteItem string, vm *VoteManager)
}

type VoteManager struct {
	//key: voter
	//value: voteItems
	//len(value): voteCount
	VoteMessage map[string][]string
	vote        Vote
}

var (
	once sync.Once
	vm   *VoteManager
)

func getVoteManagerInstance() *VoteManager {
	once.Do(func() {
		vm = &VoteManager{
			VoteMessage: make(map[string][]string),
			//初始化的状态
			vote: &NomalVote{},
		}
	})
	return vm
}

func (v *VoteManager) Vote(voter, voteItem string) {
	/*
		voteCount := len(v.VoteMessage[voter])
		比较传统的做法，虽然是引入状态模式让扩展更加方便但是还需要有if-else判断
		if voteCount == 0 {
			v.vote = NomalVote{}
		}
		if voteCount > 0 && voteCount <= 3 {
			v.vote = RepeatedVote{}
		}
		if voteCount > 3 {
			v.vote = BlackListVote{}
		}
	*/
	v.vote.Vote(voter, voteItem, v)
}

type NomalVote struct {
}

func (NomalVote) Vote(voter, voteItem string, vm *VoteManager) {
	vm.VoteMessage[voter] = append(vm.VoteMessage[voter], voteItem)
	fmt.Println("投票成功")
	//状态的切换
	vm.vote = RepeatedVote{}
}

type RepeatedVote struct{}

func (RepeatedVote) Vote(voter, voteItem string, vm *VoteManager) {
	vm.VoteMessage[voter] = append(vm.VoteMessage[voter], voteItem)
	fmt.Println("请不要重复投票,该次投票无效")
	if len(vm.VoteMessage[voter]) >= 3 {
		//状态的切换，将if分散开
		vm.vote = BlackListVote{}
	}
}

type BlackListVote struct{}

func (BlackListVote) Vote(voter, voteItem string, vm *VoteManager) {
	vm.VoteMessage[voter] = append(vm.VoteMessage[voter], voteItem)
	fmt.Println("该用户已被加入黑名单")
}
