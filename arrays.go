package main

func main() {
}

type Contact struct {
	Name  string
	Email string
	Title string
}

// 1. What is the length of the contacts slice?

var contacts = []Contact{}

// 2. Add the following people to the contacts slice.

var jake = Contact{
	Name:  "Jake Overall",
	Email: "jake.overall@boisecodeworks.com",
	Title: "founder",
}

var matt = Contact{
	Name:  "Matt Overall",
	Email: "matt.overall@boisecodeworks.com",
	Title: "founder",
}

var tony = Contact{
	Name:  "Mark Ohnsman",
	Email: "mark@boisecodeworks.com",
	Title: "instructor",
}

var andrew = Contact{
	Name:  "Darryl Kilzer",
	Email: "darryl@boisecodeworks.com",
	Title: "instructor",
}

var tom = Contact{
	Name:  "Tom Day",
	Email: "tom@boisecodeworks.com",
	Title: "instructor",
}

// 3. Woops after adding all of those people to the same contacts list you realized you need a list just the instructors.
// Create a new variable named instructors and populate it using the contacts slice.
