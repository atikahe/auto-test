# Test Result Example
- Say I have a package `set` written in Go where I can add into and remove from a set.
```go
package set

type void struct{}

var exists void

type set struct {
	val map[string]void
}

type Set interface {
	Add(str string)
	Remove(str string)
	Contains(str string) bool
	Values() []string
}

// Initialize a new set
func New() Set {
	return &set{
		val: make(map[string]void),
	}
}

// Add new value to set
func (s *set) Add(str string) {
	s.val[str] = exists
}

// Remove value from set
func (s *set) Remove(str string) {
	delete(s.val, str)
}

// Check if set contains a value
func (s *set) Contains(str string) bool {
	_, ok := s.val[str]
	return ok
}

// Return all the values from set as a slice of string
func (s *set) Values() []string {
	values := make([]string, len(s.val))

	i := 0
	for v := range s.val {
		values[i] = v
		i++
	}

	return values
}
```
- In the directory, run `auto-test -f pkg/set/set.go`
- It will generate a `pkg/set/set_test.go` file with content as the following.
```go
func TestAdd(t *testing.T) {
	s := New()
	s.Add("test")
	if !s.Contains("test") {
		t.Error("Expected set to contain 'test'")
	}
}

func TestRemove(t *testing.T) {
	s := New()
	s.Add("test")
	s.Remove("test")
	if s.Contains("test") {
		t.Error("Expected set to not contain 'test'")
	}
}

func TestContains(t *testing.T) {
	s := New()
	s.Add("test")
	if !s.Contains("test") {
		t.Error("Expected set to contain 'test'")
	}
}

func TestValues(t *testing.T) {
	s := New()
	s.Add("test")
	s.Add("test2")
	values := s.Values()
	if len(values) != 2 {
		t.Error("Expected set to contain 2 values")
	}
}

```

- Great! It looks OK. Now, sometime in the project, I want to use table-based test to cover more cases.
- Run `auto-test -f pkg/set/set.go -p "Use golang's table-based test."`
- It will rewrite  the `pkg/set/set_test.go` file into this.
```go
func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		args string
	}{
		{
			name: "add string",
			args: "test",
		},
		{
			name: "add empty string",
			args: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			s.Add(tt.args)
		})
	}
}

func TestRemove(t *testing.T) {
	tests := []struct {
		name string
		args string
	}{
		{
			name: "remove string",
			args: "test",
		},
		{
			name: "remove empty string",
			args: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			s.Remove(tt.args)
		})
	}
}

func TestContains(t *testing.T) {
	tests := []struct {
		name string
		args string
		want bool
	}{
		{
			name: "contains string",
			args: "test",
			want: true,
		},
		{
			name: "contains empty string",
			args: "",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			s.Add(tt.args)
			if got := s.Contains(tt.args); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValues(t *testing.T) {
	tests := []struct {
		name string
		want []string
	}{
		{
			name: "values",
			want: []string{"test", "test2"},
		},
		{
			name: "values",
			want: []string{"test", "test2", ""},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			s.Add("test")
			s.Add("test2")
			s.Add("")
			if got := s.Values(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Values() = %v, want %v", got, tt.want)
			}
		})
	}
}
```
- Looks OK on first sight, but it's missing `package` and `import` declaration, which I then have to add manually.
- Done!