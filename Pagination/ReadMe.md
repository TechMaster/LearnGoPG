# Code demo tính năng thực hiện LIMIT/OFFSET trong Postgresql qua Go-PG

## Demo gồm 4 bước

1. Connect vào database thông qua hàm ConnectDB() trong file model.go

2. Tạo bảng trong database vừa tạo qua hàm InitSchema()

3. Insert dữ liệu vào bảng qua hàm SaveData()

4. Trong file query.go, thực hiện câu lệnh SELECT với LIMIT/OFFSET:
- Khởi tạo 1 struct StudentFilter
- Trong struct StudentFilter, tạo 1 trường có kiểu dữ liệu orm.Pager --> orm.Pager là 1 struct chứa thông tin về limit, offset, max-limit, max-offset --> ta sẽ gán giá trị cho các trường này ở bước tiếp theo
- Khai báo method Filter gắn với struct StudentFilter: Trong method này ta sẽ gán giá trị cho các trường limit, offset, max-limit, max-offset. Method Filter này sẽ được sử dụng với cấu trúc: Db.Apply(Filter) để chạy câu lệnh SELECT với các điều kiện LIMIT, OFFSET vốn đã được thiết lập trong method Filter