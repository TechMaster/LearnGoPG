Cài đặt portainer để quản lý Docker container
```docker run -d -p 9000:9000 -v /var/run/docker.sock:/var/run/docker.sock -v portainer_data:/data portainer/portainer```

Khởi động postgresql container
```
docker run --name db -e POSTGRES_PASSWORD=123 -d -p 5432:5432 postgres:alpine
```

Cài đặt pgadmin bằng docker
```
docker run -p 8900:80 \
-e "PGADMIN_DEFAULT_EMAIL=cuong@techmaster.vn" \
-e "PGADMIN_DEFAULT_PASSWORD=tuycacchu" \
-d dpage/pgadmin4
```

Vào pgadmin tạo connection đến server. Chú ý địa chỉ IP của server sẽ không phải là localhost mà địa chỉ IPv4 mạng LAN 
gán cho network interface. Dùng lệnh ```ifconfig | grep inet``` để tìm địa chỉ.
user là postgres
pass là 123

Kết nối terminal vào docker container db vừa tạo
```
docker exec -it -u postgres db psql
```

Ở màn hình terminal, gõ lệnh ```\l``` liệt kê danh sách cơ sở dữ liệu
```
postgres=# \l
                                 List of databases
   Name    |  Owner   | Encoding |  Collate   |   Ctype    |   Access privileges
-----------+----------+----------+------------+------------+-----------------------
 postgres  | postgres | UTF8     | en_US.utf8 | en_US.utf8 |
 template0 | postgres | UTF8     | en_US.utf8 | en_US.utf8 | =c/postgres          +
           |          |          |            |            | postgres=CTc/postgres
 template1 | postgres | UTF8     | en_US.utf8 | en_US.utf8 | =c/postgres          +
           |          |          |            |            | postgres=CTc/postgres
(3 rows)
```

Tạo mới một database có tên là demo
```
postgres=# CREATE DATABASE demo;
CREATE DATABASE
```
Kết nối vào database demo
```
\c demo;
```
Tạo bảng book
```sql
CREATE TABLE book (
	id SERIAL PRIMARY KEY NOT NULL,
	title TEXT NOT NULL
);
```
Chèn một số dữ liệu vào bảng book
```sql
demo=# INSERT INTO book (title) VALUES ('Gone with wind');
INSERT 0 1
demo=# INSERT INTO book (title) VALUES ('De men phieu luu ky');
INSERT 0 1
demo=# INSERT INTO book (title) VALUES ('Bi Vo');
INSERT 0 1
demo=# table book;
 id |        title
----+---------------------
  1 | Gone with wind
  2 | De men phieu luu ky
  3 | Bi Vo
(3 rows)
```

Tạo mới một schema
```
CREATE SCHEMA auth;
```
Liệt kê danh sách schema
```
demo=# \dn;
  List of schemas
  Name  |  Owner
--------+----------
 auth   | postgres
 public | postgres
(2 rows)
```

Cho xem danh sách schema hiện thời và chuyển

```sql
CREATE TABLE blog.posts (
 id SERIAL PRIMARY KEY NOT NULL,
 content TEXT NOT NULL,
 authorid INTEGER REFERENCES auth.users(id))
```

SELECT có JOIN

```
SELECT p.id, p.content, u.name FROM blog.posts AS p
JOIN auth.users AS u
ON p.authorid = u.id;
```


