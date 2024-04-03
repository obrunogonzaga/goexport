package entity

type Category struct {
	ID       int
	Name     string
	CourseID []int
}

func (c *Category) NewCategory(id int, name string, courseID []int) *Category {
	return &Category{
		ID:       id,
		Name:     name,
		CourseID: courseID,
	}
}

type Course struct {
	ID         int
	Name       string
	CategoryID int
}
