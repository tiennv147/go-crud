# go-crud
Implement basic function for mysql

# Task 1
Learning go-lang: https://tour.golang.org/welcome/1

# Task 2
Learning Go-kit https://medium.com/swlh/getting-started-with-go-kit-f2ccf71d491f

# Task 3
Build a backend service with 4 APIs 
+ Create
+ Update
+ Delete 
+ Read

---
Model News:
 + Name (string)
 + Thumbnail (url - string)
 + Content (string)
 + Tags (string)
 + CreatedAt (Date)
 + UpdatedAt (Date)
---

# Task 4:
Research gRPC: https://grpc.io/
Research implement JWT Token: https://www.sohamkamani.com/blog/golang/2019-01-01-jwt-authentication/
Write a Authentication service base on GoKit: https://gokit.io/ but expose APIs via gRPC
 + Register - Restful
 + Login response JWT Token - Restful
 + Verify Token (response userID) - gRPC

# Task 5:
Implement authentication for service CRUD news with authentication:
 + Every request call to restful CRUD news service will have a token in header
 + NewsService will call to AuthService to verify token valid or not (if yes will have userID)
 +  Update model News with 2 fields: CreateBy & UpdateBy with userID response from AuthService

---
# Recommend:
1. For ORM (Object-relational mapping) you can consider using that library: https://gorm.io/
2. Go kit: https://gokit.io/
3. Template projects: Reference: https://github.com/golang-standards/project-layout
4. Style guideline: https://github.com/uber-go/guide/blob/master/style.md
