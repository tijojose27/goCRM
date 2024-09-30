# goCRM
Sample CRM build with go


The Customer data set has :
ID (Int) -> which is used to identify the customer. This can be used to retrive, delete customers.
Name (String) -> name property of customer
Role (String) -> role of the customer
Email (String) -> email of customer
Phone (String) -> phone of customer
Contacted (Bool True/False) -> This indicates if the customer has been contacted


The project can be started via go run main.go

The project initialises with a set of Customer data on startup
It is launched in port 3000 on localhost

The routes are :
    "localhost:3000/customers" -> GET REQUEST
        This fetches all customers
	"localhost:3000/customers/{id}" -> GET REQUEST
        The {id} is the id of the customer you want to fetch 
	"localhost:3000/customers" -> POST REQUEST
        add customer via post request the customer must have json post with similar to model
	"localhost:3000/customers/{id}" -> DELETE REQUEST
        delete customers based on id
	"localhost:3000/customers/{id}" -> PUT REQUEST
        Updates the customer with the id 