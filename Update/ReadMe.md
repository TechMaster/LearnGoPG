# Code demo tính năng thực hiện UPDATE trong Postgresql qua Go-PG

## Demo gồm 4 bước

1. Connect vào database thông qua hàm ConnectDB() trong file model.go

2. Tạo bảng trong database vừa tạo qua hàm InitSchema()

3. Insert dữ liệu vào bảng qua hàm SaveData()

4. Trong file query.go, thực hiện câu lệnh UPDATE theo 3 phong cách: Update 1 bản ghi, update nhiều bản ghi sử dụng nhiều struct, update nhiều bản ghi sử dụng 1 slice chứa nhiều struct

* Update 1 bản ghi
- Đầu tiên, khởi tạo 1 struct chứa thông tin về bản ghi cần update. Trong ví dụ này ta sử dụng primary key là cột Id trong bảng students
```go
student := &Student{Id: 1}
```
- Gán giá trị mới cho trường Name của biến student:
```go
student.Name = "Duy dau troc"
```
- Chạy câu lệnh UPDATE trong Go-PG, sử dụng biến student --> Database sẽ update giá trị cột Name của bản ghi có Id = 1 trong bảng students
```go
Db.Update(student)
```
Cú pháp SQL tương ứng:
```SQL
Update students
Set Name = "Duy dau troc"
Where Id = 1
```

* Update nhiều bản ghi sử dụng nhiều struct
- Đầu tiên, khởi tạo các struct chứa thông tin về các bản ghi cần update. Trong ví dụ này ta update cột Age của các bản ghi có Id là 2 và 3
```go
student2 := &Student{
    Id: 2,
    Age: 30,
}
student3 := &Student{
    Id: 3,
    Age: 40,
}
```
- Chạy câu lệnh UPDATE trong Go-PG, sử dụng 2 struct student2 và student3. Ta sử dụng thêm 1 option trong Go-PG là .Column("age") --> Database sẽ update giá trị cột Age của bản ghi có Id = 2 và 3 trong bảng students
```go
Db.Model(student2, student3).Column("age").Update()
```
Cú pháp SQL tương ứng:
```SQL
Update students
Set Age = 30
Where Id = 2

Update students
Set Age = 40
Where Id = 3
```

* Update nhiều bản ghi sử dụng 1 slice chứa nhiều struct
- Tương tự như ví dụ trên nhưng thay vì khai báo nhiều biến kiểu struct, ta khai báo 1 slice chứa các struct của các bản ghi cần update:
```go
students := []Student{
    {
        Id: 1,
        Name: "Duong Ukraine",
        Age: 20,
    },
    {
        Id: 2,
        Name: "Quyen Luc",
        Age: 21,
    },
}
```
- Tiến hành update các cột Name, Age của các bản ghi có Id = 1 và 2:
```go
Db.Model(&students).Column("name", "age").Update()
```
Cú pháp SQL tương ứng:
```SQL
Update students
Set Name = "Duong Ukraine", Age = 30
Where Id = 1

Update students
Set Name = "Quyen Luc", Age = 40
Where Id = 2
```
