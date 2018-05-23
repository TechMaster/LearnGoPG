# Code demo tính năng thực hiện Sub Query trong Postgresql qua Go-PG

## Demo gồm 4 bước

1. Connect vào database thông qua hàm ConnectDB() trong file model.go

2. Tạo bảng trong database vừa tạo qua hàm InitSchema()

3. Insert dữ liệu vào bảng qua hàm SaveData()

4. Trong file query.go, thực hiện câu lệnh SELECT với Sub Query:
- Đầu tiên, thực hiện lệnh SELECT 2 cột: student_id và count(number) từ bảng students:
```go
studentPhone := Db.Model(&phones).
    Column("student_id").
    ColumnExpr("count(number) AS phone_count").
    Group("student_id").
    Order("student_id")
```
Đoạn code trên tương ứng với cú pháp SQL:
```SQL
Select student_id, count(number) as phone_count
from students
Group by student_id
Order by student_id
```
- Kết quả trả về của câu lệnh Select ở trên sẽ được coi như 1 bảng và ta sẽ thực hiện lệnh tính tổng đối với cột phone_count:
```go
var sum struct {
	Total int
}
Db.Model().With("student_phone", studentPhone).Table("student_phone").
    ColumnExpr("sum(student_phone.phone_count) AS total").Select(&sum)
```
Ta đặt 1 tên alias là "student_phone" cho kết quả trả về của lệnh query đầu tiên, sau đó thực hiện lệnh tính tổng đối với cột phone_count. Kết quả tính tổng sẽ được "hứng" bởi struct sum
Cú pháp SQL tương ứng:
```SQL
Select SUM(student_phone.phone_count)
From (
    Select student_id, count(number) as phone_count
    from students
    Group by student_id
    Order by student_id
) AS student_phone
```