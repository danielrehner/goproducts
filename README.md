# GoProducts
> A sample go application showcasing how to interact with the Go AWS SDKs for DynamoDB and CloudSearch.

### Assumptions
* Full text search is the main priority to allow users to enter any string and find matching products.
* Assumes you have an AWS account with a DynamoDB product table and a CloudSearch search domain configured and an account with read and write access.
* The API is read-only for now and should be as simple as possible for the first iteration.
* Availability of the API is more important than the immediate consistency of the search results.

### API Usage
This version of the API has 4 endpoints:
* The root path 

### Things I would do given more time
* Add the ability to pass a sort field and direction (ascending or descending) into the API.
* Add a lambda function to sync data from DynamoDB to CloudSearch in real time.
* Make product ids UUIDs. 
* Add more test coverage, including benchmark and example tests.
* Consider refactoring the DynamoDB Scan and CloudSearch search functionality behind one interface so the implementation could be swapped out at runtime dynamically based on the use case or need. Also, consider splitting into two files.
* Add more fields to the products table such as a product description or keywords field to allow search terms not included in the Title to still match relevant products. 
* Add the ability to write products through the API.
* Add authentication.
* Improve the data setup / intake process.
* Add searching by category perhaps making use of CloudSearch facets.
* Implement a caching layer. I wanted to try DAX with this project, but AWS doesn't provide a Go SDK searching against DAX yet, only for configuring DAX.
* Add a user comment and/or rating concept to the system.

### Helpful Links
* [Go Gin Boilerplate](https://github.com/vsouza/go-gin-boilerplate)
* [AWS GO SDK examples for DynamoDB](https://github.com/awsdocs/aws-doc-sdk-examples/tree/master/go/example_code/dynamodb)

### Comments
I really enjoyed this project and learned a lot in the process. Thank you in advance to the people who will take time out of their lives to review this sample code. It means a lot to me.

### Contact
If you have any questions, improvement ideas, or just want to go get some chinese takeout and chat, please feel free to reach out:
<danielrehner@gmail.com>
