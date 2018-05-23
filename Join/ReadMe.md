# Code demo tính năng thực hiện Join trong Postgresql qua Go-PG

## Demo gồm 4 bước

1. Connect vào database thông qua hàm ConnectDB() trong file model.go

2. Tạo bảng trong database vừa tạo qua hàm InitSchema()

3. Insert dữ liệu vào bảng qua hàm SaveData()

4. Trong file query.go, thực hiện câu lệnh SELECT với Join:
- Khai báo biến result, biến này có kiểu dữ liệu là 1 slice các struct gồm 4 trường: PhoneId, PhoneNumber, StudentID và StudentNumber. Biến result được dùng để lưu các kết quả trả về khi SELECT các cột phone.id, phone.number, phone.student_id và student.name từ 2 bảng phone và student
- Để JOIN 2 bảng phone và student, dùng lệnh sau:
```go
    Join("JOIN students ON students.id = phone.student_id")
```
- Đoạn code trên sử dụng INNER JOIN giữa 2 bảng student và phone. Để sử dụng LEFT JOIN, RIGHT JOIN, FULL JOIN, chạy các lệnh sau:
```go
    Join("LEFT JOIN students ON students.id = phone.student_id")
```

```go
    Join("RIGHT JOIN students ON students.id = phone.student_id")
```

```go
    Join("FULL JOIN students ON students.id = phone.student_id")
```