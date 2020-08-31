我是光年实验室高级招聘经理。
我在github上访问了你的开源项目，你的代码超赞。你最近有没有在看工作机会，我们在招软件开发工程师，拉钩和BOSS等招聘网站也发布了相关岗位，有公司和职位的详细信息。
我们公司在杭州，业务主要做流量增长，是很多大型互联网公司的流量顾问。公司弹性工作制，福利齐全，发展潜力大，良好的办公环境和学习氛围。
公司官网是http://www.gnlab.com,公司地址是杭州市西湖区古墩路紫金广场B座，若你感兴趣，欢迎与我联系，
电话是0571-88839161，手机号：18668131388，微信号：echo 'bGhsaGxoMTEyNAo='|base64 -D ,静待佳音。如有打扰，还请见谅，祝生活愉快工作顺利。

# GoProducts
> A sample go application showcasing how to interact with the Go AWS SDK for DynamoDB and CloudSearch.

### Assumptions
* Full text search is the main priority to allow users to enter any string and find matching products.
* DynamoDB query operations do not support text search well. Scan operations take longer, are more expensive on large data sets, and do not make sorting easy. Using CloudSearch is likely more optimal and less expensive.
* We have an AWS account with a DynamoDB product table and a CloudSearch search domain configured and an account with read and write access.
* The API is read-only for now and should be as simple as possible for the first iteration.
* Availability of the API is more important than the immediate consistency of the search results.

### API Docs
Info on the API definition can be found in [docs/swagger/swagger.json](docs/swagger/swagger.json).

### Things I would do given more time
* Add the ability to pass a sort field and direction (ascending or descending) into the API.
* Add a lambda function to sync data from DynamoDB to CloudSearch in real time.
* Make product ids UUIDs. 
* Add more test coverage including benchmark and example tests.
* Consider refactoring the DynamoDB Scan and CloudSearch search functionality behind one interface so the implementation could be swapped out at runtime dynamically based on the use case or need. Also, consider splitting into two files.
* Add more fields to the products table such as product description or keywords to allow search terms not included in the Title to still match relevant products. 
* Add the ability to write products through the API.
* Add authentication.
* Consider security and various attack vectors.
* Add per-user API throttling.
* Improve the data setup / intake process.
* Add searching by category, perhaps making use of CloudSearch facets.
* Implement a caching layer.
* Add a user comment and/or rating concept to the system.

### Comments
I really enjoyed this project and learned a lot in the process. Thank you in advance to the people who will take time out of their lives to review this sample code. It means a lot to me.

### Contact
If you have any questions, improvement ideas, or just want to go get some chinese takeout and chat, please feel free to reach out.
