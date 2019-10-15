## Point One后端面试小项目

### API文档

#### /postWeibo

- **功能：** 记录@相关信息
- **Method：** POST
- **Params**
  - `from - string` @操作发起人ID
  - `to - [string]` @操作受用人ID列表
- **Example**
  ```sh
  # userA在一条微博上@了userB和userC
  ~ $ curl -X POST -d "from=userA&to=userB,userC" "localhost:8080/postWeibo"
  # => {"id":1,"message":"ok"}

  # userA在另一条微博上@了userD
  ~ $ curl -X POST -d "from=userA&to=userD" "localhost:8080/postWeibo"
  # => {"id":2,"message":"ok"}
  ```

#### /suggest

- **功能：** 给出一个userID，返回推荐的用户列表
- **Method：** GET
- **Params：** `user - string` 用户ID
- **Example**
  ```sh
  ~ $ curl -X GET "localhost:8080/suggest?user=userB"
  # => {"suggest":["userC","userD"]}
  ```

#### /reset

- **功能：** 重置数据
- **Method：** GET
- **Example**
  ```sh
  ~ $ curl -X GET "localhost:8080/reset"
  # => {"message": "ok"}
  ```

### 时间复杂度

- 记录的时间复杂度为 $O(n)$
- 推荐的时间复杂度为 $O(m^2+mn)$
