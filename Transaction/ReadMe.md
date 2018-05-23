# Code demo tính năng thực hiện database transaction trong Postgresql qua Go-PG

## Demo gồm 4 bước

1. Connect vào database thông qua hàm ConnectDB() trong file model.go

2. Tạo bảng trong database vừa tạo qua hàm InitSchema()

3. Insert dữ liệu vào bảng qua hàm SaveData()

4. Thực hiện transaction
Các transaction được khai báo và thực thi qua hàm RunTransaction() trong file query.go
1 transaction bao gồm:
- Khai báo bắt đầu 1 transaction: db.Begin()
- Khai báo các tác vụ được thực thi trong transaction: Trong demo này có 2 tác vụ là đọc dữ liệu từ cột counter trong bảng transactions và update dữ liệu của cột này
- Khai báo Rollback: Trong trường hợp có ít nhất 1 tác vụ bị lỗi, Rollback sẽ hủy việc thực thi transaction và đưa database quay về trạng thái trước khi bắt đầu transaction
- Sau khi tất cả các tác vụ chạy thành công, commit các thay đổi lên database

Trong demo này, chúng ta có 10 transaction chạy đồng thời (concurrent), cùng đọc và ghi dữ liệu vào cột counter trong bảng transaction. Để tránh trường hợp 1 transaction đọc dữ liệu trong khi 1 transaction khác chưa kịp commit, Postgresql sử dụng cơ chế Read Commited