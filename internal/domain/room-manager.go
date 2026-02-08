package domain

import (
	"errors"
	"sync"
)

type RoomManager struct {
	rooms map[string]*Room
	mutex sync.Mutex
}

func NewRoomManager() *RoomManager {
	return &RoomManager{
		rooms: make(map[string]*Room),
	}
}

func (rm *RoomManager) CreateRoom(roomId string) (*Room, error) {
	if _, exist := rm.rooms[roomId]; exist {
		return nil, errors.New("room already exists")
	}
	rm.mutex.Lock()
	rm.rooms[roomId] = NewRoom(roomId)
	rm.mutex.Unlock()
	return rm.rooms[roomId], nil
}

func (rm *RoomManager) GetRoom(roomId string) (*Room, error) {
	rm.mutex.Lock()
	room, exists := rm.rooms[roomId]
	rm.mutex.Unlock()
	if !exists {
		return nil, errors.New("room not found")
	}
	return room, nil
}

func (rm *RoomManager) DeleteRoom(roomId string) error {
	rm.mutex.Lock()
	if _, exist := rm.rooms[roomId]; !exist {
		return errors.New("error delete room\nroom does not exist")
	}
	delete(rm.rooms, roomId)
	rm.mutex.Unlock()
	return nil
}
