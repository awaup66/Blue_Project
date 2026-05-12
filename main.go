package main

import (
	"fmt"
	"myapp/database"
	"myapp/handler"
	"myapp/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func sayHello() {
	fmt.Println("Say Hello Go !!")
}

func greet(name string) {
	fmt.Println("Hello,", name)
}

func add(a int, b int) int {
	return a + b
}

func checkAge(age int) {
	if age >= 18 {
		fmt.Println("อายุ : ", age, "เป็นผู้ใหญ่")
	} else {
		fmt.Println("อายุ : ", age, "เป็นเด็ก")
	}
}

func checkGrade(point int) string {
	if point >= 80 {
		return "A"
	} else if point >= 70 {
		return "B"
	} else if point >= 60 {
		return "C"
	} else if point >= 50 {
		return "D"
	} else {
		return "F"
	}
}

// สร้าง Struct
type User struct {
	Name  string
	Age   int
	Email string
	Point int
}

// สร้าง Method

func (u User) Greet1() string {
	return "Hello,I'm " + u.Name
}

func (u User) IsAdult() bool {
	return u.Age >= 18
}

func (u User) CheckGrade() string {
	if u.Point >= 80 {
		return "Your're grade A"
	} else if u.Point >= 70 {
		return "Your're grade B"
	} else if u.Point >= 60 {
		return "Your're grade C"
	} else if u.Point >= 50 {
		return "Your're grade D"
	} else {
		return "Your're grade F"
	}
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("หารด้วย 0 ไม่ได้")
	}
	return a / b, nil
}

func getUser(name string, id int) (int, error) {
	if id <= 0 {
		return 0, fmt.Errorf("ID ผิดพลาด น้อยกว่าเท่ากับ 0 ไม่ได้")
	}
	return id, nil
}

func main() {
	// var name string = "Bluee"
	// var age int = 25
	// fmt.Println("ประกาศตัวแปรแบบยาว ชื่อ : ", name, "อายุ : ", age)
	// fmt.Println("อายุ : ", age)

	// name1 := "John"
	// age1 := 21
	// fmt.Println("ประกาศตัวแปรแบบสั้น : ", name1, age1)

	// sayHello()
	// greet("John")
	// greet("Jane")
	// greet(name)
	// result := add(15, 20)
	// fmt.Println(result) //35

	// checkAge(20)
	// checkAge(15)

	// grade := checkGrade(75)
	// fmt.Println("เกรดของคุณคือ :", grade)

	// fmt.Println("เกรดของคุณคือ :", checkGrade(65))

	// // แบบที่ 1 วนตามจำนวน
	// for i := 0; i < 5; i++ {
	// 	fmt.Println("แบบที่ 1 วนตามจำนวน : ", i)
	// }

	// // แบบที่ 2 เหมือน while
	// i := 0
	// for i < 5 {
	// 	fmt.Println("แบบที่ 2 — เหมือน while : ", i+1)
	// 	i++
	// }

	// // แบบที่ 3 วนใน Slice (array)
	// fruits := []string{"apple", "banana", "cherry", "mango"}

	// for i, fruit := range fruits {
	// 	fmt.Println(i, fruit)
	// }
	// // ไม่ต้องการ index ใช้ _ แทน
	// for _, fruit := range fruits {
	// 	fmt.Println(fruit)
	// }

	// // ใช้ Struct
	// user := []User{
	// 	{Name: "Blue", Age: 21, Email: "blue@gmail.com"},
	// 	{Name: "John", Age: 25, Email: "John@gmail.com"},
	// 	{Name: "Jane", Age: 29, Email: "Jane@gmail.com"},
	// 	{Name: "Bob", Age: 29, Email: "Bob@gmail.com"},
	// }

	// // fmt.Println(user.Name)
	// // fmt.Println(user.Age)
	// // fmt.Println(user.Email)

	// for i, user1 := range user {
	// 	fmt.Println(i+1, user1.Name, user1.Age, user1.Email)
	// }

	// Step 7 Method
	// user := User{Name: "Blue", Age: 21, Email: "blue@gmail.com", Point: 41}

	// fmt.Println(user.Greet1())
	// fmt.Println(user.IsAdult())
	// fmt.Println(user.CheckGrade())

	// Step 8 Error Handing
	// result, err := divide(10, 2)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// } else {
	// 	fmt.Println("ผลลัพธ์:", result)
	// }

	// result2, err2 := divide(10, 0)
	// if err2 != nil {
	// 	fmt.Println("Error:", err2)
	// } else {
	// 	fmt.Println("ผลลัพธ์:", result2)
	// }

	// name1 := "Blue"
	// result3, err3 := getUser(name1, 0)
	// if err3 != nil {
	// 	fmt.Println("คุณ : ", name1, ",Error", err3)
	// } else {
	// 	fmt.Println("คุณ :", name1, ",ID : ", result3)
	// }

	godotenv.Load()

	database.Connect()

	r := gin.Default()
	// http.HandleFunc("/users", usersHandler)
	// fmt.Println("Server running on : 8080")
	// fmt.Println("http://localhost:8080")
	// http.ListenAndServe(":8080", nil)

	r.POST("/register", handler.Register)
	r.POST("/login", handler.Login)

	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.GET("/users", handler.GetUsers)            // GET /users ดึงทั้งหมด
		auth.GET("/users/:name", handler.GetUserByname) // GET /user by name
		auth.POST("/users", handler.CreateUser)         // POST /users เพิ่มใหม่
		auth.PUT("/users/:id", handler.UpdateUser)      // PUT /users/:name แก้ไข User
		auth.DELETE("/users/:id", handler.DeleteUser)   // DELETE /users/:name ลบ user
	}
	r.Run(":8080")

}
