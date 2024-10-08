### Mobius
![banner.png](resource/banner.png)

由gin+mysql开发的数据集生成api服务。

#### 支持的厂商列表
- OpenAI
- Model Scope

#### api
- /api/v1/generate  

#### 用法
1. 确保拥有 go、mysql 环境
2. 创建数据库 `create database mobius charset=utf8mb4;`
3. 拉取项目到本地 `git clone https://github.com/TCM-GPT/Mobius.git`
4. 修改配置文件 `config/config.yaml` adapter、mysql 参数
5. 运行后端 `go run main.go`
6. 现在你可以使用任意客户端发起请求 `http://localhost:8080/api/v1/generate`
7. 请求参数（必填）：split_prefix、content、restrict、model， 其中 restrict 未截取长度，类型为int

#### 效果
以postman 为例，请求例子如下：
![postman.png](resource/postman.png)

数据库中演示数据效果：
![mysql.png](resource/mysql.png)

### 后续功能开发计划
- 支持更多厂商
- 批量导出，下载功能
- 文件处理，内容清洗
- 日志记录
- 加入用户系统
- 嵌入 LLama-Factory 训练工具

> 非专业开发人员，代:码有点乱，欢迎指正，项目开源仅供交流学习使用，如果你觉得项目不错，可以给个star，感谢！