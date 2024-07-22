package set

import (
	"math"
	"testing"
)

func TestSet_Equals(t *testing.T) {
	tests := []struct {
		name string
		self Set[int]
		args Set[int]
		want bool
	}{
		{
			name: "same",
			self: New(1, 2, 3, 42),
			args: New(1, 2, 3, 42),
			want: true,
		},
		{
			name: "different",
			self: New(1, 2, 3, 42),
			args: New(1, 5, 3, 42),
			want: false,
		},
		{
			name: "nil",
			self: New(1, 2, 3, 42),
			args: nil,
			want: false,
		},
		{
			name: "length",
			self: New(1, 2, 3, 42),
			args: New(1, 2, 3),
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.self.Equals(tt.args); got != tt.want {
				t.Errorf("Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Contains(t *testing.T) {
	tests := []struct {
		name string
		self Set[int]
		args []int
		want bool
	}{
		{
			name: "some",
			self: New(1, 2, 3, 42),
			args: []int{1, 2, 3},
			want: true,
		},
		{
			name: "all",
			self: New(1, 2, 3, 42),
			args: []int{1, 2, 3, 42},
			want: true,
		},
		{
			name: "none",
			self: New(1, 2, 3, 42),
			args: nil,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.self.Contains(tt.args...); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_ContainsAny(t *testing.T) {
	tests := []struct {
		name string
		self Set[int]
		args []int
		want bool
	}{
		{
			name: "some",
			self: New(1, 2),
			args: []int{1, 2, 3, 42},
			want: true,
		},
		{
			name: "all",
			self: New(1, 2, 3, 42),
			args: []int{1, 2, 3, 42},
			want: true,
		},
		{
			name: "none",
			self: New(1, 2, 3, 42),
			args: []int{9001, 24, 32},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.self.ContainsAny(tt.args...); got != tt.want {
				t.Errorf("ContainsAny() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_IsSubSet(t *testing.T) {
	tests := []struct {
		name string
		self Set[int]
		arg  Set[int]
		want bool
	}{
		{
			name: "is",
			self: New(1, 2, 3),
			arg:  New(1, 2, 3, 4, 5),
			want: true,
		},
		{
			name: "not",
			self: New(1, 15, 2, 3),
			arg:  New(1, 2, 3, 4, 5),
			want: false,
		},
		{
			name: "empty lhs",
			self: New[int](),
			arg:  New(1, 2, 3, 4, 5),
			want: true,
		},
		{
			name: "empty rhs",
			self: New(1, 2, 3),
			arg:  New[int](),
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.self.IsSubSet(tt.arg); got != tt.want {
				t.Errorf("IsSubSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_IsSuperSet(t *testing.T) {
	tests := []struct {
		name string
		self Set[int]
		arg  Set[int]
		want bool
	}{
		{
			name: "is",
			self: New(1, 2, 3, 4, 5),
			arg:  New(1, 2, 3),
			want: true,
		},
		{
			name: "not",
			self: New(1, 2, 3, 4, 5),
			arg:  New(1, 2, 3, 15),
			want: false,
		},
		{
			name: "empty lhs",
			self: New[int](),
			arg:  New(1, 2, 3, 15),
			want: false,
		},
		{
			name: "empty rhs",
			self: New(1, 2, 3, 4, 5),
			arg:  New[int](),
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.self.IsSuperSet(tt.arg); got != tt.want {
				t.Errorf("IsSubSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Copy(t *testing.T) {
	oldSet := New(1, 2, 3, 42)
	newSet := oldSet.Copy()

	t.Run("equals", func(t *testing.T) {
		if !oldSet.Equals(newSet) {
			t.Errorf("Copy(), new set is not equal to old set, got %v want %v", newSet, oldSet)
		}
	})

	t.Run("noref", func(t *testing.T) {
		oldSet.Add(24)
		if newSet.Contains(24) {
			t.Errorf("Copy(), mutation found newset is not a copy")
		}
	})
}

func TestSet_Merge(t *testing.T) {
	tests := []struct {
		name  string
		self  []int
		other []int
		want  []int
	}{
		{
			name:  "some",
			self:  []int{1, 2, 3},
			other: []int{42, 9001},
			want:  []int{1, 2, 3, 42, 9001},
		},
		{
			name:  "none",
			self:  []int{1, 2, 3},
			other: nil,
			want:  []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			self := New(tt.self...)
			other := New(tt.other...)
			want := New(tt.want...)

			self.Merge(other)

			if !self.Equals(want) {
				t.Errorf("Merge() = %v, want %v", self, want)
			}
		})
	}
}

func TestSet_Separate(t *testing.T) {
	tests := []struct {
		name  string
		self  []int
		other []int
		want  []int
	}{
		{
			name:  "some",
			self:  []int{1, 2, 3, 42, 9001},
			other: []int{42, 9001},
			want:  []int{1, 2, 3},
		},
		{
			name:  "none",
			self:  []int{1, 2, 3, 42, 9001},
			other: nil,
			want:  []int{1, 2, 3, 42, 9001},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			self := New(tt.self...)
			other := New(tt.other...)
			want := New(tt.want...)

			self.Separate(other)

			if !self.Equals(want) {
				t.Errorf("Separate() = %v, want %v", self, want)
			}
		})
	}
}

func TestUnion(t *testing.T) {
	tests := []struct {
		name string
		args [][]int
		want []int
	}{
		{
			name: "expected",
			args: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "nil",
			args: [][]int{
				{1, 2, 3},
				nil,
				{7, 8, 9},
			},
			want: []int{1, 2, 3, 7, 8, 9},
		},
		{
			name: "none",
			args: nil,
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sets := make([]Set[int], len(tt.args))
			for i, s := range tt.args {
				if s != nil {
					sets[i] = New(s...)
				} else {
					sets[i] = nil
				}
			}

			got := Union(sets...)
			want := New(tt.want...)

			if !got.Equals(want) {
				t.Errorf("Union() = %v, want %v", got, want)
			}
		})
	}
}

func TestDifference(t *testing.T) {
	type args struct {
		lhs    []int
		others [][]int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "expected",
			args: args{
				lhs: []int{1, 2, 3, 42, 9001},
				others: [][]int{
					{42},
					{9001},
				},
			},
			want: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lhs := New(tt.args.lhs...)

			others := make([]Set[int], len(tt.args.others))
			for i, s := range tt.args.others {
				others[i] = New(s...)
			}

			want := New(tt.want...)
			got := Difference(lhs, others...)

			if !got.Equals(want) {
				t.Errorf("Difference() = %v, want %v", got, want)
			}
		})
	}
}

func TestIntersection(t *testing.T) {
	tests := []struct {
		name string
		args [][]int
		want []int
	}{
		{
			name: "expected",
			args: [][]int{
				{1, 2, 3, 42, 9001},
				{9001, 42, 5, 7, 6},
				{42, 9, 10, 24, 9001},
			},
			want: []int{9001, 42},
		},
		{
			name: "empty",
			args: [][]int{},
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sets := make([]Set[int], len(tt.args))
			for i, s := range tt.args {
				sets[i] = New(s...)
			}

			got := Intersection(sets...)
			want := New(tt.want...)

			if !got.Equals(want) {
				t.Errorf("Intersection() = %v, want %v", got, want)
			}
		})
	}
}

func TestSymmetricDifference(t *testing.T) {
	tests := []struct {
		name string
		lhs  []int
		rhs  []int
		want []int
	}{
		{
			name: "expected",
			lhs:  []int{1, 2, 42},
			rhs:  []int{9001, 1, 2},
			want: []int{42, 9001},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lhs := New(tt.lhs...)
			rhs := New(tt.rhs...)

			want := New(tt.want...)
			got := SymmetricDifference(lhs, rhs)

			if !want.Equals(got) {
				t.Errorf("SymmetricDifference() = %v, want %v", got, want)
			}
		})
	}
}

func TestFloatingPoint(t *testing.T) {
	t.Run("Contains NaN floats", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()

		_ = New(math.NaN())
	})

	t.Run("Floating Point", func(t *testing.T) {
		a := New(math.Pi, math.Pi, math.Pi, math.E)
		b := New(math.Pi, math.E)

		if !a.Equals(b) {
			t.Errorf("Contains no NaN floats") // TODO proper error message
		}
	})
}
