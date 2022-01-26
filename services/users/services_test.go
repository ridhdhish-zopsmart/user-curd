package users

import (
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"

	"layer/user/models"
	"layer/user/stores"
)

func TestGetUserById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserStore := stores.NewMockUser(ctrl)
	testUserService := New(mockUserStore)

	tests := []struct {
		desc     string
		id       int
		expected models.User
		mockCall *gomock.Call
	}{
		{
			desc:     "Case1",
			id:       1,
			expected: models.User{Id: 1, Name: "Naruto", Email: "naruto@japan.com", Phone: "9999999999", Age: 18},
			mockCall: mockUserStore.EXPECT().GetUserById(1).Return(&models.User{Id: 1, Name: "Naruto", Email: "naruto@japan.com", Phone: "9999999999", Age: 18}, nil),
		},
		{
			desc:     "Case2",
			id:       2,
			expected: models.User{},
			mockCall: mockUserStore.EXPECT().GetUserById(2).Return(&models.User{}, errors.New("Cannot fetch user for given id")),
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			user, err := testUserService.GetUserById(test.id)

			if err != nil && !reflect.DeepEqual(test.expected, user) {
				t.Errorf("Expected: %v, Got: %v", test.expected, user)
			}
		})
	}
}

func TestGetUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserStore := stores.NewMockUser(ctrl)
	testUserService := New(mockUserStore)

	data1 := []models.User{
		{Id: 1, Name: "Naruto", Email: "naruto@gmail.com", Phone: "9999999999", Age: 18},
		{Id: 2, Name: "Itachi", Email: "itachi@gmail.com", Phone: "8320578360", Age: 24},
	}

	tests := []struct {
		desc     string
		expected []models.User
		mockCall *gomock.Call
	}{
		{
			desc:     "Case1",
			expected: data1,
			mockCall: mockUserStore.EXPECT().GetUsers().Return(data1, nil),
		},
		{
			desc:     "Case2",
			expected: []models.User{},
			mockCall: mockUserStore.EXPECT().GetUsers().Return([]models.User{}, errors.New("Cannot fetch users")),
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			users, err := testUserService.GetUsers()

			if err != nil && !reflect.DeepEqual(test.expected, users) {
				t.Errorf("Expected: %v, Got: %v", test.expected, users)
			}
		})
	}
}

func TestUpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserStore := stores.NewMockUser(ctrl)
	testUserService := New(mockUserStore)

	testUser := models.User{Name: "Ridhdhish", Email: "ridhdhish@gmail.com", Phone: "8320578360", Age: 21}

	tests := []struct {
		desc     string
		id       int
		expected int
		mockCall *gomock.Call
	}{
		{
			desc:     "Case1",
			id:       1,
			expected: 1,
			mockCall: mockUserStore.EXPECT().UpdateUser(1, testUser).Return(1, nil),
		},
		{
			desc:     "Case2",
			id:       2,
			expected: 0,
			mockCall: mockUserStore.EXPECT().UpdateUser(2, testUser).Return(0, errors.New("Invalid id")),
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			lastInsertedId, _ := testUserService.UpdateUser(test.id, testUser)

			if lastInsertedId != test.expected {
				t.Errorf("Expected: %v, Got: %v", test.expected, lastInsertedId)
			}
		})
	}
}

func TestDeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserStore := stores.NewMockUser(ctrl)
	testUserService := New(mockUserStore)

	tests := []struct {
		desc     string
		id       int
		expected int
		mockCall *gomock.Call
	}{
		{
			desc:     "Case1",
			id:       1,
			expected: 1,
			mockCall: mockUserStore.EXPECT().DeleteUser(1).Return(1, nil),
		},
		{
			desc:     "Case2",
			id:       2,
			expected: 0,
			mockCall: mockUserStore.EXPECT().DeleteUser(2).Return(0, errors.New("Invalid id")),
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			rowsAffected, _ := testUserService.DeleteUser(test.id)

			if rowsAffected != test.expected {
				t.Errorf("Expected: %v, Got: %v", test.expected, rowsAffected)
			}
		})
	}
}
