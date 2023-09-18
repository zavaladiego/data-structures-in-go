package linkedlist

import (
	"reflect"
	"testing"
)

func TestNewNode(t *testing.T) {
	type args struct {
		data int
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{
			name: "Create New Node with data 13",
			args: args{data: 13},
			want: &Node{Data: 13, next: nil},
		},
		{
			name: "Create New Node with data 9",
			args: args{data: 9},
			want: &Node{Data: 9, next: nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNode(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewLinkedList(t *testing.T) {
	tests := []struct {
		name string
		want *LinkedList
	}{
		{
			name: "Create New Linked List",
			want: &LinkedList{head: nil, len: 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLinkedList(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLinkedList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_Prepend(t *testing.T) {
	type fields struct {
		head *Node
		len  int
	}
	type args struct {
		data int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Node
	}{
		{
			name:   "Prepend to empty linked list",
			fields: fields{head: nil, len: 0},
			args:   args{data: 10},
			want:   &Node{Data: 10, next: nil},
		},
		{
			name:   "Prepend to non-empty linked list",
			fields: fields{head: &Node{Data: 10, next: nil}, len: 1},
			args:   args{data: 20},
			want:   &Node{Data: 20, next: &Node{Data: 10, next: nil}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{
				head: tt.fields.head,
				len:  tt.fields.len,
			}
			if got := l.Prepend(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LinkedList.Prepend() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_Append(t *testing.T) {
	type fields struct {
		head *Node
		len  int
	}
	type args struct {
		data int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Node
	}{
		{
			name:   "Append to empty linked list",
			fields: fields{head: nil, len: 0},
			args:   args{data: 10},
			want:   &Node{Data: 10, next: nil},
		},
		{
			name:   "Append to non-empty linked list",
			fields: fields{head: &Node{Data: 10, next: nil}, len: 1},
			args:   args{data: 20},
			want:   &Node{Data: 20, next: nil},
		},
		{
			name:   "Append to non-empty linked list with multiple nodes",
			fields: fields{head: &Node{Data: 10, next: &Node{Data: 20, next: nil}}, len: 2},
			args:   args{data: 30},
			want:   &Node{Data: 30, next: nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{
				head: tt.fields.head,
				len:  tt.fields.len,
			}
			if got := l.Append(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LinkedList.Append() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_DeleteAt(t *testing.T) {
	type args struct {
		pos int
	}
	tests := []struct {
		name string
		l    *LinkedList
		args args
	}{
		{
			name: "Delete at position 1 (empty list)",
			l:    NewLinkedList(),
			args: args{pos: 1},
		},
		{
			name: "Delete at position 1",
			l: func() *LinkedList {
				l := NewLinkedList()
				l.Append(10)
				return l
			}(),
			args: args{pos: 1},
		},
		{
			name: "Delete at position 2",
			l: func() *LinkedList {
				l := NewLinkedList()
				l.Append(10)
				l.Append(20)
				return l
			}(),
			args: args{pos: 2},
		},
		{
			name: "Delete at position 3",
			l: func() *LinkedList {
				l := NewLinkedList()
				l.Append(10)
				l.Append(20)
				l.Append(30)
				return l
			}(),
			args: args{pos: 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.l.DeleteAt(tt.args.pos)
		})
	}
}

func TestLinkedList_Display(t *testing.T) {
	tests := []struct {
		name string
		l    *LinkedList
		want string
	}{
		{
			name: "Display Empty Linked List",
			l:    NewLinkedList(),
			want: "Linked list is empty",
		},
		{
			name: "Display Non-Empty Linked List",
			l: func() *LinkedList {
				l := NewLinkedList()
				l.Append(10)
				l.Append(20)
				return l
			}(),
			want: "10 -> 20 -> nil",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Display(); got != tt.want {
				t.Errorf("LinkedList.Display() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_Len(t *testing.T) {
	tests := []struct {
		name string
		l    *LinkedList
		want int
	}{
		{
			name: "Length of empty linked list",
			l:    NewLinkedList(),
			want: 0,
		},
		{
			name: "Length of linked list with one node",
			l: func() *LinkedList {
				l := NewLinkedList()
				l.Append(10)
				return l
			}(),
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Len(); got != tt.want {
				t.Errorf("LinkedList.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_Head(t *testing.T) {
	tests := []struct {
		name string
		l    *LinkedList
		want *Node
	}{
		{
			name: "Head of empty linked list",
			l:    NewLinkedList(),
			want: nil,
		},
		{
			name: "Head of linked list with one node",
			l: func() *LinkedList {
				l := NewLinkedList()
				l.Append(10)
				return l
			}(),
			want: &Node{Data: 10, next: nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Head(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LinkedList.Head() = %v, want %v", got, tt.want)
			}
		})
	}
}
