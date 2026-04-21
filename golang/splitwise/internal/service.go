package internal

import "fmt"

func NewSplitService() *SplitService {
	return &SplitService{
		balances: make(map[int]map[int]float64),
		groups:   make(map[int]*Group),
		expenses: make(map[int]*Expense),
		users:    make(map[int]*User),
	}
}

func (s *SplitService) AddUser(id int, name string, phone string, email string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.users[id]; exists {
		return fmt.Errorf("user with id: %d already exists", id)
	}

	s.users[id] = &User{
		Id:    id,
		Name:  name,
		Phone: phone,
		Email: email,
	}
	return nil
}

func (s *SplitService) AddGroup(name string, members []int) int {
	s.mu.Lock()
	defer s.mu.Unlock()

	memberStruct := make(map[int]struct{})
	for _, member := range members {
		memberStruct[member] = struct{}{}
	}

	groupId := len(s.groups) + 1
	s.groups[groupId] = &Group{
		Id:      groupId,
		Name:    name,
		Members: memberStruct,
	}
	return groupId
}

func (s *SplitService) AddExpense(groupId int, description string, amount float64, paidBy int, splitType string, splits []Split) int {
	s.mu.Lock()
	defer s.mu.Unlock()

	expenseId := len(s.expenses) + 1
	s.expenses[expenseId] = &Expense{
		Id:          expenseId,
		Description: description,
		Amount:      amount,
		PaidBy:      paidBy,
		SplitType:   splitType,
		Split:       splits,
	}

	switch splitType {
	case "equal":
		splitAmount := amount / float64(len(splits))
		for _, split := range splits {
			if split.UserId == paidBy {
				continue
			}
			s.updateBalance(split.UserId, paidBy, splitAmount)
		}
	case "exact":
		for _, split := range splits {
			if split.UserId == paidBy {
				continue
			}
			s.updateBalance(split.UserId, paidBy, split.Amount)
		}
	}

	return expenseId
}

func (s *SplitService) AddUserToGroup(groupId int, userId int) error {
	s.mu.RLock()

	if _, exists := s.users[userId]; !exists {
		return fmt.Errorf("user with id: %d does not exists", userId)
	}

	group, exists := s.groups[groupId]
	if !exists {
		return fmt.Errorf("group with id: %d does not exists", groupId)
	}
	s.mu.Unlock()

	group.mu.Lock()
	defer group.mu.Unlock()

	_, exists = group.Members[userId]
	if exists {
		return fmt.Errorf("User with id %d already exists in group", userId)
	}

	group.Members[userId] = struct{}{}
	return nil
}

func (s *SplitService) updateBalance(fromUserId, toUserId int, amount float64) {
	if s.balances[fromUserId] == nil {
		s.balances[fromUserId] = make(map[int]float64)
	}
	s.balances[fromUserId][toUserId] += amount

	if s.balances[toUserId] == nil {
		s.balances[toUserId] = make(map[int]float64)
	}
	s.balances[toUserId][fromUserId] -= amount
}

func (s *SplitService) GetBalance(userId int) map[int]float64 {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if _, exists := s.users[userId]; !exists {
		return nil
	}

	if balances, exists := s.balances[userId]; exists {
		return balances
	}
	return nil
}
