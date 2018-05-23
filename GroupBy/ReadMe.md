# Code demo tính năng thực hiện Group By trong Postgresql qua Go-PG

## Demo gồm 4 bước

1. Connect vào database thông qua hàm ConnectDB() trong file model.go

2. Tạo bảng trong database vừa tạo qua hàm InitSchema()

3. Insert dữ liệu vào bảng qua hàm SaveData()

4. Trong file query.go, thực hiện câu lệnh SELECT với Group By:
- Khai báo biến queryResults, biến này có kiểu dữ liệu là 1 slice các struct gồm 2 trường: StudentID và PhoneCount. Biến queryResults được dùng để lưu các kết quả trả về khi SELECT các cột Student_id và Count(number) ở bảng Phone
- Để đếm được số lượng điện thoại theo student_id, dùng lệnh sau:
```go
    Column("student_id").
    ColumnExpr("count(number) AS phone_count").
    Group("student_id").
    Select(&queryResults)
```
Cú pháp SQL tương ứng:
```SQL
Select student_id, count(number) as phone_count
from students
Group by student_id
```